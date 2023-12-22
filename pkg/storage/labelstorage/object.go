package labelstorage

import (
	"strings"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/format/descriptionformat"
	"github.com/NaoNaoOnline/apiserver/pkg/format/handleformat"
	"github.com/NaoNaoOnline/apiserver/pkg/format/nameformat"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/xh3b4sd/tracer"
)

type Object struct {
	// Crtd is the time at which the label got created.
	Crtd time.Time `json:"crtd"`
	// Dltd is the time at which the label got deleted.
	Dltd time.Time `json:"dltd,omitempty"`
	// Desc is the label's description.
	Desc objectfield.String `json:"desc"`
	// Kind is the label type.
	//
	//     bltn for system labels
	//     cate for category labels
	//     host for host labels
	//
	Kind string `json:"kind"`
	// Labl is the ID of the label being created.
	Labl objectid.ID `json:"labl"`
	// Name is the label name.
	Name objectfield.String `json:"name"`
	// Prfl is the map of external accounts related to this label. These accounts
	// may point to references about this label or to the label owner on other
	// platforms.
	Prfl objectfield.MapStr `json:"prfl"`
	// User is the user ID creating this label, or the user ID owning this label
	// after ownership transferal. Because labels are transferable. user IDs are
	// not just object IDs but objectfield IDs, in order to reflect update
	// timestamps.
	User objectfield.ID `json:"user"`
}

func (o *Object) ProPat() []string {
	var pat []string

	for k := range o.Prfl.Data {
		pat = append(pat, "/prfl/data/"+k)
	}

	return pat
}

func (o *Object) Verify() error {
	{
		txt := strings.TrimSpace(o.Desc.Data)

		// Label description might just be empty.
		if txt != "" && !descriptionformat.Verify(txt) {
			return tracer.Mask(labelDescFormatError)
		}
	}

	{
		if o.Kind != "bltn" && o.Kind != "cate" && o.Kind != "host" {
			return tracer.Maskf(labelKindInvalidError, o.Kind)
		}
	}

	{
		if o.Name.Data == "" {
			return tracer.Mask(labelNameEmptyError)
		}
		if !nameformat.Verify(o.Name.Data) {
			return tracer.Maskf(labelNameFormatError, o.Name.Data)
		}
		if len(o.Name.Data) < 2 {
			return tracer.Maskf(labelNameLengthError, "%s (%d)", o.Name.Data, len(o.Name.Data))
		}
		if len(o.Name.Data) > 25 {
			return tracer.Maskf(labelNameLengthError, "%s (%d)", o.Name.Data, len(o.Name.Data))
		}
	}

	{
		for k, v := range o.Prfl.Data {
			if !vldPrfl(k) {
				return tracer.Maskf(labelPrflInvalidError, k)
			}
			if !handleformat.Verify(v) {
				return tracer.Maskf(labelPrflFormatError, v)
			}
		}
	}

	{
		if o.User.Data == "" {
			return tracer.Mask(runtime.UserIDEmptyError)
		}
	}

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
