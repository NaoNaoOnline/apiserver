package notificationstorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) CreateNoti(uid []objectid.ID, obj *Object) error {
	var err error

	for _, x := range uid {
		// TODO verify

		{
			err = r.red.Sorted().Create().Score(notObj(x), musStr(obj), obj.Noti.Float())
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}
