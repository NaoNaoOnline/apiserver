package eventstorage

import (
	"encoding/json"

	"github.com/NaoNaoOnline/apiserver/pkg/scoreid"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Search(hos scoreid.String, cat ...scoreid.String) ([]*Object, error) {
	var err error

	var lab []string
	if hos != "" {
		lab = append(lab, eveLab(hos))
	}
	for _, x := range cat {
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
