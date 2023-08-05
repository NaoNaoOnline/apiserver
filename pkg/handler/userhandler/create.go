package userhandler

import (
	"context"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/user"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, req *user.CreateI) (*user.CreateO, error) {
	var err error

	var sub string
	{
		cla, _ := ctx.Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		if cla == nil || cla.RegisteredClaims.Subject == "" {
			return nil, tracer.Mask(subjectClaimEmptyError)
		}

		sub = cla.RegisteredClaims.Subject
	}

	var img string
	var nam string
	{
		img = req.Object[0].Public.Imag
		nam = req.Object[0].Public.Name
	}

	var obj *userstorage.Object
	{
		obj, err = h.use.Create(sub, img, nam)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var out *user.CreateO
	{
		out = &user.CreateO{
			Object: []*user.CreateO_Object{
				{
					Intern: &user.CreateO_Object_Intern{
						Crtd: strconv.Itoa(int(obj.Crtd.Unix())),
						User: obj.User,
					},
				},
			},
		}
	}

	return out, nil
}
