package policystorage

import (
	"regexp"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/tracer"
)

type Object struct {
	// Acce is the SMA record level, permission or role.
	Acce int64 `json:"acce"`
	// Blck is the block height at which this record got created. If the same SMA
	// record exists on multiple chains, the list of blocks is tracked in record
	// creation order.
	Blck []int64 `json:"blck"`
	// ChID is the chain ID, the unique identifier representing the blockchain
	// network on which this record is located. If the same SMA record exists on
	// multiple chains, the list of chain IDs is tracked in record creation order.
	ChID []int64 `json:"chid"`
	// Crtd is the unix timestamp in seconds at which the record got cached
	// internally.
	Crtd time.Time `json:"crtd"`
	// From is the record creator, the sender of the transaction that submitted
	// this record. If the same SMA record exists on multiple chains, the list of
	// creators is tracked in record creation order.
	From []string `json:"from"`
	// Hash is the onchain transaction hash that submitted this record. If the
	// same SMA record exists on multiple chains, the list of transaction hashes
	// is tracked in record creation order.
	Hash []string `json:"hash"`
	// Kind is the record type.
	//
	//     CreateMember for records of members being created within a system
	//     CreateSystem for records of systems being created
	//     DeleteMember for records of members being deleted within a system
	//     DeleteSystem for records of systems being deleted
	//
	Kind string `json:"kind"`
	// Memb is the SMA record account, identity or user.
	Memb string `json:"memb"`
	// Plcy is the internal ID of the record being created.
	Plcy objectid.ID `json:"plcy"`
	// Syst is the SMA record context, resource or scope.
	Syst int64 `json:"syst"`
	// Time is the unix timestamp in seconds at which the record got created
	// externally. Note that policy records are external data objects that get
	// created somewhere else, in this case onchain, and thus must bring a
	// creation timestamp with them. So the created timestamp here originates from
	// some blockchain network. If the same SMA record exists on multiple chains,
	// the list of creation time is tracked in record creation order.
	Time []time.Time `json:"time"`
}

var (
	hexaexpr = regexp.MustCompile(`^0x[0-9a-fA-F]+$`)
)

func (o *Object) Eqlrec(x *Object) bool {
	return o.Syst == x.Syst && o.Memb == x.Memb && o.Acce == x.Acce
}

func (o *Object) Verify() error {
	{
		if o.Kind != "CreateMember" && o.Kind != "CreateSystem" && o.Kind != "DeleteMember" && o.Kind != "DeleteSystem" {
			return tracer.Maskf(policyKindInvalidError, o.Kind)
		}
	}

	{
		if o.Acce < 0 {
			return tracer.Mask(policyAcceNegativeError)
		}
	}

	{
		if len(o.Time) == 0 {
			return tracer.Mask(policyTimeEmptyError)
		}
		for _, x := range o.Time {
			if x.IsZero() {
				return tracer.Mask(policyTimeEmptyError)
			}
		}
	}

	{
		if o.Memb == "" {
			return tracer.Mask(policyMembEmptyError)
		}
		if len(o.Memb) != 42 {
			return tracer.Maskf(policyMembLengthError, "%d", len(o.Memb))
		}
		if !hexaexpr.MatchString(o.Memb) {
			return tracer.Mask(policyMembFormatError)
		}
	}

	{
		if o.Syst < 0 {
			return tracer.Mask(policySystNegativeError)
		}
	}

	return nil
}
