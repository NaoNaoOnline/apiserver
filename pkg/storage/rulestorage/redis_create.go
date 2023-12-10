package rulestorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) CreateRule(inp []*Object) ([]*Object, error) {
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

		// Create the rule specific mappings for rule specific search queries. With
		// that we can search for all lists mapped to a set of rules.
		{
			err = r.red.Sorted().Create().Score(keyfmt.LisRul(inp[i].Rule), inp[i].List.String(), inp[i].List.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Create the list specific mappings for list specific search queries. With
		// that we can search for all rules mapped to a specific list.
		{
			err = r.red.Sorted().Create().Score(keyfmt.RulLis(inp[i].List), inp[i].Rule.String(), inp[i].Rule.Float())
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

func (r *Redis) CreateWrkr(inp []*Object) ([]objectstate.String, error) {
	var err error

	var out []objectstate.String
	for i := range inp {
		{
			err = r.emi.Create(inp[i].Rule)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		{
			out = append(out, objectstate.Started)
		}
	}

	return out, nil
}
