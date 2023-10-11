package policyhandler

import (
	"github.com/NaoNaoOnline/apigocode/pkg/policy"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/policystorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/rescue"
)

func tesHan() policy.API {
	return &wrapper{
		han: NewHandler(HandlerConfig{
			Log: logger.Fake(),
			Pol: policystorage.Fake(),
			Res: rescue.Fake(),
		}),
	}
}
