package git

import (
	"errors"
	"github.com/miyazi777/git-desc/shell"
	"strings"
)

type Command interface {
	GetConfigList() ([]string, error)
	GetCurrentBranch() (string, error)
	SetConfigValue(key string, value string) error
	GetConfigValue(key string) (string, error)
	DeleteConfigValue(key string) error
}

type CommandImpl struct {
	Command shell.Command
}

func (c *CommandImpl) GetConfigList() ([]string, error) {
	result, err := c.Command.Run("git", "config", "--local", "--list")
	if err != nil {
		return nil, errors.New("Not a git repository")
	}

	return strings.Split(result, "\n"), nil
}

func (c *CommandImpl) GetCurrentBranch() (string, error) {
	result, err := c.Command.Run("git", "symbolic-ref", "--short", "HEAD")
	if err != nil {
		return "", errors.New("Not a git repository")
	}
	return result, nil
}

func (c *CommandImpl) SetConfigValue(key string, value string) error {
	_, err := c.Command.Run("git", "config", "--local", key, value)
	if err != nil {
		return errors.New("Not a git repository")
	}
	return nil
}

func (c *CommandImpl) GetConfigValue(key string) (string, error) {
	result, err := c.Command.Run("git", "config", "--local", key)
	if err != nil {
		return "", errors.New("Not a git repository")
	}
	return result, nil
}

func (c *CommandImpl) DeleteConfigValue(key string) error {
	var err error
	result, err := c.Command.Run("git", "config", key)
	if result == "" {
		return nil
	}

	_, err = c.Command.Run("git", "config", "--unset", key)
	if err != nil {
		return errors.New("Not a git repository")
	}
	return nil
}
