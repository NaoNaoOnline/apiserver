package eventstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
)

type Interface interface {
	// CreateEvnt persists new event objects.
	//
	//     @inp[0] the list of event objects providing event specific information
	//     @out[0] the list of event objects persisted internally
	//
	CreateEvnt([]*Object) ([]*Object, error)

	// CreateWrkr emits the respective worker tasks that will be processed in the
	// background for the given event objects that have just been created. Workers
	// can subscribe to certain event creation tasks and manage e.g. forms of
	// notification.
	//
	//     @inp[0] the event objects that have been created
	//     @out[0] the list of operation states related to the initialized event objects
	//
	CreateWrkr([]*Object) ([]objectstate.String, error)

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

	// DeleteWrkr initializes the asynchronous deletion process for the given
	// event objects and all of its associated data structures by setting
	// Object.Dltd and creating the respective worker tasks that will be processed
	// in the background.
	//
	//     @inp[0] the event objects to delete
	//     @out[0] the list of operation states related to the purged event objects
	//
	DeleteWrkr([]*Object) ([]objectstate.String, error)

	// SearchCrtr returns the user IDs of all users who are recorded to have added
	// events to the platform within a rolling time window. All user IDs can be
	// fetched using pagination range [0 -1]. The first 10 user IDs can be fetched
	// using pagination range [0 9].
	//
	//     @inp[0] the pagination range defining lower and upper inclusive boundaries
	//     @out[0] the list of user IDs who have added events to the platform
	//
	SearchCrtr([2]int) ([]objectid.ID, error)

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

	// SearchLabl returns the event objects grouped under all the given label IDs.
	//
	//     @inp[0] the category and/or host label IDs to search for
	//     @out[0] the list of event objects associated to all the given labels
	//
	SearchLabl([]objectid.ID) ([]*Object, error)

	// SearchLike returns the event objects the given user ID reacted to in the
	// form of description likes. All event objects can be fetched using
	// pagination range [0 -1]. The first 10 event objects can be fetched using
	// pagination range [0 9].
	//
	//     @inp[0] the user ID that reacted to events
	//     @inp[1] the pagination range defining lower and upper inclusive boundaries
	//     @out[0] the list of event objects the given user ID reacted to
	//
	SearchLike(objectid.ID, [2]int) ([]*Object, error)

	// SearchLink returns the user IDs that visited the given event ID in the form
	// of a link click. This function is mainly used for cleaning up internal user
	// related data structures in a background process.
	//
	//     @inp[0] the event ID to search users for
	//     @out[0] the list of user IDs that visited the given event ID
	//
	SearchLink(objectid.ID) ([]objectid.ID, error)

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
	//     @inp[1] the bool expressing whether the given user has a premium subscription
	//     @inp[2] the list of event objects to modify
	//     @out[0] the list of operation states related to the modified event objects
	//
	UpdateClck(objectid.ID, bool, []*Object) ([]objectstate.String, error)
}
