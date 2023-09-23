package wallethandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/wallet"
	"github.com/NaoNaoOnline/apiserver/pkg/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Update(ctx context.Context, req *wallet.UpdateI) (*wallet.UpdateO, error) {
	var err error

	if userid.FromContext(ctx) == "" {
		return nil, tracer.Mask(userIDEmptyError)
	}

	upd := map[objectid.String]walletstorage.Object{}
	for _, x := range req.Object {
		upd[objectid.String(x.Intern.Wllt)] = walletstorage.Object{
			Mess: x.Public.Mess,
			Pubk: x.Public.Pubk,
			Sign: x.Public.Sign,
		}
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

	for i, x := range obj {
		if userid.FromContext(ctx) != x.User {
			return nil, tracer.Mask(userNotOwnerError)
		}

		{
			obj[i].Mess = upd[x.Wllt].Mess
			obj[i].Pubk = upd[x.Wllt].Pubk
			obj[i].Sign = upd[x.Wllt].Sign
		}
	}

	var out []objectstate.String
	{
		out, err = h.wal.Update(obj)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *wallet.UpdateO
	{
		res = &wallet.UpdateO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &wallet.UpdateO_Object{
			Intern: &wallet.UpdateO_Object_Intern{
				Stts: x.String(),
			},
			Public: &wallet.UpdateO_Object_Public{},
		})
	}

	return res, nil
}
