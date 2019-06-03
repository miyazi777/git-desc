package terminal

import (
	"errors"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetWidth() (int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 0, errors.New("Error stty command")
	}

	cols := strings.Split(string(out), " ")
	width, _ := strconv.Atoi(strings.Trim(cols[1], "\n"))
	return width, nil
}
