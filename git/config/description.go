package config

import (
	"github.com/miyazi777/git-desc/git"
)

type Description interface {
	Get() (string, error)
	Set(desc string) error
	DeleteDescription(branchName string) error
}

type DescriptionImpl struct {
	Command git.Command
}

func (d *DescriptionImpl) Get() (string, error) {
	branchName, err := d.Command.GetCurrentBranch()
	if err != nil {
		return "", err
	}

	key := buildDescriptionKey(branchName)
	desc, err := d.Command.GetConfigValue(key)
	if err != nil {
		return "", nil
	}

	return desc, nil
}

func (d *DescriptionImpl) Set(desc string) error {
	var err error
	branchName, err := d.Command.GetCurrentBranch()
	if err != nil {
		return err
	}

	key := buildDescriptionKey(branchName)
	err = d.Command.SetConfigValue(key, desc)
	if err != nil {
		return err
	}

	return nil
}

func (d *DescriptionImpl) DeleteDescription(branchName string) error {
	descKey := buildDescriptionKey(branchName)
	err := d.Command.DeleteConfigValue(descKey)
	if err != nil {
		return err
	}

	return nil
}

func buildDescriptionKey(branchName string) string {
	return "branch." + branchName + ".description"
}
