package fakeit

import (
	"github.com/spf13/cobra"
)

const (
	use = "fakeit"
	sho = "Fill redis with fake data."
	lon = "Fill redis with fake data."
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
