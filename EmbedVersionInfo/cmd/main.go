package cmd

import (
	"github.com/zaoshi00/examples-go/EmbedVersionInfo/version"
	"os"

	"github.com/spf13/cobra"
)

var mainCmd = &cobra.Command{
	Use:   "run",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	err := mainCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	mainCmd.AddCommand(
		version.Cmd,
	)
}
