package rulehandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/rule"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
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

	var out []*rulestorage.Object
	{
		out, err = h.rul.Create(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct RPC response.
	//

	var res *rule.CreateO
	{
		res = &rule.CreateO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &rule.CreateO_Object{
			Intern: &rule.CreateO_Object_Intern{
				Crtd: strconv.FormatInt(x.Crtd.Unix(), 10),
				Rule: x.Rule.String(),
			},
		})
	}

	return res, nil
}
