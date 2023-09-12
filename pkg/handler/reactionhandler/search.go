package reactionhandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/reaction"
	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/reactionstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *reaction.SearchI) (*reaction.SearchO, error) {
	var err error

	var kin []string
	for _, x := range req.Object {
		kin = append(kin, x.Public.Kind)
	}

	var out []*reactionstorage.Object
	{
		out, err = h.rct.SearchKind(objectid.Uni(kin))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *reaction.SearchO
	{
		res = &reaction.SearchO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &reaction.SearchO_Object{
			Intern: &reaction.SearchO_Object_Intern{
				Crtd: strconv.Itoa(int(x.Crtd.Unix())),
				Rctn: x.Rctn.String(),
				User: x.User.String(),
			},
			Public: &reaction.SearchO_Object_Public{
				Html: x.Html,
				Kind: x.Kind,
				Name: x.Name,
			},
		})
	}

	return res, nil
}
