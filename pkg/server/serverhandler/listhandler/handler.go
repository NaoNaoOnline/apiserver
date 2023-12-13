package listhandler

import (
	"fmt"
	"strconv"
	"time"

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
		lis = append(lis, &liststorage.Patch{
			Frm: x.Frm,
			Ope: x.Ope,
			Pat: x.Pat,
			Val: x.Val,
		})
	}

	return lis
}

func outTim(tim time.Time) string {
	if !tim.IsZero() {
		return strconv.FormatInt(tim.Unix(), 10)
	}

	return ""
}

func preStr(pre bool, str string) string {
	if pre {
		return str
	}

	return ""
}
