package twitterclient

import (
	"context"
	"fmt"
	"os"

	"github.com/michimani/gotwi"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

const (
	OAuthTokenEnvKeyName       = "GOTWI_ACCESS_TOKEN"
	OAuthTokenSecretEnvKeyName = "GOTWI_ACCESS_TOKEN_SECRET"
)

var (
	cfg = &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           os.Getenv(OAuthTokenEnvKeyName),
		OAuthTokenSecret:     os.Getenv(OAuthTokenSecretEnvKeyName),
	}
)

type Config struct {
	Log logger.Interface
}

type Client struct {
	cli *gotwi.Client
	log logger.Interface
}

func New(c Config) *Client {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}

	var cli *Client
	{
		cli = &Client{
			cli: musCli(),
			log: c.Log,
		}
	}

	if !cli.Verify() {
		cli.log.Log(
			context.Background(),
			"level", "info",
			"message", "twitter client not initialized",
		)
	}

	return cli
}

func musCli() *gotwi.Client {
	cli, err := gotwi.NewClient(cfg)
	if err != nil {
		return nil
	}

	return cli
}
