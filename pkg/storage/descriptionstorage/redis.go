package descriptionstorage

import (
	"encoding/json"
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/emitter/descriptionemitter"
	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type RedisConfig struct {
	Emi descriptionemitter.Interface
	Log logger.Interface
	Red redigo.Interface
}

type Redis struct {
	emi descriptionemitter.Interface
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

func desEve(oid objectid.ID) string {
	return fmt.Sprintf(keyfmt.DescriptionEvent, oid)
}

func desLik(use objectid.ID, oid objectid.ID) string {
	return fmt.Sprintf(keyfmt.DescriptionLike, use, oid)
}

func desObj(oid objectid.ID) string {
	return fmt.Sprintf(keyfmt.DescriptionObject, oid)
}

func desUse(oid objectid.ID) string {
	return fmt.Sprintf(keyfmt.DescriptionUser, oid)
}

func eveObj(oid objectid.ID) string {
	return fmt.Sprintf(keyfmt.EventObject, oid)
}

func musByt(pat []*Patch) []byte {
	byt, err := json.Marshal(pat)
	if err != nil {
		tracer.Panic(tracer.Mask(err))
	}

	return byt
}

func musStr(obj *Object) string {
	byt, err := json.Marshal(obj)
	if err != nil {
		tracer.Panic(tracer.Mask(err))
	}

	return string(byt)
}
