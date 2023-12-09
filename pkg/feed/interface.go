package feed

import (
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/rulestorage"
)

// TODO
type Interface interface {
	CreateEvnt(*eventstorage.Object) error
	CreateFeed(objectid.ID) error
	CreateRule(*rulestorage.Object) error

	DeleteEvnt(*eventstorage.Object) error
	DeleteFeed(objectid.ID) error
	DeleteRule(*rulestorage.Object) error

	SearchEvnt(objectid.ID, [2]int) ([]objectid.ID, error)
	SearchFeed(objectid.ID, [2]int) ([]objectid.ID, error)
	SearchList(objectid.ID, [2]int) ([]objectid.ID, error)
	SearchRule(objectid.ID, [2]int) ([]objectid.ID, error)
}
