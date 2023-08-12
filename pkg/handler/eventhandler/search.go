package eventhandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/event"
	"github.com/NaoNaoOnline/apiserver/pkg/scoreid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *event.SearchI) (*event.SearchO, error) {
	var err error

	var hos scoreid.String
	{
		hos = scoreid.String(req.Object[0].Intern.Host)
	}

	var cat []scoreid.String
	{
		cat = inpCat(req.Object[0].Intern.Cate)
	}

	var out []*eventstorage.Object
	{
		out, err = h.eve.Search(hos, cat...)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *event.SearchO
	{
		res = &event.SearchO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &event.SearchO_Object{
			Intern: &event.SearchO_Object_Intern{
				Crtd: strconv.Itoa(int(x.Crtd.Unix())),
				Evnt: x.Evnt.String(),
				User: x.User.String(),
			},
			Public: &event.SearchO_Object_Public{
				Cate: outCat(x.Cate),
				Dura: outDur(x.Dura),
				Host: string(x.Host),
				Link: x.Link,
				Time: outTim(x.Time),
			},
		})
	}

	return res, nil
}
