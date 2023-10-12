package labelstorage

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
)

func Fake() Interface {
	return NewRedis(RedisConfig{
		Log: logger.Fake(),
		Red: redigo.Fake(),
	})
}
