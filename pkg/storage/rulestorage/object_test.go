package rulestorage

import (
	"fmt"
	"testing"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
)

func Test_Storage_RuleStorage_Object_HasRes(t *testing.T) {
	testCases := []struct {
		obj *Object
		has bool
	}{
		// Case 000
		{
			obj: &Object{},
			has: false,
		},
		// Case 001
		{
			obj: &Object{
				Incl: []objectid.ID{"1234"},
			},
			has: true,
		},
		// Case 002
		{
			obj: &Object{
				Excl: []objectid.ID{"1234"},
			},
			has: true,
		},
		// Case 003
		{
			obj: &Object{
				Excl: []objectid.ID{"1234"},
				Incl: []objectid.ID{"1234"},
			},
			has: true,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			has := tc.obj.HasRes()

			if has != tc.has {
				t.Fatalf("expected %#v got %#v", tc.has, has)
			}
		})
	}
}
