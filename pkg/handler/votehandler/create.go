package votehandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/vote"
	"github.com/NaoNaoOnline/apiserver/pkg/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/votestorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, req *vote.CreateI) (*vote.CreateO, error) {
	var err error

	if userid.FromContext(ctx) == "" {
		return nil, tracer.Mask(userIDEmptyError)
	}

	var inp []*votestorage.Object
	for _, x := range req.Object {
		inp = append(inp, &votestorage.Object{
			Desc: objectid.String(x.Public.Desc),
			Rctn: objectid.String(x.Public.Rctn),
			User: userid.FromContext(ctx),
		})
	}

	var out []*votestorage.Object
	{
		out, err = h.vot.Create(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *vote.CreateO
	{
		res = &vote.CreateO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &vote.CreateO_Object{
			Intern: &vote.CreateO_Object_Intern{
				Crtd: strconv.Itoa(int(x.Crtd.Unix())),
				Vote: x.Vote.String(),
			},
		})
	}

	return res, nil
}
