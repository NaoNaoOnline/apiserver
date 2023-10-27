package listhandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/list"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/liststorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *list.SearchI) (*list.SearchO, error) {
	var err error

	var use objectid.ID
	for _, x := range req.Object {
		if x.Intern != nil && x.Intern.User != "" {
			use = objectid.ID(x.Intern.User)
		}
	}

	{
		if userid.FromContext(ctx) != use {
			return nil, tracer.Mask(runtime.UserNotOwnerError)
		}
	}

	var out []*liststorage.Object
	{
		out, err = h.lis.SearchUser(use)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct RPC response.
	//

	var res *list.SearchO
	{
		res = &list.SearchO{}
	}

	for _, x := range out {
		// Lists marked to be deleted cannot be searched anymore.
		if !x.Dltd.IsZero() {
			continue
		}

		res.Object = append(res.Object, &list.SearchO_Object{
			Intern: &list.SearchO_Object_Intern{
				Crtd: strconv.FormatInt(x.Crtd.Unix(), 10),
				List: x.List.String(),
				User: x.User.String(),
			},
			Public: &list.SearchO_Object_Public{
				Desc: x.Desc,
			},
		})
	}

	return res, nil
}
