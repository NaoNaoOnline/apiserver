package policyhandler

import (
	"context"

	"github.com/NaoNaoOnline/apigocode/pkg/policy"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/permission"
	"github.com/NaoNaoOnline/apiserver/pkg/server/context/userid"
	fuzz "github.com/google/gofuzz"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/rescue"
)

func tesCtx() context.Context {
	var str string
	fuzz.New().NilChance(0.5).Fuzz(&str)
	return userid.NewContext(context.Background(), objectid.ID(str))
}

func tesHan() policy.API {
	return &wrapper{
		han: NewHandler(HandlerConfig{
			Cid: []int64{1},
			Log: logger.Fake(),
			Prm: permission.Fake(),
			Res: rescue.Fake(),
		}),
	}
}
