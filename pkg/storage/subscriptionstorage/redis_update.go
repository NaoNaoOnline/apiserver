package subscriptionstorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Update(inp []*Object) ([]objectstate.String, error) {
	var err error

	var sta []objectstate.String
	for i := range inp {
		// At first we need to validate the given input object and, amongst others,
		// whether the creator addresses comply with the expected format.
		{
			err := inp[i].Verify()
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Once we know the subscription object is valid, we update the normalized
		// key-value pair so that we can reflect the subscription object's internal
		// change.
		{
			err = r.red.Simple().Create().Element(subObj(inp[i].Subs), musStr(inp[i]))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		{
			sta = append(sta, objectstate.Updated)
		}
	}

	return sta, nil
}
