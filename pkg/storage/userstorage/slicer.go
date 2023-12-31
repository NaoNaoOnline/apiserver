package userstorage

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

type Slicer []*Object

func (s Slicer) User() []objectid.ID {
	var ids []objectid.ID

	for _, x := range s {
		ids = append(ids, x.User)
	}

	return ids
}
