package rulestorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
)

type Object struct {
	// Crtd is the time at which the rule got created.
	Crtd time.Time `json:"crtd"`
	// Dltd is the time at which the rule got deleted.
	Dltd time.Time `json:"dltd,omitempty"`
	// Rule is the ID of the rule being created.
	Rule objectid.ID `json:"rule"`
	// User is the user ID creating this rule.
	User objectid.ID `json:"user"`
}

func (o *Object) Verify() error {
	return nil
}
