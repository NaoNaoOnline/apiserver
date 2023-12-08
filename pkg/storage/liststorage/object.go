package liststorage

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/format/descriptionformat"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/xh3b4sd/tracer"
)

type Object struct {
	// Crtd is the time at which the list got created.
	Crtd time.Time `json:"crtd"`
	// Dltd is the time at which the list got deleted.
	Dltd time.Time `json:"dltd,omitempty"`
	// Desc is the list's description.
	Desc objectfield.String `json:"desc"`
	// Feed is the time until the list's feed got consumed the last time. If this
	// timestamp is zero, it means the list's notifications are disabled.
	Feed time.Time `json:"feed,omitempty"`
	// List is the ID of the list being created.
	List objectid.ID `json:"list"`
	// User is the user ID creating this list.
	User objectid.ID `json:"user"`
}

func (o *Object) Verify() error {
	{
		if o.Desc.Data == "" {
			return tracer.Mask(listDescEmptyError)
		}
		if !descriptionformat.Verify(o.Desc.Data) {
			return tracer.Maskf(listDescFormatError, o.Desc.Data)
		}
		if len(o.Desc.Data) < 2 {
			return tracer.Maskf(listDescLengthError, "%d", len(o.Desc.Data))
		}
		if len(o.Desc.Data) > 40 {
			return tracer.Maskf(listDescLengthError, "%d", len(o.Desc.Data))
		}
	}

	{
		if o.User == "" {
			return tracer.Mask(runtime.UserIDEmptyError)
		}
	}

	return nil
}
