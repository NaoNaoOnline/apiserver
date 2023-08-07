package userstorage

import (
	"encoding/json"

	"github.com/NaoNaoOnline/apiserver/pkg/scoreid"
	"github.com/xh3b4sd/redigo/pkg/simple"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Search(sub string, use scoreid.String) (*Object, error) {
	var err error

	// It is not valid to call Search with both inputs empty or given. Only either
	// input must be provided. Either the external subject claim, or the internal
	// user ID.
	if (sub == "" && use == "") || (sub != "" && use != "") {
		return nil, tracer.Mask(invalidInputError)
	}

	if use == "" {
		val, err := r.red.Simple().Search().Value(useCla(sub))
		if simple.IsNotFound(err) {
			return nil, tracer.Mask(notFoundError)
		} else if err != nil {
			return nil, tracer.Mask(err)
		}

		use = scoreid.String(val)
	}

	var jsn string
	{
		jsn, err = r.red.Simple().Search().Value(useObj(use))
		if simple.IsNotFound(err) {
			return nil, tracer.Mask(notFoundError)
		} else if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var out Object
	{
		err = json.Unmarshal([]byte(jsn), &out)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return &out, nil
}
