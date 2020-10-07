package goshell

import (
	"log"
	"os/exec"
	"strings"
)

//Cmd executes the given string as shell command
func Cmd(arg string) string {
	log.Printf("shell: %v", arg)
	args := strings.Split(arg, " ")
	out, err := exec.Command(args[0], args[1:]...).CombinedOutput()
	if err != nil {
		log.Printf("Command: '%v';\nOutput: %v;\nError: %v\n", arg, string(out), err)
		return err.Error()
	}
	return string(out)
}
