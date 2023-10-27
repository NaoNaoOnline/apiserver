package rulestorage

import (
	"encoding/json"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/redigo/pkg/simple"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) SearchList(lis []objectid.ID) ([]*Object, error) {
	var err error

	var out []*Object
	for _, x := range lis {
		if x == "" {
			return nil, tracer.Mask(listIDEmptyError)
		}

		// val will result in a list of all rule IDs belonging to the given list ID,
		// if any.
		var val []string
		{
			val, err = r.red.Sorted().Search().Order(rulLis(x), 0, -1)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// There might not be any values, and so we do not proceed, but instead
		// continue with the next event ID, if any.
		if len(val) == 0 {
			continue
		}

		{
			lis, err := r.SearchRule(objectid.IDs(val))
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	return out, nil
}

func (r *Redis) SearchRule(inp []objectid.ID) ([]*Object, error) {
	var err error

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(objectid.Fmt(inp, keyfmt.RuleObject)...)
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(ruleObjectNotFoundError, "%v", inp)
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
