package rulestorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
)

type Interface interface {
	// Create persists new rule objects.
	//
	//     @inp[0] the rule objects providing rule specific information
	//     @out[0] the rule objects mapped to their internal rule IDs
	//
	Create([]*Object) ([]*Object, error)

	// Delete purges the given rule objects.
	//
	//     @inp[0] the rule objects to delete
	//     @out[0] the list of operation states related to the purged rule objects
	//
	Delete([]*Object) ([]objectstate.String, error)

	// SearchList returns the rule objects belonging to the given list IDs.
	//
	//     @inp[0] the list IDs to search rules for
	//     @out[0] the list of rule objects belonging to the given list IDs
	//
	SearchList([]objectid.ID) ([]*Object, error)

	// SearchRule returns the rule objects matching the given rule IDs.
	//
	//     @inp[0] the rule IDs to search for
	//     @out[0] the list of rule objects matching the given rule IDs
	//
	SearchRule([]objectid.ID) ([]*Object, error)
}
