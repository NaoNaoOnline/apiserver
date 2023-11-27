package subscriptionstorage

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
)

func Test_Storage_Subscription_Object_Verify_Unix(t *testing.T) {
	testCases := []struct {
		uni string
		err error
	}{
		// case 000
		{
			uni: "2023-10-01T00:00:00Z",
			err: subscriptionUnixInvalidError,
		},
		// case 001
		{
			uni: "2023-11-01T00:00:00Z",
			err: nil,
		},
		// Case 002
		{
			uni: "2023-10-05T00:00:00Z",
			err: subscriptionUnixInvalidError,
		},
		// Case 003
		{
			uni: "2023-11-05T00:00:00Z",
			err: subscriptionUnixInvalidError,
		},
		// case 004
		{
			uni: "2023-10-01T15:34:03Z",
			err: subscriptionUnixInvalidError,
		},
		// case 005
		{
			uni: "2023-11-01T15:34:03Z",
			err: subscriptionUnixInvalidError,
		},
		// case 006
		{
			uni: "2023-10-21T05:30:47Z",
			err: subscriptionUnixInvalidError,
		},
		// case 007
		{
			uni: "2023-11-21T05:30:47Z",
			err: subscriptionUnixInvalidError,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			err := witUni(tc.uni).Verify()
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected %#v got %#v", tc.err, err)
			}
		})
	}
}

func witUni(uni string) *Object {
	return &Object{
		ChID: 1,
		Crtr: []string{"0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},
		Sbsc: "0x70997970C51812dc3A010C7d01b50e0d17dc79C8",
		Unix: musTim(uni),
		User: objectid.ID("1234"),

		time: &faker{
			tim: musTim("2023-11-01T15:34:03Z"),
		},
	}
}

func musTim(str string) time.Time {
	tim, err := time.Parse("2006-01-02T15:04:05.999999Z", str)
	if err != nil {
		panic(err)
	}

	return tim
}
