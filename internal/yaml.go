package internal

import (
	"gopkg.in/yaml.v3"
	"os/exec"
)

func (c *Command) UnmarshalYAML(value *yaml.Node) error {
	var raw map[string][]string
	if err := value.Decode(&raw); err != nil {
		return err
	}
	for k, v := range raw {
		c.Name = k
		c.Args = v
		break
	}
	return nil
}

func (c Command) MarshalYAML() (interface{}, error) {
	return map[string][]string{
		c.Name: c.Args,
	}, nil
}

func CommandToExec(command *Command) *exec.Cmd {
	name := command.Name
	args := command.Args
	cmd := exec.Command(name, args...)
	return cmd
}
