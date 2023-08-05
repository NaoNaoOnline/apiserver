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
	"github.com/NaoNaoOnline/apiserver/pkg/handler/labelhandler"
	"github.com/NaoNaoOnline/apiserver/pkg/handler/userhandler"
	"github.com/NaoNaoOnline/apiserver/pkg/hook/failed"
	"github.com/NaoNaoOnline/apiserver/pkg/middleware/authmiddleware"
	"github.com/NaoNaoOnline/apiserver/pkg/middleware/corsmiddleware"
	"github.com/NaoNaoOnline/apiserver/pkg/middleware/usermiddleware"
	"github.com/NaoNaoOnline/apiserver/pkg/server"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
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

	var use userstorage.Interface
	{
		use = userstorage.NewRedis(userstorage.RedisConfig{
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
				labelhandler.NewHandler(labelhandler.HandlerConfig{Log: log}),
				userhandler.NewHandler(userhandler.HandlerConfig{Log: log, Use: use}),
			},
			Lis: lis,
			Log: log,
			Mid: []mux.MiddlewareFunc{
				corsmiddleware.NewMiddleware(corsmiddleware.MiddlewareConfig{Log: log}).Handler,
				authmiddleware.NewMiddleware(authmiddleware.MiddlewareConfig{Aud: env.OauthAud, Iss: env.OauthIss, Log: log}).Handler,
				usermiddleware.NewMiddleware(usermiddleware.MiddlewareConfig{Log: log, Use: use}).Handler,
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
