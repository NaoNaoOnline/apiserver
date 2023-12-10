package fakeit

import (
	"strconv"
	"strings"

	"github.com/NaoNaoOnline/apiserver/pkg/emitter"
	"github.com/NaoNaoOnline/apiserver/pkg/envvar"
	"github.com/NaoNaoOnline/apiserver/pkg/storage"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/rescue"
	"github.com/xh3b4sd/rescue/engine"
	"github.com/xh3b4sd/tracer"
)

var (
	minRan = []int{0, 5, 10, 15, 20}
	maxRan = []int{50, 5000, 500000, 50000000}
)

type run struct{}

func (r *run) run(cmd *cobra.Command, args []string) {
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

	var pcn []string
	{
		pcn = append(pcn, splStr(env.ChainPol)...)
	}

	var rpc []string
	{
		rpc = append(rpc, splStr(env.ChainRpc)...)
	}

	var scn []string
	{
		scn = append(scn, splStr(env.ChainSub)...)
	}

	// --------------------------------------------------------------------- //

	var log logger.Interface
	{
		log = logger.Default()
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

	var emi *emitter.Emitter
	{
		emi = emitter.New(emitter.Config{
			Cid: cid,
			Log: log,
			Pcn: pcn,
			Res: res,
			Rpc: rpc,
			Scn: scn,
		})
	}

	// --------------------------------------------------------------------- //

	var fak *gofakeit.Faker
	{
		fak = gofakeit.NewCrypto()
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

	for i := 0; i < 10; i++ {
		err = r.createUser(sto, r.randomUser(sto, fak))
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	for i := 0; i < 15; i++ {
		err = r.createLabl(sto, r.randomLabl(sto, fak))
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	for i := 0; i < 10; i++ {
		err = r.createEvnt(sto, r.randomEvnt(sto, fak))
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	for i := 0; i < 50; i++ {
		err = r.createDesc(sto, r.randomDesc(sto, fak))
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	for i := 0; i < 50; i++ {
		err = r.createList(sto, r.randomList(sto, fak))
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	for i := 0; i < 100; i++ {
		err = r.createRule(sto, r.randomRule(sto, fak))
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}
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
