package fakeit

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/reactionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/votestorage"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/xh3b4sd/tracer"
)

func (r *run) createVote(sto *storage.Storage, obj ...*votestorage.Object) error {
	{
		_, err := sto.Vote().Create(obj)
		if votestorage.IsVoteEventLimit(err) {
			// fall through
		} else if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	return nil
}

func (r *run) randomVote(sto *storage.Storage, fak *gofakeit.Faker) *votestorage.Object {
	var err error

	var des descriptionstorage.Slicer
	var eve eventstorage.Slicer
	for len(des) == 0 {
		{
			eve, err = sto.Evnt().SearchLtst()
			if err != nil {
				tracer.Panic(tracer.Mask(err))
			}
		}

		{
			gofakeit.ShuffleAnySlice(eve)
		}

		{
			des, err = sto.Desc().SearchEvnt([]objectid.ID{eve.Upc().IDs()[0]})
			if err != nil {
				tracer.Panic(tracer.Mask(err))
			}
		}

		{
			gofakeit.ShuffleAnySlice(des)
		}
	}

	var rea reactionstorage.Slicer
	{
		rea, err = sto.Rctn().SearchKind([]string{"bltn"})
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	{
		gofakeit.ShuffleAnySlice(rea)
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

	var obj *votestorage.Object
	{
		obj = &votestorage.Object{
			Desc: des.IDs()[0],
			Evnt: eve.Upc().IDs()[0],
			Rctn: rea.IDs()[0],
			User: use.IDs()[0],
		}
	}

	return obj
}
