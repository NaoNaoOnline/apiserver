package descriptionstorage

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/rescue"
)

func Fake() Interface {
	return NewRedis(RedisConfig{
		Log: logger.Fake(),
		Red: redigo.Fake(),
		Res: rescue.Fake(),
	})
}
