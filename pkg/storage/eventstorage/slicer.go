package eventstorage

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

type Slicer []*Object

// Fltr returns a slicer implementation to remove certain objects from this
// list.
func (s Slicer) Fltr() Filter {
	return Filter(s)
}

// Clck returns the cumulative amount of link clicks for the underling list of
// event objects.
func (s Slicer) Clck() int64 {
	var clk int64

	for _, x := range s {
		clk += x.Clck.Data
	}

	return clk
}

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
