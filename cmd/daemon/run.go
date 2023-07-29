package daemon

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/envvar"
	"github.com/NaoNaoOnline/apiserver/pkg/handler"
	"github.com/NaoNaoOnline/apiserver/pkg/handler/label"
	"github.com/NaoNaoOnline/apiserver/pkg/hook/failed"
	"github.com/NaoNaoOnline/apiserver/pkg/middleware/auth"
	"github.com/NaoNaoOnline/apiserver/pkg/middleware/cors"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/twitchtv/twirp"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
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

	var env envvar.Env
	{
		env = envvar.Load()
	}

	var log logger.Interface
	{
		c := logger.Config{}

		log, err = logger.New(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var erh func(ctx context.Context, err twirp.Error) context.Context
	{
		erh = failed.NewHook(failed.HookConfig{Log: log}).Error()
	}

	var han []handler.Interface
	{
		han = []handler.Interface{
			label.NewHandler(label.HandlerConfig{Log: log}),
		}
	}

	var rtr *mux.Router
	{
		rtr = mux.NewRouter()
	}

	{
		rtr.Use(
			cors.NewMiddleware(cors.MiddlewareConfig{
				Log: log,
			}).Handler,
			auth.NewMiddleware(auth.MiddlewareConfig{
				Aud: env.OauthAud,
				Iss: env.OauthIss,
				Log: log,
			}).Handler,
		)
	}

	var hoo *twirp.ServerHooks
	{
		hoo = &twirp.ServerHooks{
			Error: func(ctx context.Context, err twirp.Error) context.Context {
				ctx = erh(ctx, err)
				return ctx
			},
		}
	}

	for _, x := range han {
		x.Attach(rtr, twirp.WithServerHooks(hoo), twirp.WithServerPathPrefix(""))
	}

	var ser *http.Server
	{
		ser = &http.Server{
			Handler: rtr,
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

		log.Log(context.Background(), "level", "info", "message", fmt.Sprintf("rpc server running at %s", lis.Addr().String()))

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
