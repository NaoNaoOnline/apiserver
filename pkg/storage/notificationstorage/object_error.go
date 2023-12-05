package notificationstorage

import (
	"github.com/xh3b4sd/tracer"
)

var notificationEvntEmptyError = &tracer.Error{
	Kind: "notificationEvntEmptyError",
	Desc: "The request expects the notification event not to be empty. The notification event was found to be empty. Therefore the request failed.",
}
