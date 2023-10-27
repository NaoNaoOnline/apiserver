package rulestorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
)

type Object struct {
	// Crtd is the time at which the rule got created.
	Crtd time.Time `json:"crtd"`
	// Dltd is the time at which the rule got deleted.
	Dltd time.Time `json:"dltd,omitempty"`
	// Excl is the rule's object ID to remove from the associated list, if any.
	Excl []objectid.ID `json:"excl"`
	// Incl is the rule's object ID to add to the associated list, if any.
	Incl []objectid.ID `json:"incl"`
	// Kind is the rule's object type defining the resource for included and
	// excluded object IDs. kind set to "host" includes and excludes the
	// respective host label IDs when computing the associated list.
	//
	//     cate for adding or removing events matching the given category IDs
	//     host for adding or removing events matching the given host IDs
	//     user for adding or removing events created by the given user IDs
	//
	Kind string `json:"kind"`
	// Rule is the ID of the rule being created.
	Rule objectid.ID `json:"rule"`
	// User is the user ID creating this rule.
	User objectid.ID `json:"user"`
}

func (o *Object) KeyFmt() string {
	if o.Kind == "cate" || o.Kind == "host" {
		return keyfmt.EventLabel
	}

	return keyfmt.EventUser
}

func (o *Object) Verify() error {
	return nil
}
