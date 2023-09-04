package labelstorage

type Interface interface {
	// Create persists new label objects, if none exists already with the given
	// name.
	//
	//     @inp[0] the label objects providing label specific information
	//     @out[0] the label objects mapped to their internal label ID
	//
	Create([]*Object) ([]*Object, error)
	// Search returns the label objects of the given kind.
	//
	//     @inp[0] the label kinds under which label objects are grouped together
	//     @out[0] the list of label objects of either kind category or kind host
	//
	Search([]string) ([]*Object, error)
}
