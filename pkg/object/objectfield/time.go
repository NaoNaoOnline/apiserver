package objectfield

import "time"

type Time struct {
	// Data is the time data of this object field.
	Data time.Time `json:"data"`
	// Time is the most recent time at which this object field got updated.
	Time time.Time `json:"time"`
}
