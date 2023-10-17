package walletstorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Delete(inp []*Object) ([]objectstate.String, error) {
	var err error

	var out []objectstate.String
	for i := range inp {
		// Delete the user specific mappings for user specific search queries.
		{
			err = r.red.Sorted().Delete().Score(walUse(inp[i].User), inp[i].Wllt.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Delete the wallet kind mappings for wallet kind search queries.
		{
			err = r.red.Sorted().Delete().Score(walKin(inp[i].User, inp[i].Kind), inp[i].Wllt.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Since the deletion process starts with the normalized key-value pair in
		// the handler, we delete it as the very last step, so the operation can
		// eventually be retried. Here we also delete the wallet address mappings in
		// one go.
		{
			add := walAdd(inp[i].Addr.Data)
			obj := walObj(inp[i].User, inp[i].Wllt)

			_, err = r.red.Simple().Delete().Multi(add, obj)
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
