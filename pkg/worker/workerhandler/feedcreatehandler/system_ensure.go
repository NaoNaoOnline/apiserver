package feedcreatehandler

import (
	"context"
	"fmt"
	"slices"
	"strconv"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/generic"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/feedstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/rulestorage"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

const (
	// paglmt is the amount of users that we process at a time during a single
	// cycle of task execution.
	paglmt = 5
)

func (h *SystemHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	var err error

	var eid objectid.ID
	{
		eid = objectid.ID(tas.Meta.Get(objectlabel.EvntObject))
	}

	// Fetch the event object that got created and for which we want to process
	// all relevant feeds. If the relevant event has been deleted intermittendly,
	// we stop processing here.
	var eob []*eventstorage.Object
	{
		eob, err = h.eve.SearchEvnt("", []objectid.ID{eid})
		if eventstorage.IsEventObjectNotFound(err) {
			h.log.Log(
				context.Background(),
				"level", "warning",
				"message", "stopped processing task",
				"reason", "event object could not be found",
			)

			return nil
		} else if err != nil {
			return tracer.Mask(err)
		}
	}

	var oid objectid.ID
	var kin string
	{
		oid, kin, err = objKin(tas)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var pnt int
	{
		pnt = int(musNum(tas.Sync.Get(task.Paging)))
	}

	var min int
	var max int
	{
		min = pnt
		max = pnt + bud.Claim(paglmt)
	}

	// It may happen, for whatever reason, that the given budget is exhausted and
	// the paging range could not be computed. In that case we should get to know
	// about it.
	if (max - min) == 0 {
		h.log.Log(
			context.Background(),
			"level", "warning",
			"message", "stopped processing task",
			"reason", "paging range could not be computed",
		)

		return nil
	}

	var pag [2]int
	{
		pag = [2]int{
			min,
			max,
		}
	}

	// Search for the relevant user and list IDs. We store these as ID pairs. The
	// returned lists are synchronized in a way so that each pair shares the same
	// index between both lists.
	var uid []objectid.ID
	var lid []objectid.ID
	{
		uid, lid, err = h.fee.SearchUser(kin, oid, pag)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// We searched for the next page of user IDs and may did not receive any. That
	// may happen if we processed the full amount of user IDs already in the prior
	// cycle of task execution. In that case we end the cycle here by setting the
	// synced paging pointer to zero and return, for the next cycle to begin with
	// the next task execution.
	if len(uid) == 0 {
		tas.Sync.Set(task.Paging, "0")
		return nil
	}

	// Search for all rules for all of the given list IDs. The resulting rules
	// will be grouped together by list ID in order to find out which list should
	// actually receive the current event feed based on their exclusion rules.
	var sli rulestorage.Slicer
	{
		sli, err = h.rul.SearchList(lid, rulestorage.PagAll())
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		uid, lid = slcIDs(uid, lid, filIDs(eob[0], sli))
	}

	var now time.Time
	{
		now = time.Now().UTC()
	}

	var nid objectid.ID
	{
		nid = objectid.Random(objectid.Time(now))
	}

	var obj []*feedstorage.Object
	for i := range uid {
		obj = append(obj, &feedstorage.Object{
			Crtd: now,
			Evnt: eid,
			Feed: nid,
			Kind: kin,
			List: lid[i],
			Obct: oid,
			User: uid[i],
		})
	}

	{
		err = h.fee.CreateFeed(obj)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// If the amount of user IDs we received matches the amount of user IDs we
	// asked for, then we can set the synced paging pointer to the upper end of
	// our current boundary, so that it can become the lower end of the boundary
	// during the next execution cycle of the task. Note that if the current and
	// desired amounts of user IDs does not match, then we received less than we
	// asked for, with the reason being that we reached the end of the line of all
	// iterable user IDs. And with that we can then set our synced paging pointer
	// back to zero.
	if len(uid) == int(max-min) {
		tas.Sync.Set(task.Paging, fmt.Sprintf("%d", max))
	} else {
		tas.Sync.Set(task.Paging, "0")
	}

	return nil
}

func musNum(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}

	return num
}

func filIDs(eob *eventstorage.Object, sli rulestorage.Slicer) []objectid.ID {
	var fil []objectid.ID

	for k, v := range sli.List() {
		var cat []objectid.ID
		var hos []objectid.ID
		var use []objectid.ID
		{
			cat = eob.Cate
			hos = eob.Host
			use = []objectid.ID{eob.User}
		}

		// Filter out lists that exclude category IDs specified in the given event
		// object.
		if generic.Any(cat, v.Fltr().Cate()) {
			continue
		}

		// Filter out lists that exclude host IDs specified in the given event
		// object.
		if generic.Any(hos, v.Fltr().Host()) {
			continue
		}

		// Filter out lists that exclude user IDs specified in the given event
		// object.
		if generic.Any(use, v.Fltr().User()) {
			continue
		}

		{
			fil = append(fil, k)
		}
	}

	return fil
}

func objKin(tas *task.Task) (objectid.ID, string, error) {
	if tas.Meta.Exi(objectlabel.CateObject) {
		return objectid.ID(tas.Meta.Get(objectlabel.CateObject)), "cate", nil
	}

	if tas.Meta.Exi(objectlabel.HostObject) {
		return objectid.ID(tas.Meta.Get(objectlabel.HostObject)), "host", nil
	}

	if tas.Meta.Exi(objectlabel.UserObject) {
		return objectid.ID(tas.Meta.Get(objectlabel.UserObject)), "user", nil
	}

	return "", "", tracer.Maskf(runtime.ExecutionFailedError, "object label must not be empty")
}

func slcIDs(uid []objectid.ID, lid []objectid.ID, slc []objectid.ID) ([]objectid.ID, []objectid.ID) {
	var use []objectid.ID
	var lis []objectid.ID

	for _, x := range slc {
		var ind int
		{
			ind = slices.Index(lid, x)
		}

		{
			use = append(use, uid[ind])
			lis = append(lis, lid[ind])
		}
	}

	return use, lis
}
