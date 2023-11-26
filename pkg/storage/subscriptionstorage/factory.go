package subscriptionstorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/emitter/subscriptionemitter"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
)

func Fake() Interface {
	return NewRedis(RedisConfig{
		Emi: subscriptionemitter.Fake(),
		Log: logger.Fake(),
		Red: redigo.Fake(),
	})
}
