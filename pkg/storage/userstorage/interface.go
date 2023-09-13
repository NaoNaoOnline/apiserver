package userstorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
)

type Interface interface {
	// Create persists a new user object given the provided subject claim, if none
	// does exist already. Create is therefore idempotent and yields the same
	// persisted user object given the same provided subject claim.
	//
	//     @inp[0] the user object providing user specific information
	//     @out[0] the user object mapped to the given subject claim
	//
	Create(*Object) (*Object, error)

	// SearchName returns the user objects matching the given user names.
	//
	//     @inp[0] the user names to search for
	//     @out[0] the list of user objects matching the given user names
	//
	SearchName([]string) ([]*Object, error)

	// SearchSubj returns the user object mapped to the given subject claim, it it
	// exists. SearchSubj will return an error if there is no user mapping already
	// persisted between the external subject claim and the internal user ID.
	//
	//     @inp[0] external subject claim mapped to some internal user ID
	//     @out[0] the user object mapped to the given subject claim
	//
	SearchSubj(string) (*Object, error)

	// SearchUser returns the user objects matching the given user IDs.
	//
	//     @inp[0] the user IDs to search for
	//     @out[0] the list of user objects matching the given user IDs
	//
	SearchUser([]objectid.String) ([]*Object, error)
}
