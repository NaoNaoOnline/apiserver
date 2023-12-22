package version

import (
	"fmt"
	"os"

	"github.com/NaoNaoOnline/apiserver/pkg/runtime"
	"github.com/spf13/cobra"
)

type run struct{}

func (r *run) run(cmd *cobra.Command, args []string) {
	fmt.Fprintf(os.Stdout, "Git Sha       %s\n", runtime.Sha())
	fmt.Fprintf(os.Stdout, "Git Tag       %s\n", runtime.Tag())
	fmt.Fprintf(os.Stdout, "Repository    %s\n", runtime.Src())
	fmt.Fprintf(os.Stdout, "Go Arch       %s\n", runtime.Arc())
	fmt.Fprintf(os.Stdout, "Go OS         %s\n", runtime.Gos())
	fmt.Fprintf(os.Stdout, "Go Version    %s\n", runtime.Ver())
}
