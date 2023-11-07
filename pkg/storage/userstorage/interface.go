package userstorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
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

	// SearchFake returns all user objects. This is used to create fake test data
	// during development. DO NOT USE IN PRODUCTION.
	//
	//     @out[0] the list of all user objects in redis
	//
	SearchFake() ([]*Object, error)

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
	SearchUser([]objectid.ID) ([]*Object, error)

	// Update modifies the existing user objects by applying the given RFC6902
	// JSON-Patches to the underlying JSON documents. The list items are used
	// according to their respective indices, e.g. the second patch is applied to
	// the second object.
	//
	//     @inp[0] the list of user objects to modify
	//     @inp[1] the list of RFC6902 compliant JSON-Patches
	//     @out[0] the list of operation states related to the modified user objects
	//
	Update([]*Object, [][]*Patch) ([]objectstate.String, error)
}
