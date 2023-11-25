package permission

import (
	"github.com/NaoNaoOnline/apiserver/pkg/cache/policycache"
	"github.com/NaoNaoOnline/apiserver/pkg/emitter/policyemitter"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/policystorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/locker"
	"github.com/xh3b4sd/logger"
)

func Fake() Interface {
	return New(Config{
		Cac: policycache.Fake(),
		Emi: policyemitter.Fake(),
		Loc: locker.Fake(),
		Log: logger.Fake(),
		Pol: policystorage.Fake(),
		Wal: walletstorage.Fake(),
	})
}
