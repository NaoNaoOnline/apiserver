package labelstorage

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

type Slicer []*Object

func (s Slicer) Labl() []objectid.ID {
	var ids []objectid.ID

	for _, x := range s {
		ids = append(ids, x.Labl)
	}

	return ids
}
