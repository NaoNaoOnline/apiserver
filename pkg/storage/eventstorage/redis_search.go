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

	// val will result in a list of all event IDs associated to the given label
	// IDs, if any.
	var val []string
	{
		val, err = r.red.Sorted().Search().Inter(objectid.Fmt(lab, keyfmt.EventLabel)...)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// There might not be any values, and so we do not proceed, but instead
	// return nothing.
	if len(val) == 0 {
		return nil, nil
	}

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(objectid.Fmt(val, keyfmt.EventObject)...)
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(eventObjectNotFoundError, "%v", val)
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

	var min time.Time
	var max time.Time
	{
		min = now.Add(-oneWeek)
		max = now.Add(+oneWeek)
	}

	var out []*Object
	{
		out, err = r.searchTime(min, max)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return out, nil
}

func (r *Redis) searchTime(min time.Time, max time.Time) ([]*Object, error) {
	var err error

	// val will result in a list of all event IDs indexed to happen during the
	// given time period. kind, if any.
	var val []string
	{
		val, err = r.red.Sorted().Search().Value(keyfmt.EventTime, float64(max.Unix()), float64(min.Unix()))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// There might not be any values, and so we do not proceed, but instead
	// return nothing.
	if len(val) == 0 {
		return nil, nil
	}

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(objectid.Fmt(val, keyfmt.EventObject)...)
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(eventObjectNotFoundError, "%v", val)
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

	// The user votes are indexed in a way were vote IDs are values and event IDs
	// are scores. Below we lookup all scores, resulting in a list of event IDs
	// that potentially contains duplicates.
	var sco []string
	{
		sco, err = r.red.Sorted().Search().Order(votUse(use), 0, -1, true)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// There might not be any values, and so we do not proceed, but instead
	// return nothing.
	if len(sco) == 0 {
		return nil, nil
	}

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(objectid.Fmt(objectid.Uni(sco), keyfmt.EventObject)...)
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
