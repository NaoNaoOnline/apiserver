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

	// SearchPayr returns the subscription objects created by the given user IDs.
	// That is, the users who paid for the subscriptions being searched. All
	// subscriptions can be fetched using pagination range [0 -1]. The latest
	// subscription can be fetched using pagination range [-1 -1].
	//
	//     @inp[0] the user IDs that paid for the subscriptions
	//     @inp[1] the pagination range defining lower and upper inclusive boundaries
	//     @out[0] the list of subscription objects paid by the given user IDs
	//
	SearchPayr([]objectid.ID, [2]int) ([]*Object, error)

	// SearchRecv returns the subscription objects received by the given user IDs.
	// That is, the users who received the subscriptions being searched. All
	// subscriptions can be fetched using pagination range [0 -1]. The latest
	// subscription can be fetched using pagination range [-1 -1].
	//
	//     @inp[0] the user IDs that received the subscriptions
	//     @inp[1] the pagination range defining lower and upper inclusive boundaries
	//     @out[0] the list of subscription objects received by the given user IDs
	//
	SearchRecv([]objectid.ID, [2]int) ([]*Object, error)

	// SearchSubs returns the subscription objects matching the given subscription
	// IDs.
	//
	//     @inp[0] the subscription IDs to search for
	//     @out[0] the list of subscription objects matching the given subscription IDs
	//
	SearchSubs([]objectid.ID) ([]*Object, error)
}
