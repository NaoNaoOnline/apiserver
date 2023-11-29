package labelstorage

import (
	"encoding/json"
	"testing"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/rafaeljusto/redigomock/v3"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/redigo/pool"
	"github.com/xh3b4sd/tracer"
)

func Test_Storage_Label_UpdatePtch(t *testing.T) {
	var err error

	conn := redigomock.NewConn()

	cmd := conn.GenericCommand("SET").Handle(func(args []interface{}) (interface{}, error) {
		var jsn string
		{
			jsn = args[1].(string)
		}

		var obj *Object
		{
			obj = &Object{}
		}

		if jsn != "" {
			err = json.Unmarshal([]byte(jsn), obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		if obj.Prfl.Data["Twitter"] != "vivekr" {
			t.Fatal("label profile Twitter should be added")
		}

		return "OK", nil
	})

	var red redigo.Interface
	{
		red, err = redigo.New(redigo.Config{
			Pool: pool.NewSinglePoolWithConnection(conn),
			Kind: redigo.KindSingle,
		})
		if err != nil {
			t.Fatal(err)
		}
	}

	sto := NewRedis(RedisConfig{
		Log: logger.Fake(),
		Red: red,
	})

	// The object we are about to patch has no label profiles set. Below ensure
	// that we are able to add a label profile using the fancy jsonpatch option
	// &jsonpatch.ApplyOptions{EnsurePathExistsOnAdd: true}. Without this option
	// adding new map members fails. So this test ensures that we can add new
	// label profiles.
	var obj *Object
	{
		obj = &Object{
			Kind: "host",
			Name: objectfield.String{
				Data: "Vivek Ramaswamy",
			},
			Prfl: objectfield.Map{},
			User: objectfield.ID{
				Data: objectid.ID("1234"),
			},
		}
	}

	var pat []*Patch
	{
		pat = []*Patch{
			{Ope: "add" /*****/, Pat: "/prfl/data/Twitter", Val: "vivekr"},
		}
	}

	_, err = sto.UpdatePtch([]*Object{obj}, PatchSlicer{pat})
	if err != nil {
		t.Fatal(err)
	}

	if conn.Stats(cmd) != 1 {
		t.Fatal("command SET not called as expected")
	}
}
