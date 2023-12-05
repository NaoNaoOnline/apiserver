package notificationstorage

import (
	"encoding/json"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) SearchNoti(uid objectid.ID, pag [2]int) ([]*Object, error) {
	var err error

	// val will result in a list of all notification objects recorded to notifiy
	// the given user ID, if any.
	var jsn []string
	{
		jsn, err = r.red.Sorted().Search().Order(notObj(uid), pag[0], pag[1])
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var out []*Object
	for i := range jsn {
		var obj *Object
		{
			obj = &Object{}
		}

		if jsn[i] != "" {
			err = json.Unmarshal([]byte(jsn[i]), obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		{
			out = append(out, obj)
		}
	}

	return out, nil
}

func (r *Redis) SearchUser(kin string, oid objectid.ID, pag [2]int) ([]objectid.ID, error) {
	var err error

	// val will result in a list of all user IDs opted-in to receive notifications
	// for the given resource kind/ID combination, if any.
	var val []string
	{
		val, err = r.red.Sorted().Search().Order(notKin(kin, oid), pag[0], pag[1])
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// There might not be any values, and so we do not proceed, but instead
	// return nothing.
	if len(val) == 0 {
		return nil, nil
	}

	return objectid.IDs(val), nil
}