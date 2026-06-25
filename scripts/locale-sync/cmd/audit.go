package cmd

import (
	"fmt"

	"github.com/bobadilla-tech/locale-sync/internal/locale"
	"github.com/bobadilla-tech/locale-sync/internal/source"
	"github.com/spf13/cobra"
)

var (
	showOrphaned bool
	showMissing  bool
	auditAll     bool
)

var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Cross-reference YAML keys against source t() calls",
	Long: `audit finds:
  - Orphaned keys: defined in YAML but never called in source (dead translations)
  - Missing keys:  t('key') called in source but not defined in YAML`,
	RunE: runAudit,
}

func init() {
	auditCmd.Flags().StringVarP(&rootPath, "root", "r", ".", "Rails app root directory")
	auditCmd.Flags().StringVarP(&lang, "lang", "l", "en", "Locale language to audit")
	auditCmd.Flags().BoolVar(&showOrphaned, "orphaned", false, "Show orphaned (unused) keys")
	auditCmd.Flags().BoolVar(&showMissing, "missing", false, "Show missing (undefined) keys")
	auditCmd.Flags().BoolVar(&auditAll, "all", false, "Show both orphaned and missing keys")
	rootCmd.AddCommand(auditCmd)
}

func runAudit(cmd *cobra.Command, _ []string) error {
	if auditAll {
		showOrphaned = true
		showMissing = true
	}
	if !showOrphaned && !showMissing {
		// Default: show both
		showOrphaned = true
		showMissing = true
	}

	fmt.Printf("Auditing %s locale keys in %s...\n\n", lang, rootPath)

	entries, err := locale.Scan(rootPath, lang)
	if err != nil {
		return fmt.Errorf("scan locales: %w", err)
	}

	// Build set of defined keys.
	defined := make(map[string]bool, len(entries))
	for _, e := range entries {
		defined[e.Key] = true
	}
	fmt.Printf("Loaded %d locale keys from %d files.\n", len(defined), countFiles(entries))

	result, err := source.Audit(rootPath, lang, defined)
	if err != nil {
		return fmt.Errorf("audit source: %w", err)
	}
	fmt.Printf("Scanned %d t() calls in source files.\n\n", len(result.UsedKeys))

	if showOrphaned {
		printOrphaned(result.OrphanedKeys)
	}
	if showMissing {
		printMissing(result.MissingKeys)
	}

	fmt.Println("─────────────────────────────────────────")
	fmt.Printf("Summary:\n")
	fmt.Printf("  Defined keys:    %d\n", len(defined))
	fmt.Printf("  t() calls found: %d\n", len(result.UsedKeys))
	fmt.Printf("  Orphaned keys:   %d  (defined but never called)\n", len(result.OrphanedKeys))
	fmt.Printf("  Missing keys:    %d  (called but not defined)\n", len(result.MissingKeys))
	fmt.Println()

	return nil
}

func printOrphaned(keys []string) {
	if len(keys) == 0 {
		fmt.Println("No orphaned keys found.")
		return
	}
	fmt.Printf("=== Orphaned Keys (%d) — defined in YAML but never called ===\n\n", len(keys))
	for _, k := range keys {
		fmt.Printf("  %s\n", k)
	}
	fmt.Println()
}

func printMissing(usages []source.KeyUsage) {
	if len(usages) == 0 {
		fmt.Println("No missing keys found.")
		return
	}
	fmt.Printf("=== Missing Keys (%d) — called in source but not in YAML ===\n\n", len(usages))
	// Deduplicate by key.
	seen := map[string]bool{}
	for _, u := range usages {
		if seen[u.Key] {
			continue
		}
		seen[u.Key] = true
		fmt.Printf("  %-55s  %s:%d\n", u.Key, u.File, u.Line)
	}
	fmt.Println()
}
