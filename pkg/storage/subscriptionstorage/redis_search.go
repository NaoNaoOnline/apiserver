package subscriptionstorage

import (
	"encoding/json"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/redigo/simple"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) SearchCrtr([]objectid.ID) ([]objectid.ID, error) {
	// TODO SearchCrtr
	//
	//     1. start with a given user ID
	//     2. use user storage to search for event objects that the given user ID reacted to in the form of a link click
	//     3. reduce the sorted list of events to a list of unique user IDs
	//     4. sort the event objects by link clicks, from high to low
	//     5. use wallet storage to search for wallet objects using list of unique user IDs
	//     6. reduce the list of wallet objects by removing those that are not labelled for accounting
	//     8. collect the user IDs of the remaining wallet objects
	//     7. success
	//
	return nil, nil
}

func (r *Redis) SearchPayr(use []objectid.ID, pag [2]int) ([]*Object, error) {
	var err error

	var out []*Object
	for _, x := range use {
		// val will result in a list of all subscription IDs created by the given
		// user, if any.
		var val []string
		{
			val, err = r.red.Sorted().Search().Order(subPay(x), pag[0], pag[1])
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
			lis, err := r.SearchSubs(objectid.IDs(val))
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	return out, nil
}

func (r *Redis) SearchRecv(use []objectid.ID, pag [2]int) ([]*Object, error) {
	var err error

	var out []*Object
	for _, x := range use {
		// val will result in a list of all subscription IDs created by the given
		// user, if any.
		var val []string
		{
			val, err = r.red.Sorted().Search().Order(subRec(x), pag[0], pag[1])
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
			lis, err := r.SearchSubs(objectid.IDs(val))
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	return out, nil
}

func (r *Redis) SearchSubs(inp []objectid.ID) ([]*Object, error) {
	var err error

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(objectid.Fmt(inp, keyfmt.SubscriptionObject)...)
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(subscriptionObjectNotFoundError, "%v", inp)
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

		{
			out = append(out, obj)
		}
	}

	return out, nil
}
