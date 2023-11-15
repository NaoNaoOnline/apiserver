package objectfield

import "time"

type Map struct {
	// Data is the map data of this object field.
	Data map[string]String `json:"data"`
	// Time is the most recent time at which this object field got updated.
	Time time.Time `json:"time"`
}
