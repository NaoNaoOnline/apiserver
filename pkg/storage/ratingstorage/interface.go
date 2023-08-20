package ratingstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/scoreid"
)

type Object struct {
	// Crtd is the time at which the rating got created.
	Crtd time.Time `json:"crtd"`
	// Html is the HTML of this rating icon, e.g. some svg code.
	Html string `json:"html"`
	// Name is the rating name.
	Name string `json:"name"`
	// Rtng is the ID of the rating being created.
	Rtng scoreid.String `json:"rtng"`
	// User is the user ID creating this rating.
	User scoreid.String `json:"user"`
}

type Interface interface {
	// Search returns the for now static list of curated rating icons available
	// within the platform.
	Search() ([]*Object, error)
}
