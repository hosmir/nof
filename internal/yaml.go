package internal

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"gopkg.in/yaml.v3"
)

func ProcessArgs(args []string) ([]string, error) {
	for i, arg := range args {
		re := regexp.MustCompile("(.*?)\\s(.*?)")
		found := re.Find([]byte(arg))
		if found != nil {
			return []string{}, errors.New(fmt.Sprintf("Variable can't contian any whitespace: '%s'", arg))
		}
		if string(arg[0]) == "$" {
			cmdVar := arg[1:]
			value := os.Getenv(cmdVar)
			if value != "" {
				args[i] = value
			}
		}
	}
	return args, nil
}

func (c *Command) UnmarshalYAML(value *yaml.Node) error {
	var raw map[string][]string
	if err := value.Decode(&raw); err != nil {
		return err
	}
	for k, v := range raw {
		processed_args, err := ProcessArgs(v)
		if err != nil {
			return err
		}
		c.Name = k
		c.Args = processed_args
		break
	}
	return nil
}

func (c Command) MarshalYAML() (any, error) {
	return map[string][]string{
		c.Name: c.Args,
	}, nil
}

func CommandToExec(command *Command) *exec.Cmd {
	name := command.Name
	args := command.Args
	cmd := exec.Command(name, args...)
	cmd.Env = os.Environ()
	return cmd
}
