package subscriptionstorage

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

type Slicer []*Object

func (s Slicer) Subs() []objectid.ID {
	var mem []objectid.ID

	for _, x := range s {
		mem = append(mem, x.Subs)
	}

	return mem
}
