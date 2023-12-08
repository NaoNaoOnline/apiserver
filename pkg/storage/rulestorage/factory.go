package rulestorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/storage/feedstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
)

func Fake() Interface {
	return NewRedis(RedisConfig{
		Log: logger.Fake(),
		Fee: feedstorage.Fake(),
		Red: redigo.Fake(),
	})
}
