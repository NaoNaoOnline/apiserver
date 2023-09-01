package eventstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
)

type Object struct {
	// Cate is the list of label IDs under which the event is categorized.
	Cate []objectid.String `json:"cate"`
	// Crtd is the time at which the event got created.
	Crtd time.Time `json:"crtd"`
	// Dura is the estimated duration of the event.
	Dura time.Duration `json:"dura"`
	// Evnt is the ID of the event being created.
	Evnt objectid.String `json:"evnt"`
	// Host is the list of label IDs expected to host the event.
	Host []objectid.String `json:"host"`
	// Link is the online location at which the event is expected to take place.
	// For IRL events this may just be some informational website.
	Link string `json:"link"`
	// Time is the date time at which the event is expected to start.
	Time time.Time `json:"time"`
	// User is the user ID creating this event.
	User objectid.String `json:"user"`
}

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
	SearchEvnt([]objectid.String) ([]*Object, error)
	// SearchLabl returns the event objects grouped under all the given labels.
	//
	//     @inp[0] the category and/or host labels to include in the search query, if any
	//     @out[0] the list of event objects associated to all the given labels
	//
	SearchLabl([]objectid.String) ([]*Object, error)
	// SearchTime returns the event objects indexed to happen right now.
	//
	//     @out[0] the list of event objects indexed to happen right now
	//
	SearchTime() ([]*Object, error)
}
