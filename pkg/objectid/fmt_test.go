package objectid

import (
	"fmt"
	"testing"
)

func Test_ScoreID_Fmt_string(t *testing.T) {
	testCases := []struct {
		ids []string
		str string
		key []string
	}{
		// Case 000
		{
			ids: []string{
				"foo",
				"bar",
			},
			str: "des/eve/%s",
			key: []string{
				"des/eve/foo",
				"des/eve/bar",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			key := Fmt(tc.ids, tc.str)
			if !equal(key, tc.key) {
				t.Fatalf("expected %#v got %#v", tc.key, key)
			}
		})
	}
}

func Test_ScoreID_Fmt_String(t *testing.T) {
	testCases := []struct {
		ids []String
		str string
		key []string
	}{
		// Case 000
		{
			ids: []String{
				"foo",
				"bar",
			},
			str: "des/eve/%s",
			key: []string{
				"des/eve/foo",
				"des/eve/bar",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			key := Fmt(tc.ids, tc.str)
			if !equal(key, tc.key) {
				t.Fatalf("expected %#v got %#v", tc.key, key)
			}
		})
	}
}

// equal is copied from the go1.21 source at https://pkg.go.dev/slices#Equal.
func equal[S ~[]E, E comparable](s1, s2 S) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
