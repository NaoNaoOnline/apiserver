package notificationstorage

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

type Interface interface {
	// Create TODO
	//
	//     @inp[0] TODO
	//     @inp[1] TODO
	//
	CreateNoti(uid []objectid.ID, obj *Object) error

	// Search TODO
	//
	//     @inp[0] TODO
	//     @inp[1] TODO
	//     @out[0] TODO
	//
	SearchNoti(uid objectid.ID, pag [2]int) ([]*Object, error)

	// Search TODO
	//
	//     @inp[0] TODO
	//     @inp[1] TODO
	//     @inp[2] TODO
	//     @out[0] TODO
	//
	SearchUser(string, objectid.ID, [2]int) ([]objectid.ID, error)
}
