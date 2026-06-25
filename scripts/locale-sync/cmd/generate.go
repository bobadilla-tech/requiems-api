package cmd

import (
	"fmt"
	"strings"

	"github.com/bobadilla-tech/locale-sync/internal/locale"
	"github.com/bobadilla-tech/locale-sync/internal/source"
	"github.com/spf13/cobra"
)

var (
	languages   []string
	placeholder string
	dryRun      bool
	overwrite   bool
	noShared    bool
	noSkeleton  bool
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Consolidate duplicates into shared YAML and generate language skeletons",
	Long: `generate does everything analyze does, then:
  - Writes duplicate keys with identical values into shared.{lang}.yml
  - Creates skeleton locale files for each --languages value (empty strings)
  - Prints a migration plan showing which t() keys to use`,
	RunE: runGenerate,
}

func init() {
	generateCmd.Flags().StringVarP(&rootPath, "root", "r", ".", "Rails app root directory")
	generateCmd.Flags().StringVarP(&lang, "lang", "l", "en", "Source locale language")
	generateCmd.Flags().IntVar(&minDups, "min-dups", 2, "Minimum occurrences to consolidate")
	generateCmd.Flags().StringSliceVar(&languages, "languages", nil,
		"Additional languages to generate skeletons for (e.g. es,fr,de)")
	generateCmd.Flags().StringVar(&placeholder, "placeholder", "",
		"Value used in skeleton files for untranslated strings (default: empty)")
	generateCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Preview changes without writing files")
	generateCmd.Flags().BoolVar(&overwrite, "overwrite", false, "Overwrite existing skeleton files")
	generateCmd.Flags().BoolVar(&noShared, "no-shared", false, "Skip writing the shared YAML file")
	generateCmd.Flags().BoolVar(&noSkeleton, "no-skeleton", false, "Skip generating language skeletons")
	generateCmd.Flags().BoolVar(&scanSource, "source", false,
		"Also scan Ruby/ERB source for hardcoded strings")
	rootCmd.AddCommand(generateCmd)
}

func runGenerate(cmd *cobra.Command, _ []string) error {
	fmt.Printf("Scanning %s locale files in %s...\n\n", lang, rootPath)

	entries, err := locale.Scan(rootPath, lang)
	if err != nil {
		return fmt.Errorf("scan locales: %w", err)
	}
	fmt.Printf("Found %d locale entries across %d files.\n\n", len(entries), countFiles(entries))

	dups := locale.FindDuplicates(entries, minDups)

	// --- Shared file consolidation ---
	if !noShared {
		candidates := locale.BuildMergeCandidates(dups.KeyDups, lang)
		if len(candidates) == 0 {
			fmt.Println("No auto-mergeable duplicates found — shared file unchanged.")
		} else {
			fmt.Printf("=== Shared File Consolidation (%d keys) ===\n\n", len(candidates))
			plan := locale.MigrationPlan(candidates)

			for _, c := range candidates {
				fmt.Printf("  %q → t('%s')\n",
					c.KeyName, dotAfterLang(c.SuggestedKey))
				fmt.Printf("    value: %q\n", c.Value)
				fmt.Printf("    from:  %s\n\n",
					strings.Join(sourceFiles(c.Sources), ", "))
			}

			if !dryRun {
				sharedPath, changed, err := locale.UpsertShared(rootPath, lang, candidates)
				if err != nil {
					return fmt.Errorf("write shared: %w", err)
				}
				if changed {
					fmt.Printf("  Written: %s\n\n", sharedPath)
				} else {
					fmt.Printf("  No new entries (all already present in shared file).\n\n")
				}
			} else {
				fmt.Println("  [dry-run] Would write shared file.")
			}

			// Print migration plan
			fmt.Println("=== Migration Plan ===")
			fmt.Println("Replace these keys in your existing locale files:")
			for _, oldKey := range locale.SortedKeys(plan) {
				fmt.Printf("  %-60s → t('%s')\n", oldKey, plan[oldKey])
			}
			fmt.Println()
		}
	}

	// --- Language skeleton generation ---
	if !noSkeleton && len(languages) > 0 {
		fmt.Printf("=== Generating Skeletons for: %s ===\n\n", strings.Join(languages, ", "))
		for _, targetLang := range languages {
			targetLang = strings.TrimSpace(targetLang)
			if targetLang == "" || targetLang == lang {
				continue
			}
			if dryRun {
				fmt.Printf("  [dry-run] Would generate skeleton for %q.\n", targetLang)
				continue
			}
			written, err := locale.WriteSkeleton(rootPath, lang, targetLang, placeholder, entries, overwrite)
			if err != nil {
				return fmt.Errorf("skeleton %s: %w", targetLang, err)
			}
			if len(written) == 0 {
				fmt.Printf("  %s: all files already exist (use --overwrite to regenerate).\n", targetLang)
			} else {
				for _, p := range written {
					fmt.Printf("  Created: %s\n", p)
				}
			}
		}
		fmt.Println()
	}

	// --- Source scan ---
	if scanSource {
		fmt.Println("=== Scanning Source Files for Hardcoded Strings ===")
		found, err := source.Extract(rootPath)
		if err != nil {
			return fmt.Errorf("source scan: %w", err)
		}
		if len(found) == 0 {
			fmt.Println("No hardcoded user-facing strings detected.")
		} else {
			fmt.Printf("Found %d hardcoded strings to extract:\n\n", len(found))
			for _, s := range found {
				fmt.Printf("  %s:%d  %q\n", s.File, s.Line, s.Text)
			}
		}
		fmt.Println()
	}

	return nil
}

func sourceFiles(entries []locale.Entry) []string {
	seen := map[string]bool{}
	var out []string
	for _, e := range entries {
		if !seen[e.ShortPath] {
			seen[e.ShortPath] = true
			out = append(out, e.ShortPath)
		}
	}
	return out
}
