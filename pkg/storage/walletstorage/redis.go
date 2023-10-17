package walletstorage

import (
	"encoding/json"
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type RedisConfig struct {
	Log logger.Interface
	Red redigo.Interface
}

type Redis struct {
	log logger.Interface
	red redigo.Interface
}

func NewRedis(c RedisConfig) *Redis {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Red == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Red must not be empty", c)))
	}

	return &Redis{
		log: c.Log,
		red: c.Red,
	}
}

func walAdd(add string) string {
	return fmt.Sprintf(keyfmt.WalletAddress, add)
}

func walKin(use objectid.ID, kin string) string {
	if kin == "eth" {
		return fmt.Sprintf(keyfmt.WalletEthereum, use)
	}

	panic(fmt.Sprintf("kin must be eth, got %s", kin))
}

func walObj(use objectid.ID, oid objectid.ID) string {
	return fmt.Sprintf(keyfmt.WalletObject, use, oid)
}

func walUse(use objectid.ID) string {
	return fmt.Sprintf(keyfmt.WalletUser, use)
}

func musStr(obj *Object) string {
	byt, err := json.Marshal(obj)
	if err != nil {
		tracer.Panic(tracer.Mask(err))
	}

	return string(byt)
}
