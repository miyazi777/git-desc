package config

import (
	"github.com/miyazi777/git-desc/git"
)

type Page interface {
	Get() (string, error)
	Set(page string) error
}

type PageImpl struct {
	Git git.Git
}

func (p *PageImpl) Get() (string, error) {
	branchName, err := p.Git.GetCurrentBranch()
	if err != nil {
		return "", err
	}

	key := BuildPageKey(branchName)
	description, err := p.Git.GetConfigValue(key)
	if err != nil {
		return "", nil
	}

	return description, nil
}

func (p *PageImpl) Set(page string) error {
	var err error
	branchName, err := p.Git.GetCurrentBranch()
	if err != nil {
		return err
	}

	key := BuildPageKey(branchName)
	err = p.Git.SetConfigValue(key, page)
	if err != nil {
		return err
	}

	return nil
}

func BuildPageKey(branchName string) string {
	return "branch." + branchName + ".page"
}
