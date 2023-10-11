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
		agg, del = searchAggr(obj)
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

	var out []*Object
	{
		out, err = r.SearchPlcy(objectid.IDs(val))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return out, nil
}

func (r *Redis) SearchPlcy(inp []objectid.ID) ([]*Object, error) {
	var err error

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(objectid.Fmt(inp, keyfmt.PolicyObject)...)
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(policyObjectNotFoundError, "%v", inp)
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

func searchAggr(inp []*Object) ([]*Object, []*Object) {
	var cre []*Object
	var del []*Object

	// We want to remove the create records that have equivalent delete records.
	// Therefore we prepare our data structures and create separate lists for
	// both.
	for _, o := range inp {
		if o.Kind == "CreateMember" || o.Kind == "CreateSystem" {
			cre = append(cre, o)
		}
		if o.Kind == "DeleteMember" || o.Kind == "DeleteSystem" {
			del = append(del, o)
		}
	}

	// Our aggregation considers policy contract deployments on multiple chains.
	// Any create record on any chain that is not negated by its equivalent delete
	// record on that same chain remains valid for the platform. That means we
	// have to keep track of the chain IDs that records carry with them, since we
	// will use them as indicator for negation. For instance, consider a create
	// record with the chain IDs [1 2 3] and an equivalent delete record with the
	// chain IDs [1 3]. The create record in that example remains valid since its
	// policy record on chain ID 2 remains intact.
	chn := map[*Object][]int64{}

	var agg []*Object
	for _, x := range cre {
		var neg bool

		for _, y := range del {
			// Delete record y is only allowed to negate create record x if their SMA
			// fields are equal. So if their SMA fields do not match, we ignore y and
			// move on to the next delete record.
			if !x.Eqlrec(y) {
				continue
			}

			// Since the SMA fields of x and y match we should be able to find a
			// common chain ID. The common chain ID is added to the mapping so that we
			// can further compare whether all create records on all chains got
			// negated already.
			{
				c, e := eqlChn(x, y)
				if e {
					chn[x] = append(chn[x], c)
				}
			}

			// If all create records on all chains got negated, then we can remove the
			// create record from our aggregation.
			if len(chn[x]) == len(x.ChID) {
				neg = true
				break
			}
		}

		// Any create record that got negated is not added to our aggregation list.
		if !neg {
			agg = append(agg, x)
		}
	}

	return agg, del
}

// eqlChn tries to find a common chain ID. When eqlChn is used, then chain IDs
// may or may not be shared between policy records. Therefore the second return
// value bool is returned in order to indicate that a chain ID in fact matched
// between the given objects. The returned chain ID value should not be
// considered without ensuring that the returned bool is true, since the return
// value 0 may be misleading on its own in any given case.
func eqlChn(a *Object, b *Object) (int64, bool) {
	for _, x := range a.ChID {
		for _, y := range b.ChID {
			if x == y {
				return x, true
			}
		}
	}

	return 0, false
}
