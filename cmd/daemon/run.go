package daemon

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/handler"
	"github.com/NaoNaoOnline/apiserver/pkg/handler/label"
	"github.com/NaoNaoOnline/apiserver/pkg/interceptor/failed"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type run struct{}

func (r *run) runE(cmd *cobra.Command, args []string) error {
	var err error

	var erc chan error
	var sig chan os.Signal
	{
		erc = make(chan error, 1)
		sig = make(chan os.Signal, 2)

		defer close(erc)
		defer close(sig)
	}

	// --------------------------------------------------------------------- //

	var ctx context.Context
	{
		ctx = context.Background()
	}

	var log logger.Interface
	{
		c := logger.Config{}

		log, err = logger.New(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var cep []grpc.UnaryServerInterceptor
	{
		cep = []grpc.UnaryServerInterceptor{
			failed.NewInterceptor(failed.InterceptorConfig{Log: log}).Interceptor(),
		}
	}

	var han []handler.Interface
	{
		han = []handler.Interface{
			label.NewHandler(label.HandlerConfig{Log: log}),
		}
	}

	var ser *grpc.Server
	{
		ser = grpc.NewServer(grpc.ChainUnaryInterceptor(cep...))

		reflection.Register(ser)

		for _, x := range han {
			x.Attach(ser)
		}
	}

	// --------------------------------------------------------------------- //

	go func() {
		var lis net.Listener
		{
			lis, err = net.Listen("tcp", net.JoinHostPort("127.0.0.1", "7777"))
			if err != nil {
				erc <- tracer.Mask(err)
				return
			}
		}

		log.Log(ctx, "level", "info", "message", fmt.Sprintf("grpc server running at %s", lis.Addr().String()))

		{
			err = ser.Serve(lis)
			if err != nil {
				erc <- tracer.Mask(err)
				return
			}
		}
	}()

	// --------------------------------------------------------------------- //

	{
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

		select {
		case err := <-erc:
			return tracer.Mask(err)

		case <-sig:
			select {
			case <-time.After(10 * time.Second):
				// One SIGTERM gives the daemon some time to tear down gracefully.
			case <-sig:
				// Two SIGTERMs stop the immediatelly.
			}

			return nil
		}
	}
}
