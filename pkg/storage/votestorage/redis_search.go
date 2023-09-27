package votestorage

import (
	"encoding/json"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/xh3b4sd/redigo/pkg/simple"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) SearchDesc(inp []objectid.String) ([]*Object, error) {
	var err error

	var out []*Object
	for _, x := range inp {
		// val will result in a list of all vote IDs belonging to the given
		// description ID, if any.
		var val []string
		{
			val, err = r.red.Sorted().Search().Order(votDes(x), 0, -1)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// There might not be any values, and so we do not proceed, but instead
		// continue with the next description ID, if any.
		if len(val) == 0 {
			continue
		}

		{
			lis, err := r.SearchVote(objectid.Strings(val))
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	return out, nil
}

func (r *Redis) SearchVote(inp []objectid.String) ([]*Object, error) {
	var err error

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(objectid.Fmt(inp, keyfmt.VoteObject)...)
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(voteObjectNotFoundError, "%v", inp)
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

func (r *Redis) searchDesc(des objectid.String) (*descriptionstorage.Object, error) {
	var err error

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(desObj(des))
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(descriptionObjectNotFoundError, des.String())
		} else if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var obj *descriptionstorage.Object
	{
		obj = &descriptionstorage.Object{}
	}

	{
		err = json.Unmarshal([]byte(jsn[0]), obj)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return obj, nil
}

func (r *Redis) searchEvnt(eve objectid.String) (*eventstorage.Object, error) {
	var err error

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(eveObj(eve))
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(eventObjectNotFoundError, eve.String())
		} else if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var obj *eventstorage.Object
	{
		obj = &eventstorage.Object{}
	}

	{
		err = json.Unmarshal([]byte(jsn[0]), obj)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return obj, nil
}
