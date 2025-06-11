package internal

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestCommandUnmarshalYAML(t *testing.T) {
	tests := []struct {
		name     string
		yamlStr  string
		expected Command
	}{
		{
			name: "valid command with multiple args",
			yamlStr: `
find:
  - "/var/log"
  - "*.log"
  - "-mtime"
  - "-3"`,
			expected: Command{
				Name: "find",
				Args: []string{"/var/log", "*.log", "-mtime", "-3"},
			},
		},
		{
			name: "valid command with single arg",
			yamlStr: `
ls:
  - "-la"`,
			expected: Command{
				Name: "ls",
				Args: []string{"-la"},
			},
		},
		{
			name: "valid command with environment variable in the args",
			yamlStr: `
find:
  - "/var/log"
  - "*.log"
  - "-mtime"
  - "$TEST_ENV_KEY"`,
			expected: Command{
				Name: "find",
				Args: []string{"/var/log", "*.log", "-mtime", "TEST_ENV_VALUE"},
			},
		},
		{
			name: "empty command args",
			yamlStr: `
pwd: []`,
			expected: Command{
				Name: "pwd",
				Args: []string{},
			},
		},
	}

	// set a value for the environment variable test case
	os.Setenv("TEST_ENV_KEY", "TEST_ENV_VALUE")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var cmd Command
			err := yaml.Unmarshal([]byte(tt.yamlStr), &cmd)

			assert.NoError(t, err)
			assert.Equal(t, tt.expected.Name, cmd.Name)
			assert.Equal(t, tt.expected.Args, cmd.Args)
		})
	}

	yamlStrWhiteSpace := `
find:
  - "/var/log"
  - "*.log "
  - "-mtime"
  - "$TEST_ENV_KEY"
`

	t.Run("arg contains whitespace", func(t *testing.T) {
		var cmd Command
		err := yaml.Unmarshal([]byte(yamlStrWhiteSpace), &cmd)

		assert.Error(t, err)
	})

}

func TestCommandMarshalYAML(t *testing.T) {
	os.Setenv("TEST_ENV_KEY", "TEST_ENV_VALUE")
	tests := []struct {
		name     string
		command  Command
		expected string
	}{
		{
			name: "marshal command with multiple args",
			command: Command{
				Name: "find",
				Args: []string{"/var/log", "*.log", "-mtime", "-3"},
			},
			expected: "find:\n    - /var/log\n    - '*.log'\n    - -mtime\n    - \"-3\"\n",
		},
		{
			name: "marshal command with single arg",
			command: Command{
				Name: "ls",
				Args: []string{"-la"},
			},
			expected: "ls:\n    - -la\n",
		},
		{
			name: "marshal command with no args",
			command: Command{
				Name: "pwd",
				Args: []string{},
			},
			expected: "pwd: []\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := yaml.Marshal(tt.command)

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, string(output))
		})
	}
}

func TestCommandToExec(t *testing.T) {
	tests := []struct {
		name     string
		command  Command
		wantName string
		wantArgs []string
	}{
		{
			name: "command with multiple args",
			command: Command{
				Name: "find",
				Args: []string{"/var/log", "*.log", "-mtime", "-3"},
			},
			wantName: "/usr/bin/find",
			wantArgs: []string{"find", "/var/log", "*.log", "-mtime", "-3"},
		},
		{
			name: "command with single arg",
			command: Command{
				Name: "ls",
				Args: []string{"-la"},
			},
			wantName: "/usr/bin/ls",
			wantArgs: []string{"ls", "-la"},
		},
		{
			name: "command with no args",
			command: Command{
				Name: "pwd",
				Args: []string{},
			},
			wantName: "/usr/bin/pwd",
			wantArgs: []string{"pwd"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := CommandToExec(&tt.command)

			assert.Equal(t, tt.wantName, cmd.Path)
			assert.Equal(t, tt.wantArgs, cmd.Args)
		})
	}
}
