package objectid

import (
	"fmt"
	"slices"
	"testing"
)

func Test_ObjectID_Fmt_string(t *testing.T) {
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
		// Case 001
		{
			ids: []string{
				"foo",
				"bar",
			},
			str: fmt.Sprintf("wal/use/%s/obj/%s", "1234", "%s"),
			key: []string{
				"wal/use/1234/obj/foo",
				"wal/use/1234/obj/bar",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			key := Fmt(tc.ids, tc.str)
			if !slices.Equal(key, tc.key) {
				t.Fatalf("expected %#v got %#v", tc.key, key)
			}
		})
	}
}

func Test_ObjectID_Fmt_String(t *testing.T) {
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
		// Case 001
		{
			ids: []String{
				"foo",
				"bar",
			},
			str: fmt.Sprintf("wal/use/%s/obj/%s", "1234", "%s"),
			key: []string{
				"wal/use/1234/obj/foo",
				"wal/use/1234/obj/bar",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			key := Fmt(tc.ids, tc.str)
			if !slices.Equal(key, tc.key) {
				t.Fatalf("expected %#v got %#v", tc.key, key)
			}
		})
	}
}
