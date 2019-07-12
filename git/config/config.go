package config

import (
	"github.com/miyazi777/git-desc/git"
	"regexp"
	"sort"
)

type Config interface {
	ConfigList() ([]ConfigInfo, error)
	DeleteConfig() error
}

type ConfigInfo struct {
	Branch      string
	Description string
}

type ConfigImpl struct {
	Git git.Git
}

func (c *ConfigImpl) ConfigList() ([]ConfigInfo, error) {
	configLineList, err := c.Git.GetConfigList()
	if err != nil {
		return nil, err
	}

	configList := buildConfigList(configLineList)
	return configList, nil
}

func buildConfigList(configLineList []string) []ConfigInfo {
	sort.Strings(configLineList)

	var configInfoList []ConfigInfo
	descLineReg := regexp.MustCompile(`^branch.*description=`)
	for _, configLine := range configLineList {
		if descLineReg.MatchString(configLine) {
			info := ConfigInfo{
				Branch:      extractBranchName(configLine),
				Description: extractDescription(configLine),
			}
			configInfoList = append(configInfoList, info)
		}
	}
	return configInfoList
}

func extractDescription(line string) string {
	descReg := regexp.MustCompile(`^branch.*description=`)
	return descReg.ReplaceAllString(line, "")
}

func extractBranchName(line string) string {
	reg := regexp.MustCompile(`(branch\.|\.description|=.+)`)
	return reg.ReplaceAllString(line, "")
}

func (b *ConfigImpl) DeleteConfig() error {
	var err error
	branchName, err := b.Git.GetCurrentBranch()
	if err != nil {
		return err
	}

	descKey := BuildDescriptionKey(branchName)
	err = b.Git.DeleteConfigValue(descKey)
	if err != nil {
		return err
	}

	pageKey := BuildPageKey(branchName)
	err = b.Git.DeleteConfigValue(pageKey)
	if err != nil {
		return err
	}

	return nil
}
