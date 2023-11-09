package descriptionstorage

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

type Slicer []*Object

// Amnt returns a map where the keys are event IDs and the values are the
// respective number of descriptions per said event.
func (s Slicer) Amnt() map[objectid.ID]int {
	amo := map[objectid.ID]int{}

	for _, x := range s {
		amo[x.Evnt]++
	}

	return amo
}

func (s Slicer) Desc() []objectid.ID {
	var ids []objectid.ID

	for _, x := range s {
		ids = append(ids, x.Desc)
	}

	return ids
}

func (s Slicer) Evnt() []objectid.ID {
	var ids []objectid.ID

	for _, x := range s {
		ids = append(ids, x.Evnt)
	}

	return ids
}
