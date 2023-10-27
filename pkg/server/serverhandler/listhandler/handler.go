package listhandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/list"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/liststorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Lis liststorage.Interface
	Log logger.Interface
}

type Handler struct {
	lis liststorage.Interface
	log logger.Interface
}

func NewHandler(c HandlerConfig) *Handler {
	if c.Lis == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Lis must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}

	return &Handler{
		lis: c.Lis,
		log: c.Log,
	}
}

func inpPat(upd []*list.UpdateI_Object_Update) []*liststorage.Patch {
	var lis []*liststorage.Patch

	for _, x := range upd {
		var p *liststorage.Patch
		{
			p = &liststorage.Patch{
				Ope: x.Ope,
				Pat: x.Pat,
			}
		}

		if x.Val != nil {
			p.Val = *x.Val
		}

		lis = append(lis, p)
	}

	return lis
}
