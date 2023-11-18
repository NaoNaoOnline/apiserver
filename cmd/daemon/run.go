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
	"github.com/NaoNaoOnline/apiserver/pkg/emitter"
	"github.com/NaoNaoOnline/apiserver/pkg/envvar"
	"github.com/NaoNaoOnline/apiserver/pkg/permission"
	"github.com/NaoNaoOnline/apiserver/pkg/server"
	"github.com/NaoNaoOnline/apiserver/pkg/server/interceptor/failedinterceptor"
	"github.com/NaoNaoOnline/apiserver/pkg/server/middleware/authmiddleware"
	"github.com/NaoNaoOnline/apiserver/pkg/server/middleware/corsmiddleware"
	"github.com/NaoNaoOnline/apiserver/pkg/server/middleware/usermiddleware"
	"github.com/NaoNaoOnline/apiserver/pkg/server/serverhandler"
	"github.com/NaoNaoOnline/apiserver/pkg/storage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/labelstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/worker"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/client/twitterclient"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/workerhandler"
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

	var twi twitterclient.Interface
	{
		twi = twitterclient.New()
	}

	// --------------------------------------------------------------------- //

	var emi *emitter.Emitter
	{
		emi = emitter.New(emitter.Config{
			Cid: cid,
			Cnt: cnt,
			Log: log,
			Res: res,
			Rpc: rpc,
		})
	}

	// --------------------------------------------------------------------- //

	var sto *storage.Storage
	{
		sto = storage.New(storage.Config{
			Emi: emi,
			Log: log,
			Red: red,
		})
	}

	// --------------------------------------------------------------------- //

	{
		_, err := sto.Labl().Create(sto.Labl().SearchBltn())
		if labelstorage.IsLabelObjectAlreadyExists(err) {
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

	var prm permission.Interface
	{
		prm = permission.New(permission.Config{
			Cac: cac,
			Emi: emi.Plcy(),
			Log: log,
			Pol: sto.Plcy(),
			Wal: sto.Wllt(),
		})
	}

	{
		err = prm.EnsureActv()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// --------------------------------------------------------------------- //

	var shn *serverhandler.Handler
	{
		shn = serverhandler.New(serverhandler.Config{
			Emi: emi,
			Log: log,
			Prm: prm,
			Sto: sto,
		})
	}

	// --------------------------------------------------------------------- //

	var srv *server.Server
	{
		srv = server.New(server.Config{
			Han: shn.Hand(),
			Int: []twirp.Interceptor{
				failedinterceptor.NewInterceptor(failedinterceptor.InterceptorConfig{Log: log}).Interceptor,
			},
			Lis: lis,
			Log: log,
			Mid: []mux.MiddlewareFunc{
				corsmiddleware.NewMiddleware(corsmiddleware.MiddlewareConfig{Log: log}).Handler,
				authmiddleware.NewMiddleware(authmiddleware.MiddlewareConfig{Aud: env.OauthAud, Iss: env.OauthIss, Log: log}).Handler,
				usermiddleware.NewMiddleware(usermiddleware.MiddlewareConfig{Log: log, Use: sto.User()}).Handler,
			},
		})
	}

	{
		go srv.Daemon()
	}

	// --------------------------------------------------------------------- //

	var whn *workerhandler.Handler
	{
		whn = workerhandler.New(workerhandler.Config{
			Cid: cid,
			Cnt: cnt,
			Emi: emi,
			Log: log,
			Prm: prm,
			Rpc: rpc,
			Sto: sto,
			Twi: twi,
		})
	}

	// --------------------------------------------------------------------- //

	var wrk *worker.Worker
	{
		wrk = worker.New(worker.Config{
			Han: whn.Hand(),
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
