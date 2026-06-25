package cmd

import (
	"fmt"

	"github.com/bobadilla-tech/locale-sync/internal/locale"
	"github.com/bobadilla-tech/locale-sync/internal/source"
	"github.com/spf13/cobra"
)

var (
	rootPath   string
	lang       string
	minDups    int
	scanSource bool
	showAll    bool
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Scan locale files and report duplicate keys/values",
	Long: `analyze scans config/locales/{lang}/ for:
  - Duplicate leaf key names (same key in multiple files)
  - Duplicate text values (same string under different keys)
  - (optional) Hardcoded strings in Ruby/ERB source files`,
	RunE: runAnalyze,
}

func init() {
	analyzeCmd.Flags().StringVarP(&rootPath, "root", "r", ".", "Rails app root directory")
	analyzeCmd.Flags().StringVarP(&lang, "lang", "l", "en", "Locale language to scan")
	analyzeCmd.Flags().IntVar(&minDups, "min-dups", 2, "Minimum occurrences to flag as duplicate")
	analyzeCmd.Flags().BoolVar(&scanSource, "source", false, "Also scan Ruby/ERB source for hardcoded strings")
	analyzeCmd.Flags().BoolVar(&showAll, "all", false, "Show all duplicate key names, including those with different values")
	rootCmd.AddCommand(analyzeCmd)
}

func runAnalyze(cmd *cobra.Command, _ []string) error {
	fmt.Printf("Scanning %s locale files in %s...\n\n", lang, rootPath)

	entries, err := locale.Scan(rootPath, lang)
	if err != nil {
		return fmt.Errorf("scan locales: %w", err)
	}
	fmt.Printf("Found %d locale entries across %d files.\n", len(entries), countFiles(entries))

	dups := locale.FindDuplicates(entries, minDups)

	printKeyDups(dups.KeyDups)
	printValueDups(dups.ValueDups)

	if scanSource {
		if err := runSourceScan(rootPath); err != nil {
			return err
		}
	}

	printSummary(dups, entries)
	return nil
}

func printKeyDups(groups []locale.KeyDupGroup) {
	// By default only show auto-mergeable (same value) groups; --all shows all.
	var displayed []locale.KeyDupGroup
	for _, g := range groups {
		if g.SameValue || showAll {
			displayed = append(displayed, g)
		}
	}

	autoMergeCount := 0
	for _, g := range groups {
		if g.SameValue {
			autoMergeCount++
		}
	}

	if len(displayed) == 0 {
		fmt.Printf("No duplicate key names with identical values found (%d with differing values, use --all to show).\n",
			len(groups))
		return
	}

	label := "auto-mergeable"
	if showAll {
		label = "total"
	}
	fmt.Printf("\n=== Duplicate Key Names (%d %s) ===\n", len(displayed), label)
	if !showAll && len(groups) > autoMergeCount {
		fmt.Printf("  (%d keys omitted with differing values — use --all to show)\n", len(groups)-autoMergeCount)
	}

	for _, g := range displayed {
		indicator := "~"
		if g.SameValue {
			indicator = "✓"
		}
		fmt.Printf("\n  [%s] %q (%d occurrences)\n", indicator, g.KeyName, len(g.Entries))
		for _, e := range g.Entries {
			fmt.Printf("      %s:%d  %q\n", e.ShortPath, e.Line, e.Value)
		}
		if g.SameValue {
			fmt.Printf("      → suggest: t('%s')\n",
				dotAfterLang(locale.SuggestSharedKey(lang, g.KeyName)))
		} else {
			fmt.Printf("      → values differ — manual review needed\n")
		}
	}
	fmt.Println()
}

func printValueDups(groups []locale.ValueDupGroup) {
	if len(groups) == 0 {
		fmt.Println("No duplicate text values found.")
		return
	}
	fmt.Printf("\n=== Duplicate Text Values (%d unique values) ===\n", len(groups))
	for _, g := range groups {
		fmt.Printf("\n  %q (%d keys share this text)\n", g.Value, len(g.Entries))
		for _, e := range g.Entries {
			fmt.Printf("      %s  →  %s\n", e.ShortPath, e.Key)
		}
	}
	fmt.Println()
}

func runSourceScan(root string) error {
	fmt.Println("=== Scanning Source Files for Hardcoded Strings ===")
	found, err := source.Extract(root)
	if err != nil {
		return fmt.Errorf("source scan: %w", err)
	}
	if len(found) == 0 {
		fmt.Println("No hardcoded user-facing strings found.")
		return nil
	}
	fmt.Printf("Found %d hardcoded strings:\n\n", len(found))
	for _, s := range found {
		fmt.Printf("  %s:%d\n", s.File, s.Line)
		fmt.Printf("    text:    %q\n", s.Text)
		fmt.Printf("    context: %s\n\n", s.Context)
	}
	return nil
}

func printSummary(dups locale.Duplicates, entries []locale.Entry) {
	autoMerge := 0
	for _, g := range dups.KeyDups {
		if g.SameValue {
			autoMerge++
		}
	}
	fmt.Println("─────────────────────────────────────────")
	fmt.Printf("Summary:\n")
	fmt.Printf("  Total entries:                    %d\n", len(entries))
	fmt.Printf("  Duplicate key names (total):      %d\n", len(dups.KeyDups))
	fmt.Printf("  Auto-mergeable (same value):      %d\n", autoMerge)
	fmt.Printf("  Duplicate text values:            %d\n", len(dups.ValueDups))
	fmt.Println()
	if autoMerge > 0 {
		fmt.Printf("Run `locale-sync generate --root %s` to consolidate duplicates into shared file.\n", rootPath)
	}
}

func countFiles(entries []locale.Entry) int {
	seen := map[string]bool{}
	for _, e := range entries {
		seen[e.File] = true
	}
	return len(seen)
}

// dotAfterLang strips the leading "<lang>." from a suggested key for t() usage.
func dotAfterLang(key string) string {
	for i, c := range key {
		if c == '.' {
			return key[i+1:]
		}
	}
	return key
}
