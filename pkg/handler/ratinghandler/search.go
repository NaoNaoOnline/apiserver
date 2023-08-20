package ratinghandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/rating"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/ratingstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *rating.SearchI) (*rating.SearchO, error) {
	var err error

	var out []*ratingstorage.Object
	{
		out, err = h.rat.Search()
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *rating.SearchO
	{
		res = &rating.SearchO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &rating.SearchO_Object{
			Intern: &rating.SearchO_Object_Intern{
				Crtd: strconv.Itoa(int(x.Crtd.Unix())),
				Rtng: x.Rtng.String(),
				User: x.User.String(),
			},
			Public: &rating.SearchO_Object_Public{
				Html: x.Html,
				Name: x.Name,
			},
		})
	}

	return res, nil
}
