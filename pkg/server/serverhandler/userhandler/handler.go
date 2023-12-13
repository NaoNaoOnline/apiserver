package userhandler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/NaoNaoOnline/apigocode/pkg/user"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Log logger.Interface
	Use userstorage.Interface
}

type Handler struct {
	log logger.Interface
	use userstorage.Interface
}

func NewHandler(c HandlerConfig) *Handler {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Use == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Use must not be empty", c)))
	}

	return &Handler{
		log: c.Log,
		use: c.Use,
	}
}

func inpPat(upd []*user.UpdateI_Object_Update) []*userstorage.Patch {
	var lis []*userstorage.Patch

	for _, x := range upd {
		lis = append(lis, &userstorage.Patch{
			Frm: x.Frm,
			Ope: x.Ope,
			Pat: x.Pat,
			Val: x.Val,
		})
	}

	return lis
}

func outMap(inp objectfield.MapStr) map[string]string {
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
