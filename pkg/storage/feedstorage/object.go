package feedstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/xh3b4sd/tracer"
)

type Object struct {
	// Crtd is the time at which the feed object got created.
	Crtd time.Time `json:"crtd"`
	// Evnt is the ID of the event triggering this feed object.
	Evnt objectid.ID `json:"evnt"`
	// Feed is the ID of the feed object being created.
	Feed objectid.ID `json:"noti"`
	// Kind is the feed object's object type defining the sub-resource triggering
	// this feed object. Kind set to "host" means that an event with a particular
	// host label got created, which a user wanted to get notified about.
	//
	//     cate for events matching the given category IDs (dynamic)
	//     evnt for adding or removing events matching the given event IDs (static)
	//     host for events matching the given host IDs (dynamic)
	//     user for events created by the given user IDs (dynamic)
	//
	Kind string `json:"kind"`
	// List is the ID of the user's list this feed object belongs to.
	List objectid.ID `json:"list"`
	// Obct is the ID of the sub-resource specified by kind. Kind set to "host"
	// means that obct is the host label ID.
	Obct objectid.ID `json:"obct"`
	// User is the ID of the user this feed object belongs to.
	User objectid.ID `json:"user"`
}

func (o *Object) Verify() error {
	{
		if o.Evnt == "" {
			return tracer.Mask(feedEvntEmptyError)
		}
	}

	{
		if o.Kind != "cate" && o.Kind != "evnt" && o.Kind != "host" && o.Kind != "user" {
			return tracer.Maskf(feedKindInvalidError, o.Kind)
		}
	}

	{
		if o.List == "" {
			return tracer.Mask(feedListEmptyError)
		}
	}

	{
		if o.Obct == "" {
			return tracer.Mask(feedObctEmptyError)
		}
	}

	{
		if o.User == "" {
			return tracer.Mask(runtime.UserIDEmptyError)
		}
	}

	return nil
}
