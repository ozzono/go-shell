package goshell

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"
)

func TestCMD(t *testing.T) {
	err := testCmd((t))
	if err != nil {
		t.Error((err))
		return
	}

	err = testLooseCmd((t))
	if err != nil {
		t.Error((err))
		return
	}
}

func testLooseCmd(t *testing.T) error {
	sleep := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(999) + 5
	t.Logf("starting testLooseCmd")
	pid, err := LooseCmd(fmt.Sprintf("/bin/sleep %d", sleep))
	if err != nil {
		return fmt.Errorf("LooseCmd err: %v", err)
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("FindProcess err: %v", err)
	}
	t.Logf("pid: %d", pid)
	err = proc.Kill()
	if err != nil {
		return fmt.Errorf("Kill err: %v", err)
	}
	t.Logf("successfully tested LooseCmd")
	return nil
}

func testCmd(t *testing.T) error {
	t.Logf("starting testCmd")
	out, err := Cmd("ls")
	if err != nil {
		t.Error(err)
		return err
	}
	outRows := []string{}
	for _, item := range strings.Split(cleanString(string(out)), "\n") {
		if len(item) > 0 {
			outRows = append(outRows, item)
		}
	}

	// expects the existence of two visible files
	// go-shell.go
	// go-shell_test.go
	if len(outRows) != 2 {
		return fmt.Errorf("More files than intended: %d", len(outRows))
	}

	t.Logf("successfully tested Cmd")
	return nil
}

func cleanString(input string) string {
	input = strings.Replace(input, " ", "", -1)
	input = strings.Replace(input, "\r", "", -1)
	return input
}
