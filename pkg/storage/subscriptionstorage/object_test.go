package subscriptionstorage

import (
	"errors"
	"fmt"
	"testing"
)

func Test_Storage_Subscription_Object_VerifyUnix_VerifyOnce(t *testing.T) {
	testCases := []struct {
		uni string
		now string
		err error
	}{
		// Case 000
		{
			uni: "2023-10-01T00:00:00Z",
			now: "2023-11-01T00:00:03Z",
			err: subscriptionUnixInvalidError,
		},
		// Case 001, today is the 8th of November, so creating a new subscription
		// for October should not work.
		{
			uni: "2023-10-01T00:00:00Z",
			now: "2023-11-08T15:34:03Z",
			err: subscriptionUnixInvalidError,
		},
		// Case 002, today is the 8th of November, so creating a new subscription
		// for November should work.
		{
			uni: "2023-11-01T00:00:00Z",
			now: "2023-11-08T15:34:03Z",
			err: nil,
		},
		// Case 003
		{
			uni: "2023-11-01T00:00:00Z",
			now: "2023-11-02T07:05:00Z",
			err: nil,
		},
		// Case 004, today is the 8th of November, so creating a new subscription
		// for December should not work.
		{
			uni: "2023-12-01T00:00:00Z",
			now: "2023-11-08T15:34:03Z",
			err: subscriptionUnixInvalidError,
		},
		// Case 005
		{
			uni: "2023-10-05T00:00:00Z",
			now: "2023-11-08T15:34:03Z",
			err: subscriptionUnixInvalidError,
		},
		// Case 006
		{
			uni: "2023-11-05T00:00:00Z",
			now: "2023-11-08T15:34:03Z",
			err: subscriptionUnixInvalidError,
		},
		// Case 007
		{
			uni: "2023-10-01T15:34:03Z",
			now: "2023-11-08T15:34:03Z",
			err: subscriptionUnixInvalidError,
		},
		// Case 008
		{
			uni: "2023-11-01T15:34:03Z",
			now: "2023-11-08T15:34:03Z",
			err: subscriptionUnixInvalidError,
		},
		// Case 009
		{
			uni: "2023-10-21T05:30:47Z",
			now: "2023-11-08T15:34:03Z",
			err: subscriptionUnixInvalidError,
		},
		// Case 010
		{
			uni: "2023-11-21T05:30:47Z",
			now: "2023-11-08T15:34:03Z",
			err: subscriptionUnixInvalidError,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			err := witUni(tc.uni).VerifyUnix(VerifyOnce(musTim(tc.now)))
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected %#v got %#v", tc.err, err)
			}
		})
	}
}

func Test_Storage_Subscription_Object_VerifyUnix_VerifyRenw(t *testing.T) {
	testCases := []struct {
		uni string
		now string
		err error
	}{
		// Case 000, today is the 22th of November, so creating a renewal
		// subscription for October should not work.
		{
			uni: "2023-10-01T00:00:00Z",
			now: "2023-11-22T15:34:03Z",
			err: subscriptionUnixRenewalError,
		},
		// Case 001, today is the 22th of November, so creating a renewal
		// subscription for November should not work.
		{
			uni: "2023-11-01T00:00:00Z",
			now: "2023-11-22T15:34:03Z",
			err: subscriptionUnixRenewalError,
		},
		// Case 002, today is the 22th of November, so creating a renewal
		// subscription for December should not work.
		{
			uni: "2023-12-01T00:00:00Z",
			now: "2023-11-22T15:34:03Z",
			err: subscriptionUnixRenewalError,
		},
		// Case 003, today is the 25th of November, so creating a renewal
		// subscription for October should not work.
		{
			uni: "2023-10-01T00:00:00Z",
			now: "2023-11-25T15:34:03Z",
			err: subscriptionUnixInvalidError,
		},
		// Case 004, today is the 25th of November, so creating a renewal
		// subscription for November should not work.
		{
			uni: "2023-11-01T00:00:00Z",
			now: "2023-11-25T15:34:03Z",
			err: subscriptionUnixInvalidError,
		},
		// Case 005, today is the 25th of November, so creating a renewal
		// subscription for December should work.
		{
			uni: "2023-12-01T00:00:00Z",
			now: "2023-11-25T15:34:03Z",
			err: nil,
		},
		// Case 006
		{
			uni: "2023-12-01T00:00:00Z",
			now: "2023-11-28T20:08:24Z",
			err: nil,
		},
		// Case 007
		{
			uni: "2023-10-05T00:00:00Z",
			now: "2023-11-25T15:34:03Z",
			err: subscriptionUnixInvalidError,
		},
		// Case 008
		{
			uni: "2023-11-05T00:00:00Z",
			now: "2023-11-25T15:34:03Z",
			err: subscriptionUnixInvalidError,
		},
		// Case 009
		{
			uni: "2023-10-01T15:34:03Z",
			now: "2023-11-25T15:34:03Z",
			err: subscriptionUnixInvalidError,
		},
		// Case 010
		{
			uni: "2023-11-01T15:34:03Z",
			now: "2023-11-25T15:34:03Z",
			err: subscriptionUnixInvalidError,
		},
		// Case 011
		{
			uni: "2023-10-21T05:30:47Z",
			now: "2023-11-25T15:34:03Z",
			err: subscriptionUnixInvalidError,
		},
		// Case 012
		{
			uni: "2023-11-21T05:30:47Z",
			now: "2023-11-25T15:34:03Z",
			err: subscriptionUnixInvalidError,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			err := witUni(tc.uni).VerifyUnix(VerifyRenw(musTim(tc.now)))
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected %#v got %#v", tc.err, err)
			}
		})
	}
}

func witUni(uni string) *Object {
	return &Object{
		Unix: musTim(uni),
	}
}
