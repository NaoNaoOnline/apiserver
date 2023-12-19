package systemd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
)

const (
	use = "systemd"
	sho = "Launch systemd unit files on the executing host."
	lon = "Launch systemd unit files on the executing host."
)

type Config struct{}

func New(con Config) (*cobra.Command, error) {
	var f *flag
	{
		f = &flag{}
	}

	var r *run
	{
		r = &run{
			ctx: context.Background(),
			fla: f,
			log: logger.Default(),
		}
	}

	var c *cobra.Command
	{
		c = &cobra.Command{
			Use:   use,
			Short: sho,
			Long:  lon,
			Run:   r.run,
		}
	}

	{
		f.Init(c)
	}

	return c, nil
}
