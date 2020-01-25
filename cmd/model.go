package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// modelCmd represents the `new model` command.
var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "Creates a new model template",
	Long: `Creates a new model template and names it the given argument.
For example:
  goapi new model Users`,
	Args: cobra.ExactArgs(1),
	Run:  NewModel,
}

func init() {
	newCmd.AddCommand(modelCmd)
}

// NewModel will run when the `new model` subcommand is executed.
func NewModel(cmd *cobra.Command, args []string) {
	arg := args[0]
	cleanArg := CleanArg(arg)
	fmt.Printf("Model called with %v\n", cleanArg)
}
