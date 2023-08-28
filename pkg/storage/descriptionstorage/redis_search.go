package descriptionstorage

import (
	"encoding/json"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
	"github.com/xh3b4sd/redigo/pkg/simple"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Search(evn []objectid.String) ([]*Object, error) {
	var err error

	var out []*Object
	for _, x := range evn {
		if x == "" {
			return nil, tracer.Mask(invalidEventIDError)
		}

		// key will result in a list of all description IDs belonging to the given
		// event ID, if any.
		var key []string
		{
			key, err = r.red.Sorted().Search().Order(desEve(x), 0, -1)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// There might not be any keys, and so we do not proceed, but instead
		// continue with the next event ID, if any.
		if len(key) == 0 {
			continue
		}

		var jsn []string
		{
			jsn, err = r.red.Simple().Search().Multi(objectid.Fmt(key, keyfmt.DescriptionObject)...)
			if simple.IsNotFound(err) {
				return nil, tracer.Maskf(descriptionNotFoundError, "%v", key)
			} else if err != nil {
				return nil, tracer.Mask(err)
			}
		}

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
	}

	return out, nil
}
