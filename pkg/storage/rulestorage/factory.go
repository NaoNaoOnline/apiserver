package rulestorage

import (
	"github.com/NaoNaoOnline/apiserver/pkg/emitter/ruleemitter"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
)

func Fake() Interface {
	return NewRedis(RedisConfig{
		Emi: ruleemitter.Fake(),
		Log: logger.Fake(),
		Red: redigo.Fake(),
	})
}
