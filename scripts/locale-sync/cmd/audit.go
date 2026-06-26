package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bobadilla-tech/locale-sync/internal/locale"
	"github.com/bobadilla-tech/locale-sync/internal/source"
	"github.com/spf13/cobra"
)

var (
	showOrphaned    bool
	showMissing     bool
	auditAll        bool
	showEmptyValues bool
	compareLocale   string
)

var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Cross-reference YAML keys against source t() calls",
	Long: `audit finds:
  - Orphaned keys:     defined in YAML but never called in source (dead translations)
  - Missing keys:      t('key') called in source but not defined in YAML
  - Empty values:      keys defined with an empty string value (--empty-values)
  - Missing in locale: keys present in one locale but absent in another (--compare-locale)`,
	RunE: runAudit,
}

func init() {
	auditCmd.Flags().StringVarP(&rootPath, "root", "r", ".", "Rails app root directory")
	auditCmd.Flags().StringVarP(&lang, "lang", "l", "en", "Locale language to audit")
	auditCmd.Flags().BoolVar(&showOrphaned, "orphaned", false, "Show orphaned (unused) keys")
	auditCmd.Flags().BoolVar(&showMissing, "missing", false, "Show missing (undefined) keys")
	auditCmd.Flags().BoolVar(&auditAll, "all", false, "Show both orphaned and missing keys")
	auditCmd.Flags().BoolVar(&showEmptyValues, "empty-values", false, "Show keys whose value is an empty string")
	auditCmd.Flags().StringVar(&compareLocale, "compare-locale", "", "Second locale to compare key coverage against (e.g. fr)")
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

	// Deduplicate missing keys once so the summary count matches the displayed list.
	uniqueMissing := dedupMissingKeys(result.MissingKeys)

	if showOrphaned {
		printOrphaned(result.OrphanedKeys)
	}
	if showMissing {
		printMissing(uniqueMissing)
	}
	if showEmptyValues {
		printEmptyValues(entries)
	}

	// Cross-locale key coverage comparison.
	var missingInOther, missingInBase []string
	if compareLocale != "" {
		otherEntries, err := locale.Scan(rootPath, compareLocale)
		if err != nil {
			return fmt.Errorf("scan %s locales: %w", compareLocale, err)
		}
		missingInOther, missingInBase = diffLocaleKeys(entries, otherEntries, lang, compareLocale)
		printLocaleDiff(lang, compareLocale, missingInOther, missingInBase)
	}

	fmt.Println("─────────────────────────────────────────")
	fmt.Printf("Summary:\n")
	fmt.Printf("  Defined keys:    %d\n", len(defined))
	fmt.Printf("  t() calls found: %d\n", len(result.UsedKeys))
	fmt.Printf("  Orphaned keys:   %d  (defined but never called)\n", len(result.OrphanedKeys))
	fmt.Printf("  Missing keys:    %d  (called but not defined)\n", len(uniqueMissing))
	if n := len(result.RelativeKeys); n > 0 {
		fmt.Printf("  Relative keys:   %d  (t('.key') — skipped, need view-path resolver)\n", n)
	}
	if showEmptyValues {
		emptyCount := countEmptyValues(entries)
		fmt.Printf("  Empty values:    %d  (defined but blank string)\n", emptyCount)
	}
	if compareLocale != "" {
		fmt.Printf("  Missing in %s:  %d  (in %s but not %s)\n", compareLocale, len(missingInOther), lang, compareLocale)
		fmt.Printf("  Missing in %s:   %d  (in %s but not %s)\n", lang, len(missingInBase), compareLocale, lang)
	}
	fmt.Println()

	return nil
}

// printEmptyValues reports locale entries whose value is an empty string.
func printEmptyValues(entries []locale.Entry) {
	var empty []locale.Entry
	for _, e := range entries {
		if e.Value == "" {
			empty = append(empty, e)
		}
	}
	if len(empty) == 0 {
		fmt.Println("No empty-value keys found.")
		return
	}
	fmt.Printf("=== Empty Values (%d) — defined but blank, will render as missing text ===\n\n", len(empty))
	for _, e := range empty {
		fmt.Printf("  %-60s  %s:%d\n", e.Key, e.ShortPath, e.Line)
	}
	fmt.Println()
}

func countEmptyValues(entries []locale.Entry) int {
	n := 0
	for _, e := range entries {
		if e.Value == "" {
			n++
		}
	}
	return n
}

// diffLocaleKeys compares the leaf key sets of two scanned locales and returns
// keys present in base but absent in other, and vice versa.
func diffLocaleKeys(baseEntries, otherEntries []locale.Entry, baseLang, otherLang string) (missingInOther, missingInBase []string) {
	baseKeys := make(map[string]bool, len(baseEntries))
	for _, e := range baseEntries {
		// Strip leading lang prefix so "en.admin.foo" → "admin.foo"
		baseKeys[strings.TrimPrefix(e.Key, baseLang+".")] = true
	}
	otherKeys := make(map[string]bool, len(otherEntries))
	for _, e := range otherEntries {
		otherKeys[strings.TrimPrefix(e.Key, otherLang+".")] = true
	}

	for k := range baseKeys {
		if !otherKeys[k] {
			missingInOther = append(missingInOther, k)
		}
	}
	for k := range otherKeys {
		if !baseKeys[k] {
			missingInBase = append(missingInBase, k)
		}
	}
	sort.Strings(missingInOther)
	sort.Strings(missingInBase)
	return
}

func printLocaleDiff(baseLang, otherLang string, missingInOther, missingInBase []string) {
	if len(missingInOther) == 0 && len(missingInBase) == 0 {
		fmt.Printf("Locales %s and %s have identical key coverage.\n\n", baseLang, otherLang)
		return
	}
	if len(missingInOther) > 0 {
		fmt.Printf("=== Keys in %s but missing in %s (%d) ===\n\n", baseLang, otherLang, len(missingInOther))
		for _, k := range missingInOther {
			fmt.Printf("  %s\n", k)
		}
		fmt.Println()
	}
	if len(missingInBase) > 0 {
		fmt.Printf("=== Keys in %s but missing in %s (%d) ===\n\n", otherLang, baseLang, len(missingInBase))
		for _, k := range missingInBase {
			fmt.Printf("  %s\n", k)
		}
		fmt.Println()
	}
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
	for _, u := range usages {
		fmt.Printf("  %-55s  %s:%d\n", u.Key, u.File, u.Line)
	}
	fmt.Println()
}

// dedupMissingKeys returns one KeyUsage per unique key (first occurrence wins).
func dedupMissingKeys(usages []source.KeyUsage) []source.KeyUsage {
	seen := map[string]bool{}
	out := usages[:0:0]
	for _, u := range usages {
		if !seen[u.Key] {
			seen[u.Key] = true
			out = append(out, u)
		}
	}
	return out
}
