package wallethandler

import (
	"context"
	"strconv"
	"strings"

	"github.com/NaoNaoOnline/apigocode/pkg/wallet"
	"github.com/NaoNaoOnline/apiserver/pkg/generic"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/limiter"
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
	// Construct the RPC response.
	//

	var res *wallet.SearchO
	{
		res = &wallet.SearchO{}
	}

	if limiter.Log(len(out)) {
		h.log.Log(
			context.Background(),
			"level", "warning",
			"message", "search response truncated",
			"limit", strconv.Itoa(limiter.Default),
			"resource", "wallet",
			"total", strconv.Itoa(len(out)),
		)
	}

	for _, x := range out[:limiter.Len(len(out))] {
		// Wallets marked to be deleted cannot be searched anymore.
		if !x.Dltd.IsZero() {
			continue
		}

		res.Object = append(res.Object, &wallet.SearchO_Object{
			Intern: &wallet.SearchO_Object_Intern{
				Addr: &wallet.SearchO_Object_Intern_Addr{
					Time: outTim(x.Addr.Time),
				},
				Crtd: strconv.FormatInt(x.Crtd.Unix(), 10),
				Labl: &wallet.SearchO_Object_Intern_Labl{
					Time: outTim(x.Labl.Time),
				},
				User: x.User.String(),
				Wllt: x.Wllt.String(),
			},
			Public: &wallet.SearchO_Object_Public{
				Addr: x.Addr.Data,
				Kind: x.Kind,
				Labl: strings.Join(x.Labl.Data, ","),
			},
		})
	}

	return res, nil
}
