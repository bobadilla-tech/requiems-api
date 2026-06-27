package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/bobadilla-tech/locale-sync/internal/locale"
	"github.com/bobadilla-tech/locale-sync/internal/source"
	"github.com/spf13/cobra"
)

var mergeFragmentsCmd = &cobra.Command{
	Use:   "merge-fragments",
	Short: "Merge sentence fragments split across ERB expressions into t() calls",
	Long: `merge-fragments scans ERB/Haml templates for lines where user-visible text
is split around <%= ... %> interpolations — a pattern that breaks i18n because
the sentence cannot be translated as a unit.

For each detected fragment it:
  1. Builds a merged locale key  e.g. "Your env for %{company} is ready."
  2. Suggests a dot-notation key path derived from the file path
  3. With --fix: writes the key to YAML and rewrites the source line

Lines with HTML tags embedded in the text, or with multi-argument ERB
expressions (link_to, mail_to, etc.), are flagged as complex and reported
but never auto-patched — those need manual restructuring.

Exit codes:
  0  no fragments found (or --fix applied cleanly)
  1  fragments remain that need attention

Examples:
  locale-sync merge-fragments --root ../../apps/dashboard --dry-run
  locale-sync merge-fragments --root ../../apps/dashboard --fix`,
	SilenceUsage: true,
	RunE:         runMergeFragments,
}

var (
	mergeFragmentsLang   string
	mergeFragmentsDryRun bool
	mergeFragmentsFix    bool
)

func init() {
	mergeFragmentsCmd.Flags().StringVarP(&rootPath, "root", "r", ".", "Rails app root directory")
	mergeFragmentsCmd.Flags().StringVarP(&mergeFragmentsLang, "lang", "l", "en", "Target locale language")
	mergeFragmentsCmd.Flags().BoolVar(&mergeFragmentsDryRun, "dry-run", false, "Show proposals without writing")
	mergeFragmentsCmd.Flags().BoolVar(&mergeFragmentsFix, "fix", false, "Write keys to YAML and rewrite source")
	rootCmd.AddCommand(mergeFragmentsCmd)
}

func runMergeFragments(_ *cobra.Command, _ []string) error {
	fragments, err := source.DetectFragments(rootPath)
	if err != nil {
		return fmt.Errorf("detect fragments: %w", err)
	}

	if len(fragments) == 0 {
		fmt.Println("merge-fragments: no mixed text/ERB fragments found.")
		return nil
	}

	// Split auto-patchable vs complex.
	var simple, complex []source.FragmentLine
	for _, f := range fragments {
		if f.Complex {
			complex = append(complex, f)
		} else {
			simple = append(simple, f)
		}
	}

	// --- Report simple (auto-patchable) fragments ---
	if len(simple) > 0 {
		fmt.Printf("=== Auto-patchable fragments (%d) ===\n\n", len(simple))
		for _, f := range simple {
			fmt.Printf("  %s:%d\n", f.File, f.LineNum)
			fmt.Printf("  before: %s\n", trimDisplay(f.Original))
			fmt.Printf("  key:    %s\n", f.KeySuggestion)
			fmt.Printf("  value:  %q\n", f.MergedValue)
			fmt.Printf("  after:  %s\n\n", trimDisplay(f.PatchedLine))
		}
	}

	// --- Report complex fragments ---
	if len(complex) > 0 {
		fmt.Printf("=== Complex fragments — manual restructuring needed (%d) ===\n\n", len(complex))
		for _, f := range complex {
			fmt.Printf("  %s:%d\n", f.File, f.LineNum)
			fmt.Printf("  text:   %s\n", trimDisplay(f.Original))
			fmt.Printf("  value:  %q\n", f.MergedValue)
			fmt.Printf("  key:    %s\n", f.KeySuggestion)
			fmt.Printf("  reason: %s\n\n", f.ComplexReason)
		}
	}

	if mergeFragmentsDryRun || (!mergeFragmentsDryRun && !mergeFragmentsFix) {
		if !mergeFragmentsDryRun {
			fmt.Printf("Run with --fix to apply auto-patchable changes.\n")
		}
		if len(simple) > 0 {
			os.Exit(1)
		}
		return nil
	}

	// --- Apply ---
	if len(simple) == 0 {
		fmt.Println("No auto-patchable fragments. Complex cases require manual restructuring.")
		if len(complex) > 0 {
			os.Exit(1)
		}
		return nil
	}

	fmt.Printf("Applying %d auto-patchable fragment(s)...\n\n", len(simple))

	// Load existing locale entries so we can route keys to the right file.
	entries, err := locale.Scan(rootPath, mergeFragmentsLang)
	if err != nil {
		return fmt.Errorf("scan locales: %w", err)
	}
	existing := make(map[string]bool, len(entries))
	for _, e := range entries {
		existing[e.Key] = true
	}

	// Write YAML keys and collect Fix objects for source rewrites.
	var fixes []source.Fix
	for _, f := range simple {
		fullKey := mergeFragmentsLang + "." + f.KeySuggestion

		if !existing[fullKey] {
			candidate := locale.MergeCandidate{
				KeyName:      lastSegment(f.KeySuggestion),
				Value:        f.MergedValue,
				SuggestedKey: fullKey,
			}
			topic := topicFromKey(f.KeySuggestion, mergeFragmentsLang)
			_, _, err := locale.UpsertTopicFile(rootPath, mergeFragmentsLang, topic, []locale.MergeCandidate{candidate})
			if err != nil {
				fmt.Fprintf(os.Stderr, "  warning: could not write key %s: %v\n", f.KeySuggestion, err)
				continue
			}
			fmt.Printf("  + %s = %q\n", fullKey, f.MergedValue)
		} else {
			fmt.Printf("  ~ %s (already exists)\n", fullKey)
		}

		fixes = append(fixes, source.Fix{
			File:      f.File,
			Line:      f.LineNum,
			Original:  f.Original,
			Patched:   f.PatchedLine,
			YAMLKey:   f.KeySuggestion,
			YAMLValue: f.MergedValue,
		})
	}

	if len(fixes) == 0 {
		return nil
	}

	plan := source.FixPlan{Fixes: fixes, YAMLAdds: map[string]string{}}
	written, err := source.ApplyFixes(plan, rootPath)
	if err != nil {
		return fmt.Errorf("apply source fixes: %w", err)
	}
	fmt.Printf("\nRewritten %d source file(s): %v\n", len(written), written)

	if len(complex) > 0 {
		fmt.Printf("\n%d complex fragment(s) still need manual restructuring (see above).\n", len(complex))
		os.Exit(1)
	}
	return nil
}

func trimDisplay(s string) string {
	s = strings.TrimSpace(s)
	if len(s) > 100 {
		return s[:97] + "..."
	}
	return s
}
