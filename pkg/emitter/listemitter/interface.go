package listemitter

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

type Interface interface {
	DeleteList(objectid.ID) error
}
