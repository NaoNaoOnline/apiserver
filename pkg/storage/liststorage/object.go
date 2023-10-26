package liststorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
)

type Object struct {
	// Crtd is the time at which the list got created.
	Crtd time.Time `json:"crtd"`
	// Dltd is the time at which the list got deleted.
	Dltd time.Time `json:"dltd,omitempty"`
	// Desc is the list's description.
	Desc string `json:"desc"`
	// List is the ID of the list being created.
	List objectid.ID `json:"list"`
	// User is the user ID creating this list.
	User objectid.ID `json:"user"`
}

func (o *Object) Verify() error {
	return nil
}
