package envvar

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/xh3b4sd/tracer"
)

type Env struct {
	ChainCid string `split_words:"true" required:"true"`
	ChainPol string `split_words:"true" required:"true"`
	ChainRpc string `split_words:"true" required:"true"`
	ChainSub string `split_words:"true" required:"true"`
	HttpHost string `split_words:"true" default:"127.0.0.1"`
	HttpPort string `split_words:"true" default:"7777"`
	OauthAud string `split_words:"true" required:"true"`
	OauthIss string `split_words:"true" required:"true"`
	UpremTim string `split_words:"true" required:"false"`
}

func Load() Env {
	err := godotenv.Load(".env.local")
	if err != nil {
		tracer.Panic(tracer.Mask(err))
	}

	var env Env
	err = envconfig.Process("APISERVER", &env)
	if err != nil {
		tracer.Panic(tracer.Mask(err))
	}

	return env
}
