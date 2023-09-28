package objectid

import (
	"fmt"
	"slices"
	"testing"
)

func Test_ObjectID_Uni_string(t *testing.T) {
	testCases := []struct {
		lis []string
		uni []string
	}{
		// Case 000
		{
			lis: []string{},
			uni: nil,
		},
		// Case 001
		{
			lis: []string{
				"55",
				"44",
			},
			uni: []string{
				"55",
				"44",
			},
		},
		// Case 002
		{
			lis: []string{
				"33",
				"44",
				"33",
				"33",
			},
			uni: []string{
				"33",
				"44",
			},
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
			uni: []string{
				"33",
				"44",
				"88",
				"22",
				"55",
				"66",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			uni := Uni(tc.lis)
			if !slices.Equal(uni, tc.uni) {
				t.Fatalf("expected %#v got %#v", tc.uni, uni)
			}
		})
	}
}

func Test_ObjectID_Uni_ID(t *testing.T) {
	testCases := []struct {
		lis []ID
		uni []ID
	}{
		// Case 000
		{
			lis: []ID{},
			uni: nil,
		},
		// Case 001
		{
			lis: []ID{
				"55",
				"44",
			},
			uni: []ID{
				"55",
				"44",
			},
		},
		// Case 002
		{
			lis: []ID{
				"33",
				"44",
				"33",
				"33",
			},
			uni: []ID{
				"33",
				"44",
			},
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
			uni: []ID{
				"33",
				"44",
				"88",
				"22",
				"55",
				"66",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			uni := Uni(tc.lis)
			if !slices.Equal(uni, tc.uni) {
				t.Fatalf("expected %#v got %#v", tc.uni, uni)
			}
		})
	}
}
