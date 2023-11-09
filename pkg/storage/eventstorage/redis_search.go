package eventstorage

import (
	"encoding/json"
	"fmt"
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

func (r *Redis) SearchEvnt(use objectid.ID, inp []objectid.ID) ([]*Object, error) {
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

	var lin []string
	if use != "" {
		lin, err = r.red.Simple().Search().Multi(objectid.Fmt(inp, fmt.Sprintf(keyfmt.LinkMapping, use, "%s"))...)
		if simple.IsNotFound(err) {
			// fall through
		} else if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var out []*Object
	for i := range jsn {
		var obj *Object
		{
			obj = &Object{}
		}

		if jsn[i] != "" {
			err = json.Unmarshal([]byte(jsn[i]), obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Annotate the event object with the indication whether the calling user
		// visited this event already. The linear mapping between event objects and
		// visit indicators should work reliably because Simple.Search.Multi gets
		// called with the same amount of keys for each query. And so each and every
		// JSON string should relate to each and every visit indicator for the
		// calling user.
		if len(lin) == len(jsn) && lin[i] == "1" {
			obj.Clck.User = true
		}

		{
			out = append(out, obj)
		}
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
		out, err = r.SearchEvnt("", objectid.IDs(val))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return out, nil
}

func (r *Redis) SearchLike(use objectid.ID, min int, max int) ([]*Object, error) {
	var err error

	// val will result in a list of all paired ID strings, containing the event
	// ID, that the given user reacted to in the form of a description like.
	var val []string
	{
		val, err = r.red.Sorted().Search().Order(likUse(use), min, max)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// There might not be any values, and so we do not proceed, but instead return
	// nothing.
	if len(val) == 0 {
		return nil, nil
	}

	var out []*Object
	{
		out, err = r.SearchEvnt(use, generic.Uni(objectid.Frst(val)))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return out, nil
}

func (r *Redis) SearchLink(eve objectid.ID) ([]objectid.ID, error) {
	var err error

	// val will result in a list of all user IDs having visited the given event
	// ID, if any.
	var val []string
	{
		val, err = r.red.Sorted().Search().Order(linEve(eve), 0, -1)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// There might not be any values, and so we do not proceed, but instead
	// continue with the next event ID, if any.
	if len(val) == 0 {
		return nil, nil
	}

	return objectid.IDs(val), nil
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

	var inc []string
	{
		inc = sli.Incl()
	}

	// There might not be any rules to begin with, and so we do not proceed, but
	// instead return nothing. Note that we check for the included items and not
	// the whole list of rules, because it may be that all rules only defines
	// excludes. While a list containing only of excludes does not make sense, it
	// may happen that we face such a situation.
	if len(inc) == 0 {
		return nil, nil
	}

	// val will result in a list of all event IDs to be included in the given
	// list.
	var val []string
	{
		val, err = r.red.Sorted().Search().Union(inc...)
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(eventObjectNotFoundError, "%v", inc)
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
		out, err = r.SearchEvnt("", generic.Uni(objectid.Frst(val)))
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

	// Remove the event objects that the given user IDs reacted to in the form of
	// a description like.
	if len(sli.Like()) != 0 {
		// val will result in a list of all event IDs that the given users reacted
		// to in the form of a description like.
		var val []string
		{
			val, err = r.red.Sorted().Search().Union(objectid.Fmt(sli.Like(), keyfmt.LikeUser)...)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Filter event IDs against event IDs.
		{
			out = out.Fltr().Evnt(generic.Uni(objectid.Frst(val))...)
		}
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
		out, err = r.SearchEvnt("", objectid.IDs(val))
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
			lis, err := r.SearchEvnt(x, objectid.IDs(val))
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	return out, nil
}
