package systemd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type run struct{}

func (r *run) run(cmd *cobra.Command, args []string) {
	// TODO
	fmt.Printf("%#v\n", "hello world")
}
