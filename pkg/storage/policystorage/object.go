package policystorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/format/hexformat"
	"github.com/NaoNaoOnline/apiserver/pkg/generic"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/tracer"
)

type Object struct {
	// Acce is the SMA record level, permission or role.
	Acce int64 `json:"acce"`
	// ChID is the chain ID, the unique identifier representing the blockchain
	// network on which this record is located. If the same SMA record exists on
	// multiple chains, the list of chain IDs is tracked in record creation order.
	ChID []int64 `json:"chid"`
	// Memb is the SMA record account, identity or user.
	Memb string `json:"memb"`
	// Syst is the SMA record context, resource or scope.
	Syst int64 `json:"syst"`
	// User is the user ID matched to this policy record on the fly, if available.
	// We do not persist the direct relationship between policy members and user
	// IDs because of several critical synchronization issues. The user ID will be
	// looked up on demand when searching for polices or when requesting
	// permission related resources and actions. Note that it might as well be
	// that there is no user association for a policy record intermittently.
	User objectid.ID `json:"-"`
}

func (r *Object) Verify() error {
	{
		if r == nil {
			return tracer.Mask(policyRecordEmptyError)
		}
	}

	{
		if r.Acce < 0 {
			return tracer.Mask(policyAcceNegativeError)
		}
	}

	{
		if len(r.ChID) == 0 {
			return tracer.Mask(policyChIDEmptyError)
		}
		if generic.Dup(r.ChID) {
			return tracer.Mask(policyChIDDuplicateError)
		}
	}

	{
		if r.Memb == "" {
			return tracer.Mask(policyMembEmptyError)
		}
		if len(r.Memb) != 42 {
			return tracer.Maskf(policyMembLengthError, "%d", len(r.Memb))
		}
		if !hexformat.Verify(r.Memb) {
			return tracer.Mask(policyMembFormatError)
		}
	}

	{
		if r.Syst < 0 {
			return tracer.Mask(policySystNegativeError)
		}
	}

	return nil
}
