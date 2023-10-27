package rulehandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/rule"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/liststorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/rulestorage"
	fuzz "github.com/google/gofuzz"
	"github.com/xh3b4sd/logger"
)

func tesCtx() context.Context {
	var str string
	fuzz.New().NilChance(0.5).Fuzz(&str)
	return userid.NewContext(context.Background(), objectid.ID(str))
}

func tesHan() rule.API {
	return &wrapper{
		han: NewHandler(HandlerConfig{
			Lis: liststorage.Fake(),
			Log: logger.Fake(),
			Rul: rulestorage.Fake(),
		}),
	}
}
