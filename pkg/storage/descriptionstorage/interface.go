package descriptionstorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
)

type Interface interface {
	// Create persists new description objects mapped to the referenced events.
	//
	//     @inp[0] the description objects mapped to the referenced events
	//     @out[0] the description objects mapped to their internal description IDs
	//
	Create([]*Object) ([]*Object, error)

	// DeleteDesc purges the given description objects. Note that DeleteDesc does
	// not purge associated data structures.
	//
	//     @inp[0] the description objects to delete
	//     @out[0] the list of operation states related to the purged description objects
	//
	DeleteDesc([]*Object) ([]objectstate.String, error)

	// DeleteLike purges internal data structures related to description likes,
	// given the many user IDs that have reacted to the provided description ID.
	//
	//     @inp[0] the description ID to delete
	//     @inp[1] the many user IDs that have reacted to the provided description ID
	//     @out[0] the list of operation states related to the purged data structures
	//
	DeleteLike(des objectid.ID, use []objectid.ID) ([]objectstate.String, error)

	// DeleteWrkr initializes the asynchronous deletion process for the given
	// description objects and all of its associated data structures by setting
	// Object.Dltd and creating the respective worker tasks that will be processed
	// in the background.
	//
	//     @inp[0] the description objects to delete
	//     @out[0] the list of operation states related to the purged description objects
	//
	DeleteWrkr([]*Object) ([]objectstate.String, error)

	// SearchDesc returns the description objects matching the given description
	// IDs.
	//
	//     @inp[0] the calling user
	//     @inp[1] the description IDs to search for
	//     @out[0] the list of description objects matching the given description IDs
	//
	SearchDesc(objectid.ID, []objectid.ID) ([]*Object, error)

	// SearchEvnt returns the description objects belonging to the given event
	// IDs.
	//
	//     @inp[0] the calling user
	//     @inp[1] the event IDs to search descriptions for
	//     @out[0] the list of description objects belonging to the given event IDs
	//
	SearchEvnt(objectid.ID, []objectid.ID) ([]*Object, error)

	// SearchLike returns the user IDs that reacted to the given description ID in
	// the form of a like.
	//
	//     @inp[0] the description ID to search likes for
	//     @out[0] the list of user IDs that reacted to the given description ID
	//
	SearchLike(objectid.ID) ([]objectid.ID, error)

	// UpdateLike modifies the existing description objects by tracking the
	// addition or removal of a like for the given user.
	//
	//     @inp[0] the user liking or unliking the description
	//     @inp[1] the list of description objects to modify
	//     @inp[2] the bool expressing whether to increment or decrement the like count
	//     @out[0] the list of operation states related to the modified description objects
	//
	UpdateLike(objectid.ID, []*Object, []bool) ([]objectstate.String, error)

	// UpdatePtch modifies the existing description objects by applying the given
	// RFC6902 JSON-Patches to the underlying JSON documents. The list items are
	// used according to their respective indices, e.g. the second patch is
	// applied to the second object.
	//
	//     @inp[0] the list of description objects to modify
	//     @inp[1] the list of RFC6902 compliant JSON-Patches
	//     @out[0] the list of operation states related to the modified description objects
	//
	UpdatePtch([]*Object, [][]*Patch) ([]objectstate.String, error)
}
