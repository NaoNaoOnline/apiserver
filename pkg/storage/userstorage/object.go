package userstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
)

type Object struct {
	// Crtd is the time at which the user got created.
	Crtd time.Time `json:"crtd"`
	// Imag is the URL pointing to the user's profile picture.
	Imag string `json:"imag"`
	// Name is the user name.
	Name string `json:"name"`
	// Subj is the list of external subject claims mapped to the user being
	// created.
	Subj []string `json:"subj"`
	// User is the internal ID of the user being created.
	User objectid.String `json:"user"`
}
