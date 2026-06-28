package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/bobadilla-tech/locale-sync/internal/locale"
	"github.com/bobadilla-tech/locale-sync/internal/translate"
	"github.com/spf13/cobra"
)

var (
	translateTo                []string
	translateFrom              string
	translateDryRun            bool
	translateForce             bool
	translateProvider          string
	translateCacheOnly         bool
	translateNoCache           bool
	translateCacheFile         string
	translateQuota             int
	translateBatchSize         int
	translateReportFile        string
	translateProtectWords      []string
	translateProtectFile       string
	translateSkeletonPlaceholder string
)

var translateCmd = &cobra.Command{
	Use:   "translate [lang...]",
	Short: "Translate missing locale keys using Google Cloud Translation API",
	Long: `translate finds keys present in the source locale (default: en) that are
missing in one or more target locale files, then fills them using the
Google Cloud Translation API.

The command is idempotent: it never overwrites real translations.
Placeholder values (e.g. "TODO: translate") are replaced with real translations.
Run with --dry-run to preview what would be translated without calling the API.

Use --protect-words or --protect-file to keep brand names and proper nouns
unchanged — they are replaced with tokens before calling the API and restored
after so Google never sees them.

Authentication:
  Set GOOGLE_APPLICATION_CREDENTIALS to a service account JSON key file.

Examples:
  locale-sync translate --to=es,fr --root=../../apps/dashboard
  locale-sync translate --to=all --dry-run
  locale-sync translate --to=es --protect-words=NeverBounce,IPstack,AbstractAPI
  locale-sync translate --to=es --protect-file=.translate-dictionary.txt
  locale-sync translate --to=es --report-file=reports/translate.json`,
	RunE: runTranslate,
}

func init() {
	translateCmd.Flags().StringVarP(&rootPath, "root", "r", ".", "Rails app root directory")
	translateCmd.Flags().StringSliceVar(&translateTo, "to", nil,
		`Target language codes (comma-separated) or "all" to translate every known language`)
	translateCmd.Flags().StringVar(&translateFrom, "from", "en", "Source locale language")
	translateCmd.Flags().BoolVar(&translateDryRun, "dry-run", false,
		"Show what would be translated without writing files or calling the API")
	translateCmd.Flags().BoolVar(&translateForce, "force", false,
		"Skip double confirmation prompt")
	translateCmd.Flags().StringVar(&translateProvider, "provider", "google",
		"Translation provider (currently only \"google\" is supported)")
	translateCmd.Flags().BoolVar(&translateCacheOnly, "cache-only", false,
		"Only use cached translations; skip API calls")
	translateCmd.Flags().BoolVar(&translateNoCache, "no-cache", false,
		"Bypass translation cache; always call the API")
	translateCmd.Flags().StringVar(&translateCacheFile, "cache-file", ".translation-cache.json",
		"Path to translation cache file")
	translateCmd.Flags().IntVar(&translateQuota, "quota-limit", 450_000,
		"Stop translating when this many characters have been sent to the API (safety margin)")
	translateCmd.Flags().IntVar(&translateBatchSize, "batch-size", 100,
		"Number of strings to send per API call")
	translateCmd.Flags().StringVar(&translateReportFile, "report-file", "",
		"Write a detailed JSON report of all translations to this file path")
	translateCmd.Flags().StringSliceVar(&translateProtectWords, "protect-words", nil,
		"Comma-separated words/phrases that must not be translated (e.g. brand names)")
	translateCmd.Flags().StringVar(&translateProtectFile, "protect-file", "",
		"Path to a file with one protected word/phrase per line (# lines are comments)")
	translateCmd.Flags().StringVar(&translateSkeletonPlaceholder, "skeleton-placeholder", "",
		"Treat this value as an untranslated placeholder (match the value used with generate --placeholder)")
	rootCmd.AddCommand(translateCmd)
}

// --- report types ---

type reportData struct {
	Timestamp  string                 `json:"timestamp"`
	SourceLang string                 `json:"source_lang"`
	TotalKeys  int                    `json:"total_keys"`
	TotalChars int                    `json:"total_chars_used"`
	QuotaLimit int                    `json:"quota_limit"`
	QuotaHit   bool                   `json:"quota_limit_hit"`
	Languages  map[string]*langReport `json:"languages"`
}

type langReport struct {
	KeysTranslated int           `json:"keys_translated"`
	KeysFromCache  int           `json:"keys_from_cache"`
	KeysViaAPI     int           `json:"keys_via_api"`
	KeysSkipped    int           `json:"keys_skipped,omitempty"`
	CharsUsed      int           `json:"chars_used"`
	FilesWritten   []string      `json:"files_written"`
	Entries        []reportEntry `json:"entries"`
}

type reportEntry struct {
	Key         string `json:"key"`
	Source      string `json:"source"`
	Translation string `json:"translation"`
	FromCache   bool   `json:"from_cache"`
}

type langMissing struct {
	lang    string
	missing []translate.MissingEntry
}

// --- main command ---

func runTranslate(_ *cobra.Command, args []string) error {
	// Validate flags before doing any work.
	if translateBatchSize <= 0 {
		return fmt.Errorf("--batch-size must be a positive integer (got %d)", translateBatchSize)
	}
	if translateProvider != "google" {
		return fmt.Errorf("unsupported --provider %q: only \"google\" is currently supported", translateProvider)
	}
	if translateCacheOnly && translateNoCache {
		return fmt.Errorf("--cache-only and --no-cache are mutually exclusive")
	}
	if err := validateLangCode(translateFrom); err != nil {
		return fmt.Errorf("--from: %w", err)
	}
	for _, lang := range translateTo {
		if lang == "all" {
			continue
		}
		if err := validateLangCode(lang); err != nil {
			return fmt.Errorf("--to: %w", err)
		}
	}

	// Build protector (custom dictionary + standard placeholders).
	// PluralSep is always protected so it survives the translation API unchanged.
	protectedWords, err := loadProtectedWords(translateProtectWords, translateProtectFile)
	if err != nil {
		return fmt.Errorf("load protected words: %w", err)
	}
	protector := translate.NewProtector(append(protectedWords, translate.PluralSep))

	if len(protectedWords) > 0 {
		fmt.Printf("Protected words: %s\n\n", strings.Join(protectedWords, ", "))
	}

	// Resolve target languages.
	targets, err := resolveTargetLangs(translateTo, args, translateFrom)
	if err != nil {
		return err
	}
	if len(targets) == 0 {
		return fmt.Errorf("no target languages specified; use --to=es,fr or positional args")
	}

	// Load source entries.
	fmt.Printf("Loading %s locale entries from %s...\n", translateFrom, rootPath)
	sourceEntries, err := locale.Scan(rootPath, translateFrom)
	if err != nil {
		return fmt.Errorf("scan %s locales: %w", translateFrom, err)
	}
	fmt.Printf("  %d source entries found.\n\n", len(sourceEntries))

	// For each target lang, find missing entries.
	var allLangs []langMissing
	totalCharsNeeded := 0
	totalKeysNeeded := 0

	for _, tLang := range targets {
		targetEntries, scanErr := locale.Scan(rootPath, tLang)
		if scanErr != nil {
			// Ignore "nothing to scan" (no locale dir yet); surface real errors.
			if len(targetEntries) == 0 {
				targetEntries = nil
			} else {
				return fmt.Errorf("scan %s locales: %w", tLang, scanErr)
			}
		}
		missing := translate.FindMissing(sourceEntries, targetEntries, translateFrom, tLang, rootPath, translateSkeletonPlaceholder)
		allLangs = append(allLangs, langMissing{lang: tLang, missing: missing})
		totalCharsNeeded += translate.CharCount(missing)
		totalKeysNeeded += len(missing)
	}

	// Print plan summary.
	fmt.Println("=== Translation Plan ===")
	for _, lm := range allLangs {
		if len(lm.missing) == 0 {
			fmt.Printf("  %-4s  already up to date\n", lm.lang)
		} else {
			fmt.Printf("  %-4s  %s keys, ~%s chars\n",
				lm.lang, formatInt(len(lm.missing)), formatInt(translate.CharCount(lm.missing)))
		}
	}
	fmt.Printf("  ────\n")
	fmt.Printf("  Total  %s keys, ~%s chars  (free tier: ~500,000/month)\n\n",
		formatInt(totalKeysNeeded), formatInt(totalCharsNeeded))

	// Dry-run: print detailed plan and exit.
	if translateDryRun {
		printDryRunPlan(allLangs)
		fmt.Println("No files written (--dry-run).")
		return nil
	}

	if totalKeysNeeded == 0 {
		fmt.Println("All target locales are up to date. Nothing to translate.")
		return nil
	}

	// Double confirmation (unless --force or --cache-only).
	if !translateForce && !translateCacheOnly {
		if !confirm("WARNING: This will call the Google Cloud Translation API.\nAre you sure?", "yes") {
			fmt.Println("Aborted.")
			return nil
		}
		fmt.Println()
		msg := fmt.Sprintf(
			"This will consume your Google quota (~%s chars).\nFree tier: ~500,000 chars/month.\nType YES to continue",
			formatInt(totalCharsNeeded),
		)
		if !confirm(msg, "YES") {
			fmt.Println("Aborted.")
			return nil
		}
		fmt.Println()
	}

	// Load cache.
	var cache *translate.Cache
	if !translateNoCache {
		cache, err = translate.LoadCache(translateCacheFile)
		if err != nil {
			return fmt.Errorf("load cache: %w", err)
		}
	}

	// Create API client (unless cache-only).
	var client *translate.Client
	if !translateCacheOnly {
		ctx := context.Background()
		client, err = translate.NewClient(ctx)
		if err != nil {
			return fmt.Errorf("init translation client: %w\n\nEnsure GOOGLE_APPLICATION_CREDENTIALS is set to a valid service account JSON file", err)
		}
		defer client.Close()
	}

	// Translate per language.
	report := &reportData{
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
		SourceLang: translateFrom,
		QuotaLimit: translateQuota,
		Languages:  make(map[string]*langReport),
	}
	totalCharsUsed := 0
	quotaExhausted := false

	for _, lm := range allLangs {
		if len(lm.missing) == 0 {
			continue
		}

		lr := &langReport{}
		report.Languages[lm.lang] = lr

		fmt.Printf("Translating to %s (%s keys)...\n", lm.lang, formatInt(len(lm.missing)))

		var translated []translate.TranslatedEntry
		var needsAPI []translate.MissingEntry
		cacheSkipped := 0

		for _, m := range lm.missing {
			if !translateNoCache && cache != nil {
				if cached, ok := cache.Get(translateFrom, m.Value, lm.lang); ok {
					translated = append(translated, translate.TranslatedEntry{
						MissingEntry:   m,
						TranslatedText: cached,
						FromCache:      true,
					})
					lr.KeysFromCache++
					continue
				}
			}
			if translateCacheOnly {
				cacheSkipped++
				continue
			}
			needsAPI = append(needsAPI, m)
		}

		// Group plural siblings (one/other/zero...) so they are translated together
		// for grammatical coherence. Ungrouped entries fall through to normal batching.
		langChars := 0
		singles, pluralGroups := translate.GroupPlurals(needsAPI)
		if len(pluralGroups) > 0 && !quotaExhausted && client != nil {
			ctx := context.Background()
			combTexts := make([]string, len(pluralGroups))
			combOrders := make([][]string, len(pluralGroups))
			pluralProtected := make([]translate.Protected, len(pluralGroups))
			batchChars := 0
			for i, g := range pluralGroups {
				combined, order := translate.CombinePluralGroup(g)
				combOrders[i] = order
				p := protector.Protect(combined)
				pluralProtected[i] = p
				combTexts[i] = p.Text
				batchChars += len([]rune(p.Text))
			}
			if totalCharsUsed+batchChars <= translateQuota {
				results, err := client.Translate(ctx, combTexts, lm.lang, translateFrom)
				if err != nil {
					return fmt.Errorf("API error plural groups (lang=%s): %w", lm.lang, err)
				}
				totalCharsUsed += batchChars
				langChars += batchChars
				for i, g := range pluralGroups {
					restored := pluralProtected[i].Restore(results[i])
					split := translate.SplitPluralTranslation(restored, combOrders[i])
					for j, entry := range g.Entries {
						form := combOrders[i][j]
						text := split[form]
						translated = append(translated, translate.TranslatedEntry{
							MissingEntry:   entry,
							TranslatedText: text,
							FromCache:      false,
						})
						if !translateNoCache && cache != nil {
							cache.Set(translateFrom, entry.Value, lm.lang, text)
						}
					}
				}
				if !translateNoCache && cache != nil {
					if err := cache.Save(); err != nil {
						fmt.Fprintf(os.Stderr, "  Warning: cache save failed: %v\n", err)
					}
				}
			} else {
				// Quota would be exceeded — fall back to treating plural forms as singles.
				singles = append(singles, flattenGroups(pluralGroups)...)
				quotaExhausted = true
				report.QuotaHit = true
			}
		} else {
			singles = append(singles, flattenGroups(pluralGroups)...)
		}
		needsAPI = singles

		// API translation in batches (singletons — plural groups handled above).
		if !quotaExhausted && len(needsAPI) > 0 && client != nil {
			ctx := context.Background()
			for i := 0; i < len(needsAPI); i += translateBatchSize {
				end := i + translateBatchSize
				if end > len(needsAPI) {
					end = len(needsAPI)
				}
				batch := needsAPI[i:end]

				protecteds := make([]translate.Protected, len(batch))
				apiTexts := make([]string, len(batch))
				batchChars := 0
				for j, m := range batch {
					p := protector.Protect(m.Value)
					protecteds[j] = p
					apiTexts[j] = p.Text
					batchChars += len([]rune(p.Text))
				}

				if totalCharsUsed+batchChars > translateQuota {
					fmt.Fprintf(os.Stderr,
						"  Warning: quota limit of %s chars reached. Stopping.\n",
						formatInt(translateQuota))
					quotaExhausted = true
					report.QuotaHit = true
					break
				}

				results, err := client.Translate(ctx, apiTexts, lm.lang, translateFrom)
				if err != nil {
					return fmt.Errorf("API error (lang=%s, offset=%d): %w", lm.lang, i, err)
				}

				totalCharsUsed += batchChars
				langChars += batchChars

				for j, m := range batch {
					restored := protecteds[j].Restore(results[j])
					translated = append(translated, translate.TranslatedEntry{
						MissingEntry:   m,
						TranslatedText: restored,
						FromCache:      false,
					})
					if !translateNoCache && cache != nil {
						cache.Set(translateFrom, m.Value, lm.lang, restored)
					}
				}

				if !translateNoCache && cache != nil {
					if err := cache.Save(); err != nil {
						fmt.Fprintf(os.Stderr, "  Warning: cache save failed: %v\n", err)
					}
				}
			}
		}

		lr.KeysViaAPI = len(translated) - lr.KeysFromCache
		lr.KeysSkipped = cacheSkipped
		lr.KeysTranslated = len(translated)
		lr.CharsUsed = langChars

		// Write files and print per-file summary.
		if len(translated) > 0 {
			fileCount := make(map[string]int)
			for _, e := range translated {
				fileCount[e.TargetFile]++
			}

			written, err := translate.WriteTranslations(translated)
			if err != nil {
				return fmt.Errorf("write %s translations: %w", lm.lang, err)
			}
			lr.FilesWritten = relPaths(written, rootPath)

			for _, f := range written {
				fmt.Printf("  ✓  %-35s  %s keys\n", filepath.Base(f), formatInt(fileCount[f]))
			}

			for _, e := range translated {
				lr.Entries = append(lr.Entries, reportEntry{
					Key:         dotAfterLang(e.TargetKey),
					Source:      e.Value,
					Translation: e.TranslatedText,
					FromCache:   e.FromCache,
				})
			}
		}

		if cacheSkipped > 0 {
			fmt.Printf("  ⚠  %s keys skipped (not in cache, --cache-only)\n", formatInt(cacheSkipped))
		}

		cacheTag := ""
		if lr.KeysFromCache > 0 {
			cacheTag = fmt.Sprintf("  |  %s from cache", formatInt(lr.KeysFromCache))
		}
		fmt.Printf("  ─  %s keys written%s  |  ~%s chars\n\n",
			formatInt(lr.KeysTranslated), cacheTag, formatInt(langChars))
	}

	// Final summary.
	report.TotalChars = totalCharsUsed
	report.TotalKeys = totalKeysNeeded
	fmt.Printf("Total API chars used: ~%s", formatInt(totalCharsUsed))
	if quotaExhausted {
		fmt.Fprintf(os.Stderr, "\nWarning: quota limit hit — re-run to continue remaining entries")
	}
	fmt.Println()

	// Write report file.
	if translateReportFile != "" {
		if err := writeReport(report, translateReportFile); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: could not write report: %v\n", err)
		} else {
			fmt.Printf("Report: %s\n", translateReportFile)
		}
	}

	return nil
}

// loadProtectedWords merges --protect-words flag values with words from --protect-file.
func loadProtectedWords(flagWords []string, filePath string) ([]string, error) {
	var all []string
	all = append(all, flagWords...)
	if filePath != "" {
		fileWords, err := translate.LoadDictionaryFile(filePath)
		if err != nil {
			return nil, err
		}
		all = append(all, fileWords...)
	}
	// Deduplicate.
	seen := map[string]bool{}
	var out []string
	for _, w := range all {
		w = strings.TrimSpace(w)
		if w != "" && !seen[w] {
			seen[w] = true
			out = append(out, w)
		}
	}
	return out, nil
}

func writeReport(r *reportData, path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

func relPaths(paths []string, root string) []string {
	out := make([]string, 0, len(paths))
	for _, p := range paths {
		rel, err := filepath.Rel(root, p)
		if err != nil {
			rel = p
		}
		out = append(out, rel)
	}
	return out
}

func resolveTargetLangs(flagTo []string, args []string, sourceLang string) ([]string, error) {
	raw := append(append([]string{}, flagTo...), args...)
	var expanded []string
	for _, v := range raw {
		for _, part := range strings.Split(v, ",") {
			if part = strings.TrimSpace(part); part != "" {
				expanded = append(expanded, part)
			}
		}
	}
	for _, v := range expanded {
		if v == "all" {
			return translate.DiscoverLangs(rootPath, sourceLang)
		}
	}
	seen := map[string]bool{}
	var out []string
	for _, v := range expanded {
		if v == sourceLang || seen[v] {
			continue
		}
		seen[v] = true
		out = append(out, v)
	}
	return out, nil
}

func confirm(message, expected string) bool {
	fmt.Printf("%s (%s/no): ", message, expected)
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return false
	}
	return strings.TrimSpace(scanner.Text()) == expected
}

func printDryRunPlan(allLangs []langMissing) {
	fmt.Println("=== Dry Run: Keys to Translate ===")
	for _, lm := range allLangs {
		if len(lm.missing) == 0 {
			continue
		}
		fmt.Printf("\n%s (%s keys):\n", strings.ToUpper(lm.lang), formatInt(len(lm.missing)))
		for _, m := range lm.missing {
			key := dotAfterLang(m.TargetKey)
			fmt.Printf("  %-60s  %q\n", key, truncate(m.Value, 60))
		}
	}
	fmt.Println()
}

func formatInt(n int) string {
	s := fmt.Sprintf("%d", n)
	if len(s) <= 3 {
		return s
	}
	var out []byte
	for i, c := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			out = append(out, ',')
		}
		out = append(out, byte(c))
	}
	return string(out)
}

func truncate(s string, max int) string {
	r := []rune(s)
	if len(r) <= max {
		return s
	}
	return string(r[:max]) + "…"
}

func flattenGroups(groups []translate.PluralGroup) []translate.MissingEntry {
	var out []translate.MissingEntry
	for _, g := range groups {
		out = append(out, g.Entries...)
	}
	return out
}
