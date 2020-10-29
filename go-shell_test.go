package goshell

import (
	"strings"
	"testing"
)

func TestCMD(t *testing.T) {
	cmd, err := Cmd("ls")
	if err != nil {
		t.Error(err)
		return
	}
	cmdRows := []string{}
	for _, item := range strings.Split(cleanString(cmd.output), "\n") {
		if len(item) > 0 {
			cmdRows = append(cmdRows, item)
		}
	}

	// expects the existence of two files
	if len(cmdRows) != 2 {
		t.Errorf("More files than intended: %d", len(cmdRows))
		return
	}
}

func cleanString(input string) string {
	input = strings.Replace(input, " ", "", -1)
	input = strings.Replace(input, "\r", "", -1)
	return input
}
