package rulestorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) DeleteRule(inp []*Object) ([]objectstate.String, error) {
	var err error

	var out []objectstate.String
	for i := range inp {
		// Delete the the user specific mappings for user specific search queries.
		{
			err = r.red.Sorted().Delete().Score(rulUse(inp[i].User), inp[i].Rule.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Delete the the list specific mappings for list specific search queries.
		{
			err = r.red.Sorted().Delete().Score(keyfmt.RulLis(inp[i].List), inp[i].Rule.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Delete the the rule specific mappings for rule specific search queries.
		{
			err = r.red.Sorted().Delete().Score(keyfmt.LisRul(inp[i].Rule), inp[i].List.Float())
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

func (r *Redis) DeleteWrkr(inp []*Object) ([]objectstate.String, error) {
	var err error

	var out []objectstate.String
	for i := range inp {
		// Before deleting a nested structure, we need to create a worker task for
		// ensuring the deletion of the rule object and all of its associated data
		// structures.
		{
			err = r.emi.Delete(inp[i].Rule)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Mark the rule object as deleted.
		{
			inp[i].Dltd = time.Now().UTC()
		}

		// Update the rule object with the deletion timestamp.
		{
			err = r.red.Simple().Create().Element(rulObj(inp[i].Rule), musStr(inp[i]))
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
