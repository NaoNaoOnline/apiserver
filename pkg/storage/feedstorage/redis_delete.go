package feedstorage

import (
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) DeleteFeed(obj []*Object) error {
	var err error

	for i := range obj {
		// Remove the given feed object from the given user's feed.
		{
			err = r.red.Sorted().Delete().Score(notObj(obj[i].User, obj[i].List), obj[i].Evnt.Float())
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}
