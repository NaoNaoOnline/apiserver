package subscriptionstorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
)

type Interface interface {
	// CreateSubs persists new subscription objects.
	//
	//     @inp[0] the list of subscription objects providing subscription specific information
	//     @out[0] the list of subscription objects persisted internally
	//
	CreateSubs([]*Object) ([]*Object, error)

	// SearchCrtr returns a list of wallet objects representing a set of
	// legitimate content creators that the given user IDs have consumed content
	// from, in the form of event link clicks.
	//
	//     @inp[0] the list of user IDs having clicked on event links
	//     @out[0] the list of wallet objects representing the relevant content creators
	//
	SearchCrtr([]objectid.ID) ([]*walletstorage.Object, error)

	// SearchCurr returns the subscription object for the given receiver for the
	// current month, if any. Note that "the current month" here means the valid
	// latest subscription that is active right now. This might mean the very
	// first subscription valid until the end of next month. The same is true for
	// subscription renewals. Due to several assumptions made here, SearchCurr
	// should only be used to validate subscriptions that are being processed
	// right now and thus have not been active yet.
	//
	//     @inp[0] the user ID of the receiver to search for
	//     @out[0] the subscription object for the current month, or nil
	//
	SearchCurr(objectid.ID) (*Object, error)

	// SearchPayr returns the subscription objects paid for by the given user IDs.
	// All subscriptions can be fetched using pagination range [0 -1]. The latest
	// subscription can be fetched using pagination range [-1 -1].
	//
	//     @inp[0] the user IDs that paid for the subscriptions
	//     @inp[1] the pagination range defining lower and upper inclusive boundaries
	//     @out[0] the list of subscription objects paid by the given user IDs
	//
	SearchPayr([]objectid.ID, [2]int) ([]*Object, error)

	// SearchRcvr returns the subscription objects received by the given user IDs.
	// All subscriptions can be fetched using pagination range [0 -1]. The latest
	// subscription can be fetched using pagination range [-1 -1].
	//
	//     @inp[0] the user IDs that received the subscriptions
	//     @inp[1] the pagination range defining lower and upper inclusive boundaries
	//     @out[0] the list of subscription objects received by the given user IDs
	//
	SearchRcvr([]objectid.ID, [2]int) ([]*Object, error)

	// SearchSubs returns the subscription objects matching the given subscription
	// IDs.
	//
	//     @inp[0] the subscription IDs to search for
	//     @out[0] the list of subscription objects matching the given subscription IDs
	//
	SearchSubs([]objectid.ID) ([]*Object, error)

	// UpdateObct modifies the existing subscription objects.
	//
	//     @inp[0] the list of subscription objects to modify
	//     @out[0] the list of operation states related to the modified subscription objects
	//
	UpdateObct([]*Object) ([]objectstate.String, error)

	// VerifyAddr expresses whether the given wallet addresses are owned by what
	// is being considered legitimate content creators.
	//
	//     @inp[0] the list of wallet addresses to verify
	//     @out[0] the list of validity states related to the verified wallet addresses
	//
	VerifyAddr([]string) ([]bool, error)

	// VerifyUser expresses whether the given user IDs represent what is being
	// considered legitimate content creators.
	//
	//     @inp[0] the list of user IDs to verify
	//     @out[0] the list of validity states related to the verified user IDs
	//
	VerifyUser([]objectid.ID) ([]bool, error)
}
