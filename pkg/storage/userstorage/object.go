package userstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/tracer"
)

const (
	oneWeek = time.Hour * 24 * 7
)

type Object struct {
	// Crtd is the time at which the user got created.
	Crtd time.Time `json:"crtd"`
	// Dltd is the time at which the user got deleted.
	Dltd time.Time `json:"dltd,omitempty"`
	// Home is the list ID describing a custom default view, optionally configured
	// by premium subscribers. The default "default view" is "/", that is, the
	// index page of the platform showing some derivative of the latest events
	// globally.
	Home objectfield.String `json:"home"`
	// Imag is the URL pointing to the user's profile picture.
	Imag string `json:"imag"`
	// Name is the user name.
	Name objectfield.String `json:"name"`
	// Subj is the list of external subject claims mapped to the user being
	// created.
	Subj []string `json:"subj"`
	// User is the internal ID of the user being created.
	User objectid.ID `json:"user"`
}

// UpdNam expresses whether the user name of this user object is allowed to be
// updated, based on the current time. User names are only allowed to be updated
// once within a time window of 7 days.
func (o *Object) UpdNam() bool {
	return time.Now().UTC().Sub(o.Name.Time) >= oneWeek
}

func (o *Object) Verify() error {
	{
		if o.Imag == "" {
			return tracer.Mask(userImageEmptyError)
		}
	}

	{
		if o.Name.Data == "" {
			return tracer.Mask(userNameEmptyError)
		}
		if len(o.Name.Data) < 2 {
			return tracer.Maskf(userNameLengthError, "%d", len(o.Name.Data))
		}
		if len(o.Name.Data) > 30 {
			return tracer.Maskf(userNameLengthError, "%d", len(o.Name.Data))
		}
	}

	{
		if len(o.Subj) != 1 || o.Subj[0] == "" {
			return tracer.Mask(userSubjectEmptyError)
		}
	}

	return nil
}
