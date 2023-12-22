package systemd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"
)

type flag struct {
	Restart  bool
	UserData bool
	Version  string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&f.Restart, "restart", "r", false, "Whether to restart the apiserver with the given version.")
	cmd.Flags().BoolVarP(&f.UserData, "user-data", "u", false, "Whether to generate and print the cloudinit userdata.")
	cmd.Flags().StringVarP(&f.Version, "version", "v", "", "The apiserver version to download with the userdata.")
}

func (f *flag) Validate() error {
	{
		if f.Restart && f.Version == "" {
			return tracer.Mask(fmt.Errorf("apiserver version must not be empty with -r/--restart"))
		}
	}

	{
		if f.UserData && f.Version == "" {
			return tracer.Mask(fmt.Errorf("apiserver version must not be empty with -u/--user-data"))
		}
	}

	return nil
}
