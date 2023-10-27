package rulehandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/rule"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/rulestorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Delete(ctx context.Context, req *rule.DeleteI) (*rule.DeleteO, error) {
	var err error

	var rul []objectid.ID
	for _, x := range req.Object {
		if x.Intern != nil && x.Intern.Rule != "" {
			rul = append(rul, objectid.ID(x.Intern.Rule))
		}
	}

	var inp []*rulestorage.Object
	{
		inp, err = h.rul.SearchRule(rul)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	for _, x := range inp {
		if userid.FromContext(ctx) != x.User {
			return nil, tracer.Mask(runtime.UserNotOwnerError)
		}
	}

	var out []objectstate.String
	{
		out, err = h.rul.Delete(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct RPC response.
	//

	var res *rule.DeleteO
	{
		res = &rule.DeleteO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &rule.DeleteO_Object{
			Intern: &rule.DeleteO_Object_Intern{
				Stts: x.String(),
			},
			Public: &rule.DeleteO_Object_Public{},
		})
	}

	return res, nil
}
