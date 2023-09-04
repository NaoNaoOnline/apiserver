package labelstorage

import (
	"regexp"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
	"github.com/xh3b4sd/tracer"
)

type Object struct {
	// Crtd is the time at which the label got created.
	Crtd time.Time `json:"crtd"`
	// Desc is the label's description.
	Desc string `json:"desc"`
	// Disc is the label's Discord link.
	Disc string `json:"disc"`
	// Kind is the label type, e.g. host for host labels and cate for category
	// labels.
	Kind string `json:"kind"`
	// Labl is the ID of the label being created.
	Labl objectid.String `json:"labl"`
	// Name is the label name.
	Name string `json:"name"`
	// Twit is the label's Twitter link.
	Twit string `json:"twit"`
	// User is the user ID creating this label.
	User objectid.String `json:"user"`
}

var (
	lablexpr = regexp.MustCompile(`^[A-Za-z0-9\s]+$`)
)

func (o *Object) Verify() error {
	if o.Kind != "cate" && o.Kind != "host" {
		return tracer.Maskf(invalidLabelKindError, o.Kind)
	}

	if o.Name == "" {
		return tracer.Mask(labelNameEmptyError)
	}
	if !lablexpr.MatchString(o.Name) {
		return tracer.Maskf(labelNameFormatError, o.Name)
	}

	if o.Desc != "" || o.Disc != "" || o.Twit != "" {
		return tracer.Mask(fieldUnsupportedError)
	}

	return nil
}
