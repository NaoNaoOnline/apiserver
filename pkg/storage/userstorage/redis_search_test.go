package userstorage

import (
	"fmt"
	"testing"
	"time"
)

func Test_Storage_User_Search_ovrPrm(t *testing.T) {
	testCases := []struct {
		pso string
		now string
		ovr bool
	}{
		// Case 000
		{
			pso: "2023-11-01T00:00:00Z",
			now: "2023-09-01T00:00:00Z",
			ovr: true,
		},
		// Case 001
		{
			pso: "2023-11-01T00:00:00Z",
			now: "2023-09-03T00:00:00Z",
			ovr: true,
		},
		// Case 002
		{
			pso: "2023-11-01T00:00:00Z",
			now: "2023-09-25T00:00:00Z",
			ovr: true,
		},
		// Case 003
		{
			pso: "2023-11-01T00:00:00Z",
			now: "2023-10-01T00:00:00Z",
			ovr: true,
		},
		// Case 004
		{
			pso: "2023-11-01T00:00:00Z",
			now: "2023-10-10T00:00:00Z",
			ovr: true,
		},
		// Case 005
		{
			pso: "2023-11-01T00:00:00Z",
			now: "2023-10-28T00:00:00Z",
			ovr: true,
		},
		// Case 006
		{
			pso: "2023-11-01T00:00:00Z",
			now: "2023-11-01T00:00:00Z",
			ovr: false,
		},
		// Case 007
		{
			pso: "2023-11-01T00:00:00Z",
			now: "2023-11-03T00:00:00Z",
			ovr: false,
		},
		// Case 008
		{
			pso: "2023-11-01T00:00:00Z",
			now: "2023-11-05T00:00:00Z",
			ovr: false,
		},
		// Case 009
		{
			pso: "2023-11-01T00:00:00Z",
			now: "2023-11-08T00:00:00Z",
			ovr: false,
		},
		// Case 010
		{
			pso: "2023-11-01T00:00:00Z",
			now: "2023-11-25T00:00:00Z",
			ovr: false,
		},
		// Case 011
		{
			pso: "2023-11-01T00:00:00Z",
			now: "2023-11-28T00:00:00Z",
			ovr: false,
		},
		// Case 012
		{
			pso: "2023-11-01T00:00:00Z",
			now: "2023-12-01T00:00:00Z",
			ovr: false,
		},
		// Case 013
		{
			pso: "2023-11-01T00:00:00Z",
			now: "2023-12-28T00:00:00Z",
			ovr: false,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			ovr := ovrPrm(musTim(tc.pso), musTim(tc.now))
			if ovr != tc.ovr {
				t.Fatalf("expected %#v got %#v", tc.ovr, ovr)
			}
		})
	}
}

func musTim(str string) time.Time {
	tim, err := time.Parse("2006-01-02T15:04:05.999999Z", str)
	if err != nil {
		return time.Time{}
	}

	return tim
}
