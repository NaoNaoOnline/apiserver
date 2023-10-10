package policyhandler

import (
	"fmt"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/policy"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/policystorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/rescue/engine"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Log logger.Interface
	Pol policystorage.Interface
	Res engine.Interface
}

type Handler struct {
	log logger.Interface
	pol policystorage.Interface
	res engine.Interface
}

func NewHandler(c HandlerConfig) *Handler {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Pol == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Pol must not be empty", c)))
	}
	if c.Res == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Res must not be empty", c)))
	}

	return &Handler{
		log: c.Log,
		pol: c.Pol,
		res: c.Res,
	}
}

func outExt(obj *policystorage.Object) []*policy.SearchO_Object_Extern {
	var lis []*policy.SearchO_Object_Extern

	for i := range obj.Blck {
		lis = append(lis, &policy.SearchO_Object_Extern{
			Blck: strconv.FormatInt(obj.Blck[i], 10),
			Chid: strconv.FormatInt(obj.ChID[i], 10),
			From: obj.From[i],
			Hash: obj.Hash[i],
			Time: strconv.FormatInt(obj.Time[i].Unix(), 10),
		})
	}

	return lis
}
