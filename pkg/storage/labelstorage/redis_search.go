package labelstorage

import (
	"encoding/json"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
	"github.com/xh3b4sd/redigo/pkg/simple"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Search(inp []string) ([]*Object, error) {
	var err error

	var out []*Object
	for _, x := range inp {
		if x != "cate" && x != "host" {
			return nil, tracer.Mask(labelKindInvalidError)
		}

		// val will result in a list of all label IDs grouped under the given label
		// kind, if any.
		var val []string
		{
			val, err = r.red.Sorted().Search().Order(labKin(x), 0, -1)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// There might not be any values, and so we do not proceed, but instead
		// continue with the next label kind, if any.
		if len(val) == 0 {
			continue
		}

		var jsn []string
		{
			jsn, err = r.red.Simple().Search().Multi(objectid.Fmt(val, keyfmt.LabelObject)...)
			if simple.IsNotFound(err) {
				return nil, tracer.Maskf(labelObjectNotFoundError, "%v", val)
			} else if err != nil {
				return nil, tracer.Mask(err)
			}
		}

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
	}

	return out, nil
}