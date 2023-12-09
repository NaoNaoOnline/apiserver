package rulestorage

import (
	"encoding/json"
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/emitter/ruleemitter"
	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type RedisConfig struct {
	Emi ruleemitter.Interface
	Log logger.Interface
	Red redigo.Interface
}

type Redis struct {
	emi ruleemitter.Interface
	log logger.Interface
	red redigo.Interface
}

func NewRedis(c RedisConfig) *Redis {
	if c.Emi == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Emi must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Red == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Red must not be empty", c)))
	}

	return &Redis{
		emi: c.Emi,
		log: c.Log,
		red: c.Red,
	}
}

func rulObj(rul objectid.ID) string {
	return fmt.Sprintf(keyfmt.RuleObject, rul)
}

func rulUse(use objectid.ID) string {
	return fmt.Sprintf(keyfmt.RuleUser, use)
}

func musStr(obj *Object) string {
	byt, err := json.Marshal(obj)
	if err != nil {
		tracer.Panic(tracer.Mask(err))
	}

	return string(byt)
}
