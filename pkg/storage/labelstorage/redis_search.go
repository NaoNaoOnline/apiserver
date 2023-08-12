package labelstorage

import (
	"encoding/json"

	"github.com/NaoNaoOnline/apiserver/pkg/scoreid"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Search(kin []string) ([]*Object, error) {
	var err error

	var out []*Object
	for _, x := range kin {
		if x != "cate" && x != "host" {
			return nil, tracer.Mask(invalidLabelKindError)
		}

		var key []string
		{
			key, err = r.red.Sorted().Search().Order(labKin(x), 0, -1)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		for i := range key {
			key[i] = labObj(scoreid.String(key[i]))
		}

		var jsn []string
		{
			jsn, err = r.red.Simple().Search().Multi(key...)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		for _, x := range jsn {
			var obj *Object
			{
				obj = &Object{}
			}

			err = json.Unmarshal([]byte(x), obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, obj)
		}
	}

	return out, nil
}
