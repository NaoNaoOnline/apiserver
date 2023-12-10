package feed

import (
	"fmt"

	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Log logger.Interface
	Red redigo.Interface
}

type Feed struct {
	log logger.Interface
	red redigo.Interface
}

func New(c Config) *Feed {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Red == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Red must not be empty", c)))
	}

	return &Feed{
		log: c.Log,
		red: c.Red,
	}
}
