package goshell

import (
	"os/exec"
	"strings"
)

// CMD basic structure with output and pid
type CMD struct {
	Output string
	Pid    int
}

//Cmd executes the given string as shell command
func Cmd(arg string) (CMD, error) {
	args := strings.Split(arg, " ")
	cmd := exec.Command(args[0], args[1:]...)
	out, err := cmd.CombinedOutput()
	return CMD{Output: string(out), Pid: cmd.Process.Pid}, err
}
