package userstorage

import (
	"encoding/json"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
	"github.com/xh3b4sd/redigo/pkg/simple"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) SearchSubj(sub string) (*Object, error) {
	var err error

	if sub == "" {
		return nil, tracer.Mask(subjectClaimEmptyError)
	}

	var use objectid.String
	{
		val, err := r.red.Simple().Search().Value(useCla(sub))
		if simple.IsNotFound(err) {
			return nil, tracer.Mask(subjectClaimMappingError)
		} else if err != nil {
			return nil, tracer.Mask(err)
		}

		use = objectid.String(val)
	}

	var jsn string
	{
		jsn, err = r.red.Simple().Search().Value(useObj(use))
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(userNotFoundError, use.String())
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

func (r *Redis) SearchUser(use []objectid.String) ([]*Object, error) {
	var err error

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(objectid.Fmt(use, keyfmt.UserObject)...)
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(userNotFoundError, "%v", use)
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
