package labelstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/scoreid"
)

type Object struct {
	// Crtd is the time at which the label got created.
	Crtd time.Time `json:"crtd"`
	// Desc is the label's description.
	Desc string `json:"desc"`
	// Disc is the label's Discord link.
	Disc string `json:"disc"`
	// Kind is the label type, e.g. host for host labels and cate for category
	// labels.
	Kind string `json:"kind"`
	// Labl is the ID of the label being created.
	Labl scoreid.String `json:"labl"`
	// Name is the label name.
	Name string `json:"name"`
	// Twit is the label's Twitter link.
	Twit string `json:"twit"`
	// User is the user ID creating this label.
	User scoreid.String `json:"user"`
}

type Interface interface {
	// Create persists new label objects, if none exists already with the given
	// name.
	//
	//     @inp[0] the label objects providing label specific information
	//     @out[0] the label objects mapped to its internal label ID
	//
	Create([]*Object) ([]*Object, error)
	// Search returns the label objects of the given kind.
	//
	//     @inp[0] the label kinds under which label objects are grouped together
	//     @out[0] the list of label objects of either kind category or kind host
	//
	Search([]string) ([]*Object, error)
}
