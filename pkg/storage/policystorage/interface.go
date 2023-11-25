package policystorage

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

	// DeleteBffr purges all buffered policy objects.
	DeleteBffr() error

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
}
