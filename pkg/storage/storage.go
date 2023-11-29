package storage

import (
	"fmt"
	"time"

	"github.com/NaoNaoOnline/apiserver/pkg/emitter"
	"github.com/NaoNaoOnline/apiserver/pkg/envvar"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/descriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/eventstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/labelstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/liststorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/policystorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/rulestorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/subscriptionstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/userstorage"
	"github.com/NaoNaoOnline/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Emi *emitter.Emitter
	Env envvar.Env
	Log logger.Interface
	Red redigo.Interface
}

type Storage struct {
	des descriptionstorage.Interface
	eve eventstorage.Interface
	lab labelstorage.Interface
	lis liststorage.Interface
	pol policystorage.Interface
	rul rulestorage.Interface
	sub subscriptionstorage.Interface
	use userstorage.Interface
	wal walletstorage.Interface
}

func New(c Config) *Storage {
	if c.Emi == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Emi must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Red == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Red must not be empty", c)))
	}

	var pso time.Time
	if c.Env.UpremTim != "" {
		pso = musTim(c.Env.UpremTim)
	}

	var s *Storage
	{
		s = &Storage{
			des: descriptionstorage.NewRedis(descriptionstorage.RedisConfig{Emi: c.Emi.Desc(), Log: c.Log, Red: c.Red}),
			eve: eventstorage.NewRedis(eventstorage.RedisConfig{Emi: c.Emi.Evnt(), Log: c.Log, Red: c.Red}),
			lab: labelstorage.NewRedis(labelstorage.RedisConfig{Log: c.Log, Red: c.Red}),
			lis: liststorage.NewRedis(liststorage.RedisConfig{Emi: c.Emi.List(), Log: c.Log, Red: c.Red}),
			pol: policystorage.NewRedis(policystorage.RedisConfig{Log: c.Log, Red: c.Red}),
			rul: rulestorage.NewRedis(rulestorage.RedisConfig{Log: c.Log, Red: c.Red}),
			use: userstorage.NewRedis(userstorage.RedisConfig{Log: c.Log, PSO: pso, Red: c.Red}),
			wal: walletstorage.NewRedis(walletstorage.RedisConfig{Log: c.Log, Red: c.Red}),
		}
	}

	{
		s.sub = subscriptionstorage.NewRedis(subscriptionstorage.RedisConfig{
			Emi: c.Emi.Subs(),
			Eve: s.eve,
			Log: c.Log,
			Red: c.Red,
			Use: s.use,
			Wal: s.wal,
		})
	}

	return s
}

func (s *Storage) Desc() descriptionstorage.Interface {
	return s.des
}

func (s *Storage) Evnt() eventstorage.Interface {
	return s.eve
}

func (s *Storage) Labl() labelstorage.Interface {
	return s.lab
}

func (s *Storage) List() liststorage.Interface {
	return s.lis
}

func (s *Storage) Plcy() policystorage.Interface {
	return s.pol
}

func (s *Storage) Rule() rulestorage.Interface {
	return s.rul
}

func (s *Storage) Subs() subscriptionstorage.Interface {
	return s.sub
}

func (s *Storage) User() userstorage.Interface {
	return s.use
}

func (s *Storage) Wllt() walletstorage.Interface {
	return s.wal
}

func musTim(str string) time.Time {
	tim, err := time.Parse("2006-01-02T15:04:05.999999Z", str)
	if err != nil {
		panic(err)
	}

	return tim
}
