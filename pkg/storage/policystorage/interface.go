package policystorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
)

type Interface interface {
	// Create persists new policy objects.
	//
	//     @inp[0] the policy objects providing policy specific information
	//     @out[0] the policy objects mapped to their internal policy IDs
	//
	Create([]*Object) ([]*Object, error)

	// DeletePlcy purges the given policy objects.
	//
	//     @inp[0] the policy objects to delete
	//     @out[0] the list of operation states related to the purged policy objects
	//
	DeletePlcy([]*Object) ([]objectstate.String, error)

	// ExistsSyst verifies whether the given user is part of the given system.
	//
	//     @inp[0] the SMA system
	//     @inp[1] the user ID to verify
	//     @out[0] the bool expressing whether the given user is part of the given system
	//
	ExistsSyst(int64, objectid.ID) (bool, error)

	// ExistsMemb verifies whether the given user is a policy member
	//
	//     @inp[0] the user ID to verify
	//     @out[0] the bool expressing whether the given user is a policy member
	//
	ExistsMemb(objectid.ID) (bool, error)

	// ExistsAcce verifies whether the given user has the given access within the
	// given system.
	//
	//     @inp[0] the SMA system
	//     @inp[1] the user ID to verify
	//     @inp[2] the SMA access
	//     @out[0] the bool expressing whether the given user has the given access within the given system
	//
	ExistsAcce(int64, objectid.ID, int64) (bool, error)

	// SearchAggr returns the latest aggregated version of cached policy records
	// indexed from all onchain smart contracts configured. That is, the list of
	// aggregated records representing the currently active authorization states,
	// minus the list of records that have been removed so far.
	//
	//     @out[0] the list of aggregated policy records currently cached internally
	//     @out[1] the list of removed policy records currently cached internally
	//
	SearchAggr() ([]*Object, []*Object, error)

	// SearchKind returns the policy objects matching the given policy kinds, e.g.
	// CreateMember, CreateSystem, DeleteMember, DeleteSystem.
	//
	//     @inp[0] the policy kinds under which policy objects are grouped together
	//     @out[0] the list of policy objects matching the given policy kinds
	//
	SearchKind([]string) ([]*Object, error)
}
