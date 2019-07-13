package config

import (
	"github.com/miyazi777/git-desc/git/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type commandMock struct {
	mock.Mock
}

func (mock *commandMock) GetConfigList() ([]string, error) {
	result := mock.Called()
	return nil, result.Error(1)
}

func (mock *commandMock) GetCurrentBranch() (string, error) {
	result := mock.Called()
	return result.String(0), result.Error(1)
}

func (mock *commandMock) SetConfigValue(key string, value string) error {
	result := mock.Called()
	return result.Error(0)
}

func (mock *commandMock) GetConfigValue(key string) (string, error) {
	result := mock.Called()
	return result.String(0), result.Error(1)
}

func (mock *commandMock) DeleteConfigValue(key string) error {
	result := mock.Called()
	return result.Error(0)
}

func Test(t *testing.T) {
	commandMock := new(commandMock)
	commandMock.On("GetCurrentBranch").Return("test-branch", nil)
	commandMock.On("GetConfigValue").Return("test-value", nil)

	p := &config.PageImpl{
		Command: commandMock,
	}
	result, err := p.Get()

	assert := assert.New(t)
	assert.Equal(result, "test-value")
	assert.NoError(err)
}

func Test2(t *testing.T) {
	commandMock := &commandMock{}
	commandMock.On("GetCurrentBranch").Return("test-branch", nil)
	commandMock.On("SetConfigValue").Return(nil)

	p := &config.PageImpl{
		Command: commandMock,
	}
	err := p.Set("test-value")

	assert := assert.New(t)
	assert.NoError(err)
	commandMock.AssertCalled(t, "SetConfigValue")
}

func Test3(t *testing.T) {
	commandMock := &commandMock{}
	commandMock.On("GetCurrentBranch").Return("test-branch", nil)
	commandMock.On("DeleteConfigValue").Return(nil)

	p := &config.PageImpl{
		Command: commandMock,
	}
	err := p.DeletePage()

	assert := assert.New(t)
	assert.NoError(err)
	commandMock.AssertCalled(t, "DeleteConfigValue")
}
