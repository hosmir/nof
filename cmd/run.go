package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"nof/internal"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a command",
	Long: `Run a command via the given YAML file.
	The path to the YAML file must be absolute.

	nof run /path/to/file.yaml

	`,
	Run: func(cmd *cobra.Command, args []string) {
		command, err := internal.Read(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		commandexec := internal.CommandToExec(command)
		internal.Run(*commandexec)
	},
}
