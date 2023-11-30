package descriptionhandler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/NaoNaoOnline/apigocode/pkg/description"
	"github.com/NaoNaoOnline/apiserver/pkg/permission"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Des descriptionstorage.Interface
	Eve eventstorage.Interface
	Log logger.Interface
	Prm permission.Interface
}

type Handler struct {
	des descriptionstorage.Interface
	eve eventstorage.Interface
	log logger.Interface
	prm permission.Interface
}

func NewHandler(c HandlerConfig) *Handler {
	if c.Des == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Des must not be empty", c)))
	}
	if c.Eve == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Eve must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Prm == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Prm must not be empty", c)))
	}

	return &Handler{
		des: c.Des,
		eve: c.Eve,
		log: c.Log,
		prm: c.Prm,
	}
}

func inpPat(upd []*description.UpdateI_Object_Update) []*descriptionstorage.Patch {
	var lis []*descriptionstorage.Patch

	for _, x := range upd {
		lis = append(lis, &descriptionstorage.Patch{
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
