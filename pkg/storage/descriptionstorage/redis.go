package descriptionstorage

import (
	"encoding/json"
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/rescue"
	"github.com/xh3b4sd/tracer"
)

type RedisConfig struct {
	Log logger.Interface
	Red redigo.Interface
	Res rescue.Interface
}

type Redis struct {
	log logger.Interface
	red redigo.Interface
	res rescue.Interface
}

func NewRedis(c RedisConfig) *Redis {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Red == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Red must not be empty", c)))
	}
	if c.Res == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Res must not be empty", c)))
	}

	return &Redis{
		log: c.Log,
		red: c.Red,
		res: c.Res,
	}
}

func desEve(oid objectid.ID) string {
	return fmt.Sprintf(keyfmt.DescriptionEvent, oid)
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

func musPat(pat []*Patch) jsonpatch.Patch {
	byt, err := json.Marshal(pat)
	if err != nil {
		tracer.Panic(tracer.Mask(err))
	}

	typ, err := jsonpatch.DecodePatch(byt)
	if err != nil {
		tracer.Panic(tracer.Mask(err))
	}

	return typ
}

func musStr(obj *Object) string {
	byt, err := json.Marshal(obj)
	if err != nil {
		tracer.Panic(tracer.Mask(err))
	}

	return string(byt)
}
