//go:build redis

package cnf

import (
	"testing"

	"github.com/NaoNaoOnline/apiserver/pkg/feed"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/rulestorage"
	"github.com/xh3b4sd/redigo"
)

func Test_Redis_Feed_Create_Rule_Exclude(t *testing.T) {
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

	// Create rules. This should not fail even though we are only passing a rule
	// that defines excludes instead of includes.
	var rob []*rulestorage.Object
	{
		lis := []*rulestorage.Object{
			rulExc(lisOne().List),
		}

		rob, err = rul.CreateRule(lis)
		if err != nil {
			t.Fatal(err)
		}

		for _, x := range rob {
			err = fee.CreateRule(x)
			if err != nil {
				t.Fatal(err)
			}
		}
	}

	{
		err := fee.CreateFeed(lisOne().List)
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		eid, err := fee.SearchPage(lisOne().List, feed.PagAll())
		if err != nil {
			t.Fatal(err)
		}
		if len(eid) != 0 {
			t.Fatal("expected", 0, "got", len(eid))
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

		for _, x := range rob {
			err = fee.DeleteRule(x)
			if err != nil {
				t.Fatal(err)
			}
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
		eid, err := fee.SearchPage(lisOne().List, feed.PagAll())
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
		lis := []*rulestorage.Object{
			rulOne(lisOne().List),
		}

		rob, err = rul.CreateRule(lis)
		if err != nil {
			t.Fatal(err)
		}

		for _, x := range rob {
			err = fee.CreateRule(x)
			if err != nil {
				t.Fatal(err)
			}
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
		eid, err := fee.SearchPage(lisOne().List, feed.PagAll())
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

		for _, x := range rob {
			err = fee.DeleteRule(x)
			if err != nil {
				t.Fatal(err)
			}
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
		eid, err := fee.SearchPage(lisOne().List, feed.PagAll())
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
		lis := []*rulestorage.Object{
			rulOne(lisOne().List),
			rulTwo(lisOne().List),
			rulThr(lisOne().List),
		}

		rob, err = rul.CreateRule(lis)
		if err != nil {
			t.Fatal(err)
		}

		for _, x := range rob {
			err = fee.CreateRule(x)
			if err != nil {
				t.Fatal(err)
			}
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
		eid, err := fee.SearchPage(lisOne().List, feed.PagAll())
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

		for _, x := range rob {
			err = fee.DeleteRule(x)
			if err != nil {
				t.Fatal(err)
			}
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
		lis := []*rulestorage.Object{
			rulOne(lisOne().List),
			rulTwo(lisOne().List),
			rulThr(lisOne().List),
		}

		rob, err = rul.CreateRule(lis)
		if err != nil {
			t.Fatal(err)
		}

		for _, x := range rob {
			err = fee.CreateRule(x)
			if err != nil {
				t.Fatal(err)
			}
		}
	}

	{
		err := fee.CreateFeed(lisOne().List)
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		eid, err := fee.SearchPage(lisOne().List, feed.PagAll())
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
		eid, err := fee.SearchPage(lisOne().List, feed.PagAll())
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

	// Verify we are able to search for the event IDs within a certain time range.
	// Note that event IDs are based on time, and their score refers to their time
	// of creation. So if we remember when a user saw a custom list the last time,
	// we can search for the delta of missed events later on, basically enabling
	// notification features.
	{
		eid, err := fee.SearchUnix(lisOne().List, [2]float64{eveTwo().Evnt.Float(), eveOne().Evnt.Float()}) // [426989 944148]
		if err != nil {
			t.Fatal(err)
		}
		if len(eid) != 2 {
			t.Fatal("expected", 2, "got", len(eid))
		}
		if eid[0] != eveTwo().Evnt {
			t.Fatal("expected", eveTwo().Evnt, "got", eid[0])
		}
		if eid[1] != eveOne().Evnt {
			t.Fatal("expected", eveOne().Evnt, "got", eid[1])
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

		for _, x := range rob {
			err = fee.DeleteRule(x)
			if err != nil {
				t.Fatal(err)
			}
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

func Test_Redis_Feed_Search_Empty(t *testing.T) {
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
		eid, err := fee.SearchPage(lisOne().List, feed.PagAll())
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
		eid, err := fee.SearchPage(lisOne().List, feed.PagAll())
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
		lis := []*rulestorage.Object{
			rulOne(lisOne().List),
		}

		rob, err = rul.CreateRule(lis)
		if err != nil {
			t.Fatal(err)
		}

		for _, x := range rob {
			err = fee.CreateRule(x)
			if err != nil {
				t.Fatal(err)
			}
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
		eid, err := fee.SearchPage(lisOne().List, feed.PagAll())
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

		for _, x := range rob {
			err = fee.DeleteRule(x)
			if err != nil {
				t.Fatal(err)
			}
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
		lis := []*rulestorage.Object{
			rulOne(lisOne().List),
		}

		rob, err = rul.CreateRule(lis)
		if err != nil {
			t.Fatal(err)
		}

		for _, x := range rob {
			err = fee.CreateRule(x)
			if err != nil {
				t.Fatal(err)
			}
		}
	}

	{
		err := fee.CreateFeed(lisOne().List)
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		eid, err := fee.SearchPage(lisOne().List, feed.PagAll())
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
		eid, err := fee.SearchPage(lisOne().List, feed.PagAll())
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

		for _, x := range rob {
			err = fee.DeleteRule(x)
			if err != nil {
				t.Fatal(err)
			}
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
