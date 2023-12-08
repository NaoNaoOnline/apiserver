package rulestorage

import (
	"encoding/json"
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/feedstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type RedisConfig struct {
	Log logger.Interface
	Fee feedstorage.Interface
	Red redigo.Interface
}

type Redis struct {
	fee feedstorage.Interface
	log logger.Interface
	red redigo.Interface
}

func NewRedis(c RedisConfig) *Redis {
	if c.Fee == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Fee must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Red == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Red must not be empty", c)))
	}

	return &Redis{
		fee: c.Fee,
		log: c.Log,
		red: c.Red,
	}
}

func lisObj(lis objectid.ID) string {
	return fmt.Sprintf(keyfmt.ListObject, lis)
}

func notKin(kin string, oid objectid.ID) string {
	if kin == "cate" {
		return fmt.Sprintf(keyfmt.FeedCategory, oid)
	}

	if kin == "host" {
		return fmt.Sprintf(keyfmt.FeedHost, oid)
	}

	if kin == "user" {
		return fmt.Sprintf(keyfmt.FeedUser, oid)
	}

	panic(fmt.Sprintf("kin must be cate, host or user, got %s", kin))
}

func rulEve(eve objectid.ID) string {
	return fmt.Sprintf(keyfmt.RuleEvent, eve)
}

func rulLis(lis objectid.ID) string {
	return fmt.Sprintf(keyfmt.RuleList, lis)
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
