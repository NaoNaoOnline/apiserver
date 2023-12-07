package notificationstorage

import (
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) DeleteNoti(obj []*Object) error {
	var err error

	for i := range obj {
		// Remove the given notification object from the given user's notification
		// feed.
		{
			err = r.red.Sorted().Delete().Score(notObj(obj[i].User, obj[i].List), obj[i].Evnt.Float())
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}
