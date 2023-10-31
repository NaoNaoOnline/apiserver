package descriptionstorage

import (
	"regexp"
	"strings"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/tracer"
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
	// Like is the number of likes this description received.
	Like objectfield.Integer `json:"like"`
	// Text is the description explaining what an event is about.
	Text string `json:"text"`
	// User is the user ID creating this description.
	User objectid.ID `json:"user"`
}

var (
	textexpr = regexp.MustCompile(`^([A-Za-z0-9\s,.:\-'"!$%&#]+(?:\s*,\s*[A-Za-z0-9\s,.:\-'"!$%&#]+)*)$`)
)

func (o *Object) Verify() error {
	{
		if o.Evnt == "" {
			return tracer.Mask(eventIDEmptyError)
		}
	}

	{
		if o.Like.Data < 0 {
			return tracer.Mask(descriptionLikeNegativeError)
		}
	}

	{
		txt := strings.TrimSpace(o.Text)

		if txt == "" {
			return tracer.Mask(descriptionTextEmptyError)
		}
		if !textexpr.MatchString(txt) {
			return tracer.Mask(descriptionTextFormatError)
		}
		if len(txt) < 20 {
			return tracer.Maskf(descriptionTextLengthError, "%d", len(txt))
		}
		if len(txt) > 120 {
			return tracer.Maskf(descriptionTextLengthError, "%d", len(txt))
		}
	}

	return nil
}
