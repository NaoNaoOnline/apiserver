package systemd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"
)

type flag struct {
	UserData bool
	Version  string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.Version, "version", "v", "", "The apiserver version to download with the userdata.")
	cmd.Flags().BoolVarP(&f.UserData, "user-data", "u", false, "Whether to generate and print the cloudinit userdata.")
}

func (f *flag) Validate() error {
	{
		if f.UserData && f.Version == "" {
			return tracer.Mask(fmt.Errorf("apiserver version must not be empty with -u/--user-data"))
		}
	}

	return nil
}
