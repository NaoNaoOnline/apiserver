package limiter

import (
	"fmt"
	"testing"
)

func Test_Server_Limiter_Len(t *testing.T) {
	testCases := []struct {
		lis int
		len int
	}{
		// Case 000
		{
			lis: 0,
			len: 0,
		},
		// Case 001
		{
			lis: 36,
			len: 36,
		},
		// Case 002
		{
			lis: 365,
			len: 365,
		},
		// Case 003
		{
			lis: 3652,
			len: 1000,
		},
		// Case 004
		{
			lis: 36527,
			len: 1000,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			len := Len(tc.lis)

			if len != tc.len {
				t.Fatalf("expected %#v got %#v", tc.len, len)
			}
		})
	}
}

func Test_Server_Limiter_Log(t *testing.T) {
	testCases := []struct {
		lis int
		log bool
	}{
		// Case 000
		{
			lis: 0,
			log: false,
		},
		// Case 001
		{
			lis: 36,
			log: false,
		},
		// Case 002
		{
			lis: 365,
			log: false,
		},
		// Case 003
		{
			lis: 3652,
			log: true,
		},
		// Case 004
		{
			lis: 36527,
			log: true,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			log := Log(tc.lis)

			if log != tc.log {
				t.Fatalf("expected %#v got %#v", tc.log, log)
			}
		})
	}
}
