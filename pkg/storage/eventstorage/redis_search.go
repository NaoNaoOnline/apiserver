package eventstorage

import (
	"encoding/json"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
	"github.com/xh3b4sd/redigo/pkg/simple"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) SearchEvnt(inp []objectid.String) ([]*Object, error) {
	var err error

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(objectid.Fmt(inp, keyfmt.EventObject)...)
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(eventObjectNotFoundError, "%v", inp)
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

func (r *Redis) SearchLabl(lab []objectid.String) ([]*Object, error) {
	var err error

	// key will result in a list of all event IDs associated to the given label
	// IDs, if any.
	var key []string
	{
		key, err = r.red.Sorted().Search().Inter(objectid.Fmt(lab, keyfmt.EventLabel)...)
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
		jsn, err = r.red.Simple().Search().Multi(objectid.Fmt(key, keyfmt.EventObject)...)
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(eventObjectNotFoundError, "%v", key)
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

const (
	oneWeek = time.Hour * 24 * 7
)

func (r *Redis) SearchLtst() ([]*Object, error) {
	var err error

	var now time.Time
	{
		now = time.Now().UTC()
	}

	var max float64
	var min float64
	{
		max = float64(now.Add(+oneWeek).Unix())
		min = float64(now.Add(-oneWeek).Unix())
	}

	var key []string
	{
		key, err = r.red.Sorted().Search().Score(keyfmt.EventTime, max, min)
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
		jsn, err = r.red.Simple().Search().Multi(objectid.Fmt(key, keyfmt.EventObject)...)
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(eventObjectNotFoundError, "%v", key)
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

func (r *Redis) SearchRctn(use objectid.String) ([]*Object, error) {
	var err error

	var key []string
	{
		key, err = r.red.Sorted().Search().Order(eveVot(use), 0, -1)
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
		jsn, err = r.red.Simple().Search().Multi(objectid.Fmt(key, keyfmt.EventObject)...)
		if simple.IsNotFound(err) {
			// It may happen that events get deleted, that users have reacted to. The
			// event deletion process is not atomic and so it might happen that some
			// event objects cannot be found anymore intermittently.
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
