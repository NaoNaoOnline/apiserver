package policycache

import (
	"fmt"
	"sync"

	"github.com/NaoNaoOnline/apiserver/pkg/storage/policystorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type MemoryConfig struct {
	Log logger.Interface
}

type Memory struct {
	cac []*policystorage.Object
	log logger.Interface
	mem map[string]struct{}
	mut sync.Mutex
	rec map[int64]map[string]*policystorage.Object
}

func NewMemory(c MemoryConfig) *Memory {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}

	return &Memory{
		cac: []*policystorage.Object{},
		log: c.Log,
		mem: map[string]struct{}{},
		mut: sync.Mutex{},
		rec: map[int64]map[string]*policystorage.Object{},
	}
}
