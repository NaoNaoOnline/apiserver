package subscriptionstorage

import (
	"fmt"
	"testing"
	"time"
)

func Test_Storage_Subscription_Redis_Create_subRen(t *testing.T) {
	testCases := []struct {
		exi []*Object
		uni time.Time
		now time.Time
		ren bool
	}{
		// Case 000
		{
			exi: nil,
			uni: musTim("2023-11-01T00:00:00Z"),
			now: musTim("2023-10-25T15:30:00Z"),
			ren: false,
		},
		// Case 001
		{
			exi: []*Object{
				{
					Unix: musTim("2023-08-01T00:00:00Z"),
				},
				{
					Unix: musTim("2023-09-01T00:00:00Z"),
				},
			},
			uni: musTim("2023-11-01T00:00:00Z"),
			now: musTim("2023-10-25T15:30:00Z"),
			ren: false,
		},
		// Case 002
		{
			exi: []*Object{
				{
					Unix: musTim("2023-10-01T00:00:00Z"),
				},
			},
			uni: musTim("2023-11-01T00:00:00Z"),
			now: musTim("2023-10-25T15:30:00Z"),
			ren: true,
		},
		// Case 003
		{
			exi: []*Object{
				{
					Unix: musTim("2023-08-01T00:00:00Z"),
				},
				{
					Unix: musTim("2023-09-01T00:00:00Z"),
				},
				{
					Unix: musTim("2023-10-01T00:00:00Z"),
				},
			},
			uni: musTim("2023-11-01T00:00:00Z"),
			now: musTim("2023-10-25T15:30:00Z"),
			ren: true,
		},
		// Case 004, you subscribe again for November in the middle of November.
		// This is not a "renewal", because you did not have a subscription for the
		// first 15 days of the month.
		{
			exi: []*Object{
				{
					Unix: musTim("2023-08-01T00:00:00Z"),
				},
				{
					Unix: musTim("2023-09-01T00:00:00Z"),
				},
				{
					Unix: musTim("2023-10-01T00:00:00Z"),
				},
			},
			uni: musTim("2023-11-01T00:00:00Z"),
			now: musTim("2023-11-15T15:30:00Z"),
			ren: false,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			ren := subRen(tc.exi, tc.uni, tc.now)
			if ren != tc.ren {
				t.Fatalf("expected %#v got %#v", tc.ren, ren)
			}
		})
	}
}
