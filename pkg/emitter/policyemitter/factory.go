package policyemitter

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/rescue"
)

func Fake() Interface {
	return NewEmitter(EmitterConfig{
		Cid: []int64{1},
		Cnt: []string{"0x0"},
		Log: logger.Fake(),
		Res: rescue.Fake(),
		Rpc: []string{"127.0.0.1:8545"},
	})
}
