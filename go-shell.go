package goshell

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Cmd executes the given string as shell command
//
// Besided the PID as int, returns the output as string.
func Cmd(arg string) (string, error) {
	args := strings.Split(arg, " ")
	out, err := exec.Command(args[0], args[1:]...).CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// LooseCmd executes the given string as shell command
// and send it to the background automatically
//
// Besided the PID as int, returns the output as string.
func LooseCmd(arg string) (int, error) {
	args := strings.Split(arg, " ")
	process, err := os.StartProcess(
		args[0],
		args,
		&os.ProcAttr{
			Dir: ".",
			Env: os.Environ(),
			Files: []*os.File{
				os.Stdin,
				nil,
				nil,
			},
		})
	if err != nil {
		return 0, fmt.Errorf("os.StartProcess err: %v", err)
	}
	pid := process.Pid
	err = process.Release()
	if err != nil {
		return 0, fmt.Errorf("process.Release err: %v", err)
	}
	return pid, nil
}
