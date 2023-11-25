package subscriptionhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/subscription"
	"github.com/NaoNaoOnline/apiserver/pkg/emitter/subscriptionemitter"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/subscriptionstorage"
	fuzz "github.com/google/gofuzz"
	"github.com/xh3b4sd/locker"
	"github.com/xh3b4sd/logger"
)

func tesCtx() context.Context {
	var str string
	fuzz.New().NilChance(0.5).Fuzz(&str)
	return userid.NewContext(context.Background(), objectid.ID(str))
}

func tesHan() subscription.API {
	return &wrapper{
		han: NewHandler(HandlerConfig{
			Emi: subscriptionemitter.Fake(),
			Loc: locker.Fake(),
			Log: logger.Fake(),
			Sub: subscriptionstorage.Fake(),
		}),
	}
}
