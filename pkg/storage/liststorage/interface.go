package liststorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
)

type Interface interface {
	// Create persists new list objects.
	//
	//     @inp[0] the list objects providing list specific information
	//     @out[0] the list objects mapped to their internal list IDs
	//
	Create([]*Object) ([]*Object, error)

	// DeleteList purges the given list objects. Note that DeleteList does not
	// purge associated data structures.
	//
	//     @inp[0] the list objects to delete
	//     @out[0] the list of operation states related to the purged list objects
	//
	DeleteList([]*Object) ([]objectstate.String, error)

	// DeleteWrkr initializes the asynchronous deletion process for the given list
	// objects and all of its associated data structures by setting Object.Dltd
	// and creating the respective worker tasks that will be processed in the
	// background.
	//
	//     @inp[0] the list objects to delete
	//     @out[0] the list of operation states related to the purged list objects
	//
	DeleteWrkr([]*Object) ([]objectstate.String, error)

	// SearchList returns the list objects matching the given list IDs.
	//
	//     @inp[0] the list IDs to search for
	//     @out[0] the list of list objects matching the given list IDs
	//
	SearchList([]objectid.ID) ([]*Object, error)

	// SearchUser returns the list objects for the given user.
	//
	//     @inp[0] the user ID used to search lists
	//     @out[0] the list of list objects for the given user
	//
	SearchUser(objectid.ID) ([]*Object, error)

	// Update modifies the existing list objects by applying the given RFC6902
	// JSON-Patches to the underlying JSON documents. The list items are used
	// according to their respective indices, e.g. the second patch is applied to
	// the second object.
	//
	//     @inp[0] the list of list objects to modify
	//     @inp[1] the list of RFC6902 compliant JSON-Patches
	//     @out[0] the list of operation states related to the modified list objects
	//
	Update([]*Object, [][]*Patch) ([]objectstate.String, error)
}
