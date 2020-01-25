package cmd

import (
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new model, controller, or project",
	Long: `Creates a new model template, controller template, or project and names it the given arguemnt.
For example:
  goapi new model Users`,
}

func init() {
	rootCmd.AddCommand(newCmd)
}
