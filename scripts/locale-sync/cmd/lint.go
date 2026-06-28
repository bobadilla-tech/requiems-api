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
	Short: "Exit 1 if any t() call is undefined or used in an invalid Ruby context",
	Long: `lint runs two checks and fails fast in CI:

  1. Missing keys — every t('key') call must have a matching YAML entry.
     Missing keys produce blank text or "translation missing" errors in production.

  2. Bare t() in non-helper contexts — bare t() is only available in controllers,
     mailers, and helpers (ActionController/ActionMailer inject it). Using t() in
     models, jobs, services, workers, or lib raises NoMethodError at runtime.
     Use I18n.t() in those contexts instead.

Exit codes:
  0  all checks pass
  1  one or more issues found`,
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

	bareIssues, err := source.CheckBareTCalls(rootPath)
	if err != nil {
		return fmt.Errorf("check bare t() calls: %w", err)
	}

	missing := dedupMissingKeys(result.MissingKeys)

	if len(missing) == 0 && len(bareIssues) == 0 {
		fmt.Printf("lint: OK — all %d t() call(s) resolved.\n", len(result.UsedKeys))
		return nil
	}

	for _, u := range missing {
		fmt.Fprintf(os.Stderr, "%s:%d: error: missing key %q\n", u.File, u.Line, u.Key)
	}
	if len(missing) > 0 {
		fmt.Fprintf(os.Stderr, "\nlint: %d missing key(s). Add with: locale-sync add-key --key <key> --value \"<value>\"\n", len(missing))
	}

	for _, b := range bareIssues {
		fmt.Fprintf(os.Stderr, "%s:%d: error: bare t(%q) in non-helper context — use I18n.t(%q)\n",
			b.File, b.Line, b.Key, b.Key)
	}
	if len(bareIssues) > 0 {
		fmt.Fprintf(os.Stderr, "\nlint: %d bare t() call(s) in invalid context.\n", len(bareIssues))
	}

	os.Exit(1)
	return nil
}
