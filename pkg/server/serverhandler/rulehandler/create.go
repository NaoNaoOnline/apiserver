package rulehandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/rule"
	"github.com/NaoNaoOnline/apiserver/pkg/generic"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/liststorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/rulestorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, req *rule.CreateI) (*rule.CreateO, error) {
	var err error

	var inp []*rulestorage.Object
	for _, x := range req.Object {
		if x.Public != nil {
			inp = append(inp, &rulestorage.Object{
				Excl: inpIDs(x.Public.Excl),
				Incl: inpIDs(x.Public.Incl),
				Kind: x.Public.Kind,
				List: objectid.ID(x.Public.List),
				User: userid.FromContext(ctx),
			})
		}
	}

	//
	// Verify the given input.
	//

	{
		err = h.createVrfy(ctx, inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Create the given resources.
	//

	var out []*rulestorage.Object
	{
		out, err = h.rul.CreateRule(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Create background tasks for the created resources.
	//

	{
		_, err = h.rul.CreateWrkr(out)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct the RPC response.
	//

	var res *rule.CreateO
	{
		res = &rule.CreateO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &rule.CreateO_Object{
			Intern: &rule.CreateO_Object_Intern{
				Crtd: outTim(x.Crtd),
				Rule: x.Rule.String(),
			},
		})
	}

	return res, nil
}

func (h *Handler) createVrfy(ctx context.Context, obj rulestorage.Slicer) error {
	var err error

	for _, x := range obj {
		var lis []*liststorage.Object
		{
			lis, err = h.lis.SearchList([]objectid.ID{x.List})
			if err != nil {
				return tracer.Mask(err)
			}
		}

		if len(lis) != 1 {
			return tracer.Mask(runtime.ExecutionFailedError)
		}

		// Ensure rules cannot be added to lists that have already been deleted.
		if !lis[0].Dltd.IsZero() {
			return tracer.Mask(listDeletedError)
		}

		// Prevent exludes from being added to static lists. Excluding something
		// that you add manually does not make any sense. If you want to exclude
		// something from a static list, then do not add it to that list in the
		// first place.
		if x.Kind == "evnt" && len(x.Excl) != 0 {
			return tracer.Mask(ruleStaticExclError)
		}
	}

	for _, x := range obj {
		var rul rulestorage.Slicer
		{
			rul, err = h.rul.SearchList([]objectid.ID{x.List}, rulestorage.PagAll())
			if err != nil {
				return tracer.Mask(err)
			}
		}

		// At this point a list might just have been created and the first rule is
		// being added. Then the rule kind can be freely choosen without constraint.
		if len(rul) == 0 {
			continue
		}

		// Ensure the maximum amount of rules per list is respected.
		if len(rul) >= 100 {
			return tracer.Mask(ruleListLimitError)
		}

		// Ensure dynamic rules cannot be added to static lists. If the first rule
		// kind of the list that we want to add a new rule to defines kind "evnt",
		// then all other rules added to this very list must also be of kind "evnt".
		// That is to guarantee list type consistency. Dynamic lists contain dynamic
		// rules. Static lists contain static rules.
		if rul[0].Kind == "evnt" && x.Kind != "evnt" {
			return tracer.Mask(ruleStaticListError)
		}

		// Repeat the same check from above but for dynamic lists. If the current
		// list is found to be dynamic due to a dynamic rule, then it is not allowed
		// to add a static rule to the dynamic list.
		if rul[0].Kind != "evnt" && x.Kind == "evnt" {
			return tracer.Mask(ruleDynamicListError)
		}

		// Verify that no duplicated rules can be added to lists.
		for _, y := range rul {
			// If the rule kind of the item that we want to add to the list does not
			// match the rule that we are checking against, then we move on to the
			// next existing rule, if any, because the current rule is not a duplicate
			// due to its different rule kind.
			if x.Kind != y.Kind {
				continue
			}

			// If there is no overlap across includes and excludes the current rule is
			// not a duplicate. If there is any conflict between the rule the caller
			// wants to add and the existing rule we are verifying against, then we
			// return an error.
			if generic.Any(x.Incl, y.Incl) || generic.Any(x.Excl, y.Excl) {
				return tracer.Mask(listRuleDuplicateError)
			}
		}
	}

	return nil
}
