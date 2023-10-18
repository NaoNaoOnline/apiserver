package policyhandler

import (
	"context"
	"fmt"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/policy"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Update(ctx context.Context, req *policy.UpdateI) (*policy.UpdateO, error) {
	//
	// Emit buffer tasks.
	//

	for _, x := range h.cid {
		var tas *task.Task
		{
			tas = &task.Task{
				Meta: &task.Meta{
					objectlabel.PlcyAction: objectlabel.ActionBuffer,
					objectlabel.PlcyNetwrk: strconv.FormatInt(x, 10),
					objectlabel.PlcyOrigin: objectlabel.OriginCustom,
				},
				Gate: &task.Gate{
					fmt.Sprintf(objectlabel.PlcyUnique, x): task.Trigger,
				},
			}
		}

		{
			err := h.res.Create(tas)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
	}

	//
	// Construct RPC response.
	//

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
