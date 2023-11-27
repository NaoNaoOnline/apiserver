package subscriptionstorage

import (
	"encoding/json"
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/emitter/subscriptionemitter"
	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type RedisConfig struct {
	Emi subscriptionemitter.Interface
	Log logger.Interface
	Red redigo.Interface
}

type Redis struct {
	emi subscriptionemitter.Interface
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

func subRec(use objectid.ID) string {
	return fmt.Sprintf(keyfmt.SubscriptionReceiver, use)
}

func subObj(use objectid.ID) string {
	return fmt.Sprintf(keyfmt.SubscriptionObject, use)
}

func subPay(use objectid.ID) string {
	return fmt.Sprintf(keyfmt.SubscriptionPayer, use)
}

func musStr(obj *Object) string {
	byt, err := json.Marshal(obj)
	if err != nil {
		tracer.Panic(tracer.Mask(err))
	}

	return string(byt)
}