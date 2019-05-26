package git

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

// カレントブランチ取得
func GetCurrentBranch() (string, error) {
	cmd := exec.Command("git", "symbolic-ref", "--short", "HEAD")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		return "", errors.New("Not a git repository")
	}

	branchName := stdout.String()
	return strings.Trim(branchName, "\n"), nil
}

// 説明追加
func SetDescription(branchName string, desc string) error {
	key := "branch." + branchName + ".description"
	cmd := exec.Command("git", "config", "--local", key, desc)
	err := cmd.Run()
	if err != nil {
		return errors.New("Not a git repository")
	}
	return nil
}

// 説明取得
func GetDesctiption(branchName string) (string, error) {
	key := "branch." + branchName + ".description"
	cmd := exec.Command("git", "config", "--local", key)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		return "", errors.New("Not a git repository")
	}

	desc := stdout.String()
	return strings.Trim(desc, "\n"), nil
}
