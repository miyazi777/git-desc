package shell

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunCommandSuccess(t *testing.T) {
	assert := assert.New(t)

	cmd := CommandImpl{}
	result, err := cmd.Run("echo", "'test'")
	assert.Equal(result, "'test'")
	assert.Nil(err)
}

func TestRunCommandFailed(t *testing.T) {
	assert := assert.New(t)

	cmd := CommandImpl{}
	result, err := cmd.Run("echox", "'test'")
	assert.Empty(result)
	assert.EqualError(err, "exec: \"echox\": executable file not found in $PATH")
}
