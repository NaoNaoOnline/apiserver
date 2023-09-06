package eventhandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/event"
	"github.com/NaoNaoOnline/apiserver/pkg/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, req *event.CreateI) (*event.CreateO, error) {
	var err error

	if userid.FromContext(ctx) == "" {
		return nil, tracer.Mask(userIDEmptyError)
	}

	var inp []*eventstorage.Object
	for _, x := range req.Object {
		inp = append(inp, &eventstorage.Object{
			Cate: inpLab(x.Public.Cate),
			Dura: inpDur(x.Public.Dura),
			Host: inpLab(x.Public.Host),
			Link: x.Public.Link,
			Time: inpTim(x.Public.Time),
			User: userid.FromContext(ctx),
		})
	}

	var out []*eventstorage.Object
	{
		out, err = h.eve.Create(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *event.CreateO
	{
		res = &event.CreateO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &event.CreateO_Object{
			Intern: &event.CreateO_Object_Intern{
				Crtd: strconv.Itoa(int(x.Crtd.Unix())),
				Evnt: x.Evnt.String(),
			},
		})
	}

	return res, nil
}