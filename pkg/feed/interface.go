package feed

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/rulestorage"
)

// Interface describes the unified Feed component used for cross-referencing
// event objects and rule objects. While there are many events and many rules,
// users want to have an aggregated view of all events that they describe with
// their custom lists. Our objective therefore is to do as much work as possible
// asynchronously at the time of resource creation, so that search queries can
// be done with the lowest possible latency. We do all the expensive writes
// upfront in order to efficiently fan-out all the cheap reads on demand. Some
// more key facts.
//
//   - There are many users, events, lists and rules.
//   - Lists are the user facing feature.
//   - Feeds are the underlying data structure.
//   - Lists aggregate events described by the criteria of its rules.
//   - Rules define includes and excludes of various types.
//   - Lists are static when they only define the rule kind [evnt].
//   - Lists are dynamic when they define any of rule kind [cate host user].
//   - Feeds are generated during creation time of events and rules.
//   - Clients receive all includes and have to filter out excludes.
//
// The reason for letting clients do the work of excluding events is based on
// the assumption that the vast majority of events within a custom list are
// described by includes. It follows that excludes represent the vast minority,
// allowing us to accept the negligible egress bandwidth added when simply
// responding with all includes of any given feed.
type Interface interface {
	// CreateEvnt links the given event ID with all the rules describing the given
	// event's properties. CreateEvnt must be called after the event storage
	// persisted the event object initially. CreateRule and CreateEvnt synchronize
	// by cross referencing rules and events.
	//
	//     @inp[0] the new event object
	//
	CreateEvnt(*eventstorage.Object) error

	// CreateFeed generates a single sorted set containing all event IDs that the
	// given list ID describes using its associated rules. With a maximum of 3
	// Redis calls, CreateFeed enables search queries for custom lists to be
	// highly performant. Additionally, we gain the ability to compute a delta of
	// the event objects already seen by the list owner. That is, the ability to
	// provide notification features. CreateFeed must be called after either of
	// CreateEvnt or CreateRule got called during event or rule creation.
	//
	//     @inp[0] the list ID to generate a feed for
	//
	CreateFeed(objectid.ID) error

	// CreateRule links the given rule ID with all the events that the given
	// rule's criteria describe. CreateRule must be called after the rule storage
	// persisted the rule object initially. CreateRule and CreateEvnt synchronize
	// by cross referencing rules and events.
	//
	//     @inp[0] the new rule object
	//
	CreateRule(*rulestorage.Object) error

	// DeleteEvnt reverts the effects of CreateEvnt.
	//
	//     @inp[0] the event object to delete
	//
	DeleteEvnt(*eventstorage.Object) error

	// DeleteFeed reverts the effects of CreateFeed.
	//
	//     @inp[0] the list ID to delete
	//
	DeleteFeed(objectid.ID) error

	// DeleteRule reverts the effects of CreateRule.
	//
	//     @inp[0] the rule object to delete
	//
	DeleteRule(*rulestorage.Object) error

	// SearchEvnt returns all event IDs for the given rule ID. All event IDs can be
	// fetched using pagination range [0 -1]. The first 10 event IDs can be
	// fetched using pagination range [0 9].
	//
	//     @inp[0] the rule ID to search event IDs for
	//     @inp[1] the pagination range defining lower and upper inclusive boundaries
	//     @out[0] the list of event IDs linked to the given rule ID
	//
	SearchEvnt(objectid.ID, [2]int) ([]objectid.ID, error)

	// SearchFeed returns all event IDs for the given list ID. SearchFeed
	// essentially provides the aggregated list view, the feed. All event IDs can
	// be fetched using pagination range [0 -1]. The first 10 event IDs can be
	// fetched using pagination range [0 9].
	//
	//     @inp[0] the list ID to search event IDs for
	//     @inp[1] the pagination range defining lower and upper inclusive boundaries
	//     @out[0] the list of event IDs linked to the given list ID
	//
	SearchFeed(objectid.ID, [2]int) ([]objectid.ID, error)

	// SearchList returns all list IDs for the given event ID. All list IDs can be
	// fetched using pagination range [0 -1]. The first 10 list IDs can be fetched
	// using pagination range [0 9].
	//
	//     @inp[0] the event ID to search list IDs for
	//     @inp[1] the pagination range defining lower and upper inclusive boundaries
	//     @out[0] the list of rule IDs linked to the given event ID
	//
	SearchList(objectid.ID, [2]int) ([]objectid.ID, error)

	// SearchRule returns all rule IDs for the given event ID. All rule IDs can be
	// fetched using pagination range [0 -1]. The first 10 rule IDs can be fetched
	// using pagination range [0 9].
	//
	//     @inp[0] the event ID to search rule IDs for
	//     @inp[1] the pagination range defining lower and upper inclusive boundaries
	//     @out[0] the list of rule IDs linked to the given event ID
	//
	SearchRule(objectid.ID, [2]int) ([]objectid.ID, error)
}
