package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/bobadilla-tech/locale-sync/internal/locale"
	"github.com/spf13/cobra"
)

var consolidateNoPrune bool

var consolidateCmd = &cobra.Command{
	Use:   "consolidate",
	Short: "Promote duplicate topic-file keys to shared and rewrite all callers",
	Long: `consolidate runs the full shared-file migration in one shot:

  1. Finds duplicate keys across topic files (same as 'generate')
  2. Writes missing keys into shared.{lang}.yml
  3. Rewrites every t() caller — absolute AND lazy (t('.key')) — to the new shared path
  4. Deletes the now-orphaned old keys from topic YAML files

Always run with --dry-run first to review what will change.

Examples:
  locale-sync consolidate --root ../../apps/dashboard --lang en --dry-run
  locale-sync consolidate --root ../../apps/dashboard --lang en
  locale-sync consolidate --root ../../apps/dashboard --lang en --no-prune`,
	RunE: runConsolidate,
}

func init() {
	consolidateCmd.Flags().StringVarP(&rootPath, "root", "r", ".", "Rails app root directory")
	consolidateCmd.Flags().StringVarP(&lang, "lang", "l", "en", "Source locale language")
	consolidateCmd.Flags().IntVar(&minDups, "min-dups", 2, "Minimum occurrences to consolidate")
	consolidateCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Preview changes without writing files")
	consolidateCmd.Flags().BoolVar(&consolidateNoPrune, "no-prune", false,
		"Rewrite callers but skip deleting old YAML keys (run prune manually after review)")
	rootCmd.AddCommand(consolidateCmd)
}

func runConsolidate(_ *cobra.Command, _ []string) error {
	if err := validateLangCode(lang); err != nil {
		return fmt.Errorf("--lang: %w", err)
	}

	fmt.Printf("Scanning %s locale files in %s...\n", lang, rootPath)
	entries, err := locale.Scan(rootPath, lang)
	if err != nil {
		return fmt.Errorf("scan locales: %w", err)
	}
	fmt.Printf("Found %d entries across %d files.\n\n", len(entries), countFiles(entries))

	dups := locale.FindDuplicates(entries, minDups)
	candidates := locale.BuildMergeCandidates(dups.KeyDups, lang)
	candidates = locale.FilterExistingValueCandidates(candidates, entries)

	if len(candidates) == 0 {
		fmt.Println("No consolidation candidates found.")
		return nil
	}
	fmt.Printf("Found %d consolidation candidate(s).\n\n", len(candidates))

	// Write new keys to shared file (skip in dry-run).
	sharedFilePath := filepath.Join(rootPath, "config", "locales", lang, fmt.Sprintf("shared.%s.yml", lang))
	if !dryRun {
		_, changed, err := locale.UpsertShared(rootPath, lang, candidates)
		if err != nil {
			return fmt.Errorf("write shared: %w", err)
		}
		if changed {
			fmt.Printf("Wrote new key(s) to %s\n\n", shortPath(sharedFilePath, rootPath))
		} else {
			fmt.Printf("All keys already present in %s\n\n", shortPath(sharedFilePath, rootPath))
		}
	} else {
		fmt.Printf("[dry-run] Would write candidate key(s) to %s\n\n", shortPath(sharedFilePath, rootPath))
	}

	// Build migration plan: old full key → new t() path.
	plan := locale.MigrationPlan(candidates)

	// Build entry lookup: full key → locale file path.
	keyToFile := make(map[string]string, len(entries))
	for _, e := range entries {
		keyToFile[e.Key] = e.File
	}

	totalSourceFiles := 0
	totalYAMLFiles := 0
	rewroteSourceFiles := map[string]bool{}
	rewroteYAMLFiles := map[string]bool{}

	for idx, c := range candidates {
		// Collect all old keys for this candidate.
		var oldKeys []string
		for _, src := range c.Sources {
			oldKeys = append(oldKeys, src.Key)
		}
		newTPath := plan[oldKeys[0]]
		if newTPath == "" {
			continue
		}

		fmt.Printf("[%d/%d] t('%s')  ←  %s\n", idx+1, len(candidates), newTPath, summariseOldKeys(oldKeys, lang))

		callerCount := 0
		for _, oldFullKey := range oldKeys {
			oldKeyNoLang := stripLangPrefixStr(oldFullKey, lang)

			callers, err := findTCallersExtended(rootPath, oldKeyNoLang)
			if err != nil {
				fmt.Printf("  warning: scan callers for %s: %v\n", oldKeyNoLang, err)
			}

			for _, cal := range callers {
				kind := "absolute"
				if cal.relative {
					kind = "lazy"
				}
				fmt.Printf("  %s:%d  (%s)\n", cal.relFile, cal.line, kind)
				callerCount++
			}

			if !dryRun {
				if err := applyCallerRewritesExtended(rootPath, callers, oldKeyNoLang, newTPath); err != nil {
					fmt.Printf("  warning: rewrite callers: %v\n", err)
				}
				for _, cal := range callers {
					if !rewroteSourceFiles[cal.absFile] {
						rewroteSourceFiles[cal.absFile] = true
						totalSourceFiles++
					}
				}

				if !consolidateNoPrune {
					if filePath, ok := keyToFile[oldFullKey]; ok {
						changed, err := locale.DeleteKeys(filePath, []string{oldFullKey})
						if err != nil {
							fmt.Printf("  warning: delete key %s: %v\n", oldFullKey, err)
						} else if changed && !rewroteYAMLFiles[filePath] {
							rewroteYAMLFiles[filePath] = true
							totalYAMLFiles++
						}
					}
				}
			}
		}

		if dryRun {
			if callerCount == 0 {
				fmt.Printf("  (no callers found)\n")
			}
			fmt.Printf("  [dry-run] would rewrite %d caller(s) → t('%s')\n", callerCount, newTPath)
		} else {
			action := fmt.Sprintf("  ✓ %d caller(s) rewritten", callerCount)
			if !consolidateNoPrune {
				action += ", old key(s) deleted"
			}
			fmt.Println(action)
		}
		fmt.Println()
	}

	if dryRun {
		fmt.Println("[dry-run] No files modified.")
		return nil
	}

	fmt.Printf("Done: %d key(s) consolidated · %d YAML file(s) updated · %d source file(s) rewritten.\n",
		len(candidates), totalYAMLFiles, totalSourceFiles)
	if !consolidateNoPrune {
		fmt.Println("Run `locale-sync translate --to es,fr` to sync translations.")
	} else {
		fmt.Println("Run `locale-sync prune` to delete old keys after review.")
	}
	return nil
}

func summariseOldKeys(oldKeys []string, l string) string {
	stripped := make([]string, len(oldKeys))
	for i, k := range oldKeys {
		stripped[i] = stripLangPrefixStr(k, l)
	}
	if len(stripped) == 1 {
		return stripped[0]
	}
	return stripped[0] + fmt.Sprintf(" (+%d more)", len(stripped)-1)
}
