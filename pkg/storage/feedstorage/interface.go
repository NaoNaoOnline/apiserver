package feedstorage

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

type Interface interface {
	// CreateFeed persists the given feed objects CreateFeed is is primarily used
	// in asynchronous background processes using pagination in order to keep the
	// execution time low.
	//
	//     @inp[0] the list of feed objects to persist
	//
	CreateFeed([]*Object) error

	// DeleteFeed purges the given feed objects.
	//
	//     @inp[0] the feed objects to delete
	//
	DeleteFeed([]*Object) error

	// SearchFeed returns the paginated list of feed objects for the given user /
	// list combination. All feed objects can be fetched using pagination range [0
	// -1]. The latest 10 feed objects can be fetched using pagination range [-10
	// -1].
	//
	//     @inp[0] the user ID to search for
	//     @inp[1] the list ID to search for
	//     @inp[2] the pagination range defining lower and upper inclusive boundaries
	//     @out[0] the list of feed objects for the given user ID
	//
	SearchFeed(objectid.ID, objectid.ID, [2]int) ([]*Object, error)

	// SearchUser returns the user IDs opted-in to receive feeds for the given
	// resource kind/ID combination. All user IDs can be fetched using pagination
	// range [0 -1]. The first 10 user IDs can be fetched using pagination range
	// [0 9].
	//
	//     @inp[0] the resource kind to search for, e.g. cate, host or user
	//     @inp[1] the resource ID to search for
	//     @inp[2] the pagination range defining lower and upper inclusive boundaries
	//     @out[0] the user IDs opted-in to receive feeds
	//     @out[1] the list IDs related to the respective user IDs returned
	//
	SearchUser(string, objectid.ID, [2]int) ([]objectid.ID, []objectid.ID, error)
}
