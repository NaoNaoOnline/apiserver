package votehandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/vote"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/votestorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *vote.SearchI) (*vote.SearchO, error) {
	var err error

	var des []objectid.ID
	for _, x := range req.Object {
		if x.Public != nil && x.Public.Desc != "" {
			des = append(des, objectid.ID(x.Public.Desc))
		}
	}

	var out []*votestorage.Object
	{
		out, err = h.vot.SearchDesc(des)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct RPC response.
	//

	var res *vote.SearchO
	{
		res = &vote.SearchO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &vote.SearchO_Object{
			Intern: &vote.SearchO_Object_Intern{
				Crtd: strconv.FormatInt(x.Crtd.Unix(), 10),
				User: x.User.String(),
				Vote: x.Vote.String(),
			},
			Public: &vote.SearchO_Object_Public{
				Desc: x.Desc.String(),
				Rctn: x.Rctn.String(),
			},
		})
	}

	return res, nil
}
