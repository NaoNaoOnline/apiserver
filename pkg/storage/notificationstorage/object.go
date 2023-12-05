package notificationstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
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
	// Noti is the ID of the notification being created.
	Noti objectid.ID `json:"noti"`
	// Obct is the ID of the sub-resource specified by kind. Kind set to "host"
	// means that obct is the host label ID.
	Obct objectid.ID `json:"obct"`
}

func (o *Object) Verify() error {
	// TODO complete object verify

	{
		if o.Evnt == "" {
			return tracer.Mask(notificationEvntEmptyError)
		}
	}

	// Note that Object.User is not validated here like for the other resources,
	// because notifications are private objects only relevant to a specific user.
	// Tracking user IDs in notifications would be redundant and not serve any
	// purpose.

	return nil
}
