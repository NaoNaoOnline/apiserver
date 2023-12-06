package notificationcreatehandler

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/notificationstorage"
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

	var kin string
	{
		kin, err = objKin(tas)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var oid objectid.ID
	{
		oid, err = objVal(tas)
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

	var uid []objectid.ID
	var lid []objectid.ID
	{
		uid, lid, err = h.not.SearchUser(kin, oid, pag)
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

	var now time.Time
	{
		now = time.Now().UTC()
	}

	var nid objectid.ID
	{
		nid = objectid.Random(objectid.Time(now))
	}

	var obj []*notificationstorage.Object
	for i := range uid {
		obj = append(obj, &notificationstorage.Object{
			Crtd: now,
			Evnt: eid,
			Kind: kin,
			List: lid[i],
			Noti: nid,
			Obct: oid,
			User: uid[i],
		})
	}

	{
		err = h.not.CreateNoti(obj)
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

func objKin(tas *task.Task) (string, error) {
	if tas.Meta.Exi(objectlabel.CateObject) {
		return "cate", nil
	}

	if tas.Meta.Exi(objectlabel.HostObject) {
		return "host", nil
	}

	if tas.Meta.Exi(objectlabel.UserObject) {
		return "user", nil
	}

	return "", tracer.Maskf(runtime.ExecutionFailedError, "object label must not be empty")
}

func objVal(tas *task.Task) (objectid.ID, error) {
	if tas.Meta.Exi(objectlabel.CateObject) {
		return objectid.ID(tas.Meta.Get(objectlabel.CateObject)), nil
	}

	if tas.Meta.Exi(objectlabel.HostObject) {
		return objectid.ID(tas.Meta.Get(objectlabel.HostObject)), nil
	}

	if tas.Meta.Exi(objectlabel.UserObject) {
		return objectid.ID(tas.Meta.Get(objectlabel.UserObject)), nil
	}

	return "", tracer.Maskf(runtime.ExecutionFailedError, "object label must not be empty")
}
