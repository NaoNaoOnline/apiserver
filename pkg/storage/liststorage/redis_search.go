package liststorage

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/redigo/pkg/simple"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) SearchFake() ([]*Object, error) {
	var err error

	var pat string
	{
		pat = fmt.Sprintf(keyfmt.ListObject, "*")
	}

	var res chan string
	{
		res = make(chan string, 1)
	}

	var ids []string
	go func() {
		for s := range res {
			spl := strings.Split(s, "/")
			ids = append(ids, spl[len(spl)-1])
		}
	}()

	// For our testing purposes we want to read all list IDs available. For that
	// purpose we do not need to provide a done channel, because we do not want to
	// cancel the walk through all data early. We want all users.
	{
		err = r.red.Walker().Simple(pat, nil, res)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var out []*Object
	{
		out, err = r.SearchList(objectid.IDs(ids))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return out, nil
}

func (r *Redis) SearchList(inp []objectid.ID) ([]*Object, error) {
	var err error

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(objectid.Fmt(inp, keyfmt.ListObject)...)
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(listObjectNotFoundError, "%v", inp)
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

func (r *Redis) SearchUser(use objectid.ID) ([]*Object, error) {
	var err error

	var out []*Object
	{
		// val will result in a list of all list IDs created by the given user, if
		// any.
		var val []string
		{
			val, err = r.red.Sorted().Search().Order(lisUse(use), 0, -1)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// There might not be any values, and so we do not proceed, but instead
		// return nothing.
		if len(val) == 0 {
			return nil, nil
		}

		{
			lis, err := r.SearchList(objectid.IDs(val))
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	return out, nil
}
