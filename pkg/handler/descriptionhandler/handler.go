package descriptionhandler

import (
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/description"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Des descriptionstorage.Interface
	Log logger.Interface
}

type Handler struct {
	des descriptionstorage.Interface
	log logger.Interface
}

func NewHandler(c HandlerConfig) *Handler {
	if c.Des == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Des must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}

	return &Handler{
		des: c.Des,
		log: c.Log,
	}
}

func outRat(rat map[string]descriptionstorage.Rtng) map[string]*description.SearchO_Object_Public_Rtng {
	out := map[string]*description.SearchO_Object_Public_Rtng{}

	for k, v := range rat {
		out[k] = &description.SearchO_Object_Public_Rtng{
			Amnt: int32(v.Amnt),
		}
	}

	return out
}
