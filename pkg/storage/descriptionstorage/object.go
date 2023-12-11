package descriptionstorage

import (
	"strings"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/format/descriptionformat"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/xh3b4sd/tracer"
	"mvdan.cc/xurls/v2"
)

type Object struct {
	// Crtd is the time at which the description got created.
	Crtd time.Time `json:"crtd"`
	// Dltd is the time at which the description got deleted.
	Dltd time.Time `json:"dltd,omitempty"`
	// Desc is the ID of the description being created.
	Desc objectid.ID `json:"desc"`
	// Evnt is the event ID this description is mapped to.
	Evnt objectid.ID `json:"evnt"`
	// Mtrc is a mapping of various metrics related to this event object.
	Mtrc objectfield.MapInt `json:"mtrc"`
	// Text is the description explaining what an event is about.
	Text objectfield.String `json:"text"`
	// User is the user ID creating this description.
	User objectid.ID `json:"user"`
}

var (
	relxed = xurls.Relaxed()
)

func (o *Object) Verify() error {
	{
		if o.Evnt == "" {
			return tracer.Mask(eventIDEmptyError)
		}
	}

	{
		if o.Mtrc.Data[objectlabel.DescriptionMetricPrem] < 0 {
			return tracer.Mask(descriptionPremNegativeError)
		}
		if o.Mtrc.Data[objectlabel.DescriptionMetricUser] < 0 {
			return tracer.Mask(descriptionUserNegativeError)
		}
	}

	{
		txt := strings.TrimSpace(o.Text.Data)

		if txt == "" {
			return tracer.Mask(descriptionTextEmptyError)
		}
		if !descriptionformat.Verify(txt) {
			return tracer.Mask(descriptionTextFormatError)
		}
		if len(txt) < 20 {
			return tracer.Maskf(descriptionTextLengthError, "%d", len(txt))
		}
		if len(txt) > 120 {
			return tracer.Maskf(descriptionTextLengthError, "%d", len(txt))
		}
		if relxed.FindString(o.Text.Data) != "" {
			return tracer.Mask(descriptionTextURLError)
		}
	}

	{
		if o.User == "" {
			return tracer.Mask(runtime.UserIDEmptyError)
		}
	}

	return nil
}
