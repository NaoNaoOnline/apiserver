package wallethandler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/NaoNaoOnline/apigocode/pkg/wallet"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Log logger.Interface
	Wal walletstorage.Interface
}

type Handler struct {
	log logger.Interface
	wal walletstorage.Interface
}

func NewHandler(c HandlerConfig) *Handler {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Wal == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Wal must not be empty", c)))
	}

	return &Handler{
		log: c.Log,
		wal: c.Wal,
	}
}

func inpPat(upd []*wallet.UpdateI_Object_Update) []*walletstorage.Patch {
	var lis []*walletstorage.Patch

	for _, x := range upd {
		var p *walletstorage.Patch
		{
			p = &walletstorage.Patch{
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

func outTim(tim time.Time) string {
	if !tim.IsZero() {
		return strconv.FormatInt(tim.Unix(), 10)
	}

	return ""
}
