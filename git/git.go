package git

import (
	"errors"
	"github.com/miyazi777/git-desc/shell"
	"strings"
)

func GetConfigList() ([]string, error) {
	result, err := shell.Run("git", "config", "--local", "--list")
	if err != nil {
		return nil, errors.New("Not a git repository")
	}

	return strings.Split(result, "\n"), nil
}

func GetCurrentBranch() (string, error) {
	result, err := shell.Run("git", "symbolic-ref", "--short", "HEAD")
	if err != nil {
		return "", errors.New("Not a git repository")
	}
	return result, nil
}

func SetConfigValue(key string, value string) error {
	_, err := shell.Run("git", "config", "--local", key, value)
	if err != nil {
		return errors.New("Not a git repository")
	}
	return nil
}

func GetConfigValue(key string) (string, error) {
	result, err := shell.Run("git", "config", "--local", key)
	if err != nil {
		return "", errors.New("Not a git repository")
	}
	return result, nil
}
