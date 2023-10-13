package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	GitHash   string
	GoVersion string
	BuildTime string
)

var Cmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Git hash: \t", GitHash)
		fmt.Println("Go version: \t", GoVersion)
		fmt.Println("Build time: \t", BuildTime)
	},
}
