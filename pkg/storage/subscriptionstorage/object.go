package subscriptionstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/format/hexformat"
	"github.com/NaoNaoOnline/apiserver/pkg/generic"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectstate"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
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
	// Fail is the description explaining why a subscription could not be verified
	// successfully. Most subscriptions should not be accompanied by a failure
	// message.
	Fail string `json:"fail,omitempty"`
	// Recv is the wallet address of the user getting access to premium features
	// upon asynchronous subscription verification.
	Recv string `json:"recv"`
	// Stts is the resource status expressing whether this subscription is active.
	// An active subscription is verified by comparing its offchain and onchain
	// state. Subscriptions found to be invalid will not be marked as active, but
	// will instead be accompanied by
	//
	//	created for a newly registered subscriptions
	//	failure for successfully processed subscriptions
	//	success for processed subscriptions found to be invalid
	//
	Stts objectstate.String `json:"stts"`
	// Subs is the ID of the subscription being created.
	Subs objectid.ID `json:"evnt"`
	// Unix is the timestamp of the subscription period. This timestamp must be
	// represented in unix seconds, that is in UTC, pointing to the start of any
	// given month. For instance, 1698793200 is Wed Nov 01 2023 00:00:00 UTC,
	// which would subscribe for the whole month of November 2023.
	Unix time.Time `json:"unix"`
	// User is the user ID creating this subscription.
	User objectid.ID `json:"user"`

	//

	time Timer `json:"-"`
}

func (o *Object) Verify() error {
	{
		if o.ChID == 0 {
			return tracer.Mask(subscriptionChIDEmptyError)
		}
	}

	{
		if len(o.Crtr) == 0 {
			return tracer.Mask(subscriptionCrtrEmptyError)
		}
		if len(o.Crtr) > 3 {
			return tracer.Mask(subscriptionCrtrLimitError)
		}
		if generic.Dup(o.Crtr) {
			return tracer.Mask(subscriptionCrtrDuplicateError)
		}
		for _, x := range o.Crtr {
			if x == "" {
				return tracer.Mask(subscriptionCrtrEmptyError)
			}
			if len(x) != 42 {
				return tracer.Maskf(subscriptionCrtrLengthError, "%d", len(x))
			}
			if !hexformat.Verify(x) {
				return tracer.Mask(subscriptionCrtrFormatError)
			}
		}
	}

	{
		if o.Recv == "" {
			return tracer.Mask(subscriptionRecvEmptyError)
		}
		if len(o.Recv) != 42 {
			return tracer.Maskf(subscriptionRecvLengthError, "%d", len(o.Recv))
		}
		if !hexformat.Verify(o.Recv) {
			return tracer.Mask(subscriptionRecvFormatError)
		}
	}

	{
		if o.time == nil {
			o.time = &timer{}
		}
		if o.Unix.IsZero() {
			return tracer.Mask(subscriptionUnixEmptyError)
		}
		if !o.Unix.Equal(timMon(o.time.Now())) {
			return tracer.Mask(subscriptionUnixInvalidError)
		}
	}

	{
		if o.User == "" {
			return tracer.Mask(runtime.UserIDEmptyError)
		}
	}

	return nil
}

func timMon(tim time.Time) time.Time {
	return time.Date(tim.Year(), tim.Month(), 1, 0, 0, 0, 0, time.UTC)
}
