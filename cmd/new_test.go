package cmd

import (
	"os/exec"
	"testing"
)

func TestNewCmd(t *testing.T) {
	output, err := exec.Command("goapi", "new", "dummy").Output()

	expected := (`Creates a new model template, controller template, or project and names it the given arguemnt.
For example:
  goapi new model Users

Usage:
  goapi new [command]

Available Commands:
  model       Creates a new model template
  project     Creates a new webapi project boilerplate structure

Flags:
  -h, --help   help for new

Use "goapi new [command] --help" for more information about a command.
`)

	out := string(output)
	if err != nil {
		t.Errorf("Unexpected output: %v", out)
	}

	if expected != out {
		t.Errorf("\nExpected:\n%v\ngot:\n%v", expected, out)
	}
}
