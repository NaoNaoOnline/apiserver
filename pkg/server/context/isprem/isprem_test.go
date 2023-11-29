package isprem

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Test_Context_IsPrem(t *testing.T) {
	{
		v := FromContext(context.Background())
		if v {
			t.Fatal("expected", false, "got", true)
		}
	}

	{
		v := FromContext(NewContext(context.Background(), musTim("2023-11-01T00:00:00Z")))
		if v {
			t.Fatal("expected", false, "got", true)
		}
	}
}

func Test_Context_IsPrem_Time(t *testing.T) {
	testCases := []struct {
		prm string
		now string
		isp bool
	}{
		// Case 000
		{
			prm: "2023-11-01T00:00:00Z",
			now: "2023-10-01T00:00:00Z",
			isp: true,
		},
		// Case 001
		{
			prm: "2023-11-01T00:00:00Z",
			now: "2023-10-15T08:30:00Z",
			isp: true,
		},
		// Case 002
		{
			prm: "2023-11-01T00:00:00Z",
			now: "2023-10-31T18:48:23Z",
			isp: true,
		},
		// Case 003
		{
			prm: "2023-11-01T00:00:00Z",
			now: "2023-11-01T00:00:00Z",
			isp: false,
		},
		// Case 004
		{
			prm: "2023-11-01T00:00:00Z",
			now: "2023-11-05T04:20:00Z",
			isp: false,
		},
		// Case 005
		{
			prm: "2023-11-01T00:00:00Z",
			now: "2023-12-05T04:20:00Z",
			isp: false,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			isp := FromContext(NewContext(context.Background(), musTim(tc.prm)), musTim(tc.now))
			if isp != tc.isp {
				t.Fatalf("expected %#v got %#v", tc.isp, isp)
			}
		})
	}
}

func musTim(str string) time.Time {
	tim, err := time.Parse("2006-01-02T15:04:05.999999Z", str)
	if err != nil {
		panic(err)
	}

	return tim
}
