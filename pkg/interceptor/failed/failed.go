package failed

import (
	"context"
	"fmt"

	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
	"google.golang.org/grpc"
)

type InterceptorConfig struct {
	Log logger.Interface
}

type Interceptor struct {
	log logger.Interface
}

func NewInterceptor(c InterceptorConfig) *Interceptor {
	if c.Log == nil {
		tracer.Panic(fmt.Errorf("%T.Log must not be empty", c))
	}

	return &Interceptor{
		log: c.Log,
	}
}

func (i *Interceptor) Interceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, inf *grpc.UnaryServerInfo, han grpc.UnaryHandler) (interface{}, error) {
		res, err := han(ctx, req)
		if err != nil {
			i.log.Log(
				ctx,
				"level", "error",
				"message", fmt.Sprintf("request %s failed", inf.FullMethod),
				"stack", tracer.JSON(err),
			)
		}

		return res, tracer.Mask(err)
	}
}
