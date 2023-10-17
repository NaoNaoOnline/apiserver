package permission

import (
	"github.com/NaoNaoOnline/apiserver/pkg/cache/policycache"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
)

type Interface interface {
	// ExistsAcce verifies whether the given user ID has the given access within
	// the given system.
	//
	//     @inp[0] the SMA system
	//     @inp[1] the user ID potentially representing a SMA member
	//     @inp[2] the SMA access
	//     @out[0] the bool expressing whether the given user has the given access within the given system
	//
	ExistsAcce(int64, objectid.ID, int64) (bool, error)

	// ExistsMemb verifies whether the given user ID is a policy member.
	//
	//     @inp[0] the user ID potentially representing a SMA member to verify
	//     @out[0] the bool expressing whether the given user is a policy member
	//
	ExistsMemb(objectid.ID) (bool, error)

	// ExistsSyst verifies whether the given user ID is part of the given system.
	//
	//     @inp[0] the SMA system
	//     @inp[1] the user ID potentially representing a SMA member to verify
	//     @out[0] the bool expressing whether the given user is part of the given system
	//
	ExistsSyst(int64, objectid.ID) (bool, error)

	// SearchMemb returns the user IDs for the given addresses.
	//
	//     @inp[0] the SMA members to search
	//     @out[0] the list of user IDs matching the given SMA members
	//
	SearchMemb([]string) ([]objectid.ID, error)

	// SearchRcrd returns the latest aggregated version of cached policy records
	// indexed from all onchain smart contracts configured. That is, the list of
	// aggregated records representing the currently active authorization states
	// augmented on the fly with the respective user IDs, if available.
	//
	//     @out[0] the list of aggregated policy records augmented with user IDs
	//
	SearchRcrd() ([]*policycache.Record, error)

	// SearchUser returns the SMA members for the given user ID.
	//
	//     @inp[0] the user ID to search
	//     @out[0] the list of SMA members matching the given user ID
	//
	SearchUser(objectid.ID) ([]string, error)
}
