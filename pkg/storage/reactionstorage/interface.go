package reactionstorage

type Interface interface {
	// Create persists new reaction objects.
	//
	//     @inp[0] the list of reaction objects providing reaction specific information
	//     @out[0] the list of reaction objects persisted internally
	//
	Create([]*Object) ([]*Object, error)

	// SearchBltn returns the static list of curated reaction icons natively
	// supported by the system. This method should only be used to bootstrap the
	// initial system state, never to serve RPC requests.
	SearchBltn() []*Object

	// SearchKind returns the reaction objects matching the given reaction kinds,
	// e.g. bltn or user.
	//
	//     @inp[0] the reaction kinds under which reaction objects are grouped together
	//     @out[0] the list of reaction objects matching the given reaction kinds
	//
	SearchKind([]string) ([]*Object, error)
}
