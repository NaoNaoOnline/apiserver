package liststorage

import (
	"regexp"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/tracer"
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

var (
	descexpr = regexp.MustCompile(`^([A-Za-z0-9\s,.:\-'"!$%&#]+(?:\s*,\s*[A-Za-z0-9\s,.:\-'"!$%&#]+)*)$`)
)

func (o *Object) Verify() error {
	{
		if o.Desc == "" {
			return tracer.Mask(listDescEmptyError)
		}
		if !descexpr.MatchString(o.Desc) {
			return tracer.Maskf(listDescFormatError, o.Desc)
		}
		if len(o.Desc) < 2 {
			return tracer.Maskf(listDescLengthError, "%d", len(o.Desc))
		}
		if len(o.Desc) > 40 {
			return tracer.Maskf(listDescLengthError, "%d", len(o.Desc))
		}
	}

	return nil
}
