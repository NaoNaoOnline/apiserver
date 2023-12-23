package discordclient

import (
	"context"

	"github.com/xh3b4sd/logger"
)

type Config struct {
	Log logger.Interface
	// Tkn is the authorization token, the bot's client secret required to send messages to some Discord channel.
	Tkn string
}

type Client struct {
	log logger.Interface
	tkn string
}

func New(c Config) *Client {
	var cli *Client
	{
		cli = &Client{
			log: c.Log,
			tkn: c.Tkn,
		}
	}

	if !cli.Verify() {
		cli.log.Log(
			context.Background(),
			"level", "info",
			"message", "discord client not initialized",
		)
	}

	return cli
}
