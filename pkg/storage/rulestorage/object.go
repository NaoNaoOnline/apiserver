package rulestorage

import (
	"fmt"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/xh3b4sd/tracer"
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
	// excluded object IDs. Kind set to "host" includes and excludes the
	// respective host label IDs when computing the associated list.
	//
	//     cate for adding or removing events matching the given category IDs
	//     evnt for adding or removing events matching the given event IDs
	//     host for adding or removing events matching the given host IDs
	//     user for adding or removing events created by the given user IDs
	//
	Kind string `json:"kind"`
	// List is the list ID this rule is mapped to.
	List objectid.ID `json:"list"`
	// Rule is the ID of the rule being created.
	Rule objectid.ID `json:"rule"`
	// User is the user ID creating this rule.
	User objectid.ID `json:"user"`
}

func (o *Object) HasRes() bool {
	return len(o.Excl) != 0 || len(o.Incl) != 0
}

func (o *Object) KeyFmt() string {
	if o.Kind == "cate" || o.Kind == "host" {
		return keyfmt.EventLabel
	}

	if o.Kind == "evnt" {
		return keyfmt.EventReference
	}

	if o.Kind == "user" {
		return keyfmt.EventUser
	}

	panic(fmt.Sprintf("invalid rule kind %#v in rule object %#v", o.Kind, o.Rule))
}

func (o *Object) RemRes(res objectid.ID) {
	o.Excl = remRes(o.Excl, res)
	o.Incl = remRes(o.Incl, res)
}

func (o *Object) Verify() error {
	{
		if o.Kind != "cate" && o.Kind != "evnt" && o.Kind != "host" && o.Kind != "user" {
			return tracer.Maskf(ruleKindInvalidError, o.Kind)
		}
	}

	{
		if len(o.Excl) == 0 && len(o.Incl) == 0 {
			return tracer.Mask(resourceIDEmptyError)
		}

		for _, x := range append(o.Excl, o.Incl...) {
			if x == "" {
				return tracer.Mask(resourceIDEmptyError)
			}
		}
	}

	{
		if o.List == "" {
			return tracer.Mask(ruleListEmptyError)
		}
	}

	{
		if o.User == "" {
			return tracer.Mask(runtime.UserIDEmptyError)
		}
	}

	return nil
}

func remRes(ids []objectid.ID, res objectid.ID) []objectid.ID {
	var lis []objectid.ID

	for _, x := range ids {
		if x != res {
			lis = append(lis, x)
		}
	}

	return lis
}
