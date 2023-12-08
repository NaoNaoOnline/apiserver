package feedemitter

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

type Interface interface {
	Create(objectid.ID, objectid.ID, string) error
}
