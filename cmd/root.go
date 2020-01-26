package cmd

import (
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "geaux",
	Short: "A framework that creates a boilerplate api structure.",
}

// Execute the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// CleanArg strips any whitespace from given argument from command.
func CleanArg(arg string) string {
	space := regexp.MustCompile(`\s+`)
	arg = space.ReplaceAllString(arg, "")

	return arg
}
