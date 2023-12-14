package subscriptionstorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/emitter/subscriptionemitter"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
)

func Fake() Interface {
	return NewRedis(RedisConfig{
		Emi: subscriptionemitter.Fake(),
		Eve: eventstorage.Fake(),
		Log: logger.Fake(),
		Mse: 3,
		Msl: 3,
		Red: redigo.Fake(),
		Use: userstorage.Fake(),
		Wal: walletstorage.Fake(),
	})
}
