package objectfield

import (
	"time"
)

type Strings struct {
	// Data is the string data of this object field.
	Data []string `json:"data"`
	// Time is the most recent time at which this object field got updated.
	Time time.Time `json:"time"`
}
