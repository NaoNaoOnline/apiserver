package policyhandler

import (
	"fmt"
	"strconv"

	"github.com/NaoNaoOnline/apigocode/pkg/policy"
	"github.com/NaoNaoOnline/apiserver/pkg/cache/policycache"
	"github.com/NaoNaoOnline/apiserver/pkg/permission"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/rescue"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Cid []int64
	Log logger.Interface
	Prm permission.Interface
	Res rescue.Interface
}

type Handler struct {
	cid []int64
	log logger.Interface
	prm permission.Interface
	res rescue.Interface
}

func NewHandler(c HandlerConfig) *Handler {
	if len(c.Cid) == 0 {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Cid must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Prm == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Prm must not be empty", c)))
	}
	if c.Res == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Res must not be empty", c)))
	}

	return &Handler{
		cid: c.Cid,
		log: c.Log,
		prm: c.Prm,
		res: c.Res,
	}
}

func outExt(rec *policycache.Record) []*policy.SearchO_Object_Extern {
	var lis []*policy.SearchO_Object_Extern

	for _, x := range rec.ChID {
		lis = append(lis, &policy.SearchO_Object_Extern{
			Chid: strconv.FormatInt(x, 10),
		})
	}

	return lis
}
