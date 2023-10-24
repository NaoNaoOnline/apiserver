package descriptionstorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/emitter/descriptionemitter"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
)

func Fake() Interface {
	return NewRedis(RedisConfig{
		Emi: descriptionemitter.Fake(),
		Log: logger.Fake(),
		Red: redigo.Fake(),
	})
}
