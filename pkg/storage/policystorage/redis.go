package policystorage

import (
	"encoding/json"
	"fmt"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type RedisConfig struct {
	Log logger.Interface
	Red redigo.Interface
	Wal walletstorage.Interface
}

type Redis struct {
	log logger.Interface
	red redigo.Interface
	wal walletstorage.Interface
}

func NewRedis(c RedisConfig) *Redis {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Red == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Red must not be empty", c)))
	}
	if c.Wal == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Wal must not be empty", c)))
	}

	return &Redis{
		log: c.Log,
		red: c.Red,
		wal: c.Wal,
	}
}

func polKin(kin string) string {
	if kin == "CreateMember" {
		return keyfmt.PolicyCreateMember
	}

	if kin == "CreateSystem" {
		return keyfmt.PolicyCreateSystem
	}

	if kin == "DeleteMember" {
		return keyfmt.PolicyDeleteMember
	}

	if kin == "DeleteSystem" {
		return keyfmt.PolicyDeleteSystem
	}

	panic(fmt.Sprintf("kin must be CreateMember, CreateSystem, DeleteMember or DeleteSystem, got %s", kin))
}

func polMem(kin string, mem string) string {
	if kin == "CreateMember" {
		return fmt.Sprintf(keyfmt.PolicyMember, "cre", mem)
	}

	if kin == "CreateSystem" {
		return fmt.Sprintf(keyfmt.PolicyMember, "cre", mem)
	}

	if kin == "DeleteMember" {
		return fmt.Sprintf(keyfmt.PolicyMember, "del", mem)
	}

	if kin == "DeleteSystem" {
		return fmt.Sprintf(keyfmt.PolicyMember, "del", mem)
	}

	panic(fmt.Sprintf("kin must be CreateMember, CreateSystem, DeleteMember or DeleteSystem, got %s", kin))
}

func polObj(oid objectid.ID) string {
	return fmt.Sprintf(keyfmt.PolicyObject, oid)
}

func polSys(kin string, sys int64, mem string) string {
	if kin == "CreateMember" {
		return fmt.Sprintf(keyfmt.PolicySystem, "cre", sys, mem)
	}

	if kin == "CreateSystem" {
		return fmt.Sprintf(keyfmt.PolicySystem, "cre", sys, mem)
	}

	if kin == "DeleteMember" {
		return fmt.Sprintf(keyfmt.PolicySystem, "del", sys, mem)
	}

	if kin == "DeleteSystem" {
		return fmt.Sprintf(keyfmt.PolicySystem, "del", sys, mem)
	}

	panic(fmt.Sprintf("kin must be CreateMember, CreateSystem, DeleteMember or DeleteSystem, got %s", kin))
}

func musStr(obj *Object) string {
	byt, err := json.Marshal(obj)
	if err != nil {
		tracer.Panic(tracer.Mask(err))
	}

	return string(byt)
}
