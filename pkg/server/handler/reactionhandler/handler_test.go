package reactionhandler

import (
	"github.com/NaoNaoOnline/apigocode/pkg/reaction"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/reactionstorage"
	"github.com/xh3b4sd/logger"
)

func tesHan() reaction.API {
	return &wrapper{
		han: NewHandler(HandlerConfig{
			Log: logger.Fake(),
			Rct: reactionstorage.Fake(),
		}),
	}
}
