package descriptionstorage

import (
	"encoding/json"
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/redigo/pkg/simple"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) SearchDesc(use objectid.ID, inp []objectid.ID) ([]*Object, error) {
	var err error

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(objectid.Fmt(inp, keyfmt.DescriptionObject)...)
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(descriptionObjectNotFoundError, "%v", inp)
		} else if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var lik []string
	if use != "" {
		lik, err = r.red.Simple().Search().Multi(objectid.Fmt(inp, fmt.Sprintf(keyfmt.DescriptionLike, use, "%s"))...)
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

		// Annotate the description object with the indication whether the calling
		// user liked this description already. The linear mapping between
		// description objects and like indicators should work reliably because
		// Simple.Search.Multi gets called with the same amount of keys for each
		// query. And so each and every JSON string should relate to each and every
		// like indicator for the calling user.
		if len(lik) == len(jsn) && lik[i] == "1" {
			obj.Like.User = true
		}

		out = append(out, obj)
	}

	return out, nil
}

func (r *Redis) SearchEvnt(use objectid.ID, evn []objectid.ID) ([]*Object, error) {
	var err error

	var out []*Object
	for _, x := range evn {
		if x == "" {
			return nil, tracer.Mask(eventIDEmptyError)
		}

		// val will result in a list of all description IDs belonging to the given
		// event ID, if any.
		var val []string
		{
			val, err = r.red.Sorted().Search().Order(desEve(x), 0, -1)
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
			lis, err := r.SearchDesc(use, objectid.IDs(val))
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	return out, nil
}
