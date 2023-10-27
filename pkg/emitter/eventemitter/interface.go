package eventemitter

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

type Interface interface {
	DeleteEvnt(objectid.ID) error
}
