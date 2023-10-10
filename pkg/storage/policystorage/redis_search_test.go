package policystorage

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Storage_PolicyStorage_Redis_searchAggr(t *testing.T) {
	testCases := []struct {
		obj []*Object
		agg []*Object
		del []*Object
	}{
		// Case 000
		{
			obj: []*Object{
				{Kind: "CreateSystem", Syst: 0, Memb: "0x0", Acce: 0, ChID: []int64{1}},
			},
			agg: []*Object{
				{Kind: "CreateSystem", Syst: 0, Memb: "0x0", Acce: 0, ChID: []int64{1}},
			},
			del: nil,
		},
		// Case 001
		{
			obj: []*Object{
				{Kind: "CreateSystem", Syst: 0, Memb: "0x0", Acce: 0, ChID: []int64{1, 2}},
				{Kind: "CreateMember", Syst: 0, Memb: "0x1", Acce: 1, ChID: []int64{1}},
			},
			agg: []*Object{
				{Kind: "CreateSystem", Syst: 0, Memb: "0x0", Acce: 0, ChID: []int64{1, 2}},
				{Kind: "CreateMember", Syst: 0, Memb: "0x1", Acce: 1, ChID: []int64{1}},
			},
			del: nil,
		},
		// Case 002
		{
			obj: []*Object{
				{Kind: "CreateSystem", Syst: 0, Memb: "0x0", Acce: 0, ChID: []int64{1}},
				{Kind: "DeleteSystem", Syst: 0, Memb: "0x0", Acce: 0, ChID: []int64{1}},
			},
			agg: nil,
			del: []*Object{
				{Kind: "DeleteSystem", Syst: 0, Memb: "0x0", Acce: 0, ChID: []int64{1}},
			},
		},
		// Case 003 is like 002 but with reversed record order.
		{
			obj: []*Object{
				{Kind: "DeleteSystem", Syst: 0, Memb: "0x0", Acce: 0, ChID: []int64{1}},
				{Kind: "CreateSystem", Syst: 0, Memb: "0x0", Acce: 0, ChID: []int64{1}},
			},
			agg: nil,
			del: []*Object{
				{Kind: "DeleteSystem", Syst: 0, Memb: "0x0", Acce: 0, ChID: []int64{1}},
			},
		},
		// Case 004 is unrealistic, but verifies whether unexpected system behaviour
		// causes aggregation problems. Below a delete record exists without a prior
		// create record. This should never be possible. The core of the test though
		// asserts that only delete records matching the SMA fields negate their
		// equivalent create records.
		{
			obj: []*Object{
				{Kind: "CreateSystem", Syst: 0, Memb: "0x0", Acce: 0, ChID: []int64{1}},
				{Kind: "DeleteSystem", Syst: 0, Memb: "0x0", Acce: 0, ChID: []int64{2}},
			},
			agg: []*Object{
				{Kind: "CreateSystem", Syst: 0, Memb: "0x0", Acce: 0, ChID: []int64{1}},
			},
			del: []*Object{
				{Kind: "DeleteSystem", Syst: 0, Memb: "0x0", Acce: 0, ChID: []int64{2}},
			},
		},
		// Case 005
		{
			obj: []*Object{
				{Kind: "CreateSystem", Syst: 0, Memb: "0x0", Acce: 0, ChID: []int64{1, 2}},
				{Kind: "DeleteSystem", Syst: 0, Memb: "0x0", Acce: 0, ChID: []int64{1}},
			},
			agg: []*Object{
				{Kind: "CreateSystem", Syst: 0, Memb: "0x0", Acce: 0, ChID: []int64{1, 2}},
			},
			del: []*Object{
				{Kind: "DeleteSystem", Syst: 0, Memb: "0x0", Acce: 0, ChID: []int64{1}},
			},
		},
		// Case 006 is like 005 but with reversed record order.
		{
			obj: []*Object{
				{Kind: "DeleteSystem", Syst: 0, Memb: "0x0", Acce: 0, ChID: []int64{1}},
				{Kind: "CreateSystem", Syst: 0, Memb: "0x0", Acce: 0, ChID: []int64{1, 2}},
			},
			agg: []*Object{
				{Kind: "CreateSystem", Syst: 0, Memb: "0x0", Acce: 0, ChID: []int64{1, 2}},
			},
			del: []*Object{
				{Kind: "DeleteSystem", Syst: 0, Memb: "0x0", Acce: 0, ChID: []int64{1}},
			},
		},
		// Case 007
		{
			obj: []*Object{
				{Kind: "CreateSystem", Syst: 0, Memb: "0x0", Acce: 0, ChID: []int64{1, 2}},
				{Kind: "CreateSystem", Syst: 0, Memb: "0x5", Acce: 0, ChID: []int64{3}},
				{Kind: "CreateMember", Syst: 0, Memb: "0x1", Acce: 1, ChID: []int64{2}},
				{Kind: "DeleteSystem", Syst: 0, Memb: "0x5", Acce: 0, ChID: []int64{3}},
			},
			agg: []*Object{
				{Kind: "CreateSystem", Syst: 0, Memb: "0x0", Acce: 0, ChID: []int64{1, 2}},
				{Kind: "CreateMember", Syst: 0, Memb: "0x1", Acce: 1, ChID: []int64{2}},
			},
			del: []*Object{
				{Kind: "DeleteSystem", Syst: 0, Memb: "0x5", Acce: 0, ChID: []int64{3}},
			},
		},
		// Case 008 is like 007 but with reversed record order.
		{
			obj: []*Object{
				{Kind: "DeleteSystem", Syst: 0, Memb: "0x5", Acce: 0, ChID: []int64{3}},
				{Kind: "CreateMember", Syst: 0, Memb: "0x1", Acce: 1, ChID: []int64{2}},
				{Kind: "CreateSystem", Syst: 0, Memb: "0x5", Acce: 0, ChID: []int64{3}},
				{Kind: "CreateSystem", Syst: 0, Memb: "0x0", Acce: 0, ChID: []int64{1, 2}},
			},
			agg: []*Object{
				{Kind: "CreateMember", Syst: 0, Memb: "0x1", Acce: 1, ChID: []int64{2}},
				{Kind: "CreateSystem", Syst: 0, Memb: "0x0", Acce: 0, ChID: []int64{1, 2}},
			},
			del: []*Object{
				{Kind: "DeleteSystem", Syst: 0, Memb: "0x5", Acce: 0, ChID: []int64{3}},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			agg, del := searchAggr(tc.obj)

			if !reflect.DeepEqual(agg, tc.agg) {
				t.Fatalf("expected %#v got %#v", tc.agg, agg)
			}
			if !reflect.DeepEqual(del, tc.del) {
				t.Fatalf("expected %#v got %#v", tc.del, del)
			}
		})
	}
}
