package feed

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
)

func Fake() Interface {
	return New(Config{
		Log: logger.Fake(),
		Red: redigo.Fake(),
	})
}
