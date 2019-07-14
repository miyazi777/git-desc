package config

import (
	"github.com/miyazi777/git-desc/git"
)

type Page interface {
	Get() (string, error)
	Set(page string) error
	DeletePage(branchName string) error
}

type PageImpl struct {
	Command git.Command
}

func (p *PageImpl) Get() (string, error) {
	branchName, err := p.Command.GetCurrentBranch()
	if err != nil {
		return "", err
	}

	key := BuildPageKey(branchName)
	description, err := p.Command.GetConfigValue(key)
	if err != nil {
		return "", err
	}

	return description, nil
}

func (p *PageImpl) Set(page string) error {
	var err error
	branchName, err := p.Command.GetCurrentBranch()
	if err != nil {
		return err
	}

	key := BuildPageKey(branchName)
	err = p.Command.SetConfigValue(key, page)
	if err != nil {
		return err
	}

	return nil
}

func (p *PageImpl) DeletePage(branchName string) error {
	pageKey := BuildPageKey(branchName)
	err := p.Command.DeleteConfigValue(pageKey)
	if err != nil {
		return err
	}

	return nil
}

func BuildPageKey(branchName string) string {
	return "branch." + branchName + ".page"
}
