package eventstorage

import (
	"net/url"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
	"github.com/xh3b4sd/tracer"
)

type Object struct {
	// Cate is the list of label IDs under which the event is categorized.
	Cate []objectid.String `json:"cate"`
	// Crtd is the time at which the event got created.
	Crtd time.Time `json:"crtd"`
	// Dura is the estimated duration of the event.
	Dura time.Duration `json:"dura"`
	// Evnt is the ID of the event being created.
	Evnt objectid.String `json:"evnt"`
	// Host is the list of label IDs expected to host the event.
	Host []objectid.String `json:"host"`
	// Link is the online location at which the event is expected to take place.
	// For IRL events this may just be some informational website.
	Link string `json:"link"`
	// Time is the date time at which the event is expected to start.
	Time time.Time `json:"time"`
	// User is the user ID creating this event.
	User objectid.String `json:"user"`
}

func (o *Object) Verify() error {
	{
		if objectid.Dup(append(o.Cate, o.Host...)) {
			return tracer.Mask(eventLabelDuplicateError)
		}
	}

	{
		if len(o.Cate) == 0 {
			return tracer.Maskf(eventLabelEmptyError, "cate")
		}
		if len(o.Cate) > 5 {
			return tracer.Maskf(eventLabelLimitError, "cate")
		}
	}

	{
		if o.Dura == 0 {
			return tracer.Mask(eventDurationEmptyError)
		}
		if o.Dura < 0 {
			return tracer.Mask(eventDurationNegativeError)
		}
		if o.Dura > time.Duration(4)*time.Hour {
			return tracer.Mask(eventDurationLimitError)
		}
	}

	{
		if len(o.Host) == 0 {
			return tracer.Maskf(eventLabelEmptyError, "host")
		}
		if len(o.Host) > 5 {
			return tracer.Maskf(eventLabelLimitError, "host")
		}
	}

	{
		if o.Link == "" {
			return tracer.Mask(eventLinkEmptyError)
		}
		poi, err := url.Parse(o.Link)
		if err != nil {
			return tracer.Mask(eventLinkFormatError)
		}
		if poi.Scheme != "https" {
			return tracer.Mask(eventLinkFormatError)
		}
	}

	{
		if o.Time.IsZero() {
			return tracer.Mask(eventTimeEmptyError)
		}
		if o.Time.Compare(time.Now().UTC().Add(time.Hour*24*30)) == +1 {
			return tracer.Mask(eventTimeFutureError)
		}
		if o.Time.Compare(time.Now().UTC()) != +1 {
			return tracer.Mask(eventTimePastError)
		}
	}

	return nil
}

// tmvrlp returns whether o and x have a time overlap, based on their Time and
// Dura properties.
func (o *Object) tmvrlp(x *Object) bool {
	// Check if the first time range is entirely before the second.
	if o.Time.Add(o.Dura).Before(x.Time) || o.Time.Add(o.Dura).Equal(x.Time) {
		return false
	}

	// Check if the second time range is entirely before the first.
	if x.Time.Add(x.Dura).Before(o.Time) || x.Time.Add(x.Dura).Equal(o.Time) {
		return false
	}

	// If the above conditions are not met, the time ranges overlap.
	return true
}
