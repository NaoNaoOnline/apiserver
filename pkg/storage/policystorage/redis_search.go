package policystorage

import (
	"encoding/json"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) SearchActv() ([]*Object, error) {
	var err error

	var val []string
	{
		val, err = r.red.Simple().Search().Multi(keyfmt.PolicyActive)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if len(val) != 1 {
		return nil, tracer.Mask(runtime.ExecutionFailedError)
	}

	var out []*Object
	{
		out = []*Object{}
	}

	{
		err = json.Unmarshal([]byte(val[0]), &out)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return out, nil
}

func (r *Redis) SearchBffr() ([]*Object, error) {
	var err error

	var jsn []string
	{
		jsn, err = r.red.Sorted().Search().Order(keyfmt.PolicyBuffer, 0, -1)
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
