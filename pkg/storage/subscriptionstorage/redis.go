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

const (
	// TODO
	minEve = 3
	minLin = 3
)

type RedisConfig struct {
	Emi subscriptionemitter.Interface
	Eve eventstorage.Interface
	Log logger.Interface
	Red redigo.Interface
	Use userstorage.Interface
	Wal walletstorage.Interface
}

type Redis struct {
	emi subscriptionemitter.Interface
	eve eventstorage.Interface
	log logger.Interface
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
		red: c.Red,
		use: c.Use,
		wal: c.Wal,
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
