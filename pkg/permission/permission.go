package permission

import (
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/cache/policycache"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Log logger.Interface
	Pol policycache.Interface
	Wal walletstorage.Interface
}

type Permission struct {
	log logger.Interface
	pol policycache.Interface
	wal walletstorage.Interface
}

func New(c Config) *Permission {
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
		log: c.Log,
		pol: c.Pol,
		wal: c.Wal,
	}
}
