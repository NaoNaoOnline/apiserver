package generic

import (
	"fmt"
	"slices"
	"testing"
)

func Test_Generic_Rem_string(t *testing.T) {
	testCases := []struct {
		all []string
		sub []string
		rem []string
	}{
		// Case 000
		{
			all: []string{},
			sub: []string{},
			rem: []string{},
		},
		// Case 001
		{
			all: []string{
				"33",
			},
			sub: []string{},
			rem: []string{
				"33",
			},
		},
		// Case 002
		{
			all: []string{},
			sub: []string{
				"44",
			},
			rem: []string{},
		},
		// Case 003
		{
			all: []string{
				"33",
			},
			sub: []string{
				"44",
			},
			rem: []string{
				"33",
			},
		},
		// Case 004
		{
			all: []string{
				"33",
				"55",
				"22",
			},
			sub: []string{
				"44",
			},
			rem: []string{
				"33",
				"55",
				"22",
			},
		},
		// Case 005
		{
			all: []string{
				"44",
			},
			sub: []string{
				"33",
				"55",
				"22",
			},
			rem: []string{
				"44",
			},
		},
		// Case 006
		{
			all: []string{
				"44",
				"22",
				"66",
			},
			sub: []string{
				"44",
				"22",
			},
			rem: []string{
				"66",
			},
		},
		// Case 007
		{
			all: []string{
				"44",
				"22",
				"66",
				"33",
				"55",
				"22",
			},
			sub: []string{
				"66",
			},
			rem: []string{
				"44",
				"22",
				"33",
				"55",
				"22",
			},
		},
		// Case 008
		{
			all: []string{
				"11",
				"66",
				"88",
			},
			sub: []string{
				"44",
				"11",
				"66",
			},
			rem: []string{
				"88",
			},
		},
		// Case 009
		{
			all: []string{
				"44",
				"22",
				"66",
			},
			sub: []string{
				"22",
				"66",
				"44",
			},
			rem: []string{},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			rem := Rem(tc.all, tc.sub)
			if !slices.Equal(rem, tc.rem) {
				t.Fatalf("expected %#v got %#v", tc.rem, rem)
			}
		})
	}
}

func Test_Generic_Rem_int64(t *testing.T) {
	testCases := []struct {
		all []int64
		sub []int64
		rem []int64
	}{
		// Case 000
		{
			all: []int64{},
			sub: []int64{},
			rem: []int64{},
		},
		// Case 001
		{
			all: []int64{
				33,
			},
			sub: []int64{},
			rem: []int64{
				33,
			},
		},
		// Case 002
		{
			all: []int64{},
			sub: []int64{
				44,
			},
			rem: []int64{},
		},
		// Case 003
		{
			all: []int64{
				33,
			},
			sub: []int64{
				44,
			},
			rem: []int64{
				33,
			},
		},
		// Case 004
		{
			all: []int64{
				33,
				55,
				22,
			},
			sub: []int64{
				44,
			},
			rem: []int64{
				33,
				55,
				22,
			},
		},
		// Case 005
		{
			all: []int64{
				44,
			},
			sub: []int64{
				33,
				55,
				22,
			},
			rem: []int64{
				44,
			},
		},
		// Case 006
		{
			all: []int64{
				44,
				22,
				66,
			},
			sub: []int64{
				44,
				22,
			},
			rem: []int64{
				66,
			},
		},
		// Case 007
		{
			all: []int64{
				44,
				22,
				66,
				33,
				55,
				22,
			},
			sub: []int64{
				66,
			},
			rem: []int64{
				44,
				22,
				33,
				55,
				22,
			},
		},
		// Case 008
		{
			all: []int64{
				11,
				66,
				88,
			},
			sub: []int64{
				44,
				11,
				66,
			},
			rem: []int64{
				88,
			},
		},
		// Case 009
		{
			all: []int64{
				44,
				22,
				66,
			},
			sub: []int64{
				22,
				66,
				44,
			},
			rem: []int64{},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			rem := Rem(tc.all, tc.sub)
			if !slices.Equal(rem, tc.rem) {
				t.Fatalf("expected %#v got %#v", tc.rem, rem)
			}
		})
	}
}
