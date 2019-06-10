package shell

import (
	"errors"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"strings"
)

func EditTextByEditor(initText string) (string, error) {
	var err error

	err = initFile("/tmp/git-desc.txt", initText)
	if err != nil {
		return "", errors.New("Initialize temp file error")
	}

	editor := viper.GetString("editor")
	if editor == "" {
		editor = "vi"
	}

	err = executeEditor(editor, "/tmp/git-desc.txt")
	if err != nil {
		return "", err
	}

	text, err := getFileText("/tmp/git-desc.txt")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(strings.Trim(text, "\n")), nil
}

func getFileText(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", errors.New("File open error: " + filePath)
	}

	buf := make([]byte, 1024)
	var text string
	for {
		count, err := file.Read(buf)
		if count == 0 {
			break
		}
		if err != nil {
			return "", errors.New("File read error.")
		}
		text = string(buf[:count])
	}
	return text, nil
}

func executeEditor(editor string, filePath string) error {
	cmd := exec.Command(editor, filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return errors.New("Can't started editor.")
	}
	return nil
}

func initFile(filePath string, initMessage string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return errors.New("File open error: " + filePath)
	}

	defer file.Close()

	file.Write(([]byte)(initMessage))
	return nil
}
