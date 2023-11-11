package policystorage

import "time"

type Interface interface {
	// CreateActv sets the active permission state to the provided list of policy
	// objects.
	//
	//     @inp[0] the list of policy objects for the new active permission states
	//
	CreateActv([]*Object) error

	// CreateBffr persists the given policy objects for the provided chain ID, if
	// those chain IDs are consistent.
	//
	//     @inp[0] the policy objects providing chain IDs
	//
	CreateBffr([]*Object) error

	// CreateLock creates an indicator for a new ongoing update lifecycle.
	CreateLock() error

	// CreateTime sets the timestamp of the most recent update lifecycle.
	CreateTime() error

	// DeleteBffr purges all buffered policy objects.
	DeleteBffr() error

	// DeleteLock removes the indicator for the past update lifecycle.
	DeleteLock() error

	// ExistsLock returns whether an update lifecycle is currently ongoing.
	//
	//     @out[0] the bool expressing whether an update lifecycle is currently ongoing
	//
	ExistsLock() (bool, error)

	// SearchActv returns all currently active permission states.
	//
	//     @out[0] the list of currently active permission states
	//
	SearchActv() ([]*Object, error)

	// SearchBffr returns all buffered policy objects.
	//
	//     @out[0] the list of buffered policy objects
	//
	SearchBffr() ([]*Object, error)

	// SearchTime returns the timestamp of the most recent update lifecycle.
	//
	//     @out[0] the timestamp of the most recent update lifecycle
	//
	SearchTime() (time.Time, error)
}
