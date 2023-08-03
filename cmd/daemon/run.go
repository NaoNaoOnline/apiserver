package daemon

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/envvar"
	"github.com/NaoNaoOnline/apiserver/pkg/handler"
	handlerlabl "github.com/NaoNaoOnline/apiserver/pkg/handler/label"
	handleruser "github.com/NaoNaoOnline/apiserver/pkg/handler/user"
	"github.com/NaoNaoOnline/apiserver/pkg/hook/failed"
	"github.com/NaoNaoOnline/apiserver/pkg/middleware/auth"
	"github.com/NaoNaoOnline/apiserver/pkg/middleware/cors"
	"github.com/NaoNaoOnline/apiserver/pkg/middleware/user"
	"github.com/NaoNaoOnline/apiserver/pkg/server"
	storageuser "github.com/NaoNaoOnline/apiserver/pkg/storage/user"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/twitchtv/twirp"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type run struct{}

func (r *run) runE(cmd *cobra.Command, args []string) error {
	var err error

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

	var lis net.Listener
	{
		lis, err = net.Listen("tcp", net.JoinHostPort(env.HttpHost, env.HttpPort))
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var red redigo.Interface
	{
		red = redigo.Default()
	}

	var use storageuser.Interface
	{
		use = storageuser.NewRedis(storageuser.RedisConfig{
			Log: log,
			Red: red,
		})
	}

	// --------------------------------------------------------------------- //

	var srv *server.Server
	{
		srv = server.New(server.Config{
			Erh: []func(ctx context.Context, err twirp.Error) context.Context{
				failed.NewHook(failed.HookConfig{Log: log}).Error(),
			},
			Han: []handler.Interface{
				handlerlabl.NewHandler(handlerlabl.HandlerConfig{Log: log}),
				handleruser.NewHandler(handleruser.HandlerConfig{Log: log, Use: use}),
			},
			Lis: lis,
			Log: log,
			Mid: []mux.MiddlewareFunc{
				cors.NewMiddleware(cors.MiddlewareConfig{Log: log}).Handler,
				auth.NewMiddleware(auth.MiddlewareConfig{Aud: env.OauthAud, Iss: env.OauthIss, Log: log}).Handler,
				user.NewMiddleware(user.MiddlewareConfig{Log: log, Use: use}).Handler,
			},
		})
	}

	{
		go srv.Serve()
	}

	// --------------------------------------------------------------------- //

	var sig chan os.Signal
	{
		sig = make(chan os.Signal, 2)
	}

	{
		defer close(sig)
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	}

	{
		<-sig
	}

	select {
	case <-time.After(10 * time.Second):
		// One SIGTERM gives the daemon some time to tear down gracefully.
	case <-sig:
		// Two SIGTERMs stop the immediatelly.
	}

	return nil
}
