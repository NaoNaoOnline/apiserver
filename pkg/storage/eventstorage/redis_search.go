package eventstorage

import (
	"encoding/json"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/generic"
	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/rulestorage"
	"github.com/xh3b4sd/redigo/pkg/simple"
	"github.com/xh3b4sd/tracer"
)

const (
	oneWeek = time.Hour * 24 * 7
)

func (r *Redis) SearchEvnt(inp []objectid.ID) ([]*Object, error) {
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

func (r *Redis) SearchHpnd() ([]*Object, error) {
	var err error

	var now time.Time
	{
		now = time.Now().UTC()
	}

	var min time.Time
	var max time.Time
	{
		min = now.Add(-oneWeek)
		max = now
	}

	var out []*Object
	{
		out, err = r.SearchTime(min, max)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return out, nil
}

func (r *Redis) SearchLabl(lab []objectid.ID) ([]*Object, error) {
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

	var out []*Object
	{
		out, err = r.SearchEvnt(objectid.IDs(val))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return out, nil
}

func (r *Redis) SearchLike(use objectid.ID, min int, max int) ([]*Object, error) {
	var err error

	// The user likes are indexed in a way were description IDs are values and
	// event IDs are scores. Below we search for all values and their respective
	// scores using the 4th parameter true. Note that the event IDs will
	// potentially be duplicated across the list.
	var lis []string
	{
		lis, err = r.red.Sorted().Search().Order(likUse(use), min, max, true)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// There might not be any items, and so we do not proceed, but instead return
	// nothing.
	if len(lis) == 0 {
		return nil, nil
	}

	// Here we select the event IDs from the list gathered above. Every other
	// string in that list represents an element score since we searched by order
	// using true.
	var sco []string
	for i := 1; i < len(lis); i += 2 {
		sco = append(sco, lis[i])
	}

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(objectid.Fmt(generic.Uni(sco), keyfmt.EventObject)...)
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

func (r *Redis) SearchUpcm() ([]*Object, error) {
	var err error

	var now time.Time
	{
		now = time.Now().UTC()
	}

	var min time.Time
	var max time.Time
	{
		min = now
		max = now.Add(+oneWeek)
	}

	var out []*Object
	{
		out, err = r.SearchTime(min, max)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return out, nil
}

func (r *Redis) SearchRule(rul []*rulestorage.Object) ([]*Object, error) {
	var err error

	var sli rulestorage.Slicer
	{
		sli = rulestorage.Slicer(rul)
	}

	// val will result in a list of all event IDs to be included in the given
	// list.
	var val []string
	{
		val, err = r.red.Sorted().Search().Union(sli.Incl()...)
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(eventObjectNotFoundError, "%v", sli.Incl())
		} else if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// There might not be any keys, and so we do not proceed, but instead return
	// nothing.
	if len(val) == 0 {
		return nil, nil
	}

	var out Slicer
	{
		out, err = r.SearchEvnt(objectid.IDs(val))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// Remove the event objects that match all the rule's exclude definitions.
	{
		out = out.Fltr().Cate(sli.Cate()...)
		out = out.Fltr().Host(sli.Host()...)
		out = out.Fltr().User(sli.User()...)
	}

	return out, nil
}

func (r *Redis) SearchTime(min time.Time, max time.Time) ([]*Object, error) {
	var err error

	// val will result in a list of all event IDs indexed to happen during the
	// given time period, if any.
	var val []string
	{
		val, err = r.red.Sorted().Search().Score(keyfmt.EventTime, float64(min.Unix()), float64(max.Unix()))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// There might not be any values, and so we do not proceed, but instead
	// return nothing.
	if len(val) == 0 {
		return nil, nil
	}

	var out []*Object
	{
		out, err = r.SearchEvnt(objectid.IDs(val))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return out, nil
}

func (r *Redis) SearchUser(use []objectid.ID) ([]*Object, error) {
	var err error

	var out []*Object
	for _, x := range use {
		// val will result in a list of all event IDs created by the given user ID, if
		// any.
		var val []string
		{
			val, err = r.red.Sorted().Search().Order(eveUse(x), 0, -1)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// There might not be any values, and so we do not proceed, but instead
		// continue with the next user ID, if any.
		if len(val) == 0 {
			continue
		}

		{
			lis, err := r.SearchEvnt(objectid.IDs(val))
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	return out, nil
}
