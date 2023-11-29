package objectid

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_ObjectID_Pair(t *testing.T) {
	testCases := []struct {
		key []string
		fir []ID
		sec []ID
	}{
		// Case 000
		{
			key: []string{},
			fir: nil,
			sec: nil,
		},
		// Case 001
		{
			key: []string{
				"foo",
			},
			fir: []ID{
				"foo",
			},
			sec: []ID{
				"",
			},
		},
		// Case 002
		{
			key: []string{
				"",
				"foo",
			},
			fir: []ID{
				"",
				"foo",
			},
			sec: []ID{
				"",
				"",
			},
		},
		// Case 003
		{
			key: []string{
				"foo",
				"bar",
				"baz",
			},
			fir: []ID{
				"foo",
				"bar",
				"baz",
			},
			sec: []ID{
				"",
				"",
				"",
			},
		},
		// Case 004
		{
			key: []string{
				"",
				"foo",
				"bar",
				"",
				"",
				"baz",
			},
			fir: []ID{
				"",
				"foo",
				"bar",
				"",
				"",
				"baz",
			},
			sec: []ID{
				"",
				"",
				"",
				"",
				"",
				"",
			},
		},
		// Case 005
		{
			key: []string{
				"foo,123",
			},
			fir: []ID{
				"foo",
			},
			sec: []ID{
				"123",
			},
		},
		// Case 006
		{
			key: []string{
				"",
				"foo,123",
			},
			fir: []ID{
				"",
				"foo",
			},
			sec: []ID{
				"",
				"123",
			},
		},
		// Case 007
		{
			key: []string{
				"foo,123",
				"bar,345",
				"baz,567",
			},
			fir: []ID{
				"foo",
				"bar",
				"baz",
			},
			sec: []ID{
				"123",
				"345",
				"567",
			},
		},
		// Case 008
		{
			key: []string{
				"",
				"foo,123",
				"bar,345",
				"baz,567",
				"",
				"",
			},
			fir: []ID{
				"",
				"foo",
				"bar",
				"baz",
				"",
				"",
			},
			sec: []ID{
				"",
				"123",
				"345",
				"567",
				"",
				"",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			fir := Frst(tc.key)
			if !reflect.DeepEqual(fir, tc.fir) {
				t.Fatalf("expected %#v got %#v", tc.fir, fir)
			}

			sec := Scnd(tc.key)
			if !reflect.DeepEqual(sec, tc.sec) {
				t.Fatalf("expected %#v got %#v", tc.sec, sec)
			}

			// It is critical to ensure that Frst and Scnd produce the same amount of
			// elements, since they must be iterable interchangeably.
			if len(fir) != len(tc.sec) {
				t.Fatalf("expected %#v got %#v", len(sec), len(fir))
			}
		})
	}
}
