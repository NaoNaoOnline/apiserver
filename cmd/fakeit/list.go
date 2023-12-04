package fakeit

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/liststorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/xh3b4sd/tracer"
)

func (r *run) createList(sto *storage.Storage, obj ...*liststorage.Object) error {
	for _, x := range obj {
		if x == nil {
			return nil
		}
	}

	{
		_, err := sto.List().Create(obj)
		if liststorage.IsListDescLength(err) {
			// fall through
		} else if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	return nil
}

func (r *run) randomList(sto *storage.Storage, fak *gofakeit.Faker) *liststorage.Object {
	var err error

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

	var uid objectid.ID
	{
		uid = use[0].User
	}

	// 20% of users do not have any list.
	if uid.Int()%5 == 0 {
		return nil
	}

	var obj *liststorage.Object
	{
		obj = &liststorage.Object{
			Desc: objectfield.String{
				Data: fak.Phrase(),
			},
			User: uid,
		}
	}

	return obj
}
