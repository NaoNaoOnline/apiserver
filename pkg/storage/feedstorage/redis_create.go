package feedstorage

import (
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) CreateFeed(obj []*Object) error {
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
		// state. It might happen that an event got already added to the list that
		// we are dealing with during any given iteration. In that case we continue
		// with the next feed object, if any. The existence check is also the reason
		// why the feed objects are indexed with the event object score. Because
		// multiple dynamic rules may cause the same event ID to be added to any
		// given feed, read custom list. And we want to prevent that.
		{
			exi, err := r.red.Sorted().Exists().Score(notObj(obj[i].User, obj[i].List), obj[i].Evnt.Float())
			if err != nil {
				return tracer.Mask(err)
			}

			if exi {
				continue
			}
		}

		// Add the given feed object to each of the given user's feed.
		{
			err = r.red.Sorted().Create().Score(notObj(obj[i].User, obj[i].List), musStr(obj[i]), obj[i].Evnt.Float())
			if err != nil {
				return tracer.Mask(err)
			}
		}

		// Limit the per list feed for each user to 100 of the latest updates. This
		// number is somewhat arbitrary and can be adjusted with reason.
		{
			err = r.red.Sorted().Delete().Limit(notObj(obj[i].User, obj[i].List), 1000)
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}
