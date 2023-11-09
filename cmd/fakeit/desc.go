package fakeit

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/xh3b4sd/tracer"
)

func (r *run) createDesc(sto *storage.Storage, obj ...*descriptionstorage.Object) error {
	for _, x := range obj {
		if x == nil {
			return nil
		}
	}

	{
		_, err := sto.Desc().Create(obj)
		if descriptionstorage.IsDescriptionEventLimit(err) {
			// fall through
		} else if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	return nil
}

func (r *run) randomDesc(sto *storage.Storage, fak *gofakeit.Faker) *descriptionstorage.Object {
	var err error

	var eve eventstorage.Slicer
	{
		eve, err = sto.Evnt().SearchUpcm()
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	{
		gofakeit.ShuffleAnySlice(eve)
	}

	var eid objectid.ID
	{
		eid = eve.Upcm().Evnt()[0]
	}

	var des []*descriptionstorage.Object
	{
		des, err = sto.Desc().SearchEvnt("", []objectid.ID{eid})
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	// 10% of events do not have any description.
	if len(des) == 0 && eid.Int()%10 == 0 {
		return nil
	}

	// 20% of events do not have more than 1 description.
	if len(des) == 1 && eid.Int()%5 == 0 {
		return nil
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
			Evnt: eid,
			Like: objectfield.Integer{
				Data: int64(fak.Number(minRan[fak.Number(0, 3)], maxRan[fak.Number(0, 3)])),
			},
			Text: txt,
			User: use[0].User,
		}
	}

	return obj
}
