package user

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/user"
	storageuser "github.com/NaoNaoOnline/apiserver/pkg/storage/user"
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *user.SearchI) (*user.SearchO, error) {
	var err error

	var use string
	{
		use = req.Object[0].Intern.User
	}

	var sub string
	if use == "" {
		cla, _ := ctx.Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		if cla == nil || cla.RegisteredClaims.Subject == "" {
			return nil, tracer.Mask(subjectClaimEmptyError)
		}

		sub = cla.RegisteredClaims.Subject
	}

	var obj *storageuser.Object
	{
		obj, err = h.use.Search(sub, use)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var out *user.SearchO
	{
		out = &user.SearchO{
			Object: []*user.SearchO_Object{
				{
					Intern: &user.SearchO_Object_Intern{
						Crtd: strconv.Itoa(int(obj.Crtd.Unix())),
						User: obj.User,
					},
					Public: &user.SearchO_Object_Public{
						Imag: obj.Imag,
						Name: obj.Name,
					},
				},
			},
		}
	}

	return out, nil
}
