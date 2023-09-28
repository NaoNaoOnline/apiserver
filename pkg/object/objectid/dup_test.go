package objectid

import (
	"fmt"
	"testing"
)

func Test_ObjectID_Dup_string(t *testing.T) {
	testCases := []struct {
		lis []string
		dup bool
	}{
		// Case 000
		{
			lis: []string{},
			dup: false,
		},
		// Case 001
		{
			lis: []string{
				"55",
				"44",
			},
			dup: false,
		},
		// Case 002
		{
			lis: []string{
				"33",
				"44",
				"33",
				"33",
			},
			dup: true,
		},
		// Case 003
		{
			lis: []string{
				"33",
				"44",
				"88",
				"22",
				"33",
				"55",
				"66",
				"55",
				"88",
			},
			dup: true,
		},
		// Case 004
		{
			lis: []string{
				"33",
				"44",
				"88",
				"22",
				"55",
				"66",
			},
			dup: false,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			dup := Dup(tc.lis)
			if dup != tc.dup {
				t.Fatalf("expected %#v got %#v", tc.dup, dup)
			}
		})
	}
}

func Test_ObjectID_Dup_ID(t *testing.T) {
	testCases := []struct {
		lis []ID
		dup bool
	}{
		// Case 000
		{
			lis: []ID{},
			dup: false,
		},
		// Case 001
		{
			lis: []ID{
				"55",
				"44",
			},
			dup: false,
		},
		// Case 002
		{
			lis: []ID{
				"33",
				"44",
				"33",
				"33",
			},
			dup: true,
		},
		// Case 003
		{
			lis: []ID{
				"33",
				"44",
				"88",
				"22",
				"33",
				"55",
				"66",
				"55",
				"88",
			},
			dup: true,
		},
		// Case 004
		{
			lis: []ID{
				"33",
				"44",
				"88",
				"22",
				"55",
				"66",
			},
			dup: false,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			dup := Dup(tc.lis)
			if dup != tc.dup {
				t.Fatalf("expected %#v got %#v", tc.dup, dup)
			}
		})
	}
}
