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

	{
		rec := pol.SearchRcrd()
		if len(rec) != 6 {
			t.Fatal("expected", 6, "got", len(rec))
		}

		if rec[0].Syst != 0 {
			t.Fatal("expected", 0, "got", rec[0].Syst)
		}
		if rec[0].Memb != addOne {
			t.Fatal("expected", addOne, "got", rec[0].Memb)
		}
		if rec[0].Acce != 0 {
			t.Fatal("expected", 0, "got", rec[0].Acce)
		}
		if !reflect.DeepEqual(rec[0].ChID, []int64{1, 2}) {
			t.Fatal("expected", []int64{1, 2}, "got", rec[0].ChID)
		}

		if rec[1].Syst != 2 {
			t.Fatal("expected", 2, "got", rec[1].Syst)
		}
		if rec[1].Memb != addOne {
			t.Fatal("expected", addOne, "got", rec[1].Memb)
		}
		if rec[1].Acce != 0 {
			t.Fatal("expected", 0, "got", rec[1].Acce)
		}
		if !reflect.DeepEqual(rec[1].ChID, []int64{1, 2, 3}) {
			t.Fatal("expected", []int64{1, 2, 3}, "got", rec[1].ChID)
		}

		if rec[2].Syst != 2 {
			t.Fatal("expected", 2, "got", rec[2].Syst)
		}
		if rec[2].Memb != addTwo {
			t.Fatal("expected", addTwo, "got", rec[2].Memb)
		}
		if rec[2].Acce != 1 {
			t.Fatal("expected", 1, "got", rec[2].Acce)
		}
		if !reflect.DeepEqual(rec[2].ChID, []int64{1, 3}) {
			t.Fatal("expected", []int64{1, 3}, "got", rec[2].ChID)
		}

		if rec[3].Syst != 0 {
			t.Fatal("expected", 0, "got", rec[3].Syst)
		}
		if rec[3].Memb != addTwo {
			t.Fatal("expected", addTwo, "got", rec[3].Memb)
		}
		if rec[3].Acce != 1 {
			t.Fatal("expected", 1, "got", rec[3].Acce)
		}
		if !reflect.DeepEqual(rec[3].ChID, []int64{2}) {
			t.Fatal("expected", []int64{2}, "got", rec[3].ChID)
		}

		if rec[4].Syst != 1 {
			t.Fatal("expected", 1, "got", rec[4].Syst)
		}
		if rec[4].Memb != addOne {
			t.Fatal("expected", addOne, "got", rec[4].Memb)
		}
		if rec[4].Acce != 0 {
			t.Fatal("expected", 0, "got", rec[4].Acce)
		}
		if !reflect.DeepEqual(rec[4].ChID, []int64{3}) {
			t.Fatal("expected", []int64{3}, "got", rec[4].ChID)
		}

		if rec[5].Syst != 1 {
			t.Fatal("expected", 1, "got", rec[5].Syst)
		}
		if rec[5].Memb != addTwo {
			t.Fatal("expected", addTwo, "got", rec[5].Memb)
		}
		if rec[5].Acce != 1 {
			t.Fatal("expected", 1, "got", rec[5].Acce)
		}
		if !reflect.DeepEqual(rec[5].ChID, []int64{3}) {
			t.Fatal("expected", []int64{3}, "got", rec[5].ChID)
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
		var rec *Record
		{
			rec = lis[0]
		}

		var exp *Record
		{
			exp = &Record{
				Acce: 0,
				ChID: []int64{1},
				Memb: addOne,
				Syst: 0,
			}
		}

		if !reflect.DeepEqual(rec, exp) {
			t.Fatalf("\n\n%s\n", cmp.Diff(exp, rec))
		}
	}

	{
		var rec *Record
		{
			rec = lis[1]
		}

		var exp *Record
		{
			exp = &Record{
				Acce: 0,
				ChID: []int64{1},
				Memb: addOne,
				Syst: 2,
			}
		}

		if !reflect.DeepEqual(rec, exp) {
			t.Fatalf("\n\n%s\n", cmp.Diff(exp, rec))
		}
	}

	{
		var rec *Record
		{
			rec = lis[2]
		}

		var exp *Record
		{
			exp = &Record{
				Acce: 1,
				ChID: []int64{1},
				Memb: addTwo,
				Syst: 2,
			}
		}

		if !reflect.DeepEqual(rec, exp) {
			t.Fatalf("\n\n%s\n", cmp.Diff(exp, rec))
		}
	}
}
