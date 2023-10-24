package votehandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/vote"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/votestorage"
	fuzz "github.com/google/gofuzz"
	"github.com/xh3b4sd/logger"
)

func tesCtx() context.Context {
	var str string
	fuzz.New().NilChance(0.5).Fuzz(&str)
	return userid.NewContext(context.Background(), objectid.ID(str))
}

func tesHan() vote.API {
	return &wrapper{
		han: NewHandler(HandlerConfig{
			Des: descriptionstorage.Fake(),
			Eve: eventstorage.Fake(),
			Log: logger.Fake(),
			Vot: votestorage.Fake(),
		}),
	}
}
