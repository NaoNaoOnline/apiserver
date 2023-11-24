package walletstorage

import (
	"encoding/json"
	"testing"

	"github.com/NaoNaoOnline/apiserver/pkg/object/objectfield"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectlabel"
	"github.com/rafaeljusto/redigomock/v3"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/redigo/pool"
	"github.com/xh3b4sd/tracer"
)

func Test_Storage_Wallet_UpdatePtch(t *testing.T) {
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

		if !obj.HasLab(objectlabel.WalletUnassigned) {
			t.Fatal("wallet label should be changed from accounting to unassigned")
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

	// The object we are about to patch has the wallet label accounting set. Below
	// we are going to update this label.
	var obj *Object
	{
		obj = &Object{
			Kind: "eth",
			Labl: objectfield.Strings{
				Data: []string{
					objectlabel.WalletAccounting,
				},
			},
		}
	}

	// The patch defines the existing wallet label accounting to be removed and
	// the wallet label unassigned to be set.
	var pat []*Patch
	{
		pat = []*Patch{
			{Ope: "remove" /**/, Pat: "/labl/data/0", Val: objectlabel.WalletAccounting},
			{Ope: "add" /*****/, Pat: "/labl/data/-", Val: objectlabel.WalletUnassigned},
		}
	}

	_, _, err = sto.UpdatePtch([]*Object{obj}, PatchSlicer{pat})
	if err != nil {
		t.Fatal(err)
	}

	if conn.Stats(cmd) != 1 {
		t.Fatal("command SET not called as expected")
	}
}
