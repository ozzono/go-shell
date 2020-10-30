package goshell

import (
	"fmt"
	"log"
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

	err = testLooseCmd1((t))
	if err != nil {
		t.Error((err))
		return
	}

	err = testLooseCmd2((t))
	if err != nil {
		t.Error((err))
		return
	}
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

func testLooseCmd1(t *testing.T) error {
	t.Logf("starting testLooseCmd1")
	sleep := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(999) + 5
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

func testLooseCmd2(t *testing.T) error {
	t.Logf("starting testLooseCmd2")
	emulator := fmt.Sprintf("%s/Android/Sdk/emulator/emulator", os.Getenv("HOME"))
	deviceName := "lite"

	if find("emulator", "-avd", deviceName) {
		return fmt.Errorf("emulator %s already running", deviceName)
	}

	pid, err := LooseCmd(fmt.Sprintf("%s -avd %s", emulator, deviceName))
	if err != nil {
		return err
	}
	time.Sleep(10 & time.Second)
	if !find("emulator", "-avd", deviceName) {
		return fmt.Errorf("emulator not found; command pid: %d", pid)
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("FindProcess err: %v", err)
	}
	proc.Kill()

	if find("emulator", "-avd", deviceName) {
		return fmt.Errorf("failed to kill emulator; pid: %d", pid)
	}

	t.Logf("successfully tested LooseCmd again")
	return nil
}

func cleanString(input string) string {
	input = strings.Replace(input, " ", "", -1)
	input = strings.Replace(input, "\r", "", -1)
	return input
}

func find(want ...string) bool {
	out, err := Cmd("/bin/ps -ef")
	if err != nil {
		log.Printf("Cmd err: %v", err)
		return false
	}
	for _, item := range strings.Split(out, "\n") {
		if arrCheck(want, item) {
			return true
		}
	}
	return false
}

func arrCheck(want []string, have string) bool {
	for i := range want {
		if !strings.Contains(have, want[i]) {
			return false
		}
	}
	return true
}
