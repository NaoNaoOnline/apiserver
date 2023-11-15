package labelstorage

import (
	"errors"
	"fmt"
	"testing"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
)

func Test_Storage_LabelStorage_Object_Verify_Kind(t *testing.T) {
	testCases := []struct {
		kin string
		err error
	}{
		// Case 000
		{
			kin: "bltn",
			err: nil,
		},
		// Case 001
		{
			kin: "cate",
			err: nil,
		},
		// Case 002
		{
			kin: "host",
			err: nil,
		},
		// Case 003
		{
			kin: "foo",
			err: labelKindInvalidError,
		},
		// Case 004
		{
			kin: "bar",
			err: labelKindInvalidError,
		},
		// Case 005
		{
			kin: "",
			err: labelKindInvalidError,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			err := witKin(tc.kin).Verify()
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected %#v got %#v", tc.err, err)
			}
		})
	}
}

func witKin(kin string) *Object {
	return &Object{
		Kind: kin,
		Name: objectfield.String{
			Data: "name",
		},
	}
}
