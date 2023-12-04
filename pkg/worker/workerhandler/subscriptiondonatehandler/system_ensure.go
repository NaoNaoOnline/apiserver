package subscriptiondonatehandler

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

const (
	paglmt = 10
)

func (h *SystemHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	var err error

	var pnt int
	{
		pnt = int(musNum(tas.Sync.Get(objectlabel.SubsPaging)))
	}

	var min int
	var max int
	{
		min = pnt
		max = pnt + bud.Claim(int(paglmt))
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
	{
		uid, err = h.eve.SearchCrtr(pag)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// We searched for the next page of event creators and did not receive any.
	// That may happen if we processed the full amount of content creators already
	// in the prior cycle of task execution. In that case we end the cycle here by
	// setting the synced paging pointer to zero and return, for the next cycle to
	// begin with the next task execution.
	if len(uid) == 0 {
		tas.Sync.Set(objectlabel.SubsPaging, "0")
		return nil
	}

	var uob []*userstorage.Object
	{
		uob, err = h.use.SearchUser(uid)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var vld []bool
	{
		vld, err = h.sub.VerifyUser(uid)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	if len(uob) != len(vld) {
		return tracer.Maskf(runtime.ExecutionFailedError, "%d != %d", len(uob), len(vld))
	}

	var pre time.Time
	{
		pre = timPre(time.Now().UTC())
	}

	for i := range uob {
		// Any user that has already a premium subscription does not need to be
		// processed anymore. And so we continue with the next user object, if any.
		{
			if uob[i].HasPre() {
				continue
			}
		}

		// Any user that is not recognized as a legitmate content creator is not
		// allowed to receive a donated premium subscription. And so we continue
		// with the next user object, if any.
		{
			if !vld[i] {
				continue
			}
		}

		// We update the user who represents the content creator that did not have a
		// premium subscription before, by setting the up-to-date timestamp until
		// they have now a valid premium subscription. That timestamp is the end of
		// the current month.
		{
			uob[i].Prem = pre
		}

		// Update the current user object with the new premium subscription expiry.
		{
			_, err = h.use.UpdateObct([]*userstorage.Object{uob[i]})
			if err != nil {
				return tracer.Mask(err)
			}
		}

		h.log.Log(
			context.Background(),
			"level", "info",
			"message", "donated premium subscription",
			"user", uob[i].User.String(),
		)
	}

	// If the amount of user IDs we received matches the amount of user IDs we
	// asked for, then we can set the synced paging pointer to the upper end of
	// our current boundary, so that it can become the lower end of the boundary
	// in the next execution of the task. Note that if the current and desired
	// amounts of content creators does not match, then we received less than we
	// asked for, with the reason being that we reached the end of the line of all
	// iterable user IDs. And with that we can set our synced paging pointer back
	// to zero.
	if len(uid) == int(max-min) {
		tas.Sync.Set(objectlabel.SubsPaging, fmt.Sprintf("%d", max))
	} else {
		tas.Sync.Set(objectlabel.SubsPaging, "0")
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

func timPre(tim time.Time) time.Time {
	return time.Date(tim.Year(), tim.Month(), 1, 0, 0, 0, 0, time.UTC).AddDate(0, 1, 0)
}
