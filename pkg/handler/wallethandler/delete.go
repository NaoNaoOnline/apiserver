package wallethandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/wallet"
	"github.com/NaoNaoOnline/apiserver/pkg/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Delete(ctx context.Context, req *wallet.DeleteI) (*wallet.DeleteO, error) {
	var err error

	if userid.FromContext(ctx) == "" {
		return nil, tracer.Mask(userIDEmptyError)
	}

	var wal []objectid.String
	for _, x := range req.Object {
		wal = append(wal, objectid.String(x.Intern.Wllt))
	}

	var obj []*walletstorage.Object
	{
		obj, err = h.wal.SearchWllt(userid.FromContext(ctx), wal)
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
		out, err = h.wal.Delete(obj)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *wallet.DeleteO
	{
		res = &wallet.DeleteO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &wallet.DeleteO_Object{
			Intern: &wallet.DeleteO_Object_Intern{
				Stts: x.String(),
			},
			Public: &wallet.DeleteO_Object_Public{},
		})
	}

	return res, nil
}
