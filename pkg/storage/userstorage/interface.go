package userstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/scoreid"
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
	User scoreid.String `json:"user"`
}

type Interface interface {
	// Create persists a new user object given the provided subject claim, if none
	// does exist already. Create is therefore idempotent and yields the same
	// persisted user object given the same provided subject claim.
	//
	//     @inp[0] the user object providing user specific information
	//     @out[0] the user object mapped to the given subject claim
	//
	Create(*Object) (*Object, error)
	// Search returns the user object mapped to the given subject claim, it it
	// exists. Search will return an error if there is no user mapping already
	// persisted between the external subject claim and the internal user ID.
	//
	//     @inp[0] external subject claim, if given user ID must be empty
	//     @inp[1] internal user ID, if given subject claim must be empty
	//     @out[0] the user object mapped to the given subject claim or user ID
	//
	Search(string, scoreid.String) (*Object, error)
}
