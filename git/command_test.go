package git

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type commandMock struct {
	mock.Mock
}

func (mock *commandMock) Run(name string, arg ...string) (string, error) {
	result := mock.Called()
	return result.String(0), result.Error(1)
}

func buildGitInstance(mock *commandMock) *CommandImpl {
	g := &CommandImpl{
		Command: mock,
	}
	return g
}

// git configリスト取得関数成功時のテストケース
func TestGetConfigListSuccess(t *testing.T) {
	commandMock := new(commandMock)
	commandMock.On("Run").Return("branch1\nbranch2", nil)

	g := buildGitInstance(commandMock)
	result, err := g.GetConfigList()

	assert := assert.New(t)
	assert.ElementsMatch(result, [...]string{"branch1", "branch2"})
	assert.NoError(err)
}

// git configリスト取得関数失敗時のテストケース
func TestGetConfigListFailed(t *testing.T) {
	commandMock := new(commandMock)
	commandMock.On("Run").Return("", errors.New("test error"))

	g := buildGitInstance(commandMock)
	result, err := g.GetConfigList()

	assert := assert.New(t)
	assert.Empty(result)
	assert.EqualError(err, "Not a git repository")
}

// ブランチ取得関数成功時のテストケース
func TestGetCurrentBranchSuccess(t *testing.T) {
	commandMock := new(commandMock)
	commandMock.On("Run").Return("test-branch", nil)

	g := buildGitInstance(commandMock)
	result, err := g.GetCurrentBranch()

	assert := assert.New(t)
	assert.Equal(result, "test-branch")
	assert.NoError(err)
}

// ブランチ取得関数失敗時のテストケース
func TestGetCurrentBranchFailed(t *testing.T) {
	commandMock := new(commandMock)
	commandMock.On("Run").Return("test-branch", errors.New("test error"))

	g := buildGitInstance(commandMock)
	result, err := g.GetCurrentBranch()

	assert := assert.New(t)
	assert.Empty(result)
	assert.EqualError(err, "Not a git repository")
}

// config設定関数成功時のテストケース
func TestSetConfigValueSuccess(t *testing.T) {
	commandMock := new(commandMock)
	commandMock.On("Run").Return("test config value", nil)

	g := buildGitInstance(commandMock)
	err := g.SetConfigValue("dummy-key", "duumy-value")

	assert := assert.New(t)
	assert.NoError(err)
}

// config設定関数失敗時のテストケース
func TestSetConfigValueFailed(t *testing.T) {
	commandMock := new(commandMock)
	commandMock.On("Run").Return("", errors.New("test error"))

	g := buildGitInstance(commandMock)
	err := g.SetConfigValue("dummy-key", "duumy-value")

	assert := assert.New(t)
	assert.EqualError(err, "Not a git repository")
}

// config取得関数成功時のテストケース
func TestGetConfigValueSuccess(t *testing.T) {
	commandMock := new(commandMock)
	commandMock.On("Run").Return("test config value", nil)

	g := buildGitInstance(commandMock)
	result, err := g.GetConfigValue("dummy")

	assert := assert.New(t)
	assert.Equal(result, "test config value")
	assert.NoError(err)
}

// config取得関数失敗時のテストケース
func TestGetConfigValueFailed(t *testing.T) {
	commandMock := new(commandMock)
	commandMock.On("Run").Return("test config value", errors.New("test error"))

	g := buildGitInstance(commandMock)
	result, err := g.GetConfigValue("dummy")

	assert := assert.New(t)
	assert.Empty(result)
	assert.EqualError(err, "Not a git repository")
}
