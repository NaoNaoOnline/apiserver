package policycache

import (
	"fmt"
	"sync"

	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type MemoryConfig struct {
	Log logger.Interface
}

type Memory struct {
	buf map[int64][]*Record
	cac []*Record
	log logger.Interface
	mem map[string]struct{}
	mut sync.Mutex
	rec map[int64]map[string]*Record
}

func NewMemory(c MemoryConfig) *Memory {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}

	return &Memory{
		buf: map[int64][]*Record{},
		cac: []*Record{},
		log: c.Log,
		mem: map[string]struct{}{},
		mut: sync.Mutex{},
		rec: map[int64]map[string]*Record{},
	}
}
