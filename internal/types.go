package internal

import "gopkg.in/yaml.v3"

type Command struct {
	// name of the command (e.g. curl)
	Name string `json:"command"`
	// arguments/flags that should be passed to the command
	// such as "-X POST" or "--request GET"
	Args []string `json:"args"`
}

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
