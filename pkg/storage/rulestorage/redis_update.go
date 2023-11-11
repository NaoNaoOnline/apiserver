package rulestorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Update(inp []*Object) ([]objectstate.String, error) {
	var err error

	var sta []objectstate.String
	for i := range inp {
		// At first we need to validate the given input object and, amongst others,
		// ensure whether the provided resource lists are in fact valid. For
		// instance, we cannot update a rule without any rule definition for it.
		{
			err := inp[i].Verify()
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Once we know the rule object is valid, we update the normalized key-value
		// pair so that we can reflect the rule object's internal change.
		{
			err = r.red.Simple().Create().Element(rulObj(inp[i].Rule), musStr(inp[i]))
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
