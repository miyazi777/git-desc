package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	initFileErr := initFile("/tmp/git-desc.txt", "test test test")
	if initFileErr != nil {
		fmt.Println("Initialize temp file error")
		return
	}

	editor := os.Getenv("EDITOR")
	if editor == "" {
		fmt.Println("Please set env $EDITOR.")
		return
	}

	cmd := exec.Command(editor, "/tmp/git-desc.txt")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmdErr := cmd.Run()
	if cmdErr != nil {
		fmt.Println("Can't started editor.")
		return
	}
}

func initFile(filePath string, initMessage string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return errors.New("File open error: ")
	}

	defer file.Close()

	file.Write(([]byte)(initMessage))
	return nil
}
