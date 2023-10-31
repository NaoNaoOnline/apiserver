package fakeit

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
	"github.com/NaoNaoOnline/apiserver/pkg/storage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/xh3b4sd/tracer"
)

func (r *run) createDesc(sto *storage.Storage, obj ...*descriptionstorage.Object) error {
	{
		_, err := sto.Desc().Create(obj)
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	return nil
}

func (r *run) randomDesc(sto *storage.Storage, fak *gofakeit.Faker) *descriptionstorage.Object {
	var err error

	var eve eventstorage.Slicer
	{
		eve, err = sto.Evnt().SearchLtst()
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	{
		gofakeit.ShuffleAnySlice(eve)
	}

	var use userstorage.Slicer
	{
		use, err = sto.User().SearchFake()
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	{
		gofakeit.ShuffleAnySlice(use)
	}

	var txt string
	for len(txt) < 40 && len(txt) < 80 {
		txt += fak.Phrase() + " "
	}

	var obj *descriptionstorage.Object
	{
		obj = &descriptionstorage.Object{
			Evnt: eve.Upcm().Evnt()[0],
			Like: objectfield.Integer{
				Data: int64(fak.Number(0, 10000)),
			},
			Text: txt,
			User: use[0].User,
		}
	}

	return obj
}
