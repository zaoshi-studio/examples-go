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
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Go Version:", GoVersion)
		fmt.Println("Git Hash: ", GitHash)
		fmt.Println("Build Time: ", BuildTime)
	},
}
