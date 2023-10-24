package descriptionemitter

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/rescue"
)

func Fake() Interface {
	return NewEmitter(EmitterConfig{
		Log: logger.Fake(),
		Res: rescue.Fake(),
	})
}
