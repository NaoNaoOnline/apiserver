package systemd

import (
	"github.com/spf13/cobra"
)

const (
	use = "systemd"
	sho = "Launch systemd unit files on the executing host."
	lon = "Launch systemd unit files on the executing host."
)

type Config struct{}

func New(config Config) (*cobra.Command, error) {
	var c *cobra.Command
	{
		c = &cobra.Command{
			Use:   use,
			Short: sho,
			Long:  lon,
			Run:   (&run{}).run,
		}
	}

	return c, nil
}
