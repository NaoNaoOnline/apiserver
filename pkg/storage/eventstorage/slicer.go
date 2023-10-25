package eventstorage

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

type Slicer []*Object

func (s Slicer) IDs() []objectid.ID {
	var ids []objectid.ID

	for _, x := range s {
		ids = append(ids, x.Evnt)
	}

	return ids
}

// Upc returns the subset of event objects that have not yet happened based on
// the current time of execution. That is, the subset of upcoming events.
func (s Slicer) Upc() Slicer {
	var obj []*Object

	for _, x := range s {
		if !x.Happnd() {
			obj = append(obj, x)
		}
	}

	return obj
}
