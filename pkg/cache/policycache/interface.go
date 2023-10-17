package policycache

type Interface interface {
	// Buffer stores the given chain specific policy records for the next merge
	// update without affecting the currently active permissions. The buffered
	// policy records provided here will only take affect after a call to
	// Memory.Update.
	//
	//     @inp[0] the list of chain specific policy records to store for the next merge update
	//
	Buffer([]*Record) error

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
	SearchRcrd() []*Record

	// Update merges all the buffered policy records in order to create a unified
	// set of active permissions.
	Update() error
}
