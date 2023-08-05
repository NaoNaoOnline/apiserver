package userstorage

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/keyfmt"
	"github.com/google/uuid"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/redigo/pkg/simple"
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

func (r *Redis) Create(sub string, img string, nam string) (*Object, error) {
	var err error

	var obj *Object
	{
		obj, err = r.Search(sub, "")
		if IsNotFound(err) {
			// The user does not appear to exist. So first, create the mapping between
			// external subject claim and internal user ID.
			var use string
			{
				use = uuid.NewString()
			}

			{
				err = r.red.Simple().Create().Element(fmt.Sprintf(keyfmt.SubjectClaim, sub), use)
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			// In the middle of the two step process, register the new user object for
			// the local execution scope, so that we return the user object we just
			// created.
			{
				obj = &Object{
					Crtd: time.Now().UTC(),
					Imag: img,
					Name: nam,
					User: use,
				}
			}

			// Second create the mapping between internal user ID and internal user
			// object.
			var jsn string
			{
				jsn = musStr(obj)
			}

			{
				err = r.red.Simple().Create().Element(fmt.Sprintf(keyfmt.UserObject, use), jsn)
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}
		} else if err != nil {
			return nil, tracer.Mask(err)
		} else if obj.Imag != img || obj.Name != nam {
			// The user exists and we update it due to changes in profile picture
			// and/or username.
			{
				obj.Imag = img
				obj.Name = nam
			}

			var jsn string
			{
				jsn = musStr(obj)
			}

			{
				err = r.red.Simple().Create().Element(fmt.Sprintf(keyfmt.UserObject, obj.User), jsn)
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}
		}
	}

	return obj, nil
}

func (r *Redis) Search(sub string, use string) (*Object, error) {
	var err error

	// It is not valid to call Search with both inputs empty or given. Only either
	// input must be provided. Either the external subject claim, or the internal
	// user ID.
	if (sub == "" && use == "") || (sub != "" && use != "") {
		return nil, tracer.Mask(invalidInputError)
	}

	if use == "" {
		use, err = r.red.Simple().Search().Value(fmt.Sprintf(keyfmt.SubjectClaim, sub))
		if simple.IsNotFound(err) {
			return nil, tracer.Mask(notFoundError)
		} else if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var jsn string
	{
		jsn, err = r.red.Simple().Search().Value(fmt.Sprintf(keyfmt.UserObject, use))
		if simple.IsNotFound(err) {
			return nil, tracer.Mask(notFoundError)
		} else if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var obj Object
	{
		err = json.Unmarshal([]byte(jsn), &obj)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return &obj, nil
}

func musStr(obj *Object) string {
	byt, err := json.Marshal(obj)
	if err != nil {
		tracer.Panic(tracer.Mask(err))
	}

	return string(byt)
}
