package twitterclient

import (
	"os"

	"github.com/michimani/gotwi"
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

type Client struct {
	cli *gotwi.Client
}

func New() *Client {
	var cli *Client
	{
		cli = &Client{
			cli: musCli(),
		}
	}

	return cli
}

func musCli() *gotwi.Client {
	cli, err := gotwi.NewClient(cfg)
	if err != nil {
		tracer.Panic(err)
	}

	return cli
}
