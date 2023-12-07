package notificationstorage

import (
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) CreateNoti(obj []*Object) error {
	var err error

	for i := range obj {
		// At first we need to validate the given input object.
		{
			err = obj[i].Verify()
			if err != nil {
				return tracer.Mask(err)
			}
		}

		// The event objects given to us here represent a collection of desired
		// state. It might happen that an event got already happened to the list
		// that we are dealing with during any given iteration. In that case we
		// acknowledge the already exists error and continue with the next
		// notification object, if any. This is also the reason why the notification
		// objects are indexed with the event object score. So that we can ensure
		// unique
		{
			exi, err := r.red.Sorted().Exists().Score(notObj(obj[i].User, obj[i].List), obj[i].Evnt.Float())
			if err != nil {
				return tracer.Mask(err)
			}

			if exi {
				continue
			}
		}

		// Add the given notification object to each of the given user's
		// notification feed.
		{
			err = r.red.Sorted().Create().Score(notObj(obj[i].User, obj[i].List), musStr(obj[i]), obj[i].Evnt.Float())
			if err != nil {
				return tracer.Mask(err)
			}
		}

		// Limit the per list notification feed for each user to 100 of the latest
		// updates. This number is somewhat arbitrary and can be adjusted with
		// reason.
		{
			err = r.red.Sorted().Delete().Limit(notObj(obj[i].User, obj[i].List), 100) // TODO longer lists for premium
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}
