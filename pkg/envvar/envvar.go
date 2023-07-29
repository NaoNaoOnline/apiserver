package envvar

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/xh3b4sd/tracer"
)

type Env struct {
	OauthAud string `required:"true" split_words:"true"`
	OauthIss string `required:"true" split_words:"true"`
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
