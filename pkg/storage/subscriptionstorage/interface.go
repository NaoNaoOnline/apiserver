package subscriptionstorage

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

type Interface interface {
	// Create persists new subscription objects.
	//
	//     @inp[0] the list of subscription objects providing subscription specific information
	//     @out[0] the list of subscription objects persisted internally
	//
	Create([]*Object) ([]*Object, error)

	// SearchUser returns the subscription objects created by the given user IDs.
	//
	//     @inp[0] the user IDs that created the subscriptions
	//     @out[0] the list of subscription objects created by the given user IDs
	//
	SearchUser(use []objectid.ID) ([]*Object, error)
}
