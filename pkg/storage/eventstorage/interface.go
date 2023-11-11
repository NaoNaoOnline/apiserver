package eventstorage

import (
	"time"

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

	// DeleteLink purges internal data structures related to event links, given
	// the many user IDs that have visited the provided event ID. This function is
	// mainly used for cleaning up internal event related data structures in a
	// background process.
	//
	//     @inp[0] the event ID to delete
	//     @inp[1] the many user IDs that have visited the provided event ID
	//     @out[0] the list of operation states related to the purged data structures
	//
	DeleteLink(objectid.ID, []objectid.ID) ([]objectstate.String, error)

	// DeleteRule purges internal data structures related to event rules, given
	// the many rule IDs that have referenced the provided event ID. This function
	// is mainly used for cleaning up internal event related data structures in a
	// background process.
	//
	//     @inp[0] the event ID to delete
	//     @inp[1] the many rule IDs that have visited the provided event ID
	//     @out[0] the list of operation states related to the purged data structures
	//
	DeleteRule(objectid.ID, []objectid.ID) ([]objectstate.String, error)

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
	//     @inp[0] the calling user
	//     @inp[1] the event IDs to search for
	//     @out[0] the list of event objects matching the given event IDs
	//
	SearchEvnt(objectid.ID, []objectid.ID) ([]*Object, error)

	// SearchHpnd returns the event objects known to have happened already within
	// the past week.
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

	// SearchLike returns the event objects the given user ID reacted to in the
	// form of description likes.
	//
	//     @inp[0] the user ID that reacted to events
	//     @inp[1] the lower inclusive boundary of the pagination range
	//     @inp[2] the upper inclusive boundary of the pagination range
	//     @out[0] the list of event objects the given user ID reacted to
	//
	SearchLike(objectid.ID, int, int) ([]*Object, error)

	// SearchLink returns the user IDs that visited the given event ID in the form
	// of a link. This function is mainly used for cleaning up internal user
	// related data structures in a background process.
	//
	//     @inp[0] the event ID to search users for
	//     @out[0] the list of user IDs that visited the given event ID
	//
	SearchLink(objectid.ID) ([]objectid.ID, error)

	// SearchList returns the event objects matching all the criteria specified by
	// the given rule objects.
	//
	//     @inp[0] the rule objects of a certain list
	//     @out[0] the list of event objects matching all the criteria of the given list
	//
	SearchList([]*rulestorage.Object) ([]*Object, error)

	// SearchRule returns the rule IDs that explicitely define the given event ID
	// in the form of an include or exclude reference. This function is mainly used
	// for cleaning up internal event related data structures in a background
	// process.
	//
	//     @inp[0] the event ID to search rules for
	//     @out[0] the list of rule IDs that reference the given event ID
	//
	SearchRule(objectid.ID) ([]objectid.ID, error)

	// SearchTime returns the event objects known to happen within the given time
	// range.
	//
	//     @inp[0] the lower inclusive boundary of the time range
	//     @inp[1] the upper inclusive boundary of the time range
	//     @out[0] the list of event objects known to happen right now
	//
	SearchTime(time.Time, time.Time) ([]*Object, error)

	// SearchUpcm returns the event objects known to happen within the next week.
	//
	//     @out[0] the list of event objects known to happen right now
	//
	SearchUpcm() ([]*Object, error)

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
	UpdateClck(objectid.ID, []*Object) ([]objectstate.String, error)
}
