package listhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/list"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/liststorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Delete(ctx context.Context, req *list.DeleteI) (*list.DeleteO, error) {
	var err error

	var lis []objectid.ID
	for _, x := range req.Object {
		if x.Intern != nil && x.Intern.List != "" {
			lis = append(lis, objectid.ID(x.Intern.List))
		}
	}

	var inp []*liststorage.Object
	{
		inp, err = h.lis.SearchList(lis)
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
		out, err = h.lis.DeleteWrkr(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct the RPC response.
	//

	var res *list.DeleteO
	{
		res = &list.DeleteO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &list.DeleteO_Object{
			Intern: &list.DeleteO_Object_Intern{
				Stts: x.String(),
			},
			Public: &list.DeleteO_Object_Public{},
		})
	}

	return res, nil
}
