package failedinterceptor

import (
	"context"
	"fmt"

	"github.com/twitchtv/twirp"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type InterceptorConfig struct {
	Log logger.Interface
}

type Interceptor struct {
	log logger.Interface
}

func NewInterceptor(c InterceptorConfig) *Interceptor {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}

	return &Interceptor{
		log: c.Log,
	}
}

func (i *Interceptor) Interceptor(nex twirp.Method) twirp.Method {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		res, err := nex(ctx, req)
		if err != nil {
			e, o := err.(*tracer.Error)
			if o {
				i.log.Log(
					ctx,
					"level", "error",
					"message", e.Error(),
					"code", e.Code,
					"description", e.Desc,
					"docs", e.Docs,
					"kind", e.Kind,
					"stack", tracer.Stack(e),
				)

				c := twirp.ErrorCode(e.Code)
				if c == "" {
					c = twirp.Internal
				}

				return nil, twirp.NewError(c, e.Error()).WithMeta("desc", e.Desc).WithMeta("kind", e.Kind)
			} else {
				i.log.Log(
					ctx,
					"level", "error",
					"message", err.Error(),
				)

				return nil, err
			}
		}

		return res, nil
	}
}
