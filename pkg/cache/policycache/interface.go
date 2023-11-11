package policycache

import (
	"github.com/NaoNaoOnline/apiserver/pkg/storage/policystorage"
)

type Interface interface {
	// ExistsAcce verifies whether the given member has the given access within
	// the given system.
	//
	//     @inp[0] the SMA system
	//     @inp[1] the SMA member
	//     @inp[2] the SMA access
	//     @out[0] the bool expressing whether the given member has the given access within the given system
	//
	ExistsAcce(int64, string, int64) bool

	// ExistsMemb verifies whether the given member is a policy member.
	//
	//     @inp[0] the SMA member to verify
	//     @out[0] the bool expressing whether the given member is a policy member
	//
	ExistsMemb(string) bool

	// ExistsSyst verifies whether the given member is part of the given system.
	//
	//     @inp[0] the SMA system
	//     @inp[1] the SMA member to verify
	//     @out[0] the bool expressing whether the given member is part of the given system
	//
	ExistsSyst(int64, string) bool

	// SearchRcrd returns the latest aggregated version of cached policy records
	// indexed from all onchain smart contracts configured. That is, the list of
	// aggregated records representing the currently active authorization states.
	//
	//     @out[0] the list of aggregated policy records currently cached internally
	//
	SearchRcrd() []*policystorage.Object

	// UpdateRcrd merges all the buffered policy records provided, in order to
	// create a unified set of active permissions.
	//
	//     @inp[0] the list of buffered policy records to merge and activate
	//
	UpdateRcrd([]*policystorage.Object) error
}
