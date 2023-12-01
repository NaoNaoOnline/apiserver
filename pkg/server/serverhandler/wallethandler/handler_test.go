package wallethandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/wallet"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/subscriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
	fuzz "github.com/google/gofuzz"
	"github.com/xh3b4sd/logger"
)

func tesCtx() context.Context {
	var str string
	fuzz.New().NilChance(0.5).Fuzz(&str)
	return userid.NewContext(context.Background(), objectid.ID(str))
}

func tesHan() wallet.API {
	return &wrapper{
		han: NewHandler(HandlerConfig{
			Log: logger.Fake(),
			Wal: walletstorage.Fake(),
			Sub: subscriptionstorage.Fake(),
		}),
	}
}
