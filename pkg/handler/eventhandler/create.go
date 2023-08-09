package eventhandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/event"
	"github.com/NaoNaoOnline/apiserver/pkg/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/scoreid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, req *event.CreateI) (*event.CreateO, error) {
	var err error

	var inp *eventstorage.Object
	{
		inp = &eventstorage.Object{
			Cate: inpCat(req.Object[0].Public.Cate),
			Dura: inpDur(req.Object[0].Public.Dura),
			Host: scoreid.String(req.Object[0].Public.Host),
			Link: inpLin(req.Object[0].Public.Link),
			Time: inpTim(req.Object[0].Public.Time),
			User: userid.FromContext(ctx),
		}
	}

	var out *eventstorage.Object
	{
		out, err = h.eve.Create(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *event.CreateO
	{
		res = &event.CreateO{
			Object: []*event.CreateO_Object{
				{
					Intern: &event.CreateO_Object_Intern{
						Crtd: strconv.Itoa(int(out.Crtd.Unix())),
						Evnt: out.Evnt.String(),
					},
				},
			},
		}
	}

	return res, nil
}
