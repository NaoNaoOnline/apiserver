package descriptionstorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/objectstate"
)

type Interface interface {
	// Create persists new description objects mapped to the referenced events.
	//
	//     @inp[0] the description objects mapped to the referenced events
	//     @out[0] the description objects mapped to their internal description ID
	//
	Create([]*Object) ([]*Object, error)
	// SearchDesc returns the description objects matching the given description
	// IDs.
	//
	//     @inp[0] the description IDs to search for
	//     @out[0] the list of description objects matching the given description IDs
	//
	SearchDesc([]objectid.String) ([]*Object, error)
	// SearchEvnt returns the description objects belonging to the given event
	// IDs.
	//
	//     @inp[0] the event IDs to search descriptions for
	//     @out[0] the list of description objects belonging the given event IDs
	//
	SearchEvnt([]objectid.String) ([]*Object, error)
	// Update modifies the existing description objects by applying the given
	// RFC6902 JSON-Patches to the underlying JSON documents. The list items are
	// used according to their respective indices, e.g. the second patch is
	// applied to the second object.
	//
	//     @inp[0] the list of description objects to modify
	//     @inp[1] the list of RFC6902 compliant JSON-Patches
	//     @out[0] the list of operation states related to the modified description object
	//
	Update([]*Object, [][]*Patch) ([]objectstate.String, error)
}
