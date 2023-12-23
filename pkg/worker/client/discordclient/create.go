package discordclient

import (
	"github.com/bwmarrin/discordgo"
	"github.com/xh3b4sd/tracer"
)

const (
	// cid is the channel ID of the NaoNao Discord Server #events channel.
	//
	//     https://discord.com/channels/1146911064315928667/1146911069026127887
	//
	cid = "1146911069026127887"
)

func (c *Client) Create(str string) error {
	var err error

	var cli *discordgo.Session
	{
		cli, err = discordgo.New("Bot " + c.tkn)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		defer cli.Close()
	}

	{
		_, err = cli.ChannelMessageSend(cid, str)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
