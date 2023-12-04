package round

import (
	"fmt"
	"testing"
)

func Test_RoundP(t *testing.T) {
	testCases := []struct {
		f float64
		p uint
		r float64
	}{
		// Case 0
		{
			f: 3.14,
			p: 2,
			r: 3.14,
		},
		// Case 1
		{
			f: 3.44444,
			p: 3,
			r: 3.444,
		},
		// Case 2
		{
			f: 3.5,
			p: 0,
			r: 4,
		},
		// Case 3
		{
			f: 3.487,
			p: 8,
			r: 3.487,
		},
		// Case 4
		{
			f: 3.487,
			p: 2,
			r: 3.49,
		},
		// Case 5
		{
			f: 0.4318900000000008 + 3.3644300000000001 + 3.4028000000000002 + 1.98201000000301 + 0.156800000000000,
			p: 5,
			r: 9.33793,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			r := RoundP(tc.f, tc.p)
			if r != tc.r {
				t.Fatal("expected", tc.r, "got", r)
			}
		})
	}
}

func Test_RoundN(t *testing.T) {
	testCases := []struct {
		f float64
		n uint
		r float64
	}{
		// Case 0
		{
			f: 15000,
			n: 2,
			r: 15000,
		},
		// Case 1
		{
			f: 15000,
			n: 3,
			r: 15000,
		},
		// Case 2
		{
			f: 15250,
			n: 3,
			r: 15000,
		},
		// Case 3
		{
			f: 15750,
			n: 3,
			r: 16000,
		},
		// Case 4
		{
			f: 15750,
			n: 2,
			r: 15800,
		},
		// Case 5
		{
			f: 15750,
			n: 4,
			r: 20000,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			r := RoundN(tc.f, tc.n)
			if r != tc.r {
				t.Fatal("expected", tc.r, "got", r)
			}
		})
	}
}
