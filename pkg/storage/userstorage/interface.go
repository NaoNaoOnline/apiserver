package userstorage

import "time"

type Object struct {
	// Crtd is the time at which the user got created.
	Crtd time.Time `json:"crtd"`
	// Imag is the URL pointing to the user's profile picture.
	Imag string `json:"imag"`
	// Name is the user name.
	Name string `json:"name"`
	// User is the internal ID of the user being created.
	User string `json:"user"`
}

type Interface interface {
	// Create persists a new user object given the provided subject claim, if none
	// does exist already. Create is therefore idempotent and yields the same
	// persisted user object given the same provided subject claim.
	//
	//     @inp[0] external subject claim, obtained during the authentication process
	//     @inp[1] profile picture URL, given by the requestee
	//     @inp[2] username, given by the requestee
	//     @out[0] the user object mapped to the given subject claim
	//
	Create(string, string, string) (*Object, error)
	// Search returns the user object mapped to the given subject claim, it it
	// exists. Search will return an error if there is no user mapping already
	// persisted between the external subject claim and the internal user ID.
	//
	//     @inp[0] external subject claim, if given user id must be empty
	//     @inp[1] internal user id, if given subject claim must be empty
	//     @out[0] the user object mapped to the given subject claim or user id
	//
	Search(string, string) (*Object, error)
}
