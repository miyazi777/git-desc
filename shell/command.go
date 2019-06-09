package shell

import (
	"bytes"
	"os/exec"
	"strings"
)

func Run(name string, arg ...string) (string, error) {
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
