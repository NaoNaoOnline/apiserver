package userhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/user"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	fuzz "github.com/google/gofuzz"
	"github.com/xh3b4sd/logger"
)

func tesCtx() context.Context {
	var str string
	fuzz.New().NilChance(0.5).Fuzz(&str)
	return userid.NewContext(context.Background(), objectid.ID(str))
}

func tesHan() user.API {
	return &wrapper{
		han: NewHandler(HandlerConfig{
			Log: logger.Fake(),
			Use: userstorage.Fake(),
		}),
	}
}
