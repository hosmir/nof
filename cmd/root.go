package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nof",
	Short: "Run commands using a yaml file for flags and arguments.",
	Long: `
Run your commands with less flags and arguments. Configure a yaml file 
that looks very much like the actual command you meat to run. For example, the file below 
will run the following command: "find /var/log "*.log" -mtime -3"

find:
	- "/var/log"
	- "*.log"
	- "-mtime"
	- "-3"
	`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
