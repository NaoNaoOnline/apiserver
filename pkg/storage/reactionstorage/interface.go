package reactionstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/scoreid"
)

type Object struct {
	// Crtd is the time at which the reaction got created.
	Crtd time.Time `json:"crtd"`
	// Html is the HTML of this reaction icon, e.g. some svg code.
	Html string `json:"html"`
	// Name is the reaction name.
	Name string `json:"name"`
	// Rctn is the ID of the reaction being created.
	Rctn scoreid.String `json:"rctn"`
	// User is the user ID creating this reaction.
	User scoreid.String `json:"user"`
}

type Interface interface {
	// Search returns the for now static list of curated reaction icons available
	// within the platform.
	Search() ([]*Object, error)
}
