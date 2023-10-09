package policystorage

import (
	"encoding/json"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/redigo/pkg/simple"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) SearchAggr() ([]*Object, []*Object, error) {
	var err error

	var obj []*Object
	{
		obj, err = r.SearchKind([]string{"CreateMember", "CreateSystem", "DeleteMember", "DeleteSystem"})
		if err != nil {
			return nil, nil, tracer.Mask(err)
		}
	}

	var agg []*Object
	var del []*Object
	{
		agg, del = r.searchAggr(obj)
	}

	return agg, del, nil
}

func (r *Redis) SearchKind(inp []string) ([]*Object, error) {
	var err error

	var key []string
	for _, x := range inp {
		if x != "CreateMember" && x != "CreateSystem" && x != "DeleteMember" && x != "DeleteSystem" {
			return nil, tracer.Mask(policyKindInvalidError)
		}

		{
			key = append(key, polKin(x))
		}
	}

	// val will result in a list of all policy IDs grouped under the given policy
	// kind, if any.
	var val []string
	{
		val, err = r.red.Sorted().Search().Union(key...)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// There might not be any values, and so we do not proceed, but instead
	// return the zero value.
	if len(val) == 0 {
		return nil, nil
	}

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(objectid.Fmt(val, keyfmt.PolicyObject)...)
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(policyObjectNotFoundError, "%v", val)
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

func (r *Redis) searchAggr(inp []*Object) ([]*Object, []*Object) {
	var cre []*Object
	var del []*Object

	for _, x := range inp {
		if x.Kind == "CreateMember" || x.Kind == "CreateSystem" {
			cre = append(cre, x)
		}
		if x.Kind == "DeleteMember" || x.Kind == "DeleteSystem" {
			del = append(del, x)
		}
	}

	var agg []*Object
	for _, x := range cre {
		var exi bool

		for _, y := range del {
			if x.Eqlrec(y) {
				exi = true
				break
			}
		}

		if !exi {
			agg = append(agg, x)
		}
	}

	return agg, del
}
