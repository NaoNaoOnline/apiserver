package rulestorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
)

type Filter []*Object

// Cate returns the category label IDs to be excluded from the list of events
// that the given rule set describes.
func (f Filter) Cate() []objectid.ID {
	var cat []objectid.ID

	for _, x := range f {
		if x.Kind == "cate" {
			cat = append(cat, x.Excl...)
		}
	}

	return cat
}

// Evnt returns the event IDs to be excluded from the list of events that the
// given rule set describes.
func (f Filter) Evnt() []objectid.ID {
	var cat []objectid.ID

	for _, x := range f {
		if x.Kind == "evnt" {
			cat = append(cat, x.Excl...)
		}
	}

	return cat
}

// Host returns the host label IDs to be excluded from the list of events that
// the given rule set describes.
func (f Filter) Host() []objectid.ID {
	var hos []objectid.ID

	for _, x := range f {
		if x.Kind == "host" {
			hos = append(hos, x.Excl...)
		}
	}

	return hos
}

// Like returns the user IDs to be excluded from the list of events that the
// given rule set describes.
func (f Filter) Like() []objectid.ID {
	var use []objectid.ID

	for _, x := range f {
		if x.Kind == "like" {
			use = append(use, x.Excl...)
		}
	}

	return use
}

// User returns the user IDs to be excluded from the list of events that the
// given rule set describes.
func (f Filter) User() []objectid.ID {
	var use []objectid.ID

	for _, x := range f {
		if x.Kind == "user" {
			use = append(use, x.Excl...)
		}
	}

	return use
}
