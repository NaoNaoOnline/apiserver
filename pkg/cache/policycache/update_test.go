package policycache

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Cache_Policy_Memory_Update_Multi(t *testing.T) {
	var pol Interface
	{
		pol = Fake()
	}

	// Buffer for chain ID 1.
	{
		rec := []*Record{
			tesRec(0, addOne, 0, []int64{1}),
			tesRec(2, addOne, 0, []int64{1}),
			tesRec(2, addTwo, 1, []int64{1}),
		}

		err := pol.Buffer(rec)
		if err != nil {
			t.Fatal(err)
		}
	}

	// Buffer for chain ID 2.
	{
		rec := []*Record{
			tesRec(0, addOne, 0, []int64{2}),
			tesRec(0, addTwo, 1, []int64{2}),
			tesRec(2, addOne, 0, []int64{2}),
		}

		err := pol.Buffer(rec)
		if err != nil {
			t.Fatal(err)
		}
	}

	// Buffer for chain ID 3.
	{
		rec := []*Record{
			tesRec(1, addOne, 0, []int64{3}),
			tesRec(1, addTwo, 1, []int64{3}),
			tesRec(2, addOne, 0, []int64{3}),
			tesRec(2, addTwo, 1, []int64{3}),
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

	var lis []*Record
	{
		lis = pol.SearchRcrd()
	}

	{
		if len(lis) != 6 {
			t.Fatal("expected", 6, "got", len(lis))
		}
	}

	{
		var exp []*Record
		{
			exp = []*Record{
				tesRec(0, addOne, 0, []int64{1, 2}),
				tesRec(2, addOne, 0, []int64{1, 2, 3}),
				tesRec(2, addTwo, 1, []int64{1, 3}),
				tesRec(0, addTwo, 1, []int64{2}),
				tesRec(1, addOne, 0, []int64{3}),
				tesRec(1, addTwo, 1, []int64{3}),
			}
		}

		if !reflect.DeepEqual(lis, exp) {
			t.Fatalf("\n\n%s\n", cmp.Diff(exp, lis))
		}
	}
}

func Test_Cache_Policy_Memory_Update_Single(t *testing.T) {
	var pol Interface
	{
		pol = Fake()
	}

	// Buffer for chain ID 1.
	{
		rec := []*Record{
			tesRec(0, addOne, 0, []int64{1}),
			tesRec(2, addOne, 0, []int64{1}),
			tesRec(2, addTwo, 1, []int64{1}),
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

	var lis []*Record
	{
		lis = pol.SearchRcrd()
	}

	{
		if len(lis) != 3 {
			t.Fatal("expected", 3, "got", len(lis))
		}
	}

	{
		var exp []*Record
		{
			exp = []*Record{
				tesRec(0, addOne, 0, []int64{1}),
				tesRec(2, addOne, 0, []int64{1}),
				tesRec(2, addTwo, 1, []int64{1}),
			}
		}

		if !reflect.DeepEqual(lis, exp) {
			t.Fatalf("\n\n%s\n", cmp.Diff(exp, lis))
		}
	}

	// Buffer for chain ID 2.
	{
		rec := []*Record{
			tesRec(0, addOne, 0, []int64{2}),
			tesRec(0, addTwo, 1, []int64{2}),
			tesRec(2, addOne, 0, []int64{2}),
		}

		err := pol.Buffer(rec)
		if err != nil {
			t.Fatal(err)
		}
	}

	// Verify that nothing changed. Since Memory.Buffer should not affect the
	// active permissions until Memory.Update is being executed, we expect the
	// same records we verified already above.
	{
		lis = pol.SearchRcrd()
	}

	{
		if len(lis) != 3 {
			t.Fatal("expected", 3, "got", len(lis))
		}
	}

	{
		var exp []*Record
		{
			exp = []*Record{
				tesRec(0, addOne, 0, []int64{1}),
				tesRec(2, addOne, 0, []int64{1}),
				tesRec(2, addTwo, 1, []int64{1}),
			}
		}

		if !reflect.DeepEqual(lis, exp) {
			t.Fatalf("\n\n%s\n", cmp.Diff(exp, lis))
		}
	}
}
