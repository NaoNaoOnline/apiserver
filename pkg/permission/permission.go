package permission

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/cache/policycache"
	"github.com/NaoNaoOnline/apiserver/pkg/emitter/policyemitter"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/policystorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

const (
	SystemEvnt int64 = 0
	SystemDesc int64 = 1
)

const (
	AccessDelete int64 = 1
)

type Config struct {
	Cac policycache.Interface
	Emi policyemitter.Interface
	Log logger.Interface
	Pol policystorage.Interface
	Wal walletstorage.Interface
}

type Permission struct {
	cac policycache.Interface
	emi policyemitter.Interface
	log logger.Interface
	pol policystorage.Interface
	wal walletstorage.Interface
}

func New(c Config) *Permission {
	if c.Cac == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Cac must not be empty", c)))
	}
	if c.Emi == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Emi must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Pol == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Pol must not be empty", c)))
	}
	if c.Wal == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Wal must not be empty", c)))
	}

	return &Permission{
		cac: c.Cac,
		emi: c.Emi,
		log: c.Log,
		pol: c.Pol,
		wal: c.Wal,
	}
}
