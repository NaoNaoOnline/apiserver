package cmd

import (
	"github.com/NaoNaoOnline/apiserver/cmd/daemon"
	"github.com/NaoNaoOnline/apiserver/cmd/fakeit"
	"github.com/NaoNaoOnline/apiserver/cmd/systemd"
	"github.com/NaoNaoOnline/apiserver/cmd/version"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"
)

var (
	use = "apiserver"
	sho = "Golang based RPC apiserver."
	lon = "Golang based RPC apiserver."
)

func New() (*cobra.Command, error) {
	var err error

	// --------------------------------------------------------------------- //

	var cmdDae *cobra.Command
	{
		c := daemon.Config{}

		cmdDae, err = daemon.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var cmdFak *cobra.Command
	{
		c := fakeit.Config{}

		cmdFak, err = fakeit.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var cmdSys *cobra.Command
	{
		c := systemd.Config{}

		cmdSys, err = systemd.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var cmdVer *cobra.Command
	{
		c := version.Config{}

		cmdVer, err = version.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// --------------------------------------------------------------------- //

	var c *cobra.Command
	{
		c = &cobra.Command{
			Use:   use,
			Short: sho,
			Long:  lon,
			Run:   (&run{}).run,
			CompletionOptions: cobra.CompletionOptions{
				DisableDefaultCmd: true,
			},
			// We slience errors because we do not want to see spf13/cobra printing.
			// The errors returned by the commands will be propagated to the main.go
			// anyway, where we have custom error printing for the command line
			// tool.
			SilenceErrors: true,
			SilenceUsage:  true,
		}
	}

	{
		c.SetHelpCommand(&cobra.Command{Hidden: true})
	}

	{
		c.AddCommand(cmdDae)
		c.AddCommand(cmdFak)
		c.AddCommand(cmdSys)
		c.AddCommand(cmdVer)
	}

	return c, nil
}
