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

		// Add the given notification object to each of the given user's
		// notification feed.
		{
			err = r.red.Sorted().Create().Score(notObj(obj[i].User, obj[i].List), musStr(obj[i]), obj[i].Noti.Float())
			if err != nil {
				return tracer.Mask(err)
			}
		}

		// Limit the per list notification feed for each user to 100 of the latest
		// updates. This number is somewhat arbitrary and can be adjusted with
		// reason.
		{
			err = r.red.Sorted().Delete().Limit(notObj(obj[i].User, obj[i].List), 100)
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}
