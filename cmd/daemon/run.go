package daemon

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/cache/policycache"
	"github.com/NaoNaoOnline/apiserver/pkg/emitter/policyemitter"
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
	"github.com/NaoNaoOnline/apiserver/pkg/storage/policystorage"
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

	// --------------------------------------------------------------------- //

	var cid []int64
	{
		cid = append(cid, splNum(env.ChainCid)...)
	}

	var cnt []string
	{
		cnt = append(cnt, splStr(env.ChainPol)...)
	}

	var rpc []string
	{
		rpc = append(rpc, splStr(env.ChainRpc)...)
	}

	if len(cid) != len(cnt) {
		tracer.Panic(tracer.Mask(fmt.Errorf("amount of configured chain ids and contract addresses must be equal, got %d and %d", len(cid), len(cnt))))
	}

	if len(cid) != len(rpc) {
		tracer.Panic(tracer.Mask(fmt.Errorf("amount of configured chain ids and rpc endpoints must be equal, got %d and %d", len(cid), len(rpc))))
	}

	if len(cnt) != len(rpc) {
		tracer.Panic(tracer.Mask(fmt.Errorf("amount of configured contract addresses and rpc endpoints must be equal, got %d and %d", len(cnt), len(rpc))))
	}

	// --------------------------------------------------------------------- //

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
	var pol policystorage.Interface
	var rct reactionstorage.Interface
	var use userstorage.Interface
	var vot votestorage.Interface
	var wal walletstorage.Interface
	{
		lab = labelstorage.NewRedis(labelstorage.RedisConfig{Log: log, Red: red})
		pol = policystorage.NewRedis(policystorage.RedisConfig{Log: log, Red: red})
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

	var cac policycache.Interface
	{
		cac = policycache.NewMemory(policycache.MemoryConfig{
			Log: log,
		})
	}

	var emi policyemitter.Interface
	{
		emi = policyemitter.NewEmitter(policyemitter.EmitterConfig{
			Cid: cid,
			Cnt: cnt,
			Log: log,
			Res: res,
			Rpc: rpc,
		})
	}

	var prm permission.Interface
	{
		prm = permission.New(permission.Config{
			Cac: cac,
			Emi: emi,
			Log: log,
			Pol: pol,
			Wal: wal,
		})
	}

	{
		err = prm.EnsureActv()
		if err != nil {
			return tracer.Mask(err)
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
				policyhandler.NewHandler(policyhandler.HandlerConfig{Emi: emi, Log: log, Prm: prm}),
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

	var psh []workerhandler.Interface
	for i := range rpc {
		psh = append(psh, workerpolicyhandler.NewScrapeHandler(workerpolicyhandler.ScrapeHandlerConfig{
			Cid: cid[i],
			Cnt: cnt[i],
			Log: log,
			Prm: prm,
			Rpc: rpc[i],
		}))
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
					workerpolicyhandler.NewBufferHandler(workerpolicyhandler.BufferHandlerConfig{Log: log, Prm: prm}),
					workerpolicyhandler.NewUpdateHandler(workerpolicyhandler.UpdateHandlerConfig{Cid: cid, Emi: emi, Log: log, Prm: prm}),
				},
				psh...,
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

func splNum(str string) []int64 {
	var lis []int64

	for _, x := range strings.Split(str, ",") {
		lis = append(lis, musNum(x))
	}

	return lis
}

func splStr(str string) []string {
	return strings.Split(str, ",")
}

func musNum(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		tracer.Panic(tracer.Mask(err))
	}

	return num
}
