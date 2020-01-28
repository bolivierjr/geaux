package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/bolivierjr/geaux/internal/templates"
	fs "github.com/bolivierjr/geaux/pkg/fs"
	"github.com/spf13/cobra"
)

// projectCmd represents the `new project` command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Creates a new webapi project boilerplate structure",
	Long: `Creates a new webapi project boilerplate structure and names the directory after the given argument.
For example:
  geaux new project webapi`,
	Args: cobra.ExactArgs(1),
	Run:  NewProject,
}

func init() {
	newCmd.AddCommand(projectCmd)
}

// NewProject will run when the `new project` subcommand is executed.
func NewProject(cmd *cobra.Command, args []string) {
	arg := args[0]
	cleanArg := CleanArg(arg)
	fmt.Printf("Model called with %v\n", cleanArg)

	currentDir, _ := os.Getwd()
	newPath := filepath.Join(currentDir, arg)

	// Create a new root project directory
	if ok := fs.MakeDir(newPath); !ok {
		fmt.Println("Directory already exists.")
		os.Exit(1)
	}
}
