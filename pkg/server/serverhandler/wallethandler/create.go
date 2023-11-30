package wallethandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/wallet"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, req *wallet.CreateI) (*wallet.CreateO, error) {
	var err error

	var inp []*walletstorage.Object
	for _, x := range req.Object {
		if x.Public != nil {
			inp = append(inp, &walletstorage.Object{
				Kind: x.Public.Kind,
				Mess: x.Public.Mess,
				Pubk: x.Public.Pubk,
				Sign: x.Public.Sign,
				User: userid.FromContext(ctx),
			})
		}
	}

	var out []*walletstorage.Object
	{
		out, err = h.wal.CreateXtrn(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct the RPC response.
	//

	var res *wallet.CreateO
	{
		res = &wallet.CreateO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &wallet.CreateO_Object{
			Intern: &wallet.CreateO_Object_Intern{
				Crtd: outTim(x.Crtd),
				Wllt: x.Wllt.String(),
			},
		})
	}

	return res, nil
}
