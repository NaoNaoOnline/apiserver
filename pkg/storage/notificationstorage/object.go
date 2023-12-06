package notificationstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/xh3b4sd/tracer"
)

type Object struct {
	// Crtd is the time at which the notification got created.
	Crtd time.Time `json:"crtd"`
	// Evnt is the ID of the event triggering this notification.
	Evnt objectid.ID `json:"evnt"`
	// Kind is the notification's object type defining the sub-resource triggering
	// this notification. Kind set to "host" means that an event with a particular
	// host label got created, which a user wanted to get notified about.
	//
	//     cate for events matching the given category IDs
	//     host for events matching the given host IDs
	//     user for events created by the given user IDs
	//
	Kind string `json:"kind"`
	// List is the ID of the user's list this notification belongs to.
	List objectid.ID `json:"list"`
	// Noti is the ID of the notification being created.
	Noti objectid.ID `json:"noti"`
	// Obct is the ID of the sub-resource specified by kind. Kind set to "host"
	// means that obct is the host label ID.
	Obct objectid.ID `json:"obct"`
	// User is the ID of the user this notification belongs to.
	User objectid.ID `json:"user"`
}

func (o *Object) Verify() error {
	{
		if o.Evnt == "" {
			return tracer.Mask(notificationEvntEmptyError)
		}
	}

	{
		if o.Kind != "cate" && o.Kind != "host" && o.Kind != "user" {
			return tracer.Maskf(notificationKindInvalidError, o.Kind)
		}
	}

	{
		if o.List == "" {
			return tracer.Mask(notificationListEmptyError)
		}
	}

	{
		if o.Obct == "" {
			return tracer.Mask(notificationObctEmptyError)
		}
	}

	{
		if o.User == "" {
			return tracer.Mask(runtime.UserIDEmptyError)
		}
	}

	return nil
}
