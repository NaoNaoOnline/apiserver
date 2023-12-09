package eventhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/event"
	"github.com/NaoNaoOnline/apiserver/pkg/feed"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/permission"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/rulestorage"
	fuzz "github.com/google/gofuzz"
	"github.com/xh3b4sd/logger"
)

func tesCtx() context.Context {
	var str string
	fuzz.New().NilChance(0.5).Fuzz(&str)
	return userid.NewContext(context.Background(), objectid.ID(str))
}

func tesHan() event.API {
	return &wrapper{
		han: NewHandler(HandlerConfig{
			Eve: eventstorage.Fake(),
			Fee: feed.Fake(),
			Log: logger.Fake(),
			Prm: permission.Fake(),
			Rul: rulestorage.Fake(),
		}),
	}
}
