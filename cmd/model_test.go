package cmd

import (
	"os/exec"
	"testing"
)

func TestModelCmdError(t *testing.T) {
	output, err := exec.Command("goapi", "new", "model").Output()

	if err == nil {
		t.Errorf("Unexpected output: %v", string(output))
	}

	expected := "accepts 1 arg(s), received 0\n"

	out := string(output)
	if expected != out {
		t.Errorf("\nExpected:\n%v\ngot:\n%v", expected, out)
	}
}
