package eventstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/scoreid"
)

type Object struct {
	// Cate is the list of label IDs under which the event is categorized.
	Cate []scoreid.String `json:"cate"`
	// Crtd is the time at which the event got created.
	Crtd time.Time `json:"crtd"`
	// Dura is the estimated duration of the event.
	Dura time.Duration `json:"dura"`
	// Evnt is the ID of the event being created.
	Evnt scoreid.String `json:"evnt"`
	// Host is the internal host ID expected to host the event.
	Host scoreid.String `json:"host"`
	// Link is the online location at which the event is expected to take place.
	// For IRL events this may just be some informational website.
	Link string `json:"link"`
	// Time is the date time at which the event is expected to start.
	Time time.Time `json:"time"`
	// User is the user ID creating this event.
	User scoreid.String `json:"user"`
}

type Interface interface {
	// Create persists a new event object.
	//
	//     @inp[0] the event object providing event specific information
	//     @out[0] the event object persisted internally
	//
	Create(*Object) (*Object, error)
	// Search returns the event objects grouped under all the given labels.
	//
	//     @inp[0] the host label to include in the search query, if any
	//     @inp[1] the category labels to include in the search query, if any
	//     @out[0] the list of event objects associated to all the given labels
	//
	Search(scoreid.String, ...scoreid.String) ([]*Object, error)
}
