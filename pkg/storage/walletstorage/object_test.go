package walletstorage

import (
	"errors"
	"fmt"
	"testing"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
)

func Test_Storage_Wallet_Object_Messtim(t *testing.T) {
	testCases := []struct {
		mes string
		tim string
	}{
		// Case 000
		{
			mes: "",
			tim: "0001-01-01 00:00:00 +0000 UTC",
		},
		// Case 001
		{
			mes: "foo",
			tim: "0001-01-01 00:00:00 +0000 UTC",
		},
		// Case 002
		{
			mes: "signing ownership of 0x7483••••ba5B at foo",
			tim: "0001-01-01 00:00:00 +0000 UTC",
		},
		// Case 003
		{
			mes: "signing ownership of 0x7483••••ba5B at 1695326302",
			tim: "2023-09-21 19:58:22 +0000 UTC",
		},
		// Case 004
		{
			mes: "signing ownership of 0x7483••••ba5B at 1560489846",
			tim: "2019-06-14 05:24:06 +0000 UTC",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			tim := witMes(tc.mes).Mestim().String()
			if tim != tc.tim {
				t.Fatalf("expected %#v got %#v", tc.tim, tim)
			}
		})
	}
}

func Test_Storage_Wallet_Object_Verify_Kind(t *testing.T) {
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
			err := witKin(tc.kin).VerifyObct()
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected %#v got %#v", tc.err, err)
			}
		})
	}
}

func Test_Storage_Wallet_Object_Verify_Labl(t *testing.T) {
	testCases := []struct {
		lab []string
		err error
	}{
		// Case 000
		{
			lab: nil,
			err: nil,
		},
		// Case 001
		{
			lab: []string{
				objectlabel.WalletUnassigned,
			},
			err: nil,
		},
		// Case 002
		{
			lab: []string{
				objectlabel.WalletAccounting,
			},
			err: nil,
		},
		// Case 003
		{
			lab: []string{
				objectlabel.WalletModeration,
			},
			err: nil,
		},
		// Case 004
		{
			lab: []string{
				objectlabel.WalletUnassigned,
				objectlabel.WalletUnassigned,
			},
			err: walletLablDuplicateError,
		},
		// Case 005
		{
			lab: []string{
				objectlabel.WalletAccounting,
				objectlabel.WalletAccounting,
			},
			err: walletLablDuplicateError,
		},
		// Case 006
		{
			lab: []string{
				objectlabel.WalletModeration,
				objectlabel.WalletModeration,
			},
			err: walletLablDuplicateError,
		},
		// Case 007
		{
			lab: []string{
				objectlabel.WalletUnassigned,
				objectlabel.WalletAccounting,
			},
			err: walletLablConflictError,
		},
		// Case 008
		{
			lab: []string{
				objectlabel.WalletUnassigned,
				objectlabel.WalletModeration,
			},
			err: walletLablConflictError,
		},
		// Case 009
		{
			lab: []string{
				objectlabel.WalletAccounting,
				objectlabel.WalletModeration,
			},
			err: walletLablConflictError,
		},
		// Case 010
		{
			lab: []string{
				objectlabel.WalletUnassigned,
				objectlabel.WalletAccounting,
				objectlabel.WalletModeration,
			},
			err: walletLablConflictError,
		},
		// Case 011
		{
			lab: []string{
				"",
			},
			err: walletLablInvalidError,
		},
		// Case 012
		{
			lab: []string{
				"foo",
			},
			err: walletLablInvalidError,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			err := witLab(tc.lab).VerifyPtch()
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected %#v got %#v", tc.err, err)
			}
		})
	}
}

func Test_Storage_Wallet_Object_Verify_Mess(t *testing.T) {
	testCases := []struct {
		mes string
		err error
	}{
		// Case 000
		{
			mes: "signing ownership of 0x7483••••ba5B at 1695326302",
			err: nil,
		},
		// Case 001
		{
			mes: "signing ownership of 0x7483••••ba5B at 1695",
			err: walletMessFormatError,
		},
		// Case 002
		{
			mes: "signing ownership of 0x7483••••ba5B at foo",
			err: walletMessFormatError,
		},
		// Case 003
		{
			mes: "signing ownership of foo at 1695326302",
			err: walletMessFormatError,
		},
		// Case 004
		{
			mes: "foo",
			err: walletMessFormatError,
		},
		// Case 005 ensures that a longer unix timestamp can be provided. The test
		// produces an error because the signature is invalid due to the static test
		// data. Though at the point of verifying the signature itself, the message
		// validation was already successful. So in case we get the wallet signature
		// error, the message format was successfully validated.
		{
			mes: "signing ownership of 0x7483••••ba5B at 1695326302538",
			err: walletSignatureInvalidError,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			err := witMes(tc.mes).VerifySign()
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
		User: objectid.ID("1234"),
	}
}

func witLab(lab []string) *Object {
	return &Object{
		Labl: objectfield.Strings{
			Data: lab,
		},
	}
}

func witMes(mes string) *Object {
	return &Object{
		Kind: "eth",
		Mess: mes,
		Pubk: "0x0437c4df64cdef106fe01c0c55a579d05a78bb97fc4151840ed712f154407a01e07c91b07da6d1bf5ffa4930b941f4787b44c2c7b88e1efd8da2905df5cbd59cda",
		Sign: "0xba7fc983705f2067588a0119abc2c0eee035f0b9dee47fb3a4f5603d057dc2dd0d8768a056e5a6a060aace35772f446a4f64a241a1988410e6f0ab2af28c16cb1b",
		User: objectid.ID("1234"),
	}
}
