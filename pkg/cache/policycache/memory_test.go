package policycache

import (
	"errors"
	"reflect"
	"testing"

	"github.com/NaoNaoOnline/apiserver/pkg/storage/policystorage"
	"github.com/google/go-cmp/cmp"
)

const (
	addOne = "0x0000000000000000000000000000000000000000"
	addTwo = "0x1111111111111111111111111111111111111111"
)

// Test_Cache_Policy_Memory_ExistsAcce is to ensure that access levels within
// systems are a threshold based requirement, instead of a strict definition.
// Having access level 1 should provide the associated member with access to
// level 2 and below, but not to level 0.
func Test_Cache_Policy_Memory_ExistsAcce(t *testing.T) {
	var pol Interface
	{
		pol = Fake()
	}

	var buf []*policystorage.Object
	{
		buf = []*policystorage.Object{
			tesRec(0, addOne, 1, []int64{1}),
		}
	}

	{
		err := pol.UpdateRcrd(buf)
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		one := pol.ExistsAcce(0, addOne, 0)
		if one {
			t.Fatal("expected", false, "got", true)
		}
		two := pol.ExistsAcce(0, addOne, 1)
		if !two {
			t.Fatal("expected", true, "got", false)
		}
		thr := pol.ExistsAcce(0, addOne, 2)
		if !thr {
			t.Fatal("expected", true, "got", false)
		}
	}

	{
		one := pol.ExistsAcce(1, addOne, 0)
		if one {
			t.Fatal("expected", false, "got", true)
		}
		two := pol.ExistsAcce(1, addOne, 1)
		if two {
			t.Fatal("expected", false, "got", true)
		}
		thr := pol.ExistsAcce(1, addOne, 2)
		if thr {
			t.Fatal("expected", false, "got", true)
		}
	}

	{
		one := pol.ExistsAcce(0, addTwo, 0)
		if one {
			t.Fatal("expected", false, "got", true)
		}
		two := pol.ExistsAcce(0, addTwo, 1)
		if two {
			t.Fatal("expected", false, "got", true)
		}
		thr := pol.ExistsAcce(0, addTwo, 2)
		if thr {
			t.Fatal("expected", false, "got", true)
		}
	}
}

func Test_Cache_Policy_Memory_Lifecycle(t *testing.T) {
	var pol Interface
	{
		pol = Fake()
	}

	{
		one := pol.ExistsAcce(0, addOne, 0)
		if one {
			t.Fatal("expected", false, "got", true)
		}
		two := pol.ExistsAcce(0, addTwo, 0)
		if two {
			t.Fatal("expected", false, "got", true)
		}
	}

	{
		one := pol.ExistsMemb(addOne)
		if one {
			t.Fatal("expected", false, "got", true)
		}
		two := pol.ExistsMemb(addTwo)
		if two {
			t.Fatal("expected", false, "got", true)
		}
	}

	{
		one := pol.ExistsSyst(0, addOne)
		if one {
			t.Fatal("expected", false, "got", true)
		}
		two := pol.ExistsSyst(0, addTwo)
		if two {
			t.Fatal("expected", false, "got", true)
		}
	}

	{
		rec := pol.SearchRcrd()
		if len(rec) != 0 {
			t.Fatal("expected", 0, "got", len(rec))
		}
	}

	{
		err := pol.UpdateRcrd(nil)
		if !errors.Is(err, policyBufferEmptyError) {
			t.Fatal("expected", policyBufferEmptyError, "got", err)
		}
	}

	{
		rec := pol.SearchRcrd()
		if len(rec) != 0 {
			t.Fatal("expected", 0, "got", len(rec))
		}
	}

	var buf []*policystorage.Object
	{
		buf = []*policystorage.Object{
			tesRec(0, addOne, 0, []int64{1}),
		}
	}

	{
		err := pol.UpdateRcrd(buf)
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		one := pol.ExistsAcce(0, addOne, 0)
		if !one {
			t.Fatal("expected", true, "got", false)
		}
		two := pol.ExistsAcce(0, addTwo, 0)
		if two {
			t.Fatal("expected", false, "got", true)
		}
	}

	{
		one := pol.ExistsMemb(addOne)
		if !one {
			t.Fatal("expected", true, "got", false)
		}
		two := pol.ExistsMemb(addTwo)
		if two {
			t.Fatal("expected", false, "got", true)
		}
	}

	{
		one := pol.ExistsSyst(0, addOne)
		if !one {
			t.Fatal("expected", true, "got", false)
		}
		two := pol.ExistsSyst(0, addTwo)
		if two {
			t.Fatal("expected", false, "got", true)
		}
	}

	var lis []*policystorage.Object
	{
		lis = pol.SearchRcrd()
	}

	{
		if len(lis) != 1 {
			t.Fatal("expected", 1, "got", len(lis))
		}
	}

	{
		var exp []*policystorage.Object
		{
			exp = []*policystorage.Object{
				tesRec(0, addOne, 0, []int64{1}),
			}
		}

		if !reflect.DeepEqual(lis, exp) {
			t.Fatalf("\n\n%s\n", cmp.Diff(exp, lis))
		}
	}
}

func tesRec(sys int64, mem string, acc int64, cid []int64) *policystorage.Object {
	return &policystorage.Object{Acce: acc, ChID: cid, Memb: mem, Syst: sys}
}
