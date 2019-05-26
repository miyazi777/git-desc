package git

import (
	"bytes"
	"errors"
	"os/exec"
	"regexp"
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

// 説明マップ構築
func BuildDescriptionMap() (map[string]string, error) {
	cmd := exec.Command("git", "config", "--local", "--list")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		return nil, errors.New("Not a git repository")
	}

	configList := strings.Split(stdout.String(), "\n")
	descLineReg := regexp.MustCompile(`^branch.*description=`)

	descMap := make(map[string]string)
	for _, configLine := range configList {
		if descLineReg.MatchString(configLine) {
			desc := extractDescription(configLine)

			branchName := extractBranchName(configLine)
			descMap[branchName] = desc
		}
	}

	return descMap, nil
}

func extractDescription(line string) string {
	descReg := regexp.MustCompile(`^branch.*description=`)
	return descReg.ReplaceAllString(line, "")
}

func extractBranchName(line string) string {
	reg := regexp.MustCompile(`(branch\.|\.description|=.+)`)
	return reg.ReplaceAllString(line, "")
}
