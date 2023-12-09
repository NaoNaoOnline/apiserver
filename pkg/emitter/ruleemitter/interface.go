package ruleemitter

import "github.com/NaoNaoOnline/apiserver/pkg/object/objectid"

type Interface interface {
	Create(objectid.ID) error
	Delete(objectid.ID) error
}
