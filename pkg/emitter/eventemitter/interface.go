package eventemitter

import (
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
)

type Interface interface {
	// CreateEvnt emits a task for the creation of the event object. Right at the
	// time of event creation, all feeds related to that new event must be updated
	// immediately.
	CreateEvnt(objectid.ID) error

	DeleteEvnt(objectid.ID) error

	// TickerEvnt emits a task for the time at which the event is scheduled to
	// happen. When this task is being executed, the associated event should be
	// happening right now.
	TickerEvnt(objectid.ID, time.Time) error
}
