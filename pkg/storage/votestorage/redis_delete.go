package votestorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/objectstate"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Delete(inp []*Object) ([]objectstate.String, error) {
	var err error

	var out []objectstate.String
	for i := range inp {
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
