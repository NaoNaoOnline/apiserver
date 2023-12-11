package fakeit

import (
	"fmt"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/storage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/labelstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/xh3b4sd/tracer"
)

func (r *run) createEvnt(sto *storage.Storage, obj ...*eventstorage.Object) error {
	for _, x := range obj {
		if x == nil {
			continue
		}

		out, err := sto.Evnt().CreateEvnt([]*eventstorage.Object{x})
		if eventstorage.IsEventParticipationConflict(err) {
			return nil
		} else if err != nil {
			tracer.Panic(tracer.Mask(err))
		}

		_, err = sto.Evnt().CreateWrkr([]*eventstorage.Object{out[0]})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (r *run) randomEvnt(sto *storage.Storage, fak *gofakeit.Faker) *eventstorage.Object {
	var err error

	var cat labelstorage.Slicer
	{
		cat, err = sto.Labl().SearchKind([]string{"cate"})
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	{
		gofakeit.ShuffleAnySlice(cat)
	}

	var dur int
	{
		dur = fak.RandomInt([]int{
			15,
			30,
			45,
			60,
			90,
			120,
		})
	}

	var hos labelstorage.Slicer
	{
		hos, err = sto.Labl().SearchKind([]string{"host"})
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	{
		gofakeit.ShuffleAnySlice(hos)
	}

	var min int
	{
		min = fak.RandomInt([]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			15, 15, 15, 15, 15, 15, 15, 15, 15,
			30, 30, 30, 30, 30,
			45, 45,
		})
	}

	var hou int
	{
		hou = fak.RandomInt([]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			2, 2, 2, 2, 2, 2, 2, 2, 2,
			5, 5, 5, 5, 5,
			12, 12, 12,
			24, 24,
			72,
		})
	}

	var tim time.Time
	{
		tim = time.Now().UTC()
	}

	{
		tim = tim.Add(time.Hour).Truncate(time.Hour)
	}

	{
		tim = tim.Add(time.Duration(min * int(time.Minute)))
		tim = tim.Add(time.Duration(hou * int(time.Hour)))
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

	var obj *eventstorage.Object
	{
		obj = &eventstorage.Object{
			Cate: cat.Labl()[:fak.Number(1, 4)],
			Dura: time.Duration(dur * int(time.Minute)),
			Host: hos.Labl()[:fak.Number(1, 2)],
			Link: fmt.Sprintf("https://%s.%s", fak.DomainName(), fak.DomainSuffix()),
			Mtrc: objectfield.MapInt{
				Data: map[string]int64{
					objectlabel.EventMetricUser: int64(fak.Number(minRan[fak.Number(0, 3)], maxRan[fak.Number(0, 3)])),
				},
			},
			Time: tim,
			User: use.User()[0],
		}
	}

	return obj
}
