package goshell

import (
	"os/exec"
	"strings"
)

//Cmd executes the given string as shell command
func Cmd(arg string) (string, error) {
	args := strings.Split(arg, " ")
	out, err := exec.Command(args[0], args[1:]...).CombinedOutput()
	return string(out), err
}
