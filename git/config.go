package git

import (
	"regexp"
	"sort"
)

type Config interface {
	ConfigList() ([]ConfigInfo, error)
}

type ConfigInfo struct {
	Branch      string
	Description string
}

type ConfigImpl struct {
	Git Git
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
