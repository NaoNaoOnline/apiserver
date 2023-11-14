package walletstorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/generic"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
)

type Slicer []*Object

// Wllt returns all the wallet IDs for the underling list of wallet objects.
func (s Slicer) Wllt() []objectid.ID {
	var ids []objectid.ID

	for _, x := range s {
		ids = append(ids, x.Wllt)
	}

	return ids
}

func (s Slicer) Labl(lab string) bool {
	for _, x := range s {
		if x.HasLab(lab) {
			return true
		}
	}

	return false
}

func (s Slicer) Obct() []*Object {
	return s
}

func (s Slicer) Slct(ids ...objectid.ID) []*Object {
	var obj []*Object

	for _, x := range s {
		if generic.Any([]string{x.Wllt.String()}, objectid.Strings(ids)) {
			obj = append(obj, x)
		}
	}

	return obj
}
