package shell

import (
	"bytes"
	"os/exec"
	"strings"
)

type Command interface {
	Run(name string, arg ...string) (string, error)
}

type CommandImpl struct{}

func SetupCommand() Command {
	return CommandImpl{}
}

func (c CommandImpl) Run(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	str := stdout.String()
	return strings.Trim(str, "\n"), nil
}
