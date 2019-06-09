package git

import (
	"regexp"
)

type Branch struct{}

func (b Branch) DescriptionMap() (map[string]string, error) {
	configList, err := GetConfigList()
	if err != nil {
		return nil, err
	}

	descMap := buildDescriptionMap(configList)
	return descMap, nil
}

func buildDescriptionMap(configList []string) map[string]string {
	descLineReg := regexp.MustCompile(`^branch.*description=`)
	descMap := make(map[string]string)
	for _, configLine := range configList {
		if descLineReg.MatchString(configLine) {
			desc := extractDescription(configLine)
			branchName := extractBranchName(configLine)
			descMap[branchName] = desc
		}
	}

	return descMap
}

func extractDescription(line string) string {
	descReg := regexp.MustCompile(`^branch.*description=`)
	return descReg.ReplaceAllString(line, "")
}

func extractBranchName(line string) string {
	reg := regexp.MustCompile(`(branch\.|\.description|=.+)`)
	return reg.ReplaceAllString(line, "")
}

func (b Branch) Description() (string, error) {
	branchName, err := GetCurrentBranch()
	if err != nil {
		return "", err
	}

	key := buildDescriptionKey(branchName)
	description, err := GetConfigValue(key)
	if err != nil {
		return "", nil
	}

	return description, nil
}

func (b Branch) SetDescription(desc string) error {
	var err error
	branchName, err := GetCurrentBranch()
	if err != nil {
		return err
	}

	key := buildDescriptionKey(branchName)
	err = SetConfigValue(key, desc)
	if err != nil {
		return err
	}

	return nil
}

func buildDescriptionKey(branchName string) string {
	return "branch." + branchName + ".description"
}

func (b Branch) Page() (string, error) {
	branchName, err := GetCurrentBranch()
	if err != nil {
		return "", err
	}

	key := buildPageKey(branchName)
	description, err := GetConfigValue(key)
	if err != nil {
		return "", nil
	}

	return description, nil
}

func (b Branch) SetPage(desc string) error {
	var err error
	branchName, err := GetCurrentBranch()
	if err != nil {
		return err
	}

	key := buildPageKey(branchName)
	err = SetConfigValue(key, desc)
	if err != nil {
		return err
	}

	return nil
}

func buildPageKey(branchName string) string {
	return "branch." + branchName + ".page"
}
