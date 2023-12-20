package envvar

import (
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	pat = ".env.local"
)

type Env struct {
	ChainCid string `split_words:"true" required:"true"`
	ChainPol string `split_words:"true" required:"true"`
	ChainRpc string `split_words:"true" required:"true"`
	ChainSub string `split_words:"true" required:"true"`
	HttpHost string `split_words:"true" default:"127.0.0.1"`
	HttpPort string `split_words:"true" default:"7777"`
	MsubsEve int    `split_words:"true" required:"true"`
	MsubsLin int    `split_words:"true" required:"true"`
	OauthAud string `split_words:"true" required:"true"`
	OauthIss string `split_words:"true" required:"true"`
	UpremTim string `split_words:"true" required:"false"`
}

func Load() Env {
	var err error

	var env Env

	for {
		{
			err = godotenv.Load(pat)
			if err != nil {
				fmt.Printf("could not load %s (%s)\n", pat, err)
				time.Sleep(5 * time.Second)
				continue
			}
		}

		{
			err = envconfig.Process("APISERVER", &env)
			if err != nil {
				fmt.Printf("could process envfile %s (%s)\n", pat, err)
				time.Sleep(5 * time.Second)
				continue
			}
		}

		return env
	}
}
