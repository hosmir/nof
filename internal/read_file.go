package internal

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

// Read reads a file and returns a marshalled yaml
func Read(filePath string) (*Command, error) {
	// try to open, return default if no such file
	f, err := os.Open(filePath)
	if os.IsNotExist(err) {
		return &Command{}, nil
	} else if err != nil {
		return nil, err
	}

	var cmd Command
	rawExisting, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal([]byte(rawExisting), &cmd); err != nil {
		panic(err)
	}

	return &cmd, nil
}
