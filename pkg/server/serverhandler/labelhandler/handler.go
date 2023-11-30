package labelhandler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/NaoNaoOnline/apigocode/pkg/label"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/labelstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Lab labelstorage.Interface
	Log logger.Interface
}

type Handler struct {
	lab labelstorage.Interface
	log logger.Interface
}

func NewHandler(c HandlerConfig) *Handler {
	if c.Lab == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Lab must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}

	return &Handler{
		lab: c.Lab,
		log: c.Log,
	}
}

func inpPat(upd []*label.UpdateI_Object_Update) []*labelstorage.Patch {
	var lis []*labelstorage.Patch

	for _, x := range upd {
		lis = append(lis, &labelstorage.Patch{
			Frm: x.Frm,
			Ope: x.Ope,
			Pat: x.Pat,
			Val: x.Val,
		})
	}

	return lis
}

func outMap(inp objectfield.Map) map[string]string {
	out := map[string]string{}

	for k, v := range inp.Data {
		out[k] = v
	}

	return out
}

func outTim(tim time.Time) string {
	if !tim.IsZero() {
		return strconv.FormatInt(tim.Unix(), 10)
	}

	return ""
}
