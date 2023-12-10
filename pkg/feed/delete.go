package feed

import (
	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/rulestorage"
	"github.com/xh3b4sd/tracer"
)

func (f *Feed) DeleteEvnt(eob *eventstorage.Object) error {
	var err error

	// remove event ID from event keys
	for _, x := range eveWri(eob) {
		err = f.red.Sorted().Delete().Score(x, eob.Evnt.Float())
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var rid []objectid.ID
	{
		rid, err = f.SearchRule(eob.Evnt, PagAll())
		if err != nil {
			return tracer.Mask(err)
		}
	}

	for _, x := range rid {
		// remove event ID from all rule sorted sets
		err = f.red.Sorted().Delete().Score(keyfmt.EveRul(x), eob.Evnt.Float())
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// delete whole event sorted set
	{
		_, err = f.red.Simple().Delete().Multi(keyfmt.RulEve(eob.Evnt))
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (f *Feed) DeleteFeed(lid objectid.ID) error {
	var err error

	{
		_, err = f.red.Simple().Delete().Multi(keyfmt.EveFee(lid))
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (f *Feed) DeleteRule(rob *rulestorage.Object) error {
	var err error

	// remove rule ID from rule keys
	for _, x := range rulWri(rob) {
		err = f.red.Sorted().Delete().Score(x, rob.Rule.Float())
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var eid []objectid.ID
	{
		eid, err = f.SearchEvnt(rob.Rule, PagAll())
		if err != nil {
			return tracer.Mask(err)
		}
	}

	for _, x := range eid {
		// remove rule ID from all event sorted sets
		err = f.red.Sorted().Delete().Score(keyfmt.RulEve(x), rob.Rule.Float())
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// delete whole rule sorted set
	{
		_, err = f.red.Simple().Delete().Multi(keyfmt.EveRul(rob.Rule))
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
