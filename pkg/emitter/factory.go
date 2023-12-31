package emitter

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/rescue"
)

func Fake() *Emitter {
	return New(Config{
		Cid: []int64{1},
		Log: logger.Fake(),
		Pcn: []string{"0x0"},
		Res: rescue.Fake(),
		Rpc: []string{"127.0.0.1:8545"},
		Scn: []string{"0x0"},
	})
}
