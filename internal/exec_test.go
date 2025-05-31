package internal

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	// Do not print the output to stdout (cleaner test runs)
	oldStdout := os.Stdout
	os.Stdout = nil
	defer func() { os.Stdout = oldStdout }()

	tests := []struct {
		name    string
		command *exec.Cmd
	}{
		{
			name:    "successful echo command",
			command: exec.Command("echo", "hello"),
		},
		{
			name:    "successful ls command",
			command: exec.Command("ls", "-la"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Run(*tt.command)
			assert.NoError(t, err)
		})
	}
}

func TestRunWithOutput(t *testing.T) {
	// Do not print the output to stdout (cleaner test runs)
	oldStdout := os.Stdout
	os.Stdout = nil
	defer func() { os.Stdout = oldStdout }()

	cmd := exec.Command("echo", "test output")
	err := Run(*cmd)
	assert.NoError(t, err)
}

func TestRunWithStderr(t *testing.T) {
	// Do not print the output to stdout (cleaner test runs)
	oldStdout := os.Stdout
	os.Stdout = nil
	defer func() { os.Stdout = oldStdout }()

	cmd := exec.Command("ls", "/nonexistentdirectory")
	err := Run(*cmd)
	assert.NoError(t, err)
}
