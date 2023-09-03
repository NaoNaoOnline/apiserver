package descriptionhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/description"
	"github.com/NaoNaoOnline/apiserver/pkg/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Update(ctx context.Context, req *description.UpdateI) (*description.UpdateO, error) {
	var err error

	if userid.FromContext(ctx) == "" {
		return nil, tracer.Mask(userIDEmptyError)
	}

	var des []objectid.String
	var pat [][]*descriptionstorage.Patch
	for _, x := range req.Object {
		des = append(des, objectid.String(x.Intern.Desc))
		pat = append(pat, inpPat(x.Update))
	}

	var obj []*descriptionstorage.Object
	{
		obj, err = h.des.SearchDesc(des)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	for _, x := range obj {
		if userid.FromContext(ctx) != x.User {
			return nil, tracer.Mask(userNotOwnerError)
		}
	}

	var out []objectstate.String
	{
		out, err = h.des.Update(obj, pat)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

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
