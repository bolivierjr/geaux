package cmd

import (
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
)

// modelCmd represents the model subcommand.
var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "Creates a new model template",
	Long: `Creates a new model template and names it the given argument.
For example:
  goapi new model Users`,
	Args: cobra.ExactArgs(1),
	Run:  Model,
}

func init() {
	newCmd.AddCommand(modelCmd)
}

// Model will run when the model subcommand is executed.
func Model(cmd *cobra.Command, args []string) {
	arg := args[0]
	space := regexp.MustCompile(`\s+`)
	arg = space.ReplaceAllString(arg, "")
	fmt.Printf("Model called with %v\n", arg)
}
