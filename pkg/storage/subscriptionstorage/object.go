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
}

func (o *Object) VerifyObct() error {
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
		if o.Unix.IsZero() {
			return tracer.Mask(subscriptionUnixEmptyError)
		}
	}

	{
		if o.User == "" {
			return tracer.Mask(runtime.UserIDEmptyError)
		}
	}

	return nil
}

func (o *Object) VerifyUnix(vld func(time.Time) error) error {
	{
		err := vld(o.Unix)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

// VerifyOnce is used to validate the subscription timestamp for subscriptions
// that are created the first time, or without prior active subscription. New
// subscriptions can only be created for the current month.
func VerifyOnce(now time.Time) func(time.Time) error {
	return func(uni time.Time) error {
		if !uni.Equal(timMon(now)) {
			return tracer.Mask(subscriptionUnixInvalidError)
		}

		return nil
	}
}

// VerifyRenw is used to validate the subscription timestamp for subscriptions
// that are effectively renewals of already active subscriptions. While
// subscriptions can only be created for the current month, renewals can be
// created up to 7 days before the new subscription period starts.
func VerifyRenw(now time.Time) func(time.Time) error {
	return func(uni time.Time) error {
		var sta time.Time
		var end time.Time
		{
			sta = timMon(now).AddDate(0, 1, -7)
			end = timMon(now).AddDate(0, 1, 0)
		}

		if now.Before(sta) || now.After(end) {
			return tracer.Mask(subscriptionUnixRenewalError)
		}

		if !uni.Equal(end) {
			return tracer.Mask(subscriptionUnixInvalidError)
		}

		return nil
	}
}

func timMon(tim time.Time) time.Time {
	return time.Date(tim.Year(), tim.Month(), 1, 0, 0, 0, 0, time.UTC)
}
