package labelstorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
)

type Interface interface {
	// Create persists new label objects, if none exists already with the given
	// name.
	//
	//     @inp[0] the label objects providing label specific information
	//     @out[0] the label objects mapped to their internal label IDs
	//
	Create([]*Object) ([]*Object, error)

	// SearchBltn returns the static list of curated event labels natively
	// supported by the system. This method should only be used to bootstrap the
	// initial system state, never to serve RPC requests. Note that the names of
	// builtin labels need to match the eventstorage.Object.Pltfrm implementation
	// of parsing a platforms hostname.
	//
	//     @out[0] static list of curated event labels natively supported by the system
	//
	SearchBltn() []*Object

	// SearchKind returns the label objects matching the given label kinds, e.g.
	// bltn, cate or host.
	//
	//     @inp[0] the label kinds under which label objects are grouped together
	//     @out[0] the list of label objects matching the given label kinds
	//
	SearchKind([]string) ([]*Object, error)

	// SearchLabl returns the label objects matching the given label IDs.
	//
	//     @inp[0] the label IDs to search for
	//     @out[0] the list of label objects matching the given label IDs
	//
	SearchLabl([]objectid.ID) ([]*Object, error)

	// SearchName returns the label objects matching the given label names.
	//
	//     @inp[0] the label kinds matching their respective label names
	//     @inp[1] the label names to search for
	//     @out[0] the list of label objects matching the given label names
	//
	SearchName([]string, []string) ([]*Object, error)

	// SearchUser returns the label objects created by the given user.
	//
	//     @inp[0] the user ID used to search labels
	//     @out[0] the list of label objects for the given user
	//
	SearchUser(objectid.ID) ([]*Object, error)

	// UpdatePtch modifies the existing label objects by applying the given
	// RFC6902 JSON-Patches to the underlying JSON documents. The list items are
	// used according to their respective indices, e.g. the second patch is
	// applied to the second object.
	//
	//     @inp[0] the list of label objects to modify
	//     @inp[1] the list of RFC6902 compliant JSON-Patches
	//     @out[0] the list of operation states related to the modified label objects
	//
	UpdatePtch([]*Object, PatchSlicer) ([]objectstate.String, error)
}
