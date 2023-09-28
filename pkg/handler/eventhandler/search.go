package eventhandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/event"
	"github.com/NaoNaoOnline/apiserver/pkg/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *event.SearchI) (*event.SearchO, error) {
	//
	// Validate the RPC integrity.
	//

	for _, x := range req.Object {
		if x.Intern.Evnt != "" && (x.Intern.User != "" || x.Public.Cate != "" || x.Public.Host != "") {
			return nil, tracer.Mask(searchEvntConflictError)
		}
		if x.Intern.User != "" && (x.Intern.Evnt != "" || x.Public.Cate != "" || x.Public.Host != "") {
			return nil, tracer.Mask(searchUserConflictError)
		}
	}

	var out []*eventstorage.Object

	//
	// Search events by ID.
	//

	var evn []objectid.String
	for _, x := range req.Object {
		if x.Intern.Evnt != "" {
			evn = append(evn, objectid.String(x.Intern.Evnt))
		}
	}

	if len(evn) != 0 {
		lis, err := h.eve.SearchEvnt(evn)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Search events by label.
	//

	var lab [][]objectid.String
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

	//
	// Search events by reactions.
	//

	if userid.FromContext(ctx) != "" {
		var rct bool
		for _, x := range req.Object {
			if x.Symbol.Rctn == "default" {
				rct = true
			}
		}

		if rct {
			lis, err := h.eve.SearchRctn(userid.FromContext(ctx))
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	//
	// Search events by time.
	//

	var lts bool
	for _, x := range req.Object {
		if x.Symbol.Ltst == "default" {
			lts = true
		}
	}

	if lts {
		lis, err := h.eve.SearchLtst()
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Search events by user.
	//

	var use []objectid.String
	for _, x := range req.Object {
		if x.Intern.User != "" {
			use = append(use, objectid.String(x.Intern.User))
		}
	}

	if len(use) != 0 {
		lis, err := h.eve.SearchUser(use)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Construct RPC response.
	//

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
