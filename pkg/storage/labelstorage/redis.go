package labelstorage

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/NaoNaoOnline/apiserver/pkg/scoreid"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type RedisConfig struct {
	Log logger.Interface
	Red redigo.Interface
}

type Redis struct {
	log logger.Interface
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
		red: c.Red,
	}
}

func (r *Redis) Create(inp *Object) (*Object, error) {
	var err error

	{
		inp.Crtd = time.Now().UTC()
		inp.Labl = scoreid.New(inp.Crtd)
	}

	if inp.Kind != "cate" && inp.Kind != "host" {
		return nil, tracer.Mask(invalidInputError)
	}

	// At first we create the normalized key-value pair so that we can search for
	// label objects using their IDs.
	{
		err = r.red.Simple().Create().Element(keyObj(inp.Labl), musStr(inp))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// Now we create the global and user specific mappings for global and user
	// specific search queries.
	{
		err = r.red.Sorted().Create().Element(keyKin(inp.Kind), inp.Labl.String(), inp.Labl.Float())
		if err != nil {
			return nil, tracer.Mask(err)
		}

		err = r.red.Sorted().Create().Element(keyUse(inp.User), inp.Labl.String(), inp.Labl.Float())
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return nil, nil
}

func keyKin(kin string) string {
	if kin == "cate" {
		return keyfmt.LabelCategory
	}

	if kin == "host" {
		return keyfmt.LabelHost
	}

	panic(fmt.Sprintf("kin must be cate or host, got %s", kin))
}

func keyObj(sid scoreid.String) string {
	return fmt.Sprintf(keyfmt.LabelObject, sid)
}

func keyUse(use string) string {
	return fmt.Sprintf(keyfmt.LabelUser, use)
}

func musStr(obj *Object) string {
	byt, err := json.Marshal(obj)
	if err != nil {
		tracer.Panic(tracer.Mask(err))
	}

	return string(byt)
}
