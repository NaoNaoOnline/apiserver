package objectfield

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
)

type ID struct {
	// Data is the object ID data of this object field.
	Data objectid.ID `json:"data"`
	// Time is the most recent time at which this object field got updated.
	Time time.Time `json:"time"`
}
