package reactionstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/tracer"
)

type Object struct {
	// Crtd is the time at which the reaction got created.
	Crtd time.Time `json:"crtd"`
	// Dltd is the time at which the reaction got deleted.
	Dltd time.Time `json:"dltd,omitempty"`
	// Html is the HTML of this reaction icon, e.g. some svg code.
	Html string `json:"html"`
	// Kind is the reaction type.
	//
	//     bltn for system reactions
	//     user for custom reactions
	//
	Kind string `json:"kind"`
	// Name is the reaction name.
	Name string `json:"name"`
	// Rctn is the ID of the reaction being created.
	Rctn objectid.ID `json:"rctn"`
	// User is the user ID creating this reaction.
	User objectid.ID `json:"user"`
}

func (o *Object) Verify() error {
	{
		if o.Html == "" {
			return tracer.Mask(reactionHtmlEmptyError)
		}
	}

	{
		if o.Kind != "bltn" && o.Kind != "user" {
			return tracer.Maskf(reactionKindInvalidError, o.Kind)
		}
	}

	{
		if o.Name == "" {
			return tracer.Mask(reactionNameEmptyError)
		}
	}

	return nil
}
