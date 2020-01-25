package cmd

import (
	"os/exec"
	"testing"
)

func TestRootCmd(t *testing.T) {
	output, err := exec.Command("goapi").Output()

	expected := (`A framework that creates a boilerplate api structure.

Usage:
  goapi [command]

Available Commands:
  help        Help about any command
  new         Creates a new model, controller, or project

Flags:
  -h, --help   help for goapi

Use "goapi [command] --help" for more information about a command.
`)

	out := string(output)
	if err != nil {
		t.Errorf("Unexpected err: %v", err)
	}

	if expected != out {
		t.Errorf("Expected:\n %v\n got:\n %v", expected, out)
	}
}

func TestRootDummyCmd(t *testing.T) {
	// Expect an error to occur
	output, err := exec.Command("goapi", "dummy").Output()

	if err == nil {
		t.Errorf("Unexpected output: %v", string(output))
	}

}
