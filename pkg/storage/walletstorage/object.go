package walletstorage

import (
	"regexp"
	"strconv"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/generic"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/xh3b4sd/tracer"
)

type Object struct {
	// Addr is the hex encoded wallet address.
	Addr objectfield.String `json:"addr"`
	// Crtd is the time at which the wallet got created.
	Crtd time.Time `json:"crtd"`
	// Dltd is the time at which the wallet got deleted.
	Dltd time.Time `json:"dltd,omitempty"`
	// Kind is the wallet type.
	//
	//     eth for ethereum wallets
	//
	Kind string `json:"kind"`
	// Labl is the designated purpose this wallet serves. Kept empty at the time
	// of creation. Can be set after the signature verification challenge.
	//
	//     unassigned for opting out
	//     accounting for revenue sharing
	//     moderation for content moderation
	//
	Labl objectfield.Strings `json:"labl"`
	// User is the user ID creating this wallet.
	User objectid.ID `json:"user"`
	// Wllt is the ID of the wallet being created.
	Wllt objectid.ID `json:"wllt"`

	// Message, public key and signature are only ephemeral data we use in
	// transit. The properties below are part of the cryptographic verification
	// process to ensure that a user does in fact own the wallet they create.

	// Mess is the message to sign.
	Mess string `json:"-"`
	// Pubk is the hex encoded public key expected to be recovered from the given
	// signature.
	Pubk string `json:"-"`
	// Sign is the signature of the given message.
	Sign string `json:"-"`
}

func (o *Object) Comadd() common.Address {
	poi, err := crypto.UnmarshalPubkey(o.Pubkey())
	if err != nil {
		return common.BytesToAddress(nil)
	}

	return crypto.PubkeyToAddress(*poi)
}

func (o *Object) Digest() []byte {
	return accounts.TextHash([]byte(o.Mess))
}

func (o *Object) HasLab(lab string) bool {
	for _, x := range o.Labl.Data {
		if x == lab {
			return true
		}
	}

	return false
}

func (o *Object) Mestim() time.Time {
	var err error

	var sub []string
	{
		sub = unixexpr.FindStringSubmatch(o.Mess)
	}

	if len(sub) != 1 {
		return time.Time{}
	}

	var uni int64
	{
		uni, err = strconv.ParseInt(sub[0], 10, 64)
		if err != nil {
			panic(err)
		}
	}

	return time.Unix(uni, 0).UTC()
}

func (o *Object) Pubkey() []byte {
	pub, err := hexutil.Decode(o.Pubk)
	if err != nil {
		return nil
	}

	return pub
}

func (o *Object) Signtr() []byte {
	sig, err := hexutil.Decode(o.Sign)
	if err != nil {
		return nil
	}

	return sig[:len(sig)-1]
}

var (
	hexaexpr = regexp.MustCompile(`^0x[0-9a-fA-F]+$`)
	messexpr = regexp.MustCompile(`^signing ownership of 0x[0-9a-fA-F]{4}••••[0-9a-fA-F]{4} at [0-9]{10,}$`)
	unixexpr = regexp.MustCompile(`[0-9]{10,}$`)
)

func (o *Object) VerifyObct() error {
	{
		if o.Kind != "eth" {
			return tracer.Maskf(walletKindInvalidError, o.Kind)
		}
	}

	return nil
}

func (o *Object) VerifyPtch() error {
	{
		// Wallet labels cannot be set multiple times.
		if generic.Dup(o.Labl.Data) {
			return tracer.Mask(walletLablDuplicateError)
		}

		// Wallet labels cannot be arbitrary.
		for _, x := range o.Labl.Data {
			if x != objectlabel.WalletUnassigned && x != objectlabel.WalletAccounting && x != objectlabel.WalletModeration {
				return tracer.Maskf(walletLablInvalidError, x)
			}
		}

		// Wallet labels cannot be used together.
		for _, x := range o.Labl.Data {
			if o.HasLab(objectlabel.WalletUnassigned) && (x == objectlabel.WalletAccounting || x == objectlabel.WalletModeration) {
				return tracer.Maskf(walletLablConflictError, x)
			}
			if o.HasLab(objectlabel.WalletAccounting) && (x == objectlabel.WalletUnassigned || x == objectlabel.WalletModeration) {
				return tracer.Maskf(walletLablConflictError, x)
			}
			if o.HasLab(objectlabel.WalletModeration) && (x == objectlabel.WalletUnassigned || x == objectlabel.WalletAccounting) {
				return tracer.Maskf(walletLablConflictError, x)
			}
		}
	}

	return nil
}

func (o *Object) VerifySign() error {
	{
		if o.Mess == "" {
			return tracer.Mask(walletMessEmptyError)
		}
		if !messexpr.MatchString(o.Mess) {
			return tracer.Mask(walletMessFormatError)
		}
	}

	{
		if o.Pubk == "" {
			return tracer.Mask(walletPubkEmptyError)
		}
		if len(o.Pubk) != 132 {
			return tracer.Maskf(walletPubkLengthError, "%d", len(o.Pubk))
		}
		if !hexaexpr.MatchString(o.Pubk) {
			return tracer.Mask(walletPubkFormatError)
		}
	}

	{
		if o.Sign == "" {
			return tracer.Mask(walletSignEmptyError)
		}
		if len(o.Sign) != 132 {
			return tracer.Maskf(walletSignLengthError, "%d", len(o.Sign))
		}
		if !hexaexpr.MatchString(o.Sign) {
			return tracer.Mask(walletSignFormatError)
		}
	}

	{
		if !crypto.VerifySignature(o.Pubkey(), o.Digest(), o.Signtr()) {
			return tracer.Mask(walletSignatureInvalidError)
		}
	}

	// In case of an update process we have the wallet address of the existing
	// wallet object already available. After the wallet signature got verified
	// again, we need to ensure that the recovered address does not change across
	// the lifetime of the wallet object.
	{
		if o.Addr.Data != "" && o.Addr.Data != o.Comadd().Hex() {
			return tracer.Mask(walletAddrChangeError)
		}
	}

	return nil
}
