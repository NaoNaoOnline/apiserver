package userstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/tracer"
)

type Object struct {
	// Crtd is the time at which the user got created.
	Crtd time.Time `json:"crtd"`
	// Imag is the URL pointing to the user's profile picture.
	Imag string `json:"imag"`
	// Name is the user name.
	Name string `json:"name"`
	// Subj is the list of external subject claims mapped to the user being
	// created.
	Subj []string `json:"subj"`
	// User is the internal ID of the user being created.
	User objectid.String `json:"user"`
}

func (o *Object) Verify() error {
	{
		if o.Imag == "" {
			return tracer.Mask(userImageEmptyError)
		}
	}

	{
		if o.Name == "" {
			return tracer.Mask(userNameEmptyError)
		}
		if len(o.Name) < 2 {
			return tracer.Maskf(userNameLengthError, "%d", len(o.Name))
		}
		if len(o.Name) > 30 {
			return tracer.Maskf(userNameLengthError, "%d", len(o.Name))
		}
	}

	{
		if len(o.Subj) != 1 || o.Subj[0] == "" {
			return tracer.Mask(userSubjectEmptyError)
		}
	}

	return nil
}
