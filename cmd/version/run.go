package version

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	rep = "https://github.com/NaoNaoOnline/apiserver"
	sha = "n/a"
	tag = "n/a"
)

type run struct{}

func (r *run) run(cmd *cobra.Command, args []string) {
	fmt.Fprintf(os.Stdout, "Git Sha       %s\n", sha)
	fmt.Fprintf(os.Stdout, "Git Tag       %s\n", tag)
	fmt.Fprintf(os.Stdout, "Repository    %s\n", rep)
	fmt.Fprintf(os.Stdout, "Go Arch       %s\n", runtime.GOARCH)
	fmt.Fprintf(os.Stdout, "Go OS         %s\n", runtime.GOOS)
	fmt.Fprintf(os.Stdout, "Go Version    %s\n", runtime.Version())
}
