package notificationstorage

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

type Interface interface {
	// CreateNoti persists the given notification objects CreateNoti is is
	// primarily used in asynchronous background processes using pagination in
	// order to keep the execution time low.
	//
	//     @inp[0] the list of notification objects to persist
	//
	CreateNoti([]*Object) error

	// SearchNoti returns the paginated list of notification objects for the given
	// user / list combination. All notification objects can be fetched using
	// pagination range [0 -1]. The latest 10 notification objects can be fetched
	// using pagination range [-10 -1].
	//
	//     @inp[0] the user ID to search for
	//     @inp[1] the list ID to search for
	//     @inp[2] the pagination range defining lower and upper inclusive boundaries
	//     @out[0] the list of notification objects for the given user ID
	//
	SearchNoti(objectid.ID, objectid.ID, [2]int) ([]*Object, error)

	// SearchUser returns the user IDs opted-in to receive notifications for the
	// given resource kind/ID combination. All user IDs can be fetched using
	// pagination range [0 -1]. The first 10 user IDs can be fetched using
	// pagination range [0 9].
	//
	//     @inp[0] the resource kind to search for, e.g. cate, host or user
	//     @inp[1] the resource ID to search for
	//     @inp[2] the pagination range defining lower and upper inclusive boundaries
	//     @out[0] the user IDs opted-in to receive notifications
	//     @out[1] the list IDs related to the respective user IDs returned
	//
	SearchUser(string, objectid.ID, [2]int) ([]objectid.ID, []objectid.ID, error)
}
