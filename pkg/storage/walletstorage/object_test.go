package walletstorage

import (
	"errors"
	"fmt"
	"testing"
)

func Test_Storage_walletstorage_Object_Verify_Kind(t *testing.T) {
	testCases := []struct {
		kin string
		err error
	}{
		// Case 000
		{
			kin: "eth",
			err: nil,
		},
		// Case 001
		{
			kin: "foo",
			err: walletKindInvalidError,
		},
		// Case 002
		{
			kin: "bar",
			err: walletKindInvalidError,
		},
		// Case 003
		{
			kin: "",
			err: walletKindInvalidError,
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
		Mess: "signing ownership of 0x7483••••ba5B at 1695326302",
		Pubk: "0x0437c4df64cdef106fe01c0c55a579d05a78bb97fc4151840ed712f154407a01e07c91b07da6d1bf5ffa4930b941f4787b44c2c7b88e1efd8da2905df5cbd59cda",
		Sign: "0xba7fc983705f2067588a0119abc2c0eee035f0b9dee47fb3a4f5603d057dc2dd0d8768a056e5a6a060aace35772f446a4f64a241a1988410e6f0ab2af28c16cb1b",
	}
}
