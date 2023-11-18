package twitterclient

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
	"github.com/xh3b4sd/tracer"
)

func (c *Client) Create(str string) error {
	var err error

	var inp *types.CreateInput
	{
		inp = &types.CreateInput{
			Text: gotwi.String(str),
		}
	}

	{
		_, err = managetweet.Create(context.Background(), c.cli, inp)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
