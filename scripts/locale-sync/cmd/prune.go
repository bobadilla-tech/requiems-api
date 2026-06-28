package cmd

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/bobadilla-tech/locale-sync/internal/locale"
	"github.com/bobadilla-tech/locale-sync/internal/source"
	"github.com/spf13/cobra"
)

var (
	pruneLang    string
	pruneDryRun  bool
	prunePattern string
)

var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Remove orphaned locale keys that are no longer used in source",
	Long: `prune cross-references all t() calls in source with the locale YAML files and
removes keys that are defined but never called anywhere in the codebase.

Always run with --dry-run first to review what will be deleted.

Examples:
  locale-sync prune --root=../../apps/dashboard --lang=en --dry-run
  locale-sync prune --root=../../apps/dashboard --lang=en`,
	RunE: runPrune,
}

func init() {
	pruneCmd.Flags().StringVarP(&rootPath, "root", "r", ".", "Rails app root directory")
	pruneCmd.Flags().StringVarP(&pruneLang, "lang", "l", "en", "Locale language to prune")
	pruneCmd.Flags().BoolVar(&pruneDryRun, "dry-run", false,
		"Print orphaned keys without deleting them")
	pruneCmd.Flags().StringVar(&prunePattern, "pattern", "",
		"Only prune keys whose full path matches this regex (e.g. 'shared\\.common\\.')")
	rootCmd.AddCommand(pruneCmd)
}

func runPrune(_ *cobra.Command, _ []string) error {
	// Load all defined keys for the target language.
	entries, err := locale.Scan(rootPath, pruneLang)
	if err != nil {
		return fmt.Errorf("scan locales: %w", err)
	}

	definedKeys := make(map[string]bool, len(entries))
	for _, e := range entries {
		definedKeys[e.Key] = true
	}

	// Audit source for used keys.
	result, err := source.Audit(rootPath, pruneLang, definedKeys)
	if err != nil {
		return fmt.Errorf("audit: %w", err)
	}

	orphaned := result.OrphanedKeys
	if prunePattern != "" {
		re, err := regexp.Compile(prunePattern)
		if err != nil {
			return fmt.Errorf("--pattern: %w", err)
		}
		filtered := orphaned[:0]
		for _, k := range orphaned {
			if re.MatchString(k) {
				filtered = append(filtered, k)
			}
		}
		orphaned = filtered
	}

	if len(orphaned) == 0 {
		fmt.Println("No orphaned keys found.")
		return nil
	}

	fmt.Printf("Found %d orphaned key(s):\n\n", len(orphaned))
	for _, k := range orphaned {
		fmt.Printf("  %s\n", k)
	}
	fmt.Println()

	if pruneDryRun {
		fmt.Printf("--dry-run: no files modified.\n")
		return nil
	}

	// Group orphaned keys by their locale file.
	fileKeys := map[string][]string{}
	keyToEntry := make(map[string]locale.Entry, len(entries))
	for _, e := range entries {
		keyToEntry[e.Key] = e
	}

	for _, k := range orphaned {
		e, ok := keyToEntry[k]
		if !ok {
			// Key may be stored without lang prefix in orphaned list.
			e, ok = keyToEntry[pruneLang+"."+k]
		}
		if !ok {
			continue
		}
		fileKeys[e.File] = append(fileKeys[e.File], k)
	}

	totalDeleted := 0
	for filePath, keys := range fileKeys {
		changed, err := locale.DeleteKeys(filePath, keys)
		if err != nil {
			fmt.Printf("  Warning: could not update %s: %v\n", shortPath(filePath, rootPath), err)
			continue
		}
		if changed {
			totalDeleted += len(keys)
			fmt.Printf("  Removed %d key(s) from %s\n", len(keys), shortPath(filePath, rootPath))
		}
	}

	fmt.Printf("\nPruned %d key(s) across %d file(s).\n", totalDeleted, len(fileKeys))
	return nil
}

func shortPath(absPath, root string) string {
	rel, err := filepath.Rel(root, absPath)
	if err != nil {
		return absPath
	}
	return strings.ReplaceAll(rel, "\\", "/")
}
