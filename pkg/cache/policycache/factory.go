package policycache

import (
	"github.com/xh3b4sd/logger"
)

func Fake() Interface {
	return NewMemory(MemoryConfig{
		Log: logger.Fake(),
	})
}
