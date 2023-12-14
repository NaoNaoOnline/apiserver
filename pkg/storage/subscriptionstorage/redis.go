package subscriptionstorage

import (
	"encoding/json"
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/emitter/subscriptionemitter"
	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type RedisConfig struct {
	Emi subscriptionemitter.Interface
	Eve eventstorage.Interface
	Log logger.Interface
	// Mse is the minimum amount of events created required for users to be
	// considered legitimate content creators. This amount of events must have
	// been created for content creators to receive subscription fees.
	Mse int
	// Msl is the minimum amount of links clicked required for users to be
	// considered legitimate content creators. This amount of clicks must have
	// been generated for content creators to receive subscription fees.
	Msl int
	Red redigo.Interface
	Use userstorage.Interface
	Wal walletstorage.Interface
}

type Redis struct {
	emi subscriptionemitter.Interface
	eve eventstorage.Interface
	log logger.Interface
	mse int
	msl int
	red redigo.Interface
	use userstorage.Interface
	wal walletstorage.Interface
}

func NewRedis(c RedisConfig) *Redis {
	if c.Emi == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Emi must not be empty", c)))
	}
	if c.Eve == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Eve must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Mse == 0 {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Mse must not be empty", c)))
	}
	if c.Msl == 0 {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Msl must not be empty", c)))
	}
	if c.Red == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Red must not be empty", c)))
	}
	if c.Use == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Use must not be empty", c)))
	}
	if c.Wal == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Wal must not be empty", c)))
	}

	return &Redis{
		emi: c.Emi,
		eve: c.Eve,
		log: c.Log,
		mse: c.Mse,
		msl: c.Msl,
		red: c.Red,
		use: c.Use,
		wal: c.Wal,
	}
}

func subObj(sid objectid.ID) string {
	return fmt.Sprintf(keyfmt.SubscriptionObject, sid)
}

func subPay(use objectid.ID) string {
	return fmt.Sprintf(keyfmt.SubscriptionPayer, use)
}

func subRec(use objectid.ID) string {
	return fmt.Sprintf(keyfmt.SubscriptionReceiver, use)
}

func subUse(use objectid.ID) string {
	return fmt.Sprintf(keyfmt.SubscriptionUser, use)
}

func musStr(obj *Object) string {
	byt, err := json.Marshal(obj)
	if err != nil {
		tracer.Panic(tracer.Mask(err))
	}

	return string(byt)
}
