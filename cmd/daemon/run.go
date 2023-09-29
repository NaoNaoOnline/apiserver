package daemon

import (
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/envvar"
	"github.com/NaoNaoOnline/apiserver/pkg/handler"
	"github.com/NaoNaoOnline/apiserver/pkg/handler/descriptionhandler"
	"github.com/NaoNaoOnline/apiserver/pkg/handler/eventhandler"
	"github.com/NaoNaoOnline/apiserver/pkg/handler/labelhandler"
	"github.com/NaoNaoOnline/apiserver/pkg/handler/reactionhandler"
	"github.com/NaoNaoOnline/apiserver/pkg/handler/userhandler"
	"github.com/NaoNaoOnline/apiserver/pkg/handler/votehandler"
	"github.com/NaoNaoOnline/apiserver/pkg/handler/wallethandler"
	"github.com/NaoNaoOnline/apiserver/pkg/interceptor/failedinterceptor"
	"github.com/NaoNaoOnline/apiserver/pkg/middleware/authmiddleware"
	"github.com/NaoNaoOnline/apiserver/pkg/middleware/corsmiddleware"
	"github.com/NaoNaoOnline/apiserver/pkg/middleware/usermiddleware"
	"github.com/NaoNaoOnline/apiserver/pkg/server"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/labelstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/reactionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/votestorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/twitchtv/twirp"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/rescue/engine"
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
		log = logger.Default()
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

	var res engine.Interface
	{
		res = engine.New(engine.Config{
			Logger: log,
			Queue:  "api.naonao.io", // rescue.io:api.naonao.io
			Redigo: red,
		})
	}

	// --------------------------------------------------------------------- //

	var des descriptionstorage.Interface
	var eve eventstorage.Interface
	var lab labelstorage.Interface
	var rct reactionstorage.Interface
	var use userstorage.Interface
	var vot votestorage.Interface
	var wal walletstorage.Interface
	{
		des = descriptionstorage.NewRedis(descriptionstorage.RedisConfig{Log: log, Red: red, Res: res})
		eve = eventstorage.NewRedis(eventstorage.RedisConfig{Log: log, Red: red})
		lab = labelstorage.NewRedis(labelstorage.RedisConfig{Log: log, Red: red})
		rct = reactionstorage.NewRedis(reactionstorage.RedisConfig{Log: log, Red: red})
		use = userstorage.NewRedis(userstorage.RedisConfig{Log: log, Red: red})
		vot = votestorage.NewRedis(votestorage.RedisConfig{Log: log, Red: red})
		wal = walletstorage.NewRedis(walletstorage.RedisConfig{Log: log, Red: red})
	}

	// --------------------------------------------------------------------- //

	{
		_, err := lab.Create(lab.SearchBltn())
		if labelstorage.IsLabelObjectAlreadyExists(err) {
			// fall through
		} else if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		_, err := rct.Create(rct.SearchBltn())
		if reactionstorage.IsReactionObjectAlreadyExists(err) {
			// fall through
		} else if err != nil {
			return tracer.Mask(err)
		}
	}

	// --------------------------------------------------------------------- //

	var srv *server.Server
	{
		srv = server.New(server.Config{
			Han: []handler.Interface{
				descriptionhandler.NewHandler(descriptionhandler.HandlerConfig{Des: des, Log: log}),
				eventhandler.NewHandler(eventhandler.HandlerConfig{Eve: eve, Log: log}),
				labelhandler.NewHandler(labelhandler.HandlerConfig{Lab: lab, Log: log}),
				reactionhandler.NewHandler(reactionhandler.HandlerConfig{Log: log, Rct: rct}),
				userhandler.NewHandler(userhandler.HandlerConfig{Log: log, Use: use}),
				votehandler.NewHandler(votehandler.HandlerConfig{Eve: eve, Log: log, Vot: vot}),
				wallethandler.NewHandler(wallethandler.HandlerConfig{Log: log, Wal: wal}),
			},
			Int: []twirp.Interceptor{
				failedinterceptor.NewInterceptor(failedinterceptor.InterceptorConfig{Log: log}).Interceptor,
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
