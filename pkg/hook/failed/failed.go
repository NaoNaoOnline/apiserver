package failed

import (
	"context"
	"fmt"

	"github.com/twitchtv/twirp"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HookConfig struct {
	Log logger.Interface
}

type Hook struct {
	log logger.Interface
}

func NewHook(c HookConfig) *Hook {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}

	return &Hook{
		log: c.Log,
	}
}

func (h *Hook) Error() func(ctx context.Context, err twirp.Error) context.Context {
	return func(ctx context.Context, err twirp.Error) context.Context {
		h.log.Log(
			ctx,
			"level", "error",
			"code", string(err.Code()),
			"message", err.Msg(),
		)

		return ctx
	}
}
