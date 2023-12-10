package rulestorage

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

type Slicer []*Object

// Rule returns all the rule IDs for the underling list of rule objects.
func (s Slicer) Rule() []objectid.ID {
	var ids []objectid.ID

	for _, x := range s {
		ids = append(ids, x.Rule)
	}

	return ids
}
