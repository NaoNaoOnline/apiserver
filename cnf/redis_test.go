//go:build redis

package cnf

import (
	"github.com/NaoNaoOnline/apiserver/pkg/emitter/ruleemitter"
	"github.com/NaoNaoOnline/apiserver/pkg/feed"
	"github.com/NaoNaoOnline/apiserver/pkg/object/objectid"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/liststorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/rulestorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

func eveOne() *eventstorage.Object {
	return &eventstorage.Object{
		Evnt: "944148",
		Cate: []objectid.ID{"944148", "944149"},
		Host: []objectid.ID{"944148"},
		User: "944148",
	}
}

func eveTwo() *eventstorage.Object {
	return &eventstorage.Object{
		Evnt: "426989",
		Cate: []objectid.ID{"426989", "426990", "426991"},
		Host: []objectid.ID{"426989"},
		User: "426989",
	}
}

func eveThr() *eventstorage.Object {
	return &eventstorage.Object{
		Evnt: "206767",
		Cate: []objectid.ID{"206767", "206768"},
		Host: []objectid.ID{"206767", "206768", "206769", "206770"},
		User: "206767",
	}
}

func lisOne() *liststorage.Object {
	return &liststorage.Object{
		List: "752645",
	}
}

func rulOne(lid objectid.ID) *rulestorage.Object {
	return &rulestorage.Object{
		Incl: []objectid.ID{"000001", "426990", "000002"}, // none, eveTwo:426989, none
		List: lid,
		Kind: "cate",
		User: "295301",
	}
}

func rulTwo(lid objectid.ID) *rulestorage.Object {
	return &rulestorage.Object{
		Incl: []objectid.ID{"944148", "206770"}, // eveOne:944148, eveThr:206767
		List: lid,
		Kind: "host",
		User: "295301",
	}
}

func rulThr(lid objectid.ID) *rulestorage.Object {
	return &rulestorage.Object{
		Incl: []objectid.ID{"426990", "944148"}, // eveTwo:426989, eveOne:944148
		List: lid,
		Kind: "cate",
		User: "295301",
	}
}

func rulExc(lid objectid.ID) *rulestorage.Object {
	return &rulestorage.Object{
		Excl: []objectid.ID{"426990", "944148"}, // eveTwo:426989, eveOne:944148
		List: lid,
		Kind: "cate",
		User: "295301",
	}
}

func newSrv() (redigo.Interface, feed.Interface, rulestorage.Interface) {
	var red redigo.Interface
	{
		red = prgAll(redigo.Default())
	}

	var fee feed.Interface
	{
		fee = feed.New(feed.Config{
			Log: logger.Default(),
			Red: red,
		})
	}

	var rul rulestorage.Interface
	{
		rul = rulestorage.NewRedis(rulestorage.RedisConfig{
			Emi: ruleemitter.Fake(),
			Log: logger.Default(),
			Red: red,
		})
	}

	return red, fee, rul
}

// prgAll is a convenience function for calling FLUSHALL. The provided redigo
// interface is returned as is.
func prgAll(red redigo.Interface) redigo.Interface {
	{
		err := red.Purge()
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	return red
}
