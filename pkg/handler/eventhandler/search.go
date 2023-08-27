package eventhandler

import (
	"context"
	"fmt"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/event"
	"github.com/NaoNaoOnline/apiserver/pkg/scoreid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *event.SearchI) (*event.SearchO, error) {
	for _, x := range req.Object {
		if x.Intern.Evnt != "" && (x.Public.Cate != "" || x.Public.Host != "") {
			return nil, tracer.Mask(fmt.Errorf("request object must not contain evnt if either cate or host is given"))
		}
	}

	var out []*eventstorage.Object

	var evn []scoreid.String
	for _, x := range req.Object {
		if x.Intern.Evnt != "" {
			evn = append(evn, scoreid.String(x.Intern.Evnt))
		}
	}

	{
		lis, err := h.eve.SearchEvnt(evn)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	var lab [][]scoreid.String
	for _, x := range req.Object {
		if x.Public.Cate != "" || x.Public.Host != "" {
			lab = append(lab, append(inpLab(x.Public.Cate), inpLab(x.Public.Host)...))
		}
	}

	for _, x := range lab {
		lis, err := h.eve.SearchLabl(x)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
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
				Cate: outLab(x.Cate),
				Dura: outDur(x.Dura),
				Host: outLab(x.Host),
				Link: x.Link,
				Time: outTim(x.Time),
			},
		})
	}

	return res, nil
}
