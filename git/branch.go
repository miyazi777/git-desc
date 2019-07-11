package git

type Branch interface {
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
