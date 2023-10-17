package permission

import (
	"github.com/NaoNaoOnline/apiserver/pkg/cache/policycache"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/logger"
)

func Fake() Interface {
	return New(Config{
		Log: logger.Fake(),
		Pol: policycache.Fake(),
		Wal: walletstorage.Fake(),
	})
}
