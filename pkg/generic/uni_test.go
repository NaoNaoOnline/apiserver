package generic

import (
	"fmt"
	"slices"
	"testing"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
)

func Test_Generic_Uni_string(t *testing.T) {
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

func Test_Generic_Uni_ID(t *testing.T) {
	testCases := []struct {
		lis []objectid.ID
		uni []objectid.ID
	}{
		// Case 000
		{
			lis: []objectid.ID{},
			uni: nil,
		},
		// Case 001
		{
			lis: []objectid.ID{
				"55",
				"44",
			},
			uni: []objectid.ID{
				"55",
				"44",
			},
		},
		// Case 002
		{
			lis: []objectid.ID{
				"33",
				"44",
				"33",
				"33",
			},
			uni: []objectid.ID{
				"33",
				"44",
			},
		},
		// Case 003
		{
			lis: []objectid.ID{
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
			uni: []objectid.ID{
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

func Test_Generic_Uni_int64(t *testing.T) {
	testCases := []struct {
		lis []int64
		uni []int64
	}{
		// Case 000
		{
			lis: []int64{},
			uni: nil,
		},
		// Case 001
		{
			lis: []int64{
				55,
				44,
			},
			uni: []int64{
				55,
				44,
			},
		},
		// Case 002
		{
			lis: []int64{
				33,
				44,
				33,
				33,
			},
			uni: []int64{
				33,
				44,
			},
		},
		// Case 003
		{
			lis: []int64{
				33,
				44,
				88,
				22,
				33,
				55,
				66,
				55,
				88,
			},
			uni: []int64{
				33,
				44,
				88,
				22,
				55,
				66,
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
