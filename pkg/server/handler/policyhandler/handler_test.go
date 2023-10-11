package policyhandler

import (
	"github.com/NaoNaoOnline/apigocode/pkg/policy"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/policystorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo/pkg/fake"
	"github.com/xh3b4sd/rescue/engine"
)

func tesHan() policy.API {
	return &wrapper{
		han: NewHandler(HandlerConfig{
			Log: logger.Fake(),
			Pol: policystorage.NewRedis(policystorage.RedisConfig{
				Log: logger.Fake(),
				Red: fake.New(),
				Wal: walletstorage.NewRedis(walletstorage.RedisConfig{
					Log: logger.Fake(),
					Red: fake.New(),
				}),
			}),
			Res: engine.New(engine.Config{
				Redigo: fake.New(),
			}),
		}),
	}
}
