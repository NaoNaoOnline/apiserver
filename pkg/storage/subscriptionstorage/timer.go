package subscriptionstorage

import "time"

type Timer interface {
	Now() time.Time
}

type faker struct {
	tim time.Time
}

func (f *faker) Now() time.Time {
	return f.tim
}

type timer struct{}

func (f *timer) Now() time.Time {
	return time.Now().UTC()
}
