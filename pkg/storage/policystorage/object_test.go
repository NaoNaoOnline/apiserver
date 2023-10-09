package policystorage

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func Test_Storage_PolicyStorage_Object_Verify_Kind(t *testing.T) {
	testCases := []struct {
		kin string
		err error
	}{
		// Case 000
		{
			kin: "CreateMember",
			err: nil,
		},
		// Case 001
		{
			kin: "CreateSystem",
			err: nil,
		},
		// Case 002
		{
			kin: "DeleteMember",
			err: nil,
		},
		// Case 003
		{
			kin: "DeleteSystem",
			err: nil,
		},
		// Case 004
		{
			kin: "foo",
			err: policyKindInvalidError,
		},
		// Case 005
		{
			kin: "bar",
			err: policyKindInvalidError,
		},
		// Case 006
		{
			kin: "",
			err: policyKindInvalidError,
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
		Acce: 2,
		Blck: []int64{18312712},
		ChID: []int64{42161},
		From: []string{"0x87f4CF42Ab88C9634d0EC942622991f02920d429"},
		Hash: []string{"0x4eff1cc0f9ec6ca9fc5c0a2f2d5d4496c017dfd50a15483e9a18bed627c34e64"},
		Kind: kin,
		Memb: "0x87f4CF42Ab88C9634d0EC942622991f02920d429",
		Syst: 0,
		Time: []time.Time{time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)},
	}
}
