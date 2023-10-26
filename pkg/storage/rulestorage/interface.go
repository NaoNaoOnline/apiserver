package rulestorage

type Interface interface {
	// Create persists new rule objects, if none exists already with the given
	// name.
	//
	//     @inp[0] the rule objects providing rule specific information
	//     @out[0] the rule objects mapped to their internal rule IDs
	//
	Create([]*Object) ([]*Object, error)
}
