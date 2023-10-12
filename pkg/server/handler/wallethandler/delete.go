package wallethandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/wallet"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/handler"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Delete(ctx context.Context, req *wallet.DeleteI) (*wallet.DeleteO, error) {
	var err error

	var wal []objectid.ID
	for _, x := range req.Object {
		if x.Intern != nil && x.Intern.Wllt != "" {
			wal = append(wal, objectid.ID(x.Intern.Wllt))
		}
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
			return nil, tracer.Mask(handler.UserNotOwnerError)
		}
	}

	var out []objectstate.String
	{
		out, err = h.wal.Delete(obj)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct RPC response.
	//

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
