package policyhandler

import (
	"fmt"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/policy"
	"github.com/NaoNaoOnline/apiserver/pkg/emitter/policyemitter"
	"github.com/NaoNaoOnline/apiserver/pkg/permission"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/policystorage"
	"github.com/xh3b4sd/locker"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Emi policyemitter.Interface
	Loc locker.Interface
	Log logger.Interface
	Prm permission.Interface
}

type Handler struct {
	emi policyemitter.Interface
	loc locker.Interface
	log logger.Interface
	prm permission.Interface
}

func NewHandler(c HandlerConfig) *Handler {
	if c.Emi == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Emi must not be empty", c)))
	}
	if c.Loc == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Loc must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Prm == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Prm must not be empty", c)))
	}

	return &Handler{
		emi: c.Emi,
		loc: c.Loc,
		log: c.Log,
		prm: c.Prm,
	}
}

func outExt(obj *policystorage.Object) []*policy.SearchO_Object_Extern {
	var lis []*policy.SearchO_Object_Extern

	for _, x := range obj.ChID {
		lis = append(lis, &policy.SearchO_Object_Extern{
			Chid: strconv.FormatInt(x, 10),
		})
	}

	return lis
}
