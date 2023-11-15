package labelstorage

import (
	"encoding/json"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/redigo/pkg/simple"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) SearchBltn() []*Object {
	return []*Object{
		{
			Kind: "bltn",
			Name: objectfield.String{
				Data: "Discord",
			},
			User: objectfield.ID{
				Data: objectid.System(),
			},
		},
		{
			Kind: "bltn",
			Name: objectfield.String{
				Data: "Google",
			},
			User: objectfield.ID{
				Data: objectid.System(),
			},
		},
		{
			Kind: "bltn",
			Name: objectfield.String{
				Data: "Twitch",
			},
			User: objectfield.ID{
				Data: objectid.System(),
			},
		},
		{
			Kind: "bltn",
			Name: objectfield.String{
				Data: "Twitter",
			},
			User: objectfield.ID{
				Data: objectid.System(),
			},
		},
		{
			Kind: "bltn",
			Name: objectfield.String{
				Data: "YouTube",
			},
			User: objectfield.ID{
				Data: objectid.System(),
			},
		},
	}
}

func (r *Redis) SearchKind(inp []string) ([]*Object, error) {
	var err error

	var out []*Object
	for _, x := range inp {
		if x != "bltn" && x != "cate" && x != "host" {
			return nil, tracer.Mask(labelKindInvalidError)
		}

		// val will result in a list of all label IDs grouped under the given label
		// kind, if any.
		var val []string
		{
			val, err = r.red.Sorted().Search().Order(labKin(x), 0, -1)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// There might not be any values, and so we do not proceed, but instead
		// continue with the next label kind, if any.
		if len(val) == 0 {
			continue
		}

		{
			lis, err := r.SearchLabl(objectid.IDs(val))
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	return out, nil
}

func (r *Redis) SearchLabl(inp []objectid.ID) ([]*Object, error) {
	var err error

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(objectid.Fmt(inp, keyfmt.LabelObject)...)
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(labelObjectNotFoundError, "%v", inp)
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
		// val will result in a list of all label IDs created by the given user, if
		// any.
		var val []string
		{
			val, err = r.red.Sorted().Search().Order(labUse(use), 0, -1)
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
			lis, err := r.SearchLabl(objectid.IDs(val))
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	return out, nil
}
