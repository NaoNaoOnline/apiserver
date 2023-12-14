package userstorage

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type RedisConfig struct {
	Log logger.Interface
	// Pso is a global premium subscription overwrite. Setting this value provides
	// every user returned by the user storage with a premium subscription until
	// the timestamp provided here has expired. That is, if Pso were set to the
	// 1st of November and today would be the 3rd of October, then every user on
	// the platform would have a premium subscription for the rest of the month.
	Pso time.Time
	Red redigo.Interface
}

type Redis struct {
	log logger.Interface
	pso time.Time
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
		pso: c.Pso,
		red: c.Red,
	}
}

func linUse(oid objectid.ID) string {
	return fmt.Sprintf(keyfmt.LinkUser, oid)
}

func useCla(str string) string {
	return fmt.Sprintf(keyfmt.UserClaim, str)
}

func useNam(str string) string {
	return fmt.Sprintf(keyfmt.UserName, str)
}

func useObj(oid objectid.ID) string {
	return fmt.Sprintf(keyfmt.UserObject, oid)
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
