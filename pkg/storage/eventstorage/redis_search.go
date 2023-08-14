package eventstorage

import (
	"encoding/json"

	"github.com/NaoNaoOnline/apiserver/pkg/scoreid"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) SearchEvnt(evn []scoreid.String) ([]*Object, error) {
	var err error

	var key []string
	for _, x := range evn {
		key = append(key, eveObj(x))
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

func (r *Redis) SearchLabl(inp []scoreid.String) ([]*Object, error) {
	var err error

	var lab []string
	for _, x := range inp {
		lab = append(lab, eveLab(x))
	}

	var key []string
	{
		key, err = r.red.Sorted().Search().Inter(lab...)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	for i := range key {
		key[i] = eveObj(scoreid.String(key[i]))
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
