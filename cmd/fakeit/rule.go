package fakeit

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/labelstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/liststorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/rulestorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/xh3b4sd/tracer"
)

func (r *run) createRule(sto *storage.Storage, obj ...*rulestorage.Object) error {
	for _, x := range obj {
		if x == nil {
			continue
		}

		out, err := sto.Rule().CreateRule([]*rulestorage.Object{x})
		if rulestorage.IsResourceIDEmpty(err) {
			return nil
		} else if err != nil {
			tracer.Panic(tracer.Mask(err))
		}

		_, err = sto.Rule().CreateWrkr([]*rulestorage.Object{out[0]})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (r *run) randomRule(sto *storage.Storage, fak *gofakeit.Faker) *rulestorage.Object {
	var err error

	var lis liststorage.Slicer
	{
		lis, err = sto.List().SearchFake()
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	{
		gofakeit.ShuffleAnySlice(lis)
	}

	var lid objectid.ID
	{
		lid = lis[0].List
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

	var uid objectid.ID
	{
		uid = use[0].User
	}

	{
		gofakeit.ShuffleAnySlice(use)
	}

	// 20% of users do not have any list.
	if uid.Int()%5 == 0 {
		return nil
	}

	var lab labelstorage.Slicer
	{
		lab, err = sto.Labl().SearchKind([]string{"bltn", "cate", "host"})
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	var kin string
	{
		kin = fak.RandomString([]string{
			"cate",
			"host",
			"user",
		})
	}

	var min func() int
	{
		min = func() int {
			return fak.RandomInt([]int{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				1, 1, 1, 1, 1, 1, 1, 1, 1,
			})
		}
	}

	var max func() int
	{
		max = func() int {
			return fak.RandomInt([]int{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
				2, 2, 2, 2, 2, 2, 2, 2, 2,
			})
		}
	}

	var exc []objectid.ID
	var inc []objectid.ID
	{
		if kin == "cate" || kin == "host" {
			if fak.Number(0, 9) == 0 {
				exc = lab.Labl()
			} else {
				inc = lab.Labl()
			}
		}

		if kin == "user" {
			if fak.Number(0, 9) == 0 {
				exc = use.User()
			} else {
				inc = use.User()
			}
		}

		if len(exc) != 0 {
			exc = exc[min():max()]
		}

		if len(inc) != 0 {
			inc = inc[min():max()]
		}
	}

	var obj *rulestorage.Object
	{
		obj = &rulestorage.Object{
			Excl: exc,
			Incl: inc,
			Kind: kin,
			List: lid,
			User: uid,
		}
	}

	return obj
}
