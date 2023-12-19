package systemd

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/coreos/go-systemd/v22/dbus"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

const (
	PathRedis   = "/home/ubuntu/redis/data"
	PathSystemd = "/etc/systemd/system"
)

var (
	uni = []Unit{
		{
			cou: 1,
			nam: "redis.server.service",
			tem: RedisServerService,
		},
		{
			cou: 1,
			nam: "apiserver.daemon.service",
			tem: ApiserverDaemonService,
		},
	}
)

type run struct {
	ctx context.Context
	fla *flag
	log logger.Interface
}

func (r *run) run(cmd *cobra.Command, args []string) {
	var err error

	{
		err = r.fla.Validate()
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	{
		err = r.runE()
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}
}

func (r *run) runE() error {
	var err error

	if r.fla.UserData {
		{
			err = r.userData()
			if err != nil {
				return tracer.Mask(err)
			}
		}
	} else {
		{
			r.log.Log(r.ctx, "level", "info", "message", "starting")
		}

		{
			err = r.redisConf()
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			err = r.unitFiles()
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			r.log.Log(r.ctx, "level", "info", "message", "complete")
		}
	}

	return nil
}

func (r *run) unitFiles() error {
	var err error

	{
		r.log.Log(r.ctx, "level", "info", "message", "replacing systemd units")
	}

	{
		for _, u := range uni {
			for i := 0; i < u.Cou(); i++ {
				p := filepath.Join(PathSystemd, u.Nam(i))

				t, err := template.New(p).Parse(u.Tem())
				if err != nil {
					return tracer.Mask(err)
				}

				var b bytes.Buffer
				err = t.ExecuteTemplate(&b, p, r.dat(runtime.Tag()))
				if err != nil {
					return tracer.Mask(err)
				}

				err = os.WriteFile(p, b.Bytes(), 0600)
				if err != nil {
					return tracer.Mask(err)
				}
			}
		}
	}

	var con *dbus.Conn
	{
		con, err = dbus.NewSystemConnectionContext(context.Background())
		if err != nil {
			return tracer.Mask(err)
		}

		defer con.Close()
	}

	{
		err = con.ReloadContext(context.Background())
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		for _, u := range uni {
			for i := 0; i < u.Cou(); i++ {
				_, err = con.StartUnitContext(context.Background(), u.Nam(i), "replace", nil)
				if err != nil {
					return tracer.Mask(err)
				}
			}
		}
	}

	return nil
}

func (r *run) redisConf() error {
	var err error

	{
		r.log.Log(r.ctx, "level", "info", "message", "persisting redis config")
	}

	var buf bytes.Buffer
	{
		t, err := template.New(PathRedis).Parse(RedisServerConf)
		if err != nil {
			return tracer.Mask(err)
		}

		err = t.ExecuteTemplate(&buf, PathRedis, nil)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err := os.MkdirAll(PathRedis, os.ModePerm)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = os.WriteFile(filepath.Join(PathRedis, "redis.conf"), buf.Bytes(), 0644)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (r *run) userData() error {
	var buf bytes.Buffer
	{
		t, err := template.New(PathRedis).Parse(UserData)
		if err != nil {
			return tracer.Mask(err)
		}

		err = t.ExecuteTemplate(&buf, PathRedis, r.dat(r.fla.Version))
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		fmt.Printf("%s\n", buf.String())
	}

	return nil
}

func (r *run) dat(v string) interface{} {
	type ApiServer struct {
		Version string
	}

	type RedisServer struct {
		Version string
	}

	type Data struct {
		ApiServer   ApiServer
		RedisServer RedisServer
	}

	return Data{
		ApiServer: ApiServer{
			Version: v,
		},
		RedisServer: RedisServer{
			Version: "6.2.0",
		},
	}
}