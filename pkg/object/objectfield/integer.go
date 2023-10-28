package objectfield

import (
	"time"
)

type Integer struct {
	// Data is the integer data of this object field.
	Data int `json:"data"`
	// User a contextual flag set for the calling user on the fly if a certain
	// condition was found to be true.
	User bool `json:"-"`
	// Time is the most recent time at which this object field got updated.
	Time time.Time `json:"time"`
}
