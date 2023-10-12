package descriptionhandler

import (
	"context"
	"time"

	"github.com/NaoNaoOnline/apigocode/pkg/description"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/handler"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Update(ctx context.Context, req *description.UpdateI) (*description.UpdateO, error) {
	var err error

	if userid.FromContext(ctx) == "" {
		return nil, tracer.Mask(handler.UserIDEmptyError)
	}

	var des []objectid.ID
	var pat [][]*descriptionstorage.Patch
	for _, x := range req.Object {
		des = append(des, objectid.ID(x.Intern.Desc))
		pat = append(pat, inpPat(x.Update))
	}

	var inp []*descriptionstorage.Object
	{
		inp, err = h.des.SearchDesc(des)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	for _, x := range inp {
		if userid.FromContext(ctx) != x.User {
			return nil, tracer.Mask(handler.UserNotOwnerError)
		}
		// Ensure descriptions cannot be updated after 5 minutes of their creation.
		if x.Crtd.Add(5 * time.Minute).Before(time.Now().UTC()) {
			return nil, tracer.Mask(descriptionUpdatePeriodError)
		}
	}

	var out []objectstate.String
	{
		out, err = h.des.Update(inp, pat)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct RPC response.
	//

	var res *description.UpdateO
	{
		res = &description.UpdateO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &description.UpdateO_Object{
			Intern: &description.UpdateO_Object_Intern{
				Stts: x.String(),
			},
			Public: &description.UpdateO_Object_Public{},
		})
	}

	return res, nil
}
