package reactionstorage

import (
	"errors"
	"fmt"
	"testing"
)

func Test_Storage_ReactionStorage_Object_Verify_Kind(t *testing.T) {
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
			kin: "user",
			err: nil,
		},
		// Case 002
		{
			kin: "foo",
			err: reactionKindInvalidError,
		},
		// Case 003
		{
			kin: "bar",
			err: reactionKindInvalidError,
		},
		// Case 004
		{
			kin: "",
			err: reactionKindInvalidError,
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
		Html: "html",
		Kind: kin,
		Name: "name",
	}
}
