package reactionstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
)

type Object struct {
	// Crtd is the time at which the reaction got created.
	Crtd time.Time `json:"crtd"`
	// Html is the HTML of this reaction icon, e.g. some svg code.
	Html string `json:"html"`
	// Name is the reaction name.
	Name string `json:"name"`
	// Rctn is the ID of the reaction being created.
	Rctn objectid.String `json:"rctn"`
	// User is the user ID creating this reaction.
	User objectid.String `json:"user"`
}

type Interface interface {
	// Exists verifies whether the given reaction ID exists in the curated
	// whitelist.
	Exists(objectid.String) bool
	// Search returns the for now static list of curated reaction icons available
	// within the platform.
	Search() ([]*Object, error)
}
