package objectfield

import "time"

type MapStr struct {
	// Data is the map data of this object field.
	Data map[string]string `json:"data"`
	// Time is the most recent time at which this object field got updated.
	Time time.Time `json:"time"`
	// User a contextual flag set for the calling user on the fly if a certain
	// condition was found to be true.
	User map[string]bool `json:"-"`
}
