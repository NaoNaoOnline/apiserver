package daemon

import (
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/cache/policycache"
	"github.com/NaoNaoOnline/apiserver/pkg/envvar"
	"github.com/NaoNaoOnline/apiserver/pkg/permission"
	"github.com/NaoNaoOnline/apiserver/pkg/server"
	serverhandler "github.com/NaoNaoOnline/apiserver/pkg/server/handler"
	"github.com/NaoNaoOnline/apiserver/pkg/server/handler/descriptionhandler"
	"github.com/NaoNaoOnline/apiserver/pkg/server/handler/eventhandler"
	"github.com/NaoNaoOnline/apiserver/pkg/server/handler/labelhandler"
	"github.com/NaoNaoOnline/apiserver/pkg/server/handler/policyhandler"
	"github.com/NaoNaoOnline/apiserver/pkg/server/handler/reactionhandler"
	"github.com/NaoNaoOnline/apiserver/pkg/server/handler/userhandler"
	"github.com/NaoNaoOnline/apiserver/pkg/server/handler/votehandler"
	"github.com/NaoNaoOnline/apiserver/pkg/server/handler/wallethandler"
	"github.com/NaoNaoOnline/apiserver/pkg/server/interceptor/failedinterceptor"
	"github.com/NaoNaoOnline/apiserver/pkg/server/middleware/authmiddleware"
	"github.com/NaoNaoOnline/apiserver/pkg/server/middleware/corsmiddleware"
	"github.com/NaoNaoOnline/apiserver/pkg/server/middleware/usermiddleware"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/labelstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/reactionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/votestorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/worker"
	workerhandler "github.com/NaoNaoOnline/apiserver/pkg/worker/handler"
	workerdescriptionhandler "github.com/NaoNaoOnline/apiserver/pkg/worker/handler/descriptionhandler"
	workereventhandler "github.com/NaoNaoOnline/apiserver/pkg/worker/handler/eventhandler"
	workerpolicyhandler "github.com/NaoNaoOnline/apiserver/pkg/worker/handler/policyhandler"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/twitchtv/twirp"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/rescue"
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

	var pol policycache.Interface
	{
		pol = policycache.NewMemory(policycache.MemoryConfig{
			Log: log,
		})
	}

	var red redigo.Interface
	{
		red = redigo.Default()
	}

	var res rescue.Interface
	{
		res = engine.New(engine.Config{
			Logger: log,
			Queue:  "api.naonao.io", // rescue.io/api.naonao.io
			Redigo: red,
			Sepkey: "/",
		})
	}

	// --------------------------------------------------------------------- //

	var lab labelstorage.Interface
	var rct reactionstorage.Interface
	var use userstorage.Interface
	var vot votestorage.Interface
	var wal walletstorage.Interface
	{
		lab = labelstorage.NewRedis(labelstorage.RedisConfig{Log: log, Red: red})
		rct = reactionstorage.NewRedis(reactionstorage.RedisConfig{Log: log, Red: red})
		use = userstorage.NewRedis(userstorage.RedisConfig{Log: log, Red: red})
		vot = votestorage.NewRedis(votestorage.RedisConfig{Log: log, Red: red})
		wal = walletstorage.NewRedis(walletstorage.RedisConfig{Log: log, Red: red})
	}

	var des descriptionstorage.Interface
	var eve eventstorage.Interface
	{
		des = descriptionstorage.NewRedis(descriptionstorage.RedisConfig{Log: log, Red: red, Res: res})
		eve = eventstorage.NewRedis(eventstorage.RedisConfig{Log: log, Red: red, Res: res})
	}

	// --------------------------------------------------------------------- //

	var prm permission.Interface
	{
		prm = permission.New(permission.Config{
			Log: log,
			Pol: pol,
			Wal: wal,
		})
	}

	// --------------------------------------------------------------------- //

	// TODO the bootstrapping of resources should be worker tasks like we already
	// handle policy records

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

	var pbh []workerhandler.Interface
	var cid []int64
	for _, x := range strings.Split(env.ChainRpc, ",") {
		var h workerhandler.Interface
		var c int64
		{
			h, c = workerpolicyhandler.NewBufferHandler(workerpolicyhandler.BufferHandlerConfig{
				Cnt: env.PolicyContract,
				Log: log,
				Pol: pol,
				Rpc: x,
			})
		}

		{
			pbh = append(pbh, h)
			cid = append(cid, c)
		}
	}

	// --------------------------------------------------------------------- //

	var srv *server.Server
	{
		srv = server.New(server.Config{
			Han: []serverhandler.Interface{
				descriptionhandler.NewHandler(descriptionhandler.HandlerConfig{Eve: eve, Des: des, Log: log}),
				eventhandler.NewHandler(eventhandler.HandlerConfig{Eve: eve, Log: log}),
				labelhandler.NewHandler(labelhandler.HandlerConfig{Lab: lab, Log: log}),
				policyhandler.NewHandler(policyhandler.HandlerConfig{Cid: cid, Log: log, Prm: prm, Res: res}),
				reactionhandler.NewHandler(reactionhandler.HandlerConfig{Log: log, Rct: rct}),
				userhandler.NewHandler(userhandler.HandlerConfig{Log: log, Use: use}),
				votehandler.NewHandler(votehandler.HandlerConfig{Des: des, Eve: eve, Log: log, Vot: vot}),
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
		go srv.Daemon()
	}

	// --------------------------------------------------------------------- //

	var wrk *worker.Worker
	{
		wrk = worker.New(worker.Config{
			Han: append(
				[]workerhandler.Interface{
					workerdescriptionhandler.NewCustomHandler(workerdescriptionhandler.CustomHandlerConfig{Des: des, Log: log, Vot: vot}),
					workereventhandler.NewCustomHandler(workereventhandler.CustomHandlerConfig{Eve: eve, Des: des, Log: log, Vot: vot}),
					workereventhandler.NewSystemHandler(workereventhandler.SystemHandlerConfig{Eve: eve, Log: log}),
					workerpolicyhandler.NewUpdateHandler(workerpolicyhandler.UpdateHandlerConfig{Cid: cid, Log: log, Pol: pol}),
				},
				pbh...,
			),
			Log: log,
			Res: res,
		})
	}

	{
		go wrk.Daemon()
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
