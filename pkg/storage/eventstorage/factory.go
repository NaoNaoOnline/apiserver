package eventstorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/emitter/eventemitter"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
)

func Fake() Interface {
	return NewRedis(RedisConfig{
		Emi: eventemitter.Fake(),
		Log: logger.Fake(),
		Mse: 3,
		Msl: 3,
		Red: redigo.Fake(),
	})
}
