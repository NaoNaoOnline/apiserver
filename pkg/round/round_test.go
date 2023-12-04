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
		// Case 000
		{
			f: 3.1,
			p: 0,
			r: 3.0,
		},
		// Case 001
		{
			f: 3.4,
			p: 0,
			r: 3.0,
		},
		// Case 002
		{
			f: 3.5,
			p: 0,
			r: 4.0,
		},
		// Case 003
		{
			f: 3.8,
			p: 0,
			r: 4.0,
		},

		// Case 004
		{
			f: 3.1,
			p: 1,
			r: 3.1,
		},
		// Case 005
		{
			f: 3.4,
			p: 1,
			r: 3.4,
		},
		// Case 006
		{
			f: 3.5,
			p: 1,
			r: 3.5,
		},
		// Case 007
		{
			f: 3.8,
			p: 1,
			r: 3.8,
		},

		// Case 008
		{
			f: 3.1,
			p: 2,
			r: 3.10,
		},
		// Case 009
		{
			f: 3.4,
			p: 2,
			r: 3.40,
		},
		// Case 010
		{
			f: 3.5,
			p: 2,
			r: 3.50,
		},
		// Case 011
		{
			f: 3.8,
			p: 2,
			r: 3.80,
		},

		// Case 012
		{
			f: 3.14,
			p: 2,
			r: 3.14,
		},
		// Case 013
		{
			f: 3.44444,
			p: 3,
			r: 3.444,
		},
		// Case 014
		{
			f: 3.5,
			p: 0,
			r: 4,
		},
		// Case 015
		{
			f: 3.487,
			p: 8,
			r: 3.487,
		},
		// Case 016
		{
			f: 3.487,
			p: 2,
			r: 3.49,
		},
		// Case 017
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
		// Case 000
		{
			f: 15000,
			n: 2,
			r: 15000,
		},
		// Case 001
		{
			f: 15000,
			n: 3,
			r: 15000,
		},
		// Case 002
		{
			f: 15250,
			n: 3,
			r: 15000,
		},
		// Case 003
		{
			f: 15750,
			n: 3,
			r: 16000,
		},
		// Case 004
		{
			f: 15750,
			n: 2,
			r: 15800,
		},
		// Case 005
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
