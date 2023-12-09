//go:build redis

package cnf

import (
	"testing"

	"github.com/NaoNaoOnline/apiserver/pkg/feed"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/rulestorage"
	"github.com/xh3b4sd/redigo"
)

func Test_Redis_Feed_Duplicate(t *testing.T) {
	var err error

	var red redigo.Interface
	var fee feed.Interface
	var rul rulestorage.Interface
	{
		red, fee, rul = newSrv()
	}

	{
		emp, err := red.Empty()
		if err != nil {
			t.Fatal(err)
		}
		if !emp {
			t.Fatal("expected redis to be empty before test start")
		}
	}

	// Create events. Execute creation multiple times to ensure our business logic
	// is idempotent.
	{
		err = fee.CreateEvnt(eveOne())
		if err != nil {
			t.Fatal(err)
		}
		err = fee.CreateEvnt(eveOne())
		if err != nil {
			t.Fatal(err)
		}
		err = fee.CreateEvnt(eveTwo())
		if err != nil {
			t.Fatal(err)
		}
		err = fee.CreateEvnt(eveTwo())
		if err != nil {
			t.Fatal(err)
		}
		err = fee.CreateEvnt(eveThr())
		if err != nil {
			t.Fatal(err)
		}
		err = fee.CreateEvnt(eveThr())
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		err := fee.CreateFeed(lisOne().List)
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		eid, err := fee.SearchFeed(lisOne().List, feed.PagAll())
		if err != nil {
			t.Fatal(err)
		}
		if len(eid) != 0 {
			t.Fatal("expected", 0, "got", len(eid))
		}
	}

	// Create rules.
	var rob []*rulestorage.Object
	{
		rob, err = rul.CreateRule([]*rulestorage.Object{rulOne(lisOne().List)})
		if err != nil {
			t.Fatal(err)
		}
		err = fee.CreateRule(rob[0])
		if err != nil {
			t.Fatal(err)
		}
	}

	// Create feed. After events, lists and rules have been created, a must be
	// created in order to yield a single sorted set with all event IDs for the
	// respective list.
	{
		err := fee.CreateFeed(lisOne().List)
		if err != nil {
			t.Fatal(err)
		}
	}

	// Verify the correct event IDs got persisted in the list specific feed.
	{
		eid, err := fee.SearchFeed(lisOne().List, feed.PagAll())
		if err != nil {
			t.Fatal(err)
		}
		if len(eid) != 1 {
			t.Fatal("expected", 1, "got", len(eid))
		}
		if eid[0] != eveTwo().Evnt {
			t.Fatal("expected", eveTwo().Evnt, "got", eid[0])
		}
	}

	// Delete events.
	{
		err = fee.DeleteEvnt(eveOne())
		if err != nil {
			t.Fatal(err)
		}
		err = fee.DeleteEvnt(eveOne())
		if err != nil {
			t.Fatal(err)
		}
		err = fee.DeleteEvnt(eveTwo())
		if err != nil {
			t.Fatal(err)
		}
		err = fee.DeleteEvnt(eveTwo())
		if err != nil {
			t.Fatal(err)
		}
		err = fee.DeleteEvnt(eveThr())
		if err != nil {
			t.Fatal(err)
		}
		err = fee.DeleteEvnt(eveThr())
		if err != nil {
			t.Fatal(err)
		}
	}

	// Delete rules.
	{
		_, err = rul.DeleteRule(rob)
		if err != nil {
			t.Fatal(err)
		}
		err = fee.DeleteRule(rob[0])
		if err != nil {
			t.Fatal(err)
		}
	}

	// Delete feed.
	{
		err := fee.DeleteFeed(lisOne().List)
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		emp, err := red.Empty()
		if err != nil {
			t.Fatal(err)
		}
		if !emp {
			t.Fatal("expected redis to be empty after test end")
		}
	}
}

func Test_Redis_Feed_Empty_Search(t *testing.T) {
	var red redigo.Interface
	var fee feed.Interface
	{
		red, fee, _ = newSrv()
	}

	{
		emp, err := red.Empty()
		if err != nil {
			t.Fatal(err)
		}
		if !emp {
			t.Fatal("expected redis to be empty before test start")
		}
	}

	// Searching for a not existing or empty feed should not fail, but instead
	// yield an empty list.

	{
		lid, err := fee.SearchList(eveOne().Evnt, feed.PagAll())
		if err != nil {
			t.Fatal(err)
		}
		if len(lid) != 0 {
			t.Fatal("expected", 0, "got", len(lid))
		}
	}

	{
		eid, err := fee.SearchFeed(lisOne().List, feed.PagAll())
		if err != nil {
			t.Fatal(err)
		}
		if len(eid) != 0 {
			t.Fatal("expected", 0, "got", len(eid))
		}
	}

	{
		emp, err := red.Empty()
		if err != nil {
			t.Fatal(err)
		}
		if !emp {
			t.Fatal("expected redis to be empty after test end")
		}
	}
}

func Test_Redis_Feed_Multi_Event_First(t *testing.T) {
	var err error

	var red redigo.Interface
	var fee feed.Interface
	var rul rulestorage.Interface
	{
		red, fee, rul = newSrv()
	}

	{
		emp, err := red.Empty()
		if err != nil {
			t.Fatal(err)
		}
		if !emp {
			t.Fatal("expected redis to be empty before test start")
		}
	}

	// Create events.
	{
		err = fee.CreateEvnt(eveOne())
		if err != nil {
			t.Fatal(err)
		}
		err = fee.CreateEvnt(eveTwo())
		if err != nil {
			t.Fatal(err)
		}
		err = fee.CreateEvnt(eveThr())
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		err := fee.CreateFeed(lisOne().List)
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		eid, err := fee.SearchFeed(lisOne().List, feed.PagAll())
		if err != nil {
			t.Fatal(err)
		}
		if len(eid) != 0 {
			t.Fatal("expected", 0, "got", len(eid))
		}
	}

	// Create rules.
	var rob []*rulestorage.Object
	{
		rob, err = rul.CreateRule([]*rulestorage.Object{rulOne(lisOne().List), rulTwo(lisOne().List), rulThr(lisOne().List)})
		if err != nil {
			t.Fatal(err)
		}
		err = fee.CreateRule(rob[0])
		if err != nil {
			t.Fatal(err)
		}
		err = fee.CreateRule(rob[1])
		if err != nil {
			t.Fatal(err)
		}
	}

	// Create feed. After events, lists and rules have been created, a must be
	// created in order to yield a single sorted set with all event IDs for the
	// respective list.
	{
		err := fee.CreateFeed(lisOne().List)
		if err != nil {
			t.Fatal(err)
		}
	}

	// Verify the correct event IDs got persisted in the list specific feed. Note
	// that the event order is based on the ascending order of event IDs.
	{
		eid, err := fee.SearchFeed(lisOne().List, feed.PagAll())
		if err != nil {
			t.Fatal(err)
		}
		if len(eid) != 3 {
			t.Fatal("expected", 3, "got", len(eid))
		}
		if eid[0] != eveThr().Evnt {
			t.Fatal("expected", eveThr().Evnt, "got", eid[0])
		}
		if eid[1] != eveTwo().Evnt {
			t.Fatal("expected", eveTwo().Evnt, "got", eid[1])
		}
		if eid[2] != eveOne().Evnt {
			t.Fatal("expected", eveOne().Evnt, "got", eid[2])
		}
	}

	// Delete events.
	{
		err = fee.DeleteEvnt(eveOne())
		if err != nil {
			t.Fatal(err)
		}
		err = fee.DeleteEvnt(eveTwo())
		if err != nil {
			t.Fatal(err)
		}
		err = fee.DeleteEvnt(eveThr())
		if err != nil {
			t.Fatal(err)
		}
	}

	// Delete rules.
	{
		_, err = rul.DeleteRule(rob)
		if err != nil {
			t.Fatal(err)
		}
		err = fee.DeleteRule(rob[0])
		if err != nil {
			t.Fatal(err)
		}
		err = fee.DeleteRule(rob[1])
		if err != nil {
			t.Fatal(err)
		}
	}

	// Delete feed.
	{
		err := fee.DeleteFeed(lisOne().List)
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		emp, err := red.Empty()
		if err != nil {
			t.Fatal(err)
		}
		if !emp {
			t.Fatal("expected redis to be empty after test end")
		}
	}
}

func Test_Redis_Feed_Multi_Rule_First(t *testing.T) {
	var err error

	var red redigo.Interface
	var fee feed.Interface
	var rul rulestorage.Interface
	{
		red, fee, rul = newSrv()
	}

	{
		emp, err := red.Empty()
		if err != nil {
			t.Fatal(err)
		}
		if !emp {
			t.Fatal("expected redis to be empty before test start")
		}
	}

	// Create rules.
	var rob []*rulestorage.Object
	{
		rob, err = rul.CreateRule([]*rulestorage.Object{rulOne(lisOne().List), rulTwo(lisOne().List), rulThr(lisOne().List)})
		if err != nil {
			t.Fatal(err)
		}
		err = fee.CreateRule(rob[0])
		if err != nil {
			t.Fatal(err)
		}
		err = fee.CreateRule(rob[1])
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		err := fee.CreateFeed(lisOne().List)
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		eid, err := fee.SearchFeed(lisOne().List, feed.PagAll())
		if err != nil {
			t.Fatal(err)
		}
		if len(eid) != 0 {
			t.Fatal("expected", 0, "got", len(eid))
		}
	}

	// Create events.
	{
		err = fee.CreateEvnt(eveOne())
		if err != nil {
			t.Fatal(err)
		}
		err = fee.CreateEvnt(eveTwo())
		if err != nil {
			t.Fatal(err)
		}
		err = fee.CreateEvnt(eveThr())
		if err != nil {
			t.Fatal(err)
		}
	}

	// Create feed. After events, lists and rules have been created, a must be
	// created in order to yield a single sorted set with all event IDs for the
	// respective list.
	{
		err := fee.CreateFeed(lisOne().List)
		if err != nil {
			t.Fatal(err)
		}
	}

	// Verify the correct event IDs got persisted in the list specific feed. Note
	// that the event order is based on the ascending order of event IDs.
	{
		eid, err := fee.SearchFeed(lisOne().List, feed.PagAll())
		if err != nil {
			t.Fatal(err)
		}
		if len(eid) != 3 {
			t.Fatal("expected", 3, "got", len(eid))
		}
		if eid[0] != eveThr().Evnt {
			t.Fatal("expected", eveThr().Evnt, "got", eid[0])
		}
		if eid[1] != eveTwo().Evnt {
			t.Fatal("expected", eveTwo().Evnt, "got", eid[1])
		}
		if eid[2] != eveOne().Evnt {
			t.Fatal("expected", eveOne().Evnt, "got", eid[2])
		}
	}

	// Delete events.
	{
		err = fee.DeleteEvnt(eveOne())
		if err != nil {
			t.Fatal(err)
		}
		err = fee.DeleteEvnt(eveTwo())
		if err != nil {
			t.Fatal(err)
		}
		err = fee.DeleteEvnt(eveThr())
		if err != nil {
			t.Fatal(err)
		}
	}

	// Delete rules.
	{
		_, err = rul.DeleteRule(rob)
		if err != nil {
			t.Fatal(err)
		}
		err = fee.DeleteRule(rob[0])
		if err != nil {
			t.Fatal(err)
		}
		err = fee.DeleteRule(rob[1])
		if err != nil {
			t.Fatal(err)
		}
	}

	// Delete feed.
	{
		err := fee.DeleteFeed(lisOne().List)
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		emp, err := red.Empty()
		if err != nil {
			t.Fatal(err)
		}
		if !emp {
			t.Fatal("expected redis to be empty after test end")
		}
	}
}

func Test_Redis_Feed_Single_Event_First(t *testing.T) {
	var err error

	var red redigo.Interface
	var fee feed.Interface
	var rul rulestorage.Interface
	{
		red, fee, rul = newSrv()
	}

	{
		emp, err := red.Empty()
		if err != nil {
			t.Fatal(err)
		}
		if !emp {
			t.Fatal("expected redis to be empty before test start")
		}
	}

	// Create events.
	{
		err = fee.CreateEvnt(eveOne())
		if err != nil {
			t.Fatal(err)
		}
		err = fee.CreateEvnt(eveTwo())
		if err != nil {
			t.Fatal(err)
		}
		err = fee.CreateEvnt(eveThr())
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		err := fee.CreateFeed(lisOne().List)
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		eid, err := fee.SearchFeed(lisOne().List, feed.PagAll())
		if err != nil {
			t.Fatal(err)
		}
		if len(eid) != 0 {
			t.Fatal("expected", 0, "got", len(eid))
		}
	}

	// Create rules.
	var rob []*rulestorage.Object
	{
		rob, err = rul.CreateRule([]*rulestorage.Object{rulOne(lisOne().List)})
		if err != nil {
			t.Fatal(err)
		}
		err = fee.CreateRule(rob[0])
		if err != nil {
			t.Fatal(err)
		}
	}

	// Create feed. After events, lists and rules have been created, a must be
	// created in order to yield a single sorted set with all event IDs for the
	// respective list.
	{
		err := fee.CreateFeed(lisOne().List)
		if err != nil {
			t.Fatal(err)
		}
	}

	// Verify the correct event IDs got persisted in the list specific feed.
	{
		eid, err := fee.SearchFeed(lisOne().List, feed.PagAll())
		if err != nil {
			t.Fatal(err)
		}
		if len(eid) != 1 {
			t.Fatal("expected", 1, "got", len(eid))
		}
		if eid[0] != eveTwo().Evnt {
			t.Fatal("expected", eveTwo().Evnt, "got", eid[0])
		}
	}

	// Delete events.
	{
		err = fee.DeleteEvnt(eveOne())
		if err != nil {
			t.Fatal(err)
		}
		err = fee.DeleteEvnt(eveTwo())
		if err != nil {
			t.Fatal(err)
		}
		err = fee.DeleteEvnt(eveThr())
		if err != nil {
			t.Fatal(err)
		}
	}

	// Delete rules.
	{
		_, err = rul.DeleteRule(rob)
		if err != nil {
			t.Fatal(err)
		}
		err = fee.DeleteRule(rob[0])
		if err != nil {
			t.Fatal(err)
		}
	}

	// Delete feed.
	{
		err := fee.DeleteFeed(lisOne().List)
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		emp, err := red.Empty()
		if err != nil {
			t.Fatal(err)
		}
		if !emp {
			t.Fatal("expected redis to be empty after test end")
		}
	}
}

func Test_Redis_Feed_Single_Rule_First(t *testing.T) {
	var err error

	var red redigo.Interface
	var fee feed.Interface
	var rul rulestorage.Interface
	{
		red, fee, rul = newSrv()
	}

	{
		emp, err := red.Empty()
		if err != nil {
			t.Fatal(err)
		}
		if !emp {
			t.Fatal("expected redis to be empty before test start")
		}
	}

	// Create rules.
	var rob []*rulestorage.Object
	{
		rob, err = rul.CreateRule([]*rulestorage.Object{rulOne(lisOne().List)})
		if err != nil {
			t.Fatal(err)
		}
		err = fee.CreateRule(rob[0])
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		err := fee.CreateFeed(lisOne().List)
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		eid, err := fee.SearchFeed(lisOne().List, feed.PagAll())
		if err != nil {
			t.Fatal(err)
		}
		if len(eid) != 0 {
			t.Fatal("expected", 0, "got", len(eid))
		}
	}

	// Create events.
	{
		err = fee.CreateEvnt(eveOne())
		if err != nil {
			t.Fatal(err)
		}
		err = fee.CreateEvnt(eveTwo())
		if err != nil {
			t.Fatal(err)
		}
		err = fee.CreateEvnt(eveThr())
		if err != nil {
			t.Fatal(err)
		}
	}

	// Create feed. After events, lists and rules have been created, a must be
	// created in order to yield a single sorted set with all event IDs for the
	// respective list.
	{
		err := fee.CreateFeed(lisOne().List)
		if err != nil {
			t.Fatal(err)
		}
	}

	// Verify the correct event IDs got persisted in the list specific feed.
	{
		eid, err := fee.SearchFeed(lisOne().List, feed.PagAll())
		if err != nil {
			t.Fatal(err)
		}
		if len(eid) != 1 {
			t.Fatal("expected", 1, "got", len(eid))
		}
		if eid[0] != eveTwo().Evnt {
			t.Fatal("expected", eveTwo().Evnt, "got", eid[0])
		}
	}

	// Delete events.
	{
		err = fee.DeleteEvnt(eveOne())
		if err != nil {
			t.Fatal(err)
		}
		err = fee.DeleteEvnt(eveTwo())
		if err != nil {
			t.Fatal(err)
		}
		err = fee.DeleteEvnt(eveThr())
		if err != nil {
			t.Fatal(err)
		}
	}

	// Delete rules.
	{
		_, err = rul.DeleteRule(rob)
		if err != nil {
			t.Fatal(err)
		}
		err = fee.DeleteRule(rob[0])
		if err != nil {
			t.Fatal(err)
		}
	}

	// Delete feed.
	{
		err := fee.DeleteFeed(lisOne().List)
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		emp, err := red.Empty()
		if err != nil {
			t.Fatal(err)
		}
		if !emp {
			t.Fatal("expected redis to be empty after test end")
		}
	}
}
