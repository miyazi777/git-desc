package config

import (
	"github.com/miyazi777/git-desc/git"
)

type Description interface {
	Get() (string, error)
	Set(desc string) error
}

type DescriptionImpl struct {
	Git git.Git
}

func (d *DescriptionImpl) Get() (string, error) {
	branchName, err := d.Git.GetCurrentBranch()
	if err != nil {
		return "", err
	}

	key := buildDescriptionKey(branchName)
	desc, err := d.Git.GetConfigValue(key)
	if err != nil {
		return "", nil
	}

	return desc, nil
}

func (d *DescriptionImpl) Set(desc string) error {
	var err error
	branchName, err := d.Git.GetCurrentBranch()
	if err != nil {
		return err
	}

	key := buildDescriptionKey(branchName)
	err = d.Git.SetConfigValue(key, desc)
	if err != nil {
		return err
	}

	return nil
}

func buildDescriptionKey(branchName string) string {
	return "branch." + branchName + ".description"
}
