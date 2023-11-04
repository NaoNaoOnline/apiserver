package rulestorage

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

type Slicer []*Object

// Cate returns the category label IDs to be excluded from the list of events
// that the given rule set describes.
func (s Slicer) Cate() []objectid.ID {
	var cat []objectid.ID

	for _, x := range s {
		if x.Kind == "cate" {
			cat = append(cat, x.Excl...)
		}
	}

	return cat
}

// Host returns the host label IDs to be excluded from the list of events that
// the given rule set describes.
func (s Slicer) Host() []objectid.ID {
	var hos []objectid.ID

	for _, x := range s {
		if x.Kind == "host" {
			hos = append(hos, x.Excl...)
		}
	}

	return hos
}

// Incl returns the storage keys pointing to the event IDs meant to be included
// in the list associated to the underlying rules.
func (s Slicer) Incl() []string {
	var inc []string

	for _, x := range s {
		inc = append(inc, objectid.Fmt(x.Incl, x.KeyFmt())...)
	}

	return inc
}

// Like returns the user IDs to be excluded from the list of events that the
// given rule set describes.
func (s Slicer) Like() []objectid.ID {
	var use []objectid.ID

	for _, x := range s {
		if x.Kind == "like" {
			use = append(use, x.Excl...)
		}
	}

	return use
}

func (s Slicer) Rule() []objectid.ID {
	var ids []objectid.ID

	for _, x := range s {
		ids = append(ids, x.Rule)
	}

	return ids
}

// User returns the user IDs to be excluded from the list of events that the
// given rule set describes.
func (s Slicer) User() []objectid.ID {
	var use []objectid.ID

	for _, x := range s {
		if x.Kind == "user" {
			use = append(use, x.Excl...)
		}
	}

	return use
}
