package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "locale-sync",
	Short: "Extract and deduplicate i18n strings in Rails apps",
	Long: `locale-sync scans a Rails application and:
  - Finds duplicate YAML locale keys across multiple files
  - Finds identical string values stored under different keys
  - Extracts hardcoded strings from Ruby/ERB source files
  - Generates shared locale files and language skeletons`,
}

// Execute is the entry point called from main.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
