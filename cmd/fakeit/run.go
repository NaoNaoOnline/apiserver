package fakeit

import (
	"github.com/NaoNaoOnline/apiserver/pkg/emitter"
	"github.com/NaoNaoOnline/apiserver/pkg/storage"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type run struct{}

func (r *run) run(cmd *cobra.Command, args []string) {
	var err error

	// --------------------------------------------------------------------- //

	var emi *emitter.Emitter
	{
		emi = emitter.Fake()
	}

	var log logger.Interface
	{
		log = logger.Default()
	}

	var red redigo.Interface
	{
		red = redigo.Default()
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

	for i := 0; i < 5; i++ {
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
}
