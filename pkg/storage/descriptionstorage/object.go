package descriptionstorage

import (
	"regexp"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
	"github.com/xh3b4sd/tracer"
)

var (
	textexpr = regexp.MustCompile(`^([A-Za-z0-9\s,.:\-'"!$%&#]+(?:\s*,\s*[A-Za-z0-9\s,.:\-'"!$%&#]+)*)$`)
)

type Object struct {
	// Crtd is the time at which the description got created.
	Crtd time.Time `json:"crtd"`
	// Desc is the ID of the description being created.
	Desc objectid.String `json:"desc"`
	// Evnt is the event ID this description is mapped to.
	Evnt objectid.String `json:"evnt"`
	// Text is the description explaining what an event is about.
	Text string `json:"text"`
	// User is the user ID creating this description.
	User objectid.String `json:"user"`
}

func (o *Object) Verify() error {
	if o.Evnt == "" {
		return tracer.Mask(eventIDEmptyError)
	}

	if o.Text == "" {
		return tracer.Mask(descriptionTextEmptyError)
	}
	if !textexpr.MatchString(o.Text) {
		return tracer.Mask(descriptionTextFormatError)
	}
	if len(o.Text) < 20 {
		return tracer.Maskf(descriptionTextLengthError, "%d", len(o.Text))
	}
	if len(o.Text) > 120 {
		return tracer.Maskf(descriptionTextLengthError, "%d", len(o.Text))
	}

	return nil
}
