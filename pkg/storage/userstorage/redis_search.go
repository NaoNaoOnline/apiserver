package userstorage

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
		pat = fmt.Sprintf(keyfmt.UserObject, "*")
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

	// For our testing purposes we want to read all user IDs available. For that
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
		out, err = r.SearchUser(objectid.IDs(ids))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return out, nil
}

func (r *Redis) SearchName(nam []string) ([]*Object, error) {
	var err error

	// We need to generate the storage keys for the user name lookups. The first
	// step is to clean the given user names themselves so they can be matched
	// against our indexed format. The formatted user names are then used to
	// create the clean storage key representation.
	var key []string
	{
		key = objectid.Fmt(keyfmt.Strings(nam, keyfmt.Indx), keyfmt.UserName)
	}

	// val will result in a list of all user IDs mapped to the given user names,
	// if any.
	var val []string
	{
		val, err = r.red.Simple().Search().Multi(key...)
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(userNotFoundError, "%v", nam)
		} else if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// There might not be any values, and so we do not proceed, but instead
	// return nothing.
	if len(val) == 0 {
		return nil, nil
	}

	var out []*Object
	{
		out, err = r.SearchUser(objectid.IDs(val))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return out, nil
}

func (r *Redis) SearchSubj(sub string) (*Object, error) {
	var err error

	if sub == "" {
		return nil, tracer.Mask(userSubjectEmptyError)
	}

	var use []string
	{
		use, err = r.red.Simple().Search().Multi(useCla(sub))
		if simple.IsNotFound(err) {
			return nil, tracer.Mask(subjectClaimMappingError)
		} else if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(useObj(objectid.ID(use[0])))
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(userNotFoundError, use[0])
		} else if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var out Object
	{
		err = json.Unmarshal([]byte(jsn[0]), &out)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return &out, nil
}

func (r *Redis) SearchUser(use []objectid.ID) ([]*Object, error) {
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
