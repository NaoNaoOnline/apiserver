package policyhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/policy"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Update(ctx context.Context, req *policy.UpdateI) (*policy.UpdateO, error) {
	var err error

	if userid.FromContext(ctx) == "" {
		return nil, tracer.Mask(userIDEmptyError)
	}

	if len(req.Object) != 1 && req.Object[0].Symbol.Sync != "default" {
		return nil, tracer.Mask(updateSyncInvalidError)
	}

	var exi bool
	{
		exi, err = h.pol.ExistsMemb(userid.FromContext(ctx))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if !exi {
		return nil, tracer.Mask(policyMemberError)
	}

	var tas *task.Task
	{
		tas = &task.Task{
			Meta: &task.Meta{
				objectlabel.PlcyAction: objectlabel.ActionUpdate,
				objectlabel.PlcyObject: "*",
				objectlabel.PlcyOrigin: objectlabel.OriginCustom,
			},
		}
	}

	{
		err := h.res.Create(tas)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *policy.UpdateO
	{
		res = &policy.UpdateO{
			Object: []*policy.UpdateO_Object{
				{
					Intern: &policy.UpdateO_Object_Intern{
						Stts: objectstate.Started.String(),
					},
				},
			},
		}
	}

	return res, nil
}
