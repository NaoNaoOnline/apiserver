package descriptionstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/scoreid"
)

type Object struct {
	// Crtd is the time at which the description got created.
	Crtd time.Time `json:"crtd"`
	// Desc is the ID of the description being created.
	Desc scoreid.String `json:"desc"`
	// Evnt is the event ID this description is mapped to.
	Evnt scoreid.String `json:"evnt"`
	// Text is the description explaining what an event is about.
	Text string `json:"text"`
	// User is the user ID creating this description.
	User scoreid.String `json:"user"`
	// Vote is the aggregated quality measurement for this description based on
	// user likes and dislikes.
	Vote int `json:"vote"`
}

type Interface interface {
	// Create persists a new description object mapped to the referenced event.
	//
	//     @inp[0] the description object providing description specific information
	//     @out[0] the description object mapped to its associated event object
	//
	Create(*Object) (*Object, error)
	Search() ([]*Object, error)
}
