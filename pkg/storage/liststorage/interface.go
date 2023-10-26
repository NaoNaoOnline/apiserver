package liststorage

type Interface interface {
	// Create persists new list objects.
	//
	//     @inp[0] the list objects providing list specific information
	//     @out[0] the list objects mapped to their internal list IDs
	//
	Create([]*Object) ([]*Object, error)
}
