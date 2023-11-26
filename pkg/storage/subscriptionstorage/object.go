package subscriptionstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/tracer"
)

type Object struct {
	// ChID is the chain ID, the unique identifier representing the blockchain
	// network on which this subscription is located.
	ChID int64 `json:"chid"`
	// Crtd is the time at which the subscription got created.
	Crtd time.Time `json:"crtd"`
	// Crtr is the wallet address of a content creator designated for the purpose
	// of accounting. These are the addresses getting paid peer-to-peer by users
	// subscribing for accessing premium features.
	Crtr []string `json:"crtr"`
	// Sbsc is the wallet address of the user getting access to premium features
	// upon asynchronous subscription verification.
	Sbsc string `json:"sbsc"`
	// Subs is the ID of the subscription being created.
	Subs objectid.ID `json:"evnt"`
	// Unix is the timestamp of the subscription period. This timestamp must be
	// represented in unix seconds, that is in UTC, pointing to the start of any
	// given month. For instance, 1698793200 is Wed Nov 01 2023 00:00:00 UTC,
	// which would subscribe for the whole month of November 2023.
	Unix time.Time `json:"unix"`
	// User is the user ID creating this subscription.
	User objectid.ID `json:"user"`
}

// TODO verify
func (r *Object) Verify() error {
	{
		if r.ChID == 0 {
			return tracer.Mask(subscriptionChIDEmptyError)
		}
	}

	return nil
}
