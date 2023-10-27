package liststorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/emitter/listemitter"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
)

func Fake() Interface {
	return NewRedis(RedisConfig{
		Emi: listemitter.Fake(),
		Log: logger.Fake(),
		Red: redigo.Fake(),
	})
}
