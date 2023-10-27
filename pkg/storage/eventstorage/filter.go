package eventstorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/generic"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
)

type Filter []*Object

func (f Filter) Cat(ids ...objectid.ID) []*Object {
	var obj []*Object

	for _, x := range f {
		if !generic.Any(objectid.Strings(x.Cate), objectid.Strings(ids)) {
			obj = append(obj, x)
		}
	}

	return obj
}

func (f Filter) Hos(ids ...objectid.ID) []*Object {
	var obj []*Object

	for _, x := range f {
		if !generic.Any(objectid.Strings(x.Host), objectid.Strings(ids)) {
			obj = append(obj, x)
		}
	}

	return obj
}

func (f Filter) Use(ids ...objectid.ID) []*Object {
	var obj []*Object

	for _, x := range f {
		if !generic.Any([]string{string(x.User)}, objectid.Strings(ids)) {
			obj = append(obj, x)
		}
	}

	return obj
}
