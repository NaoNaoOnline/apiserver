package descriptionstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
)

type Object struct {
	// Crtd is the time at which the description got created.
	Crtd time.Time `json:"crtd"`
	// Desc is the ID of the description being created.
	Desc objectid.String `json:"desc"`
	// Evnt is the event ID this description is mapped to.
	Evnt objectid.String `json:"evnt"`
	// Text is the description explaining what an event is about.
	Text string `json:"text"`
	// User is the user ID creating this description.
	User objectid.String `json:"user"`
}

type Interface interface {
	// Create persists a new description object mapped to the referenced event.
	//
	//     @inp[0] the description object providing description specific information
	//     @out[0] the description object mapped to its associated event object
	//
	Create(*Object) (*Object, error)
	// Search returns the description objects belonging to the given event IDs.
	//
	//     @inp[0] the event IDs to search descriptions for
	//     @out[0] the list of description objects belonging the given event IDs
	//
	Search([]objectid.String) ([]*Object, error)
}
