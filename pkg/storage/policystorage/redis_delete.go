package policystorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) DeletePlcy(inp []*Object) ([]objectstate.String, error) {
	var err error

	var out []objectstate.String
	for i := range inp {
		// Delete the system specific mappings for system specific search queries.
		{
			err = r.red.Sorted().Delete().Score(polSys(inp[i].Kind, inp[i].Syst, inp[i].Memb), inp[i].Plcy.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Delete the member specific mappings for member specific search queries.
		{
			err = r.red.Sorted().Delete().Score(polMem(inp[i].Kind, inp[i].Memb), inp[i].Plcy.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Delete the kind specific mappings for kind specific search queries.
		{
			err = r.red.Sorted().Delete().Score(polKin(inp[i].Kind), inp[i].Plcy.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Since the deletion process starts with the normalized key-value pair in
		// the handler, we delete it as the very last step, so the operation can
		// eventually be retried.
		{
			_, err = r.red.Simple().Delete().Multi(polObj(inp[i].Plcy))
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
