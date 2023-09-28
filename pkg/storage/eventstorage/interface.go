package eventstorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
)

type Interface interface {
	// Create persists new event objects.
	//
	//     @inp[0] the list of event objects providing event specific information
	//     @out[0] the list of event objects persisted internally
	//
	Create([]*Object) ([]*Object, error)

	// SearchEvnt returns the event objects matching the given event IDs.
	//
	//     @inp[0] the event IDs to search for
	//     @out[0] the list of event objects matching the given event IDs
	//
	SearchEvnt([]objectid.ID) ([]*Object, error)

	// SearchLabl returns the event objects grouped under all the given labels.
	//
	//     @inp[0] the category and/or host labels to include in the search query, if any
	//     @out[0] the list of event objects associated to all the given labels
	//
	SearchLabl([]objectid.ID) ([]*Object, error)

	// SearchLtst returns the event objects known to happen right now.
	//
	//     @out[0] the list of event objects known to happen right now
	//
	SearchLtst() ([]*Object, error)

	// SearchRctn returns the event objects the given user ID reacted to.
	//
	//     @inp[0] the user ID that reacted to events
	//     @out[0] the list of event objects the given user ID reacted to
	//
	SearchRctn(objectid.ID) ([]*Object, error)

	// SearchUser returns the event objects created by the given user IDs.
	//
	//     @inp[0] the user IDs that created the events
	//     @out[0] the list of event objects created by the given user IDs
	//
	SearchUser([]objectid.ID) ([]*Object, error)
}
