package git

import (
	"errors"
	"github.com/miyazi777/git-desc/shell"
	"strings"
)

var command = shell.SetupCommand()

type Git interface {
	GetConfigList() ([]string, error)
	GetCurrentBranch() (string, error)
	SetConfigValue(key string, value string) error
	GetConfigValue(key string) (string, error)
}

type GitImpl struct {
	command shell.Command
}

func SetupGit() Git {
	return GitImpl{}
}

func (g GitImpl) GetConfigList() ([]string, error) {
	result, err := command.Run("git", "config", "--local", "--list")
	if err != nil {
		return nil, errors.New("Not a git repository")
	}

	return strings.Split(result, "\n"), nil
}

func (g GitImpl) GetCurrentBranch() (string, error) {
	result, err := command.Run("git", "symbolic-ref", "--short", "HEAD")
	if err != nil {
		return "", errors.New("Not a git repository")
	}
	return result, nil
}

func (g GitImpl) SetConfigValue(key string, value string) error {
	_, err := command.Run("git", "config", "--local", key, value)
	if err != nil {
		return errors.New("Not a git repository")
	}
	return nil
}

func (g GitImpl) GetConfigValue(key string) (string, error) {
	result, err := command.Run("git", "config", "--local", key)
	if err != nil {
		return "", errors.New("Not a git repository")
	}
	return result, nil
}
