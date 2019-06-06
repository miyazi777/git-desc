package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var err error
	err = initFile("/tmp/git-desc.txt", "test xxx")
	if err != nil {
		fmt.Println("Initialize temp file error")
		return
	}

	editor := os.Getenv("EDITOR")
	if editor == "" {
		fmt.Println("Please set env $EDITOR.")
		return
	}

	err = executeEditor(editor, "/tmp/git-desc.txt")
	if err != nil {
		return
	}
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
