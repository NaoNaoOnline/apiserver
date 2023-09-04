package votestorage

import (
	"encoding/json"
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/reactionstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type RedisConfig struct {
	Log logger.Interface
	Rct reactionstorage.Interface
	Red redigo.Interface
}

type Redis struct {
	log logger.Interface
	rct reactionstorage.Interface
	red redigo.Interface
}

func NewRedis(c RedisConfig) *Redis {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Rct == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Rct must not be empty", c)))
	}
	if c.Red == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Red must not be empty", c)))
	}

	return &Redis{
		log: c.Log,
		rct: c.Rct,
		red: c.Red,
	}
}

func desObj(oid objectid.String) string {
	return fmt.Sprintf(keyfmt.DescriptionObject, oid)
}

func eveVot(oid objectid.String) string {
	return fmt.Sprintf(keyfmt.EventUserVote, oid)
}

func votDes(oid objectid.String) string {
	return fmt.Sprintf(keyfmt.VoteDescription, oid)
}

func votObj(oid objectid.String) string {
	return fmt.Sprintf(keyfmt.VoteObject, oid)
}

func votUse(ida objectid.String, idb objectid.String) string {
	return fmt.Sprintf(keyfmt.VoteEventUser, ida, idb)
}

func musStr(obj *Object) string {
	byt, err := json.Marshal(obj)
	if err != nil {
		tracer.Panic(tracer.Mask(err))
	}

	return string(byt)
}
