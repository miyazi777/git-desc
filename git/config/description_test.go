package config

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type commandDescMock struct {
	mock.Mock
}

func (mock *commandDescMock) GetConfigList() ([]string, error) {
	result := mock.Called()
	return nil, result.Error(1)
}

func (mock *commandDescMock) GetCurrentBranch() (string, error) {
	result := mock.Called()
	return result.String(0), result.Error(1)
}

func (mock *commandDescMock) SetConfigValue(key string, value string) error {
	result := mock.Called(key, value)
	return result.Error(0)
}

func (mock *commandDescMock) GetConfigValue(key string) (string, error) {
	result := mock.Called()
	return result.String(0), result.Error(1)
}

func (mock *commandDescMock) DeleteConfigValue(key string) error {
	result := mock.Called(key)
	return result.Error(0)
}

func TestGetDescSuccess(t *testing.T) {
	commandMock := new(commandDescMock)
	commandMock.On("GetCurrentBranch").Return("test-branch", nil)
	commandMock.On("GetConfigValue").Return("test-value", nil)

	d := &DescriptionImpl{
		Command: commandMock,
	}
	result, err := d.GetDesc()

	assert := assert.New(t)
	assert.Equal(result, "test-value")
	assert.NoError(err)
}

func TestGetDescFailed1(t *testing.T) {
	commandMock := new(commandDescMock)
	commandMock.On("GetCurrentBranch").Return("test-branch", errors.New("test error"))
	commandMock.On("GetConfigValue").Return("test-value", nil)

	p := &DescriptionImpl{
		Command: commandMock,
	}
	result, err := p.GetDesc()

	assert := assert.New(t)
	assert.Empty(result)
	assert.EqualError(err, "test error")
}

func TestGetDescFailed2(t *testing.T) {
	commandMock := new(commandDescMock)
	commandMock.On("GetCurrentBranch").Return("test-branch", nil)
	commandMock.On("GetConfigValue").Return("test-value", errors.New("test error"))

	p := &DescriptionImpl{
		Command: commandMock,
	}
	result, err := p.GetDesc()

	assert := assert.New(t)
	assert.Empty(result)
	assert.NoError(err)
}

func TestSetDescSuccess(t *testing.T) {
	commandMock := &commandDescMock{}
	commandMock.On("GetCurrentBranch").Return("test-branch", nil)
	commandMock.On("SetConfigValue", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)

	p := &DescriptionImpl{
		Command: commandMock,
	}
	err := p.SetDesc("test-value")

	assert := assert.New(t)
	assert.NoError(err)
	commandMock.AssertCalled(t, "SetConfigValue", "branch.test-branch.description", "test-value")
}

func TestSetDescFailed1(t *testing.T) {
	commandMock := &commandDescMock{}
	commandMock.On("GetCurrentBranch").Return("test-branch", errors.New("test error"))
	commandMock.On("SetConfigValue", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)

	p := &DescriptionImpl{
		Command: commandMock,
	}
	err := p.SetDesc("test-value")

	assert := assert.New(t)
	assert.EqualError(err, "test error")
}

func TestSetDescFailed2(t *testing.T) {
	commandMock := &commandDescMock{}
	commandMock.On("GetCurrentBranch").Return("test-branch", nil)
	commandMock.On("SetConfigValue", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(errors.New("test error"))

	p := &DescriptionImpl{
		Command: commandMock,
	}
	err := p.SetDesc("test-value")

	assert := assert.New(t)
	assert.EqualError(err, "test error")
	commandMock.AssertCalled(t, "GetCurrentBranch")
}

func TestDeleteDescSuccess(t *testing.T) {
	commandMock := &commandDescMock{}
	commandMock.On("DeleteConfigValue", mock.AnythingOfType("string")).Return(nil)

	p := &DescriptionImpl{
		Command: commandMock,
	}
	err := p.DeleteDescription("test-branch")

	assert := assert.New(t)
	assert.NoError(err)
	commandMock.AssertCalled(t, "DeleteConfigValue", "branch.test-branch.description")
}

func TestDeleteDescFailed(t *testing.T) {
	commandMock := &commandDescMock{}
	commandMock.On("DeleteConfigValue", mock.AnythingOfType("string")).Return(errors.New("test error"))

	p := &DescriptionImpl{
		Command: commandMock,
	}
	err := p.DeleteDescription("")

	assert := assert.New(t)
	assert.EqualError(err, "test error")
}
