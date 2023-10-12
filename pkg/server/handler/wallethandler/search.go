package wallethandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/wallet"
	"github.com/NaoNaoOnline/apiserver/pkg/generic"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *wallet.SearchI) (*wallet.SearchO, error) {
	var out []*walletstorage.Object

	//
	// Search wallets by kind.
	//

	var kin []string
	for _, x := range req.Object {
		if x.Public != nil && x.Public.Kind != "" {
			kin = append(kin, x.Public.Kind)
		}
	}

	if len(kin) != 0 {
		lis, err := h.wal.SearchKind(userid.FromContext(ctx), generic.Uni(kin))
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Search wallets by ID.
	//

	var wal []objectid.ID
	for _, x := range req.Object {
		if x.Intern != nil && x.Intern.Wllt != "" {
			wal = append(wal, objectid.ID(x.Intern.Wllt))
		}
	}

	if len(wal) != 0 {
		lis, err := h.wal.SearchWllt(userid.FromContext(ctx), wal)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Construct RPC response.
	//

	var res *wallet.SearchO
	{
		res = &wallet.SearchO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &wallet.SearchO_Object{
			Intern: &wallet.SearchO_Object_Intern{
				Addr: &wallet.SearchO_Object_Intern_Addr{
					Time: strconv.Itoa(int(x.Addr.Time.Unix())),
				},
				Crtd: strconv.FormatInt(x.Crtd.Unix(), 10),
				User: x.User.String(),
				Wllt: x.Wllt.String(),
			},
			Public: &wallet.SearchO_Object_Public{
				Addr: x.Addr.Data,
				Kind: x.Kind,
			},
		})
	}

	return res, nil
}
