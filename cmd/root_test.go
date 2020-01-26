package cmd

import (
	"os/exec"
	"testing"
)

func TestRootCmd(t *testing.T) {
	output, err := exec.Command("geaux").Output()

	expected := (`A framework that creates a boilerplate api structure.

Usage:
  geaux [command]

Available Commands:
  help        Help about any command
  new         Creates a new model, controller, or project

Flags:
  -h, --help   help for goapi

Use "geaux [command] --help" for more information about a command.
`)

	if err != nil {
		t.Errorf("Unexpected err: %v", err)
	}

	out := string(output)
	if expected != out {
		t.Errorf("\nExpected:\n%v\ngot:\n%v", expected, out)
	}
}

func TestRootCmdError(t *testing.T) {
	// Expect an error to occur
	output, err := exec.Command("goapi", "dummy").Output()

	if err == nil {
		t.Errorf("Unexpected output: %v", string(output))
	}
}
