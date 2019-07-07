package shell

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// コマンドが成功した時のテストケース
func TestRunCommandSuccess(t *testing.T) {
	assert := assert.New(t)

	cmd := CommandImpl{}
	result, err := cmd.Run("echo", "'test'")
	assert.Equal(result, "'test'")
	assert.NoError(err)
}

// コマンドが失敗した時のテストケース
func TestRunCommandFailed(t *testing.T) {
	assert := assert.New(t)

	cmd := CommandImpl{}
	result, err := cmd.Run("echox", "'test'")
	assert.Empty(result)
	assert.EqualError(err, "exec: \"echox\": executable file not found in $PATH")
}
