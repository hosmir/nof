package internal

import (
	"fmt"
	"os/exec"
)

func Run(command exec.Cmd) error {
	stdout, err := command.Output()

	if err != nil {
		fmt.Printf("Error running the command: %s", command.Stderr)
	}

	fmt.Println(string(stdout))
	return nil
}
