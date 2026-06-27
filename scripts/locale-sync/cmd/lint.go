package cmd

import (
	"fmt"
	"os"

	"github.com/bobadilla-tech/locale-sync/internal/locale"
	"github.com/bobadilla-tech/locale-sync/internal/source"
	"github.com/spf13/cobra"
)

var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Exit 1 if any t() call references an undefined locale key",
	Long: `lint checks that every t('key') call in source files has a matching
YAML entry in the locale files. Missing keys silently break in production
(blank text or translation-missing errors). Use this in CI to catch them early.

Exit codes:
  0  all t() calls resolved
  1  one or more keys missing from YAML`,
	SilenceUsage: true,
	RunE:         runLint,
}

var lintLang string

func init() {
	lintCmd.Flags().StringVarP(&rootPath, "root", "r", ".", "Rails app root directory")
	lintCmd.Flags().StringVarP(&lintLang, "lang", "l", "en", "Locale language to check against")
	rootCmd.AddCommand(lintCmd)
}

func runLint(_ *cobra.Command, _ []string) error {
	entries, err := locale.Scan(rootPath, lintLang)
	if err != nil {
		return fmt.Errorf("scan locales: %w", err)
	}

	defined := make(map[string]bool, len(entries))
	for _, e := range entries {
		defined[e.Key] = true
	}

	result, err := source.Audit(rootPath, lintLang, defined)
	if err != nil {
		return fmt.Errorf("audit source: %w", err)
	}

	missing := dedupMissingKeys(result.MissingKeys)
	if len(missing) == 0 {
		fmt.Printf("lint: OK — all %d t() call(s) resolved.\n", len(result.UsedKeys))
		return nil
	}

	for _, u := range missing {
		fmt.Fprintf(os.Stderr, "%s:%d: error: missing key %q\n", u.File, u.Line, u.Key)
	}
	fmt.Fprintf(os.Stderr, "\nlint: %d missing key(s). Add with: locale-sync add-key --key <key> --value \"<value>\"\n", len(missing))
	os.Exit(1)
	return nil
}
