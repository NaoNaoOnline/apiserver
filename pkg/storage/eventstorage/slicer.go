package eventstorage

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

type Slicer []*Object

func (s Slicer) Fltr() Filter {
	return Filter(s)
}

func (s Slicer) Evnt() []objectid.ID {
	var ids []objectid.ID

	for _, x := range s {
		ids = append(ids, x.Evnt)
	}

	return ids
}

func (s Slicer) Func(fun func(objectid.ID) bool) Slicer {
	var obj []*Object

	for _, x := range s {
		if fun(x.Evnt) {
			obj = append(obj, x)
		}
	}

	return obj
}

// Upcm returns the subset of event objects that have not yet happened based on
// the current time of execution. That is, the subset of upcoming events.
func (s Slicer) Upcm() Slicer {
	var obj []*Object

	for _, x := range s {
		if !x.Happnd() {
			obj = append(obj, x)
		}
	}

	return obj
}
