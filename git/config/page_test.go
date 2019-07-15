package config

import (
	"errors"
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
	result := mock.Called(key, value)
	return result.Error(0)
}

func (mock *commandMock) GetConfigValue(key string) (string, error) {
	result := mock.Called()
	return result.String(0), result.Error(1)
}

func (mock *commandMock) DeleteConfigValue(key string) error {
	result := mock.Called(key)
	return result.Error(0)
}

func TestGetPageSuccess(t *testing.T) {
	commandMock := new(commandMock)
	commandMock.On("GetCurrentBranch").Return("test-branch", nil)
	commandMock.On("GetConfigValue").Return("test-value", nil)

	p := &PageImpl{
		Command: commandMock,
	}
	result, err := p.GetPage()

	assert := assert.New(t)
	assert.Equal(result, "test-value")
	assert.NoError(err)
}

func TestGetPageFailed1(t *testing.T) {
	commandMock := new(commandMock)
	commandMock.On("GetCurrentBranch").Return("test-branch", errors.New("test error"))
	commandMock.On("GetConfigValue").Return("test-value", nil)

	p := &PageImpl{
		Command: commandMock,
	}
	result, err := p.GetPage()

	assert := assert.New(t)
	assert.Empty(result)
	assert.EqualError(err, "test error")
}

func TestGetPageFailed2(t *testing.T) {
	commandMock := new(commandMock)
	commandMock.On("GetCurrentBranch").Return("test-branch", nil)
	commandMock.On("GetConfigValue").Return("test-value", errors.New("test error"))

	p := &PageImpl{
		Command: commandMock,
	}
	result, err := p.GetPage()

	assert := assert.New(t)
	assert.Empty(result)
	assert.NoError(err)
}

func TestSetPageSuccess(t *testing.T) {
	commandMock := &commandMock{}
	commandMock.On("GetCurrentBranch").Return("test-branch", nil)
	commandMock.On("SetConfigValue", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)

	p := &PageImpl{
		Command: commandMock,
	}
	err := p.SetPage("test-value")

	assert := assert.New(t)
	assert.NoError(err)
	commandMock.AssertCalled(t, "SetConfigValue", "branch.test-branch.page", "test-value")
}

func TestSetPageFailed1(t *testing.T) {
	commandMock := &commandMock{}
	commandMock.On("GetCurrentBranch").Return("test-branch", errors.New("test error"))
	commandMock.On("SetConfigValue", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)

	p := &PageImpl{
		Command: commandMock,
	}
	err := p.SetPage("test-value")

	assert := assert.New(t)
	assert.EqualError(err, "test error")
}

func TestSetPageFailed2(t *testing.T) {
	commandMock := &commandMock{}
	commandMock.On("GetCurrentBranch").Return("test-branch", nil)
	commandMock.On("SetConfigValue", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(errors.New("test error"))

	p := &PageImpl{
		Command: commandMock,
	}
	err := p.SetPage("test-value")

	assert := assert.New(t)
	assert.EqualError(err, "test error")
	commandMock.AssertCalled(t, "GetCurrentBranch")
}

func TestDeletePageSuccess(t *testing.T) {
	commandMock := &commandMock{}
	commandMock.On("DeleteConfigValue", mock.AnythingOfType("string")).Return(nil)

	p := &PageImpl{
		Command: commandMock,
	}
	err := p.DeletePage("test-branch")

	assert := assert.New(t)
	assert.NoError(err)
	commandMock.AssertCalled(t, "DeleteConfigValue", "branch.test-branch.page")
}

func TestDeletePageFailed(t *testing.T) {
	commandMock := &commandMock{}
	commandMock.On("DeleteConfigValue", mock.AnythingOfType("string")).Return(errors.New("test error"))

	p := &PageImpl{
		Command: commandMock,
	}
	err := p.DeletePage("")

	assert := assert.New(t)
	assert.EqualError(err, "test error")
}
