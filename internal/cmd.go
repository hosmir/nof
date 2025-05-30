package internal

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nof FILENAME",
	Short: "Run commands using a yaml file for flags and arguments.",
	Long: `Run your commands with less flags and arguments. Configure a yaml file 
	that looks very much like the actual command you meat to run. For example, the file below 
	will run the following command: "find /var/log "*.log" -mtime -3"

	find:
		- "/var/log"
		- "*.log"
		- "-mtime"
		- "-3"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		command, err := Read(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		commandexec := CommandToExec(command)
		Run(*commandexec)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
