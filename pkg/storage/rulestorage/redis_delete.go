package rulestorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/feedstorage"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Delete(inp []*Object) ([]objectstate.String, error) {
	var err error

	var out []objectstate.String
	for i := range inp {
		// Remove the rule owner from the feed for the resources specified by the
		// given rules.
		if inp[i].Kind == "cate" || inp[i].Kind == "host" || inp[i].Kind == "user" {
			for _, y := range inp[i].Incl {
				err = r.red.Sorted().Delete().Score(notKin(inp[i].Kind, y), inp[i].User.Float())
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}
		}

		if inp[i].Kind == "evnt" {
			for _, y := range inp[i].Incl {
				// Delete the event specific mappings for event specific search queries.
				{
					err = r.red.Sorted().Delete().Score(rulEve(y), inp[i].Rule.Float())
					if err != nil {
						return nil, tracer.Mask(err)
					}
				}

				// Remove the event from the static list.
				var obj []*feedstorage.Object
				{
					obj = append(obj, &feedstorage.Object{
						Evnt: y,
						List: inp[i].List,
						User: inp[i].User,
					})
				}

				{
					err = r.fee.DeleteFeed(obj)
					if err != nil {
						return nil, tracer.Mask(err)
					}
				}
			}
		}

		// Delete the the user specific mappings for user specific search queries.
		{
			err = r.red.Sorted().Delete().Score(rulUse(inp[i].User), inp[i].Rule.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Delete the the list specific mappings for list specific search queries.
		{
			err = r.red.Sorted().Delete().Score(rulLis(inp[i].List), inp[i].Rule.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Since the deletion process starts with the normalized key-value pair in
		// the handler, we delete it as the very last step, so the operation can
		// eventually be retried.
		{
			_, err = r.red.Simple().Delete().Multi(rulObj(inp[i].Rule))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		{
			out = append(out, objectstate.Deleted)
		}
	}

	return out, nil
}
