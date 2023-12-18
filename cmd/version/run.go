package version

import (
	"fmt"
	"os"
	"runtime"

	project "github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/spf13/cobra"
)

type run struct{}

func (r *run) run(cmd *cobra.Command, args []string) {
	fmt.Fprintf(os.Stdout, "Git Sha       %s\n", project.Sha())
	fmt.Fprintf(os.Stdout, "Git Tag       %s\n", project.Tag())
	fmt.Fprintf(os.Stdout, "Repository    %s\n", project.Src())
	fmt.Fprintf(os.Stdout, "Go Arch       %s\n", runtime.GOARCH)
	fmt.Fprintf(os.Stdout, "Go OS         %s\n", runtime.GOOS)
	fmt.Fprintf(os.Stdout, "Go Version    %s\n", runtime.Version())
}
