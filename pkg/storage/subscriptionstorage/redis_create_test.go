package subscriptionstorage

import (
	"fmt"
	"testing"
)

func Test_Storage_Subscription_Redis_Create_subRen(t *testing.T) {
	testCases := []struct {
		exi []*Object
		cur *Object
		ren bool
	}{
		// Case 000
		{
			exi: nil,
			cur: &Object{
				Unix: musTim("2023-11-01T00:00:00Z"),
			},
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
			cur: &Object{
				Unix: musTim("2023-11-01T00:00:00Z"),
			},
			ren: false,
		},
		// Case 002
		{
			exi: []*Object{
				{
					Unix: musTim("2023-10-01T00:00:00Z"),
				},
			},
			cur: &Object{
				Unix: musTim("2023-11-01T00:00:00Z"),
			},
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
			cur: &Object{
				Unix: musTim("2023-11-01T00:00:00Z"),
			},
			ren: true,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			ren := subRen(tc.exi, tc.cur)
			if ren != tc.ren {
				t.Fatalf("expected %#v got %#v", tc.ren, ren)
			}
		})
	}
}
