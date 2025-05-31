package internal

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	yamlAbsPathExists, _ := filepath.Abs("../examples/find.yaml")
	cmd, err := Read(yamlAbsPathExists)
	assert.NoError(t, err)
	assert.IsType(t, &Command{}, cmd)
	assert.Equal(t, "find", cmd.Name)

	yamlAbsPathNonExistent := "/path/to/nonexistent/file.yaml"
	cmd, err = Read(yamlAbsPathNonExistent)
	assert.Nil(t, err)
	assert.IsType(t, &Command{}, cmd)
	assert.Equal(t, "", cmd.Name)

	assert.Panics(t, func() {
		yamlAbsPathBadFile := "../examples/bad.yaml"
		cmd, err = Read(yamlAbsPathBadFile)
		assert.Nil(t, err)
		assert.IsType(t, &Command{}, cmd)
		assert.Equal(t, "", cmd.Name)
	})
}
