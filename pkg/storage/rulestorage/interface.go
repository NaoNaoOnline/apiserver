package rulestorage

type Interface interface {
	// Create persists new rule objects.
	//
	//     @inp[0] the rule objects providing rule specific information
	//     @out[0] the rule objects mapped to their internal rule IDs
	//
	Create([]*Object) ([]*Object, error)
}
