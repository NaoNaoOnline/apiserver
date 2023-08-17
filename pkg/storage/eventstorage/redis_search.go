package eventstorage

import (
	"encoding/json"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/scoreid"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) SearchEvnt(evn []scoreid.String) ([]*Object, error) {
	var err error

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(scoreid.Fmt(evn, keyfmt.EventObject)...)
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

func (r *Redis) SearchLabl(lab []scoreid.String) ([]*Object, error) {
	var err error

	// key will result in a list of all event IDs associated to the given label
	// IDs, if any.
	var key []string
	{
		key, err = r.red.Sorted().Search().Inter(scoreid.Fmt(lab, keyfmt.EventLabel)...)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// There might not be any keys, and so we do not proceed, but instead
	// return nothing.
	if len(key) == 0 {
		return nil, nil
	}

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(scoreid.Fmt(key, keyfmt.EventObject)...)
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
