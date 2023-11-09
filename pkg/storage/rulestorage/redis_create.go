package rulestorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Create(inp []*Object) ([]*Object, error) {
	var err error

	for i := range inp {
		// At first we need to validate the given input object and, amongst others,
		// ensure whether the list mapped to the rule does already exist. For
		// instance, we cannot create a rule for a list that is not there.
		{
			err := inp[i].Verify()
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		var key string
		{
			key = lisObj(inp[i].List)
		}

		// Ensure the referenced list object does in fact exist.
		{
			cou, err := r.red.Simple().Exists().Multi(key)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			if cou != 1 {
				return nil, tracer.Maskf(listObjectNotFoundError, "%#v", key)
			}
		}

		// Ensure the maximum allowed amount of rules per list.
		{
			cou, err := r.red.Sorted().Metric().Count(rulLis(inp[i].List))
			if err != nil {
				return nil, tracer.Mask(err)
			}

			if cou >= 100 {
				return nil, tracer.Mask(ruleListLimitError)
			}
		}

		var now time.Time
		{
			now = time.Now().UTC()
		}

		{
			inp[i].Crtd = now
			inp[i].Rule = objectid.Random(objectid.Time(now))
		}

		// Once we know the associated list exists, we create the normalized
		// key-value pair for the rule object, so that we can search for rule
		// objects using their IDs.
		{
			err = r.red.Simple().Create().Element(rulObj(inp[i].Rule), musStr(inp[i]))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Create the list specific mappings for list specific search queries. With
		// that we can search for all rules mapped to a specific list.
		{
			err = r.red.Sorted().Create().Score(rulLis(inp[i].List), inp[i].Rule.String(), inp[i].Rule.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Create the user specific mappings for user specific search queries. With
		// that we can show the user all rules they created.
		{
			err = r.red.Sorted().Create().Score(rulUse(inp[i].User), inp[i].Rule.String(), inp[i].Rule.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
	}

	return inp, nil
}
