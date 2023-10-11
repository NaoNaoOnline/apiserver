package labelhandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/label"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/labelstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, req *label.CreateI) (*label.CreateO, error) {
	var err error

	//
	// Validate the RPC integrity.
	//

	if userid.FromContext(ctx) == "" {
		return nil, tracer.Mask(userIDEmptyError)
	}

	// It is super important to validate the label kind for labels that can be
	// created via RPC, since we have natively supported and individually created
	// label kinds. Label kind bltn must not be created via RPC, since the system
	// itself is the only authority to manage those labels internally.
	for _, x := range req.Object {
		if x.Public.Kind != "cate" && x.Public.Kind != "host" {
			return nil, tracer.Mask(createKindInvalidError)
		}
	}

	var inp []*labelstorage.Object
	for _, x := range req.Object {
		inp = append(inp, &labelstorage.Object{
			Desc: x.Public.Desc,
			Disc: x.Public.Disc,
			Kind: x.Public.Kind,
			Name: x.Public.Name,
			Twit: x.Public.Twit,
			User: userid.FromContext(ctx),
		})
	}

	var out []*labelstorage.Object
	{
		out, err = h.lab.Create(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *label.CreateO
	{
		res = &label.CreateO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &label.CreateO_Object{
			Intern: &label.CreateO_Object_Intern{
				Crtd: strconv.FormatInt(x.Crtd.Unix(), 10),
				Labl: x.Labl.String(),
			},
		})
	}

	return res, nil
}
