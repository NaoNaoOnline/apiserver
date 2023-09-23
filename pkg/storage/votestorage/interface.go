package votestorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/objectstate"
)

type Interface interface {
	// Create persists new vote objects to keep track of voting relationships
	// between internal resources.
	//
	//     @inp[0] the vote objects providing vote specific information
	//     @out[0] the vote objects mapped to their internal vote IDs
	//
	Create([]*Object) ([]*Object, error)

	// Delete purges the given vote objects.
	//
	//     @inp[0] the vote objects to delete
	//     @out[0] the list of operation states related to the purged vote objects
	//
	Delete([]*Object) ([]objectstate.String, error)

	// SearchDesc returns all vote objects associated to the given description IDs.
	//
	//     @inp[0] the description IDs any vote object might be associated with
	//     @out[0] the list of all vote objects associated with the given description IDs
	//
	SearchDesc([]objectid.String) ([]*Object, error)

	// SearchVote returns the vote objects matching the given vote IDs.
	//
	//     @inp[0] the vote IDs to search for
	//     @out[0] the list of vote objects matching the given vote IDs
	//
	SearchVote([]objectid.String) ([]*Object, error)
}
