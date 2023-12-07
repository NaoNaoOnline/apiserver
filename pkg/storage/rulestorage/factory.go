package rulestorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/storage/notificationstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
)

func Fake() Interface {
	return NewRedis(RedisConfig{
		Log: logger.Fake(),
		Not: notificationstorage.Fake(),
		Red: redigo.Fake(),
	})
}
