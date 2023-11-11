package rulestorage

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

type Slicer []*Object

// Fltr returns a slicer implementation to remove certain objects from this
// list.
func (s Slicer) Fltr() Filter {
	return Filter(s)
}

// Incl returns the storage keys pointing to the event IDs meant to be included
// in the list associated to the underlying rules.
func (s Slicer) Incl() []string {
	var inc []string

	for _, x := range s {
		inc = append(inc, objectid.Fmt(x.Incl, x.KeyFmt())...)
	}

	return inc
}

// Rule returns all the rule IDs for the underling list of rule objects.
func (s Slicer) Rule() []objectid.ID {
	var ids []objectid.ID

	for _, x := range s {
		ids = append(ids, x.Rule)
	}

	return ids
}
