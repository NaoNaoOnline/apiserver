package walletstorage

import (
	"encoding/json"
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/redigo/pkg/simple"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) SearchKind(use objectid.ID, kin []string) ([]*Object, error) {
	var err error

	var out []*Object
	for _, x := range kin {
		if x != "eth" {
			return nil, tracer.Mask(walletKindInvalidError)
		}

		// val will result in a list of all wallet IDs for the given user, grouped
		// under the given wallet kind, if any.
		var val []string
		{
			val, err = r.red.Sorted().Search().Order(walKin(use, x), 0, -1)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// There might not be any values, and so we do not proceed, but instead
		// continue with the next wallet kind, if any.
		if len(val) == 0 {
			continue
		}

		{
			lis, err := r.SearchWllt(use, objectid.Strings(val))
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	return out, nil
}

func (r *Redis) SearchWllt(use objectid.ID, wal []objectid.ID) ([]*Object, error) {
	var err error

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(objectid.Fmt(wal, fmt.Sprintf(keyfmt.WalletObject, use, "%s"))...)
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(walletObjectNotFoundError, "%v", wal)
		} else if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var out []*Object
	for _, x := range jsn {
		var obj *Object
		{
			obj = &Object{}
		}

		if x != "" {
			err = json.Unmarshal([]byte(x), obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		out = append(out, obj)
	}

	return out, nil
}
