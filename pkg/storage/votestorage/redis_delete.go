package votestorage

import (
	"encoding/json"

	"github.com/NaoNaoOnline/apiserver/pkg/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/xh3b4sd/redigo/pkg/simple"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Delete(inp []*Object) ([]objectstate.String, error) {
	var err error

	var out []objectstate.String
	for i := range inp {
		// We need the event ID to delete the event/user specific mappings. So we
		// search for the description object, which provides the event ID.
		var jsn string
		{
			jsn, err = r.red.Simple().Search().Value(desObj(inp[i].Desc))
			if simple.IsNotFound(err) {
				return nil, tracer.Maskf(descriptionNotFoundError, inp[i].Desc.String())
			} else if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		var obj *descriptionstorage.Object
		{
			obj = &descriptionstorage.Object{}
		}

		{
			err = json.Unmarshal([]byte(jsn), obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Delete the event/user specific mappings for event/user specific search
		// queries.
		{
			err = r.red.Sorted().Delete().Value(votUse(obj.Evnt, inp[i].User), inp[i].Vote.String())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Delete the description specific mappings for description specific search
		// queries.
		{
			err = r.red.Sorted().Delete().Value(votDes(inp[i].Desc), inp[i].Vote.String())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Since deletion starts with the normalized key-value pair, we delete it as
		// the very last step, so the operation can eventually be retried.
		{
			_, err = r.red.Simple().Delete().Multi(votObj(inp[i].Vote))
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
