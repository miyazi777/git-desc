package config

import (
	"github.com/miyazi777/git-desc/git"
	"regexp"
	"sort"
)

type ConfigInfo struct {
	Branch      string
	Description string
}

type List interface {
	GetConfigList() ([]ConfigInfo, error)
}

type ListImpl struct {
	Command git.Command
}

func (c *ListImpl) GetConfigList() ([]ConfigInfo, error) {
	configLineList, err := c.Command.GetConfigList()
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
