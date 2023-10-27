package descriptionemitter

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

type Interface interface {
	DeleteDesc(objectid.ID) error
}
