package descriptionhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/description"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Delete(ctx context.Context, req *description.DeleteI) (*description.DeleteO, error) {
	var err error

	if userid.FromContext(ctx) == "" {
		return nil, tracer.Mask(userIDEmptyError)
	}

	var des []objectid.ID
	for _, x := range req.Object {
		des = append(des, objectid.ID(x.Intern.Desc))
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
		out, err = h.des.DeleteWrkr(obj)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *description.DeleteO
	{
		res = &description.DeleteO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &description.DeleteO_Object{
			Intern: &description.DeleteO_Object_Intern{
				Stts: x.String(),
			},
			Public: &description.DeleteO_Object_Public{},
		})
	}

	return res, nil
}
