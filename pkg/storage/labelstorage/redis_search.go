package labelstorage

import (
	"encoding/json"

	"github.com/NaoNaoOnline/apiserver/pkg/scoreid"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Search(kin string) ([]*Object, error) {
	var err error

	if kin != "cate" && kin != "host" {
		return nil, tracer.Mask(invalidInputError)
	}

	var key []string
	{
		key, err = r.red.Sorted().Search().Order(keyKin(kin), 0, -1)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	for i := range key {
		key[i] = keyObj(scoreid.String(key[i]))
	}

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(key...)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var out []*Object
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

	return out, nil
}
