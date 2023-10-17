package policycache

import (
	"errors"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	addOne = "0x0000000000000000000000000000000000000000"
	addTwo = "0x1111111111111111111111111111111111111111"
)

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
		err := pol.Update()
		if !errors.Is(err, policyBufferEmptyError) {
			t.Fatal("expected", policyBufferEmptyError, "got", err)
		}
	}

	{
		rec := []*Record{
			tesRec(0, addOne, 0, []int64{1}),
		}

		err := pol.Buffer(rec)
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		rec := pol.SearchRcrd()
		if len(rec) != 0 {
			t.Fatal("expected", 0, "got", len(rec))
		}
	}

	{
		err := pol.Update()
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

	var lis []*Record
	{
		lis = pol.SearchRcrd()
	}

	{
		if len(lis) != 1 {
			t.Fatal("expected", 1, "got", len(lis))
		}
	}

	{
		var exp []*Record
		{
			exp = []*Record{
				tesRec(0, addOne, 0, []int64{1}),
			}
		}

		if !reflect.DeepEqual(lis, exp) {
			t.Fatalf("\n\n%s\n", cmp.Diff(exp, lis))
		}
	}
}

func tesRec(sys int64, mem string, acc int64, cid []int64) *Record {
	return &Record{Acce: acc, ChID: cid, Memb: mem, Syst: sys}
}
