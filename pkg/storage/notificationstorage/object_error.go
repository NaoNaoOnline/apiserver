package notificationstorage

import (
	"github.com/xh3b4sd/tracer"
)

var notificationEvntEmptyError = &tracer.Error{
	Kind: "notificationEvntEmptyError",
	Desc: "The request expects the notification event not to be empty. The notification event was found to be empty. Therefore the request failed.",
}

var notificationKindInvalidError = &tracer.Error{
	Kind: "notificationKindInvalidError",
	Desc: "The request expects the notification kind to be one of [cate host user]. The notification kind was not found to be one of [cate host user]. Therefore the request failed.",
}

var notificationListEmptyError = &tracer.Error{
	Kind: "notificationListEmptyError",
	Desc: "The request expects the notification list not to be empty. The notification list was found to be empty. Therefore the request failed.",
}

var notificationObctEmptyError = &tracer.Error{
	Kind: "notificationObctEmptyError",
	Desc: "The request expects the notification object not to be empty. The notification object was found to be empty. Therefore the request failed.",
}
