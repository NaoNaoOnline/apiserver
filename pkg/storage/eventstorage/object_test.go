package eventstorage

import (
	"fmt"
	"testing"
)

func Test_Storage_EventStorage_Object_Pltfrm(t *testing.T) {
	testCases := []struct {
		lin string
		lab string
	}{
		// Case 000
		{
			lin: "https://discord.gg",
			lab: "discord",
		},
		// Case 001
		{
			lin: "https://discord.com/invite/HESeTrU3",
			lab: "discord",
		},
		// Case 002
		{
			lin: "https://meet.google.com/cfz-zjkr-njr",
			lab: "google",
		},
		// Case 003
		{
			lin: "https://www.twitter.com",
			lab: "twitter",
		},
		// Case 004
		{
			lin: "https://twitter.com",
			lab: "twitter",
		},
		// Case 005
		{
			lin: "https://twitter.com/i/spaces/47lHrzOTWKnKF?s=20",
			lab: "twitter",
		},
		// Case 006
		{
			lin: "https://www.twitch.tv",
			lab: "twitch",
		},
		// Case 007
		{
			lin: "https://twitch.tv",
			lab: "twitch",
		},
		// Case 008
		{
			lin: "https://www.youtube.com/",
			lab: "youtube",
		},
		// Case 009
		{
			lin: "https://youtube.com/watch?v=T3uHkIOlxwv",
			lab: "youtube",
		},
		// Case 010
		{
			lin: "https://go.dev/play",
			lab: "go",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			lab := (&Object{Link: tc.lin}).Pltfrm()
			if lab != tc.lab {
				t.Fatalf("expected %#v got %#v", tc.lab, lab)
			}
		})
	}
}
