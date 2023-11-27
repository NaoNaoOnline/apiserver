package subscriptionstorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
)

type Interface interface {
	// CreateSubs persists new subscription objects.
	//
	//     @inp[0] the list of subscription objects providing subscription specific information
	//     @out[0] the list of subscription objects persisted internally
	//
	CreateSubs([]*Object) ([]*Object, error)

	// CreateWrkr emits the respective worker tasks that will be processed in the
	// background for the given subscription objects that have just been created.
	// Workers can e.g. verify subscriptions asynchronously between onchain and
	// offchain state.
	//
	//     @inp[0] the subscription objects that have been created
	//     @out[0] the list of operation states related to the initialized subscription objects
	//
	CreateWrkr(inp []*Object) ([]objectstate.String, error)

	// SearchLtst returns the most recent subscription object for the given user.
	//
	//     @inp[0] the user ID to search for
	//     @out[0] the most recent subscription object for the given user
	//
	SearchLtst(objectid.ID) (*Object, error)

	// SearchSubs returns the subscription objects matching the given subscription
	// IDs.
	//
	//     @inp[0] the subscription IDs to search for
	//     @out[0] the list of subscription objects matching the given subscription IDs
	//
	SearchSubs([]objectid.ID) ([]*Object, error)

	// SearchUser returns the subscription objects created by the given user IDs.
	//
	//     @inp[0] the user IDs that created the subscriptions
	//     @out[0] the list of subscription objects created by the given user IDs
	//
	SearchUser([]objectid.ID) ([]*Object, error)
}
