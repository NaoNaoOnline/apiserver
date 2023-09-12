package labelstorage

type Interface interface {
	// Create persists new label objects, if none exists already with the given
	// name.
	//
	//     @inp[0] the label objects providing label specific information
	//     @out[0] the label objects mapped to their internal label ID
	//
	Create([]*Object) ([]*Object, error)
	// SearchBltn returns the static list of curated event labels natively
	// supported by the system. This method should only be used to bootstrap the
	// initial system state, never to serve RPC requests.
	SearchBltn() []*Object
	// SearchKind returns the label objects of the given kind, e.g. bltn, cate or
	// host.
	//
	//     @inp[0] the label kinds under which label objects are grouped together
	//     @out[0] the list of label objects of either kind category or kind host
	//
	SearchKind([]string) ([]*Object, error)
}
