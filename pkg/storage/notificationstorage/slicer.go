package notificationstorage

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

type Slicer []*Object

func (s Slicer) Evnt() []objectid.ID {
	var ids []objectid.ID

	for _, x := range s {
		ids = append(ids, x.Evnt)
	}

	return ids
}

func (s Slicer) Noti() []objectid.ID {
	var ids []objectid.ID

	for _, x := range s {
		ids = append(ids, x.Noti)
	}

	return ids
}
