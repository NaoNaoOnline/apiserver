package generic

import (
	"fmt"
	"testing"
)

func Test_Generic_Any_string(t *testing.T) {
	testCases := []struct {
		all []string
		sub []string
		any bool
	}{
		// Case 000
		{
			all: []string{},
			sub: []string{},
			any: false,
		},
		// Case 001
		{
			all: []string{
				"33",
			},
			sub: []string{},
			any: false,
		},
		// Case 002
		{
			all: []string{},
			sub: []string{
				"44",
			},
			any: false,
		},
		// Case 003
		{
			all: []string{
				"33",
			},
			sub: []string{
				"44",
			},
			any: false,
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
			any: false,
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
			any: false,
		},
		// Case 006
		{
			all: []string{
				"44",
				"22",
				"66",
			},
			sub: []string{
				"33",
				"55",
				"22",
			},
			any: true,
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
				"11",
				"66",
				"88",
			},
			any: true,
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
				"22",
				"66",
				"33",
				"55",
				"22",
			},
			any: true,
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
				"66",
			},
			any: true,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			any := Any(tc.all, tc.sub)
			if any != tc.any {
				t.Fatalf("expected %#v got %#v", tc.any, any)
			}
		})
	}
}

func Test_Generic_Any_int64(t *testing.T) {
	testCases := []struct {
		all []int64
		sub []int64
		any bool
	}{
		// Case 000
		{
			all: []int64{},
			sub: []int64{},
			any: false,
		},
		// Case 001
		{
			all: []int64{
				33,
			},
			sub: []int64{},
			any: false,
		},
		// Case 002
		{
			all: []int64{},
			sub: []int64{
				44,
			},
			any: false,
		},
		// Case 003
		{
			all: []int64{
				33,
			},
			sub: []int64{
				44,
			},
			any: false,
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
			any: false,
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
			any: false,
		},
		// Case 006
		{
			all: []int64{
				44,
				22,
				66,
			},
			sub: []int64{
				33,
				55,
				22,
			},
			any: true,
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
				11,
				66,
				88,
			},
			any: true,
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
				22,
				66,
				33,
				55,
				22,
			},
			any: true,
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
				66,
			},
			any: true,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			any := Any(tc.all, tc.sub)
			if any != tc.any {
				t.Fatalf("expected %#v got %#v", tc.any, any)
			}
		})
	}
}
