package cmd

import (
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

// langCodeRe matches valid BCP 47 / ISO 639-1 language codes: "en", "es", "zh-CN".
var langCodeRe = regexp.MustCompile(`^[a-z]{2,3}(-[A-Za-z]{2,4})?$`)

// validateLangCode returns an error if code is not a valid language code.
func validateLangCode(code string) error {
	if !langCodeRe.MatchString(code) {
		return fmt.Errorf("invalid language code %q: expected ISO 639-1 format like \"en\", \"es\", \"zh-CN\"", code)
	}
	return nil
}

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
