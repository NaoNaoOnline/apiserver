package subscriptionupdatehandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/subscriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (h *UpdateHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	var err error

	var sid string
	{
		sid = tas.Meta.Get(objectlabel.SubsObject)
	}

	var sob []*subscriptionstorage.Object
	{
		sob, err = h.sub.SearchSubs([]objectid.ID{objectid.ID(sid)})
		if subscriptionstorage.IsSubscriptionObjectNotFound(err) {
			h.log.Log(
				context.Background(),
				"level", "warning",
				"message", "stopped processing task",
				"object", sid,
				"reason", "subscription object not found",
			)

			return nil
		} else if err != nil {
			return tracer.Mask(err)
		}
	}

	if len(sob) != 1 {
		return tracer.Mask(runtime.ExecutionFailedError)
	}

	// Only if the verification process was successful we can mark the respective
	// user objects to have a valid premium subscription. If the subscription at
	// hand was found to be invalid, then we skip the user modification and go
	// straight to deleting the distributed lock.
	if sob[0].Stts == objectstate.Success {
		var uid []objectid.ID
		{
			uid, _, err = h.wal.SearchAddr([]string{sob[0].Recv})
			if err != nil {
				return tracer.Mask(err)
			}
		}

		if len(uid) != 1 {
			return tracer.Mask(runtime.ExecutionFailedError)
		}

		var uob []*userstorage.Object
		{
			uob, err = h.use.SearchUser(uid)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		if len(uob) != 1 {
			return tracer.Mask(runtime.ExecutionFailedError)
		}

		// We update the user who represents the subscription receiver with the
		// timestamp until they have a valid premium subscription. That timestamp is
		// the subscription timestamp plus one month. The subscription timestamp
		// defines the beginning of the subscription period, which is the beginning of
		// any given month. And since the premium subscription is for a whole month,
		// it is valid until that month is over.
		{
			uob[0].Prem = sob[0].Unix.AddDate(0, 1, 0)
		}

		{
			_, err = h.use.UpdateObct(uob)
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	// Once the user is updated to have a valid premium subscription we remove the
	// distributed lock for the subscription object which we have just
	// successfully processed.
	var key string
	{
		key = fmt.Sprintf(objectlabel.SubsLocker, sid)
	}

	{
		err = h.loc.Delete(key)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
