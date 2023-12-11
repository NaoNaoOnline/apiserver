package fakeit

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
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

	var evo *eventstorage.Object
	{
		evo = eve.Upcm().Obct()[0]
	}

	var eid objectid.ID
	{
		eid = evo.Evnt
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

	var uid objectid.ID
	if len(des) == 0 {
		uid = evo.User
	} else {
		gofakeit.ShuffleAnySlice(use)
		uid = use[0].User
	}

	var txt string
	for len(txt) < 40 && len(txt) < 80 {
		txt += fak.Phrase() + " "
	}

	var obj *descriptionstorage.Object
	{
		obj = &descriptionstorage.Object{
			Evnt: eid,
			Mtrc: objectfield.MapInt{
				Data: map[string]int64{
					objectlabel.DescriptionMetricUser: int64(fak.Number(minRan[fak.Number(0, 3)], maxRan[fak.Number(0, 3)])),
				},
			},
			Text: objectfield.String{
				Data: txt,
			},
			User: uid,
		}
	}

	return obj
}
