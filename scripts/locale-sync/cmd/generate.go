package cmd

import (
	"fmt"
	"strings"

	"github.com/bobadilla-tech/locale-sync/internal/locale"
	"github.com/bobadilla-tech/locale-sync/internal/source"
	"github.com/spf13/cobra"
)

var (
	languages          []string
	placeholder        string
	dryRun             bool
	overwrite          bool
	noShared           bool
	noSkeleton         bool
	fixSource          bool
	erbOnly            bool
	safeOnly           bool
	generateChangedFiles string
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Consolidate duplicates into shared YAML and generate language skeletons",
	Long: `generate does everything analyze does, then:
  - Writes duplicate keys with identical values into shared.{lang}.yml
  - Creates skeleton locale files for each --languages value (empty strings)
  - Prints a migration plan showing which t() keys to use
  - (with --fix) Replaces hardcoded strings in ERB/Ruby source with t() calls`,
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
		"Scan Ruby/ERB source for hardcoded strings (report only)")
	generateCmd.Flags().BoolVar(&fixSource, "fix", false,
		"Auto-replace hardcoded strings in ERB/Ruby source with t() calls and add keys to YAML")
	generateCmd.Flags().BoolVar(&erbOnly, "erb-only", false,
		"With --fix: only rewrite .erb/.haml/.slim files, skip Ruby .rb files")
	generateCmd.Flags().BoolVar(&safeOnly, "safe-only", false,
		"With --fix: only apply replacements that reuse an existing YAML key (no new keys generated)")
	generateCmd.Flags().StringVar(&generateChangedFiles, "changed-files", "",
		`Only scan the listed source files. Pass a file path or "-" to read from stdin.
Useful in CI: git diff --name-only origin/main | locale-sync generate --source --changed-files -`)
	rootCmd.AddCommand(generateCmd)
}

func runGenerate(cmd *cobra.Command, _ []string) error {
	if minDups < 2 {
		return fmt.Errorf("--min-dups must be at least 2 (got %d); singleton entries are not duplicates", minDups)
	}
	if err := validateLangCode(lang); err != nil {
		return fmt.Errorf("--lang: %w", err)
	}
	for _, l := range languages {
		if err := validateLangCode(l); err != nil {
			return fmt.Errorf("--languages: %w", err)
		}
	}
	if fixSource {
		scanSource = true
	}

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
		candidates = locale.FilterExistingValueCandidates(candidates, entries)
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

	// --- Source scan / fix ---
	if scanSource || fixSource {
		cf, err := loadChangedFiles(generateChangedFiles)
		if err != nil {
			return fmt.Errorf("--changed-files: %w", err)
		}
		found, err := source.Extract(rootPath, cf)
		if err != nil {
			return fmt.Errorf("source scan: %w", err)
		}

		// Scope to ERB/template files if requested — applies to both report and fix modes.
		if erbOnly {
			filtered := found[:0]
			for _, h := range found {
				if strings.HasSuffix(h.File, ".erb") ||
					strings.HasSuffix(h.File, ".haml") ||
					strings.HasSuffix(h.File, ".slim") {
					filtered = append(filtered, h)
				}
			}
			found = filtered
		}

		if len(found) == 0 {
			fmt.Println("No hardcoded user-facing strings detected.")
			return nil
		}

		if !fixSource {
			// Report-only mode.
			fmt.Printf("=== Hardcoded Strings Found (%d) ===\n\n", len(found))
			for _, s := range found {
				fmt.Printf("  [%s] %s:%d\n", s.Category, s.File, s.Line)
				fmt.Printf("    %q\n\n", s.Text)
			}
			return nil
		}

		// Build value→key index from current YAML.
		valueToKey := buildValueToKey(entries)

		plan := source.BuildFixPlan(found, valueToKey, lang, rootPath)

		// --safe-only: drop any fix that would add a new YAML key.
		if safeOnly {
			kept := plan.Fixes[:0]
			for _, f := range plan.Fixes {
				if !f.NewInYAML {
					kept = append(kept, f)
				}
			}
			plan.Fixes = kept
			plan.YAMLAdds = nil
		}

		fmt.Printf("=== Source Auto-Fix (%d replacements planned) ===\n\n", len(plan.Fixes))

		for _, fix := range plan.Fixes {
			status := "~"
			if fix.NewInYAML {
				status = "+"
			}
			fmt.Printf("  [%s] %s:%d\n", status, fix.File, fix.Line)
			fmt.Printf("    before: %s\n", strings.TrimSpace(fix.Original))
			fmt.Printf("    after:  %s\n", strings.TrimSpace(fix.Patched))
			fmt.Printf("    key:    t('%s')\n\n", fix.YAMLKey)
		}

		if len(plan.YAMLAdds) > 0 {
			fmt.Printf("=== New YAML Keys to Add (%d) ===\n\n", len(plan.YAMLAdds))
			for _, k := range locale.SortedKeys(plan.YAMLAdds) {
				fmt.Printf("  %s: %q\n", k, plan.YAMLAdds[k])
			}
			fmt.Println()
		}

		if dryRun {
			fmt.Println("[dry-run] No files written.")
			return nil
		}

		// Apply YAML additions — route each key to its correct topic file.
		if len(plan.YAMLAdds) > 0 {
			byTopic := make(map[string]map[string]string)
			for fullKey, val := range plan.YAMLAdds {
				topic := topicFromKey(fullKey, lang)
				if byTopic[topic] == nil {
					byTopic[topic] = make(map[string]string)
				}
				byTopic[topic][fullKey] = val
			}
			for topic, adds := range byTopic {
				candidates := yamlAddsToMergeCandidates(adds, lang)
				path, changed, err := locale.UpsertTopicFile(rootPath, lang, topic, candidates)
				if err != nil {
					return fmt.Errorf("write YAML (%s): %w", topic, err)
				}
				if changed {
					fmt.Printf("YAML: wrote %d new keys to %s\n", len(candidates), path)
				}
			}
		}

		// Apply source file fixes.
		writtenFiles, err := source.ApplyFixes(plan, rootPath)
		if err != nil {
			return fmt.Errorf("apply fixes: %w", err)
		}
		fmt.Printf("Source: rewrote %d files (%d replacements).\n",
			len(writtenFiles), len(plan.Fixes))
	}

	return nil
}

// buildValueToKey inverts the locale entries into a value → full YAML key map.
// Preference order: shared keys > per-topic keys; ties broken lexicographically.
func buildValueToKey(entries []locale.Entry) map[string]string {
	m := make(map[string]string)
	for _, e := range entries {
		if e.Value == "" {
			continue
		}
		existing, has := m[e.Value]
		if !has {
			m[e.Value] = e.Key
			continue
		}
		newIsShared := strings.Contains(e.Key, ".shared.")
		existingIsShared := strings.Contains(existing, ".shared.")
		if newIsShared && !existingIsShared {
			m[e.Value] = e.Key // shared beats per-topic
		} else if !newIsShared && !existingIsShared && e.Key < existing {
			m[e.Value] = e.Key // stable: always pick lexicographically smaller key
		}
	}
	return m
}

// yamlAddsToMergeCandidates converts the plan's YAMLAdds into MergeCandidate
// objects that can be passed to UpsertShared. New keys go under shared.generated.
func yamlAddsToMergeCandidates(adds map[string]string, lang string) []locale.MergeCandidate {
	var out []locale.MergeCandidate
	for fullKey, value := range adds {
		parts := strings.SplitN(fullKey, ".", 2)
		tKey := fullKey
		if len(parts) == 2 {
			tKey = parts[1]
		}
		out = append(out, locale.MergeCandidate{
			KeyName:      tKey,
			Value:        value,
			SuggestedKey: fullKey,
		})
	}
	return out
}

// topicFromKey extracts the YAML topic (second dot-segment after lang prefix).
// "en.tools.email_normalizer.hero.heading" → "tools"
// "en.devise.mailer.reset_password.body"   → "devise"
func topicFromKey(fullKey, lang string) string {
	stripped := strings.TrimPrefix(fullKey, lang+".")
	if idx := strings.Index(stripped, "."); idx >= 0 {
		return stripped[:idx]
	}
	return "shared"
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
