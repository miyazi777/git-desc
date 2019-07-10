package git

import (
	"regexp"
	"sort"
)

type Branch interface {
	DescriptionList() ([]BranchInfo, error)
	Description() (string, error)
	SetDescription(desc string) error
	Page() (string, error)
	SetPage(desc string) error
	DeleteConfig() error
}

type BranchInfo struct {
	Branch      string
	Description string
}

type BranchImpl struct {
	Git Git
}

func (b *BranchImpl) DescriptionList() ([]BranchInfo, error) {
	configList, err := b.Git.GetConfigList()
	if err != nil {
		return nil, err
	}

	descList := buildDescriptionList(configList)
	return descList, nil
}

func buildDescriptionList(configList []string) []BranchInfo {

	sort.Strings(configList)

	var descList []BranchInfo
	descLineReg := regexp.MustCompile(`^branch.*description=`)
	for _, configLine := range configList {
		if descLineReg.MatchString(configLine) {
			info := BranchInfo{
				Branch:      extractBranchName(configLine),
				Description: extractDescription(configLine),
			}
			descList = append(descList, info)
		}
	}
	return descList
}

func extractDescription(line string) string {
	descReg := regexp.MustCompile(`^branch.*description=`)
	return descReg.ReplaceAllString(line, "")
}

func extractBranchName(line string) string {
	reg := regexp.MustCompile(`(branch\.|\.description|=.+)`)
	return reg.ReplaceAllString(line, "")
}

func (b *BranchImpl) Description() (string, error) {
	branchName, err := b.Git.GetCurrentBranch()
	if err != nil {
		return "", err
	}

	key := buildDescriptionKey(branchName)
	description, err := b.Git.GetConfigValue(key)
	if err != nil {
		return "", nil
	}

	return description, nil
}

func (b *BranchImpl) SetDescription(desc string) error {
	var err error
	branchName, err := b.Git.GetCurrentBranch()
	if err != nil {
		return err
	}

	key := buildDescriptionKey(branchName)
	err = b.Git.SetConfigValue(key, desc)
	if err != nil {
		return err
	}

	return nil
}

func buildDescriptionKey(branchName string) string {
	return "branch." + branchName + ".description"
}

func (b *BranchImpl) Page() (string, error) {
	branchName, err := b.Git.GetCurrentBranch()
	if err != nil {
		return "", err
	}

	key := buildPageKey(branchName)
	description, err := b.Git.GetConfigValue(key)
	if err != nil {
		return "", nil
	}

	return description, nil
}

func (b *BranchImpl) SetPage(desc string) error {
	var err error
	branchName, err := b.Git.GetCurrentBranch()
	if err != nil {
		return err
	}

	key := buildPageKey(branchName)
	err = b.Git.SetConfigValue(key, desc)
	if err != nil {
		return err
	}

	return nil
}

func buildPageKey(branchName string) string {
	return "branch." + branchName + ".page"
}

func (b *BranchImpl) DeleteConfig() error {
	var err error
	branchName, err := b.Git.GetCurrentBranch()
	if err != nil {
		return err
	}

	descKey := buildDescriptionKey(branchName)
	err = b.Git.DeleteConfigValue(descKey)
	if err != nil {
		return err
	}

	pageKey := buildPageKey(branchName)
	err = b.Git.DeleteConfigValue(pageKey)
	if err != nil {
		return err
	}

	return nil
}
