package rulestorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
)

type Interface interface {
	// CreateRule persists new rule objects.
	//
	//     @inp[0] the rule objects providing rule specific information
	//     @out[0] the rule objects mapped to their internal rule IDs
	//
	CreateRule([]*Object) ([]*Object, error)

	// CreateWrkr emits the respective worker tasks that will be processed in the
	// background for the given rule objects that have just been created.
	//
	//     @inp[0] the rule objects that have been created
	//     @out[0] the list of operation states related to the initialized rule objects
	//
	CreateWrkr([]*Object) ([]objectstate.String, error)

	// DeleteRule purges the given rule objects.
	//
	//     @inp[0] the rule objects to delete
	//     @out[0] the list of operation states related to the purged rule objects
	//
	DeleteRule([]*Object) ([]objectstate.String, error)

	// DeleteWrkr initializes the asynchronous deletion process for the given
	// rule objects and all of its associated data structures by setting
	// Object.Dltd and creating the respective worker tasks that will be processed
	// in the background.
	//
	//     @inp[0] the rule objects to delete
	//     @out[0] the list of operation states related to the purged rule objects
	//
	DeleteWrkr([]*Object) ([]objectstate.String, error)

	// SearchList returns the rule objects belonging to the given list IDs. All
	// rule objects can be fetched using pagination range [0 -1]. The first rule
	// object can be fetched using pagination range [0 0].
	//
	//     @inp[0] the list IDs to search rules for
	//     @inp[1] the pagination range defining lower and upper inclusive boundaries
	//     @out[0] the list of rule objects belonging to the given list IDs
	//
	SearchList([]objectid.ID, [2]int) ([]*Object, error)

	// SearchRule returns the rule objects matching the given rule IDs.
	//
	//     @inp[0] the rule IDs to search for
	//     @out[0] the list of rule objects matching the given rule IDs
	//
	SearchRule([]objectid.ID) ([]*Object, error)

	// Update modifies the existing rule objects.
	//
	//     @inp[0] the list of rule objects to modify
	//     @out[0] the list of operation states related to the modified rule objects
	//
	Update([]*Object) ([]objectstate.String, error)
}
