package votestorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/tracer"
)

type Object struct {
	// Crtd is the time at which the vote got created.
	Crtd time.Time `json:"crtd"`
	// Desc is the ID of the description the user voted on.
	Desc objectid.ID `json:"desc"`
	// Evnt is the ID of the event the user voted on.
	Evnt objectid.ID `json:"evnt"`
	// Rctn is the ID of the reaction the user used to vote.
	Rctn objectid.ID `json:"rctn"`
	// User is the user ID creating this vote.
	User objectid.ID `json:"user"`
	// Vote is the ID of the vote being created.
	Vote objectid.ID `json:"vote"`
}

func (o *Object) Verify() error {
	if o.Desc == "" {
		return tracer.Mask(descriptionIDEmptyError)
	}

	if o.Rctn == "" {
		return tracer.Mask(reactionIDEmptyError)
	}

	return nil
}
