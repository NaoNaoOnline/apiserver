package liststorage

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

type Slicer []*Object

func (s Slicer) List() []objectid.ID {
	var ids []objectid.ID

	for _, x := range s {
		ids = append(ids, x.List)
	}

	return ids
}
