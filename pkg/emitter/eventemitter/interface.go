package eventemitter

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

type Interface interface {
	CreateEvnt(objectid.ID) error
	CreateTwtr(objectid.ID) error
	DeleteEvnt(objectid.ID) error
}
