package wallethandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/wallet"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Update(ctx context.Context, req *wallet.UpdateI) (*wallet.UpdateO, error) {
	var err error

	var use objectid.ID
	{
		use = userid.FromContext(ctx)
	}

	var out []*walletstorage.Object
	var sta []objectstate.String

	//
	// Update by JSON-Patch.
	//

	{
		var wal []objectid.ID
		var pat [][]*walletstorage.Patch
		for _, x := range req.Object {
			if x.Intern != nil && x.Update != nil && x.Intern.Wllt != "" {
				wal = append(wal, objectid.ID(x.Intern.Wllt))
				pat = append(pat, inpPat(x.Update))
			}
		}

		if len(wal) != 0 {
			var sli walletstorage.Slicer
			{
				sli, err = h.wal.SearchKind(use, []string{"eth"})
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			//
			// Verify the given input.
			//

			{
				err = h.updateVrfyPtch(ctx, sli, wal, pat)
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			//
			// Update the given resources.
			//

			{
				ols, sls, err := h.wal.UpdatePtch(sli.Slct(wal...), pat)
				if err != nil {
					return nil, tracer.Mask(err)
				}

				out = append(out, ols...)
				sta = append(sta, sls...)
			}
		}
	}

	//
	// Update wallet signature.
	//

	{
		upd := map[objectid.ID]walletstorage.Object{}
		var wal []objectid.ID
		for _, x := range req.Object {
			if x.Intern != nil && x.Public != nil && x.Intern.Wllt != "" {
				upd[objectid.ID(x.Intern.Wllt)] = walletstorage.Object{
					Mess: x.Public.Mess,
					Pubk: x.Public.Pubk,
					Sign: x.Public.Sign,
				}

				{
					wal = append(wal, objectid.ID(x.Intern.Wllt))
				}
			}
		}

		if len(wal) != 0 {
			var obj []*walletstorage.Object
			{
				obj, err = h.wal.SearchWllt(userid.FromContext(ctx), wal)
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			//
			// Verify the given input.
			//

			{
				err = h.updateVrfySign(ctx, obj, upd)
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			//
			// Update the given resources.
			//

			{
				ols, sls, err := h.wal.UpdateSign(obj)
				if err != nil {
					return nil, tracer.Mask(err)
				}

				out = append(out, ols...)
				sta = append(sta, sls...)
			}
		}
	}

	//
	// Construct the RPC response.
	//

	var res *wallet.UpdateO
	{
		res = &wallet.UpdateO{}
	}

	for i := range out {
		res.Object = append(res.Object, &wallet.UpdateO_Object{
			Intern: &wallet.UpdateO_Object_Intern{
				Addr: &wallet.UpdateO_Object_Intern_Addr{
					Time: outTim(out[i].Addr.Time),
				},
				Labl: &wallet.UpdateO_Object_Intern_Labl{
					Time: outTim(out[i].Labl.Time),
				},
				Stts: sta[i].String(),
			},
			Public: &wallet.UpdateO_Object_Public{},
		})
	}

	return res, nil
}

func (h *Handler) updateVrfyPtch(
	ctx context.Context,
	sli walletstorage.Slicer,
	wal []objectid.ID,
	pat walletstorage.PatchSlicer,
) error {
	var use objectid.ID
	{
		use = userid.FromContext(ctx)
	}

	// Here we iterate through all existing user wallets in order to find out
	// whether the user does already have an accounting wallet, while trying to
	// designate another one.
	if sli.Labl(objectlabel.WalletAccounting) {
		for i := range sli.Slct(wal...) {
			if pat.AddLab(i, objectlabel.WalletAccounting) {
				return tracer.Mask(walletLabelAccountingError)
			}
		}
	}

	// Here we iterate through the subset of user wallets subject to the patch RPC
	// in order to find out whether any conflicting operation was sked to be
	// processed.
	for i, x := range sli.Slct(wal...) {
		if use != x.User {
			return tracer.Mask(runtime.UserNotOwnerError)
		}

		for _, y := range x.Labl.Data {
			// Ensure wallet labels can only be added if they do not already exist. If
			// RemLab returns false, given the existing label y, then it means that a
			// patch defines a label to be added that does not already exist.
			if pat.AddLab(i, y) {
				return tracer.Mask(walletLabelAlreadyExistsError)
			}

			// Ensure wallet labels can only be removed if they do already exist. If
			// RemLab returns true, given the existing label y, then it means that a
			// patch defines the existing label y to be removed.
			if !pat.RemLab(i, y) {
				return tracer.Mask(walletLabelNotFoundError)
			}
		}
	}

	return nil
}

func (h *Handler) updateVrfySign(
	ctx context.Context,
	obj []*walletstorage.Object,
	upd map[objectid.ID]walletstorage.Object,
) error {
	var use objectid.ID
	{
		use = userid.FromContext(ctx)
	}

	for i, x := range obj {
		if use != x.User {
			return tracer.Mask(runtime.UserNotOwnerError)
		}

		{
			obj[i].Mess = upd[x.Wllt].Mess
			obj[i].Pubk = upd[x.Wllt].Pubk
			obj[i].Sign = upd[x.Wllt].Sign
		}
	}

	return nil
}
