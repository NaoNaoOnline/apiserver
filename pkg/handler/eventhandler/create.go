package eventhandler

import (
	"context"
	"net/url"
	"strconv"
	"strings"
	"time"

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
			Cate: silCat(req.Object[0].Public.Cate),
			Dura: silDur(req.Object[0].Public.Dura),
			Host: scoreid.String(req.Object[0].Public.Host),
			Link: silLin(req.Object[0].Public.Link),
			Time: silTim(req.Object[0].Public.Time),
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

func silCat(str string) []scoreid.String {
	var lis []scoreid.String

	for _, x := range strings.Split(str, ",") {
		lis = append(lis, scoreid.String(x))
	}

	return lis
}

func silDur(str string) time.Duration {
	sec, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	return time.Duration(sec) * time.Second
}

func silLin(str string) *url.URL {
	poi, err := url.Parse(str)
	if err != nil {
		return nil
	}

	return poi
}

func silTim(str string) time.Time {
	sec, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return time.Time{}
	}

	return time.Unix(sec, 0)
}
