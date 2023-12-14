package eventstorage

import (
	"encoding/json"
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/emitter/eventemitter"
	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type RedisConfig struct {
	Emi eventemitter.Interface
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
}

type Redis struct {
	emi eventemitter.Interface
	log logger.Interface
	mse int
	msl int
	red redigo.Interface
}

func NewRedis(c RedisConfig) *Redis {
	if c.Emi == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Emi must not be empty", c)))
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

	return &Redis{
		emi: c.Emi,
		log: c.Log,
		mse: c.Mse,
		msl: c.Msl,
		red: c.Red,
	}
}

func eveLab(oid objectid.ID) string {
	return fmt.Sprintf(keyfmt.EventLabel, oid)
}

func eveObj(oid objectid.ID) string {
	return fmt.Sprintf(keyfmt.EventObject, oid)
}

func eveUse(oid objectid.ID) string {
	return fmt.Sprintf(keyfmt.EventUser, oid)
}

func labObj(oid objectid.ID) string {
	return fmt.Sprintf(keyfmt.LabelObject, oid)
}

func likUse(oid objectid.ID) string {
	return fmt.Sprintf(keyfmt.LikeUser, oid)
}

func linEve(oid objectid.ID) string {
	return fmt.Sprintf(keyfmt.LinkEvent, oid)
}

func linMap(use objectid.ID, oid objectid.ID) string {
	return fmt.Sprintf(keyfmt.LinkMapping, use, oid)
}

func linUse(oid objectid.ID) string {
	return fmt.Sprintf(keyfmt.LinkUser, oid)
}

func musStr(obj *Object) string {
	byt, err := json.Marshal(obj)
	if err != nil {
		tracer.Panic(tracer.Mask(err))
	}

	return string(byt)
}
