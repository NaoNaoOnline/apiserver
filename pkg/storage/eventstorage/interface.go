package eventstorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/rulestorage"
)

type Interface interface {
	// Create persists new event objects.
	//
	//     @inp[0] the list of event objects providing event specific information
	//     @out[0] the list of event objects persisted internally
	//
	Create([]*Object) ([]*Object, error)

	// DeleteEvnt purges the given event objects. Note that DeleteEvnt does not
	// purge associated data structures.
	//
	//     @inp[0] the event objects to delete
	//     @out[0] the list of operation states related to the purged event objects
	//
	DeleteEvnt([]*Object) ([]objectstate.String, error)

	// DeleteWrkr initializes the asynchronous deletion process for the given
	// event objects and all of its associated data structures by setting
	// Object.Dltd and creating the respective worker tasks that will be processed
	// in the background.
	//
	//     @inp[0] the event objects to delete
	//     @out[0] the list of operation states related to the purged event objects
	//
	DeleteWrkr([]*Object) ([]objectstate.String, error)

	// SearchEvnt returns the event objects matching the given event IDs.
	//
	//     @inp[0] the event IDs to search for
	//     @out[0] the list of event objects matching the given event IDs
	//
	SearchEvnt([]objectid.ID) ([]*Object, error)

	// SearchHpnd returns the event objects that happened over a week ago. This
	// function is mainly used for cleaning up old events in a background process.
	//
	//     @out[0] the list of event objects that happened over a week ago
	//
	SearchHpnd() ([]*Object, error)

	// SearchLabl returns the event objects grouped under all the given labels.
	//
	//     @inp[0] the category and/or host labels to include in the search query, if any
	//     @out[0] the list of event objects associated to all the given labels
	//
	SearchLabl([]objectid.ID) ([]*Object, error)

	// SearchLtst returns the event objects known to happen right now.
	// Specifically, these are the latest events within a time range of -1 and +1
	// week, relative to time of execution, read "now".
	//
	//     @out[0] the list of event objects known to happen right now
	//
	SearchLtst() ([]*Object, error)

	// SearchLike returns the event objects the given user ID reacted to in the
	// form of description likes.
	//
	//     @inp[0] the user ID that reacted to events
	//     @out[0] the list of event objects the given user ID reacted to
	//
	SearchLike(objectid.ID) ([]*Object, error)

	// SearchRule returns the event objects matching all the criteria specified by
	// the given rule objects.
	//
	//     @inp[0] the rule objects of a certain list
	//     @out[0] the list of event objects matching all the criteria of the given list
	//
	SearchRule([]*rulestorage.Object) ([]*Object, error)

	// SearchUser returns the event objects created by the given user IDs.
	//
	//     @inp[0] the user IDs that created the events
	//     @out[0] the list of event objects created by the given user IDs
	//
	SearchUser([]objectid.ID) ([]*Object, error)

	// UpdateClck modifies the existing event objects by tracking the addition of
	// a link click for the given user.
	//
	//     @inp[0] the user clicking the event link
	//     @inp[1] the list of event objects to modify
	//     @out[0] the list of operation states related to the modified event objects
	//
	UpdateClck([]*Object) ([]objectstate.String, error)
}
