package feed

import (
	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/rulestorage"
	"github.com/xh3b4sd/redigo/pkg/sorted"
	"github.com/xh3b4sd/tracer"
)

// TODO lock with distributed lock
func (f *Feed) CreateEvnt(eob *eventstorage.Object) error {
	var err error

	// Write the given event ID to all of the event keys that we can generate
	// based on the given event's configuration. With that CreateRule can fetch
	// all relevant event IDs for any rule being created that describes its own
	// criteria to match the event created here. Note that in a lot of cases we
	// make a call for an event's criteria that already exists. In such a case we
	// simply continue with the next key, if any.
	for _, x := range eveWri(eob) {
		err = f.red.Sorted().Create().Score(x, eob.Evnt.String(), eob.Evnt.Float())
		if sorted.IsAlreadyExistsError(err) {
			continue
		} else if err != nil {
			return tracer.Mask(err)
		}
	}

	// get rule IDs using rule keys
	var val []string
	{
		val, err = f.red.Sorted().Search().Union(rulRea(eob)...)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// link event ID with rule IDs
	for _, x := range objectid.IDs(val) {
		// add event ID to all rule sorted sets
		err = f.red.Sorted().Create().Score(keyfmt.EveRul(x), eob.Evnt.String(), eob.Evnt.Float())
		if sorted.IsAlreadyExistsError(err) {
			continue
		} else if err != nil {
			return tracer.Mask(err)
		}

		// add all rule IDs to event sorted set
		err = f.red.Sorted().Create().Score(keyfmt.RulEve(eob.Evnt), x.String(), x.Float())
		if sorted.IsAlreadyExistsError(err) {
			continue
		} else if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

// TODO lock with distributed lock
func (f *Feed) CreateFeed(lid objectid.ID) error {
	var err error

	var pag [2]int
	{
		pag = PagAll()
	}

	// get all rule IDs for the given list
	var rid []string
	{
		rid, err = f.red.Sorted().Search().Order(keyfmt.RulLis(lid), pag[0], pag[1])
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// there might not be any rule IDs for this feed
	if len(rid) == 0 {
		return nil
	}

	// store event IDs using event keys
	var cou int64
	{
		cou, err = f.red.Sorted().Create().Union(keyfmt.EveFee(lid), fmtFnc(rid, eveRul)...)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// limit the maximum amount of events per feed
	if cou > 1000 {
		err = f.red.Sorted().Delete().Limit(keyfmt.EveFee(lid), 1000)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

// TODO lock with distributed lock
func (f *Feed) CreateRule(rob *rulestorage.Object) error {
	var err error

	// write rule ID to rule keys
	for _, x := range rulWri(rob) {
		err = f.red.Sorted().Create().Score(x, rob.Rule.String(), rob.Rule.Float())
		if sorted.IsAlreadyExistsError(err) {
			continue
		} else if err != nil {
			return tracer.Mask(err)
		}
	}

	// get event IDs using event keys
	var val []string
	{
		val, err = f.red.Sorted().Search().Union(eveRea(rob)...)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// link rule ID with event IDs
	for _, x := range objectid.IDs(val) {
		// add rule ID to all event sorted sets
		err = f.red.Sorted().Create().Score(keyfmt.RulEve(x), rob.Rule.String(), rob.Rule.Float())
		if sorted.IsAlreadyExistsError(err) {
			continue
		} else if err != nil {
			return tracer.Mask(err)
		}

		// add all event IDs to rule sorted set
		err = f.red.Sorted().Create().Score(keyfmt.EveRul(rob.Rule), x.String(), x.Float())
		if sorted.IsAlreadyExistsError(err) {
			continue
		} else if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
