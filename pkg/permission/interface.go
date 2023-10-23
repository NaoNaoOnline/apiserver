package permission

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/policystorage"
	"github.com/NaoNaoOnline/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
)

type Interface interface {
	// BufferActv tries to read the active permission state from the underlying
	// sorted set in order to track these in the policy cache's memory. If no
	// active permission state can be found, then BufferActv is a noop and false
	// is returned.
	BufferActv() (bool, error)

	// EnsureActv executes BufferActv, and if BufferActv returns false, then an
	// update cycle will be initiated by emitting all necessary scrape tasks. At
	// the end of the initiated update cycle, BufferActv will be called again, and
	// the active permission state should be provisioned and be readily available
	// in process memory for all workers participating in the network.
	EnsureActv() error

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

	// ScrapeRcrd stores the given chain specific policy records for the next
	// merge update without affecting the currently active permissions. The
	// scraped policy records provided here will only take affect after a call to
	// Permission.Update.
	//
	//     @inp[0] the task definition specifying the chain to scrape from
	//     @inp[1] the task budget specifying an execution limit
	//
	ScrapeRcrd(*task.Task, *budget.Budget) error

	// SearchActv returns the latest aggregated version of the active permission
	// state indexed from all onchain smart contracts configured. That is, the
	// list of aggregated policy records representing the currently active
	// authorization states augmented on the fly with the respective user IDs, if
	// available.
	//
	//     @out[0] the list of aggregated policy records augmented with user IDs
	//
	SearchActv() ([]*policystorage.Object, error)

	// SearchUser returns the SMA members for the given user ID.
	//
	//     @inp[0] the user ID to search
	//     @out[0] the list of SMA members matching the given user ID
	//
	SearchUser(objectid.ID) ([]string, error)

	// UpdateActv merges all the scraped policy records in order to create a
	// unified set of active permissions. The dedicated sorted set for buffering
	// the scraped policy records is removed once the dedicated sorted set for the
	// active permissions was finalized. Further, the memory implementation of the
	// policy cache is synchronized with the new active permission states.
	UpdateActv() error
}
