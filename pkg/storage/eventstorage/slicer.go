package eventstorage

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

type Slicer []*Object

// Evnt returns all the event IDs for the underling list of event objects.
func (s Slicer) Evnt() []objectid.ID {
	var ids []objectid.ID

	for _, x := range s {
		ids = append(ids, x.Evnt)
	}

	return ids
}

// Func returns a slicer containing the underlying objects that match the given
// functions criteria. That is, each object for which fun returns true will be
// returned with the new slicer.
func (s Slicer) Func(fun func(*Object) bool) Slicer {
	var obj []*Object

	for _, x := range s {
		if fun(x) {
			obj = append(obj, x)
		}
	}

	return obj
}

// Muse returns the cumulative amount for metrics described by the given metric
// label, for the underling list of event objects.
func (s Slicer) Mtrc(lab string) int64 {
	var cou int64

	for _, x := range s {
		cou += x.Mtrc.Data[lab]
	}

	return cou
}

func (s Slicer) Obct() []*Object {
	return s
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

// User returns all the user IDs for the underling list of event objects.
func (s Slicer) User() []objectid.ID {
	var ids []objectid.ID

	for _, x := range s {
		ids = append(ids, x.User)
	}

	return ids
}
