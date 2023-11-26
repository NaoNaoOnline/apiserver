package subscriptionemitter

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

type Interface interface {
	Scrape(objectid.ID) error
}
