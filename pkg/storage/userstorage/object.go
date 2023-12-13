package userstorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/format/nameformat"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
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
	// Prem is the time until the user got a valid premium subscription, if any.
	Prem time.Time `json:"prem,omitempty"`
	// Prfl is the map of external accounts related to this user. These accounts
	// may point to references about this user on other platforms.
	Prfl objectfield.MapStr `json:"prfl"`
	// Sclm is the list of external subject claims mapped to the user being
	// created.
	Sclm []string `json:"sclm"`
	// User is the internal ID of the user being created.
	User objectid.ID `json:"user"`
}

func (o *Object) HasPre() bool {
	return !o.Prem.IsZero() && time.Now().UTC().Before(o.Prem)
}

func (o *Object) ProPat() []string {
	var pat []string

	for k := range o.Prfl.Data {
		pat = append(pat, "/prfl/data/"+k)
	}

	return pat
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
		for k, v := range o.Prfl.Data {
			if !vldPrfl(k) {
				return tracer.Maskf(userPrflInvalidError, k)
			}
			if !nameformat.Verify(v) {
				return tracer.Maskf(userPrflFormatError, v)
			}
		}
	}

	{
		if len(o.Sclm) != 1 || o.Sclm[0] == "" {
			return tracer.Mask(userSubjectEmptyError)
		}
	}

	// Note that Object.User is not validated here like for the other resources,
	// because the user ID is the resource ID for the user object. The user ID is
	// set in UserStorage.Create, as opposed to being injected from the outside
	// like most of the other storage implementations work.

	return nil
}

func vldPrfl(key string) bool {
	for _, x := range objectlabel.SearchPrfl() {
		if key == x {
			return true
		}
	}

	return false
}
