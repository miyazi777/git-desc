package git

import (
	"github.com/miyazi777/git-desc/shell"
	"regexp"
)

var git = SetupGit(shell.NewCommand())

type Branch interface {
	DescriptionMap() (map[string]string, error)
	Description() (string, error)
	SetDescription(desc string) error
	Page() (string, error)
	SetPage(desc string) error
}

type BranchImpl struct{}

func SetupBranch() Branch {
	return BranchImpl{}
}

func (b BranchImpl) DescriptionMap() (map[string]string, error) {
	configList, err := git.GetConfigList()
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

func (b BranchImpl) Description() (string, error) {
	branchName, err := git.GetCurrentBranch()
	if err != nil {
		return "", err
	}

	key := buildDescriptionKey(branchName)
	description, err := git.GetConfigValue(key)
	if err != nil {
		return "", nil
	}

	return description, nil
}

func (b BranchImpl) SetDescription(desc string) error {
	var err error
	branchName, err := git.GetCurrentBranch()
	if err != nil {
		return err
	}

	key := buildDescriptionKey(branchName)
	err = git.SetConfigValue(key, desc)
	if err != nil {
		return err
	}

	return nil
}

func buildDescriptionKey(branchName string) string {
	return "branch." + branchName + ".description"
}

func (b BranchImpl) Page() (string, error) {
	branchName, err := git.GetCurrentBranch()
	if err != nil {
		return "", err
	}

	key := buildPageKey(branchName)
	description, err := git.GetConfigValue(key)
	if err != nil {
		return "", nil
	}

	return description, nil
}

func (b BranchImpl) SetPage(desc string) error {
	var err error
	branchName, err := git.GetCurrentBranch()
	if err != nil {
		return err
	}

	key := buildPageKey(branchName)
	err = git.SetConfigValue(key, desc)
	if err != nil {
		return err
	}

	return nil
}

func buildPageKey(branchName string) string {
	return "branch." + branchName + ".page"
}
