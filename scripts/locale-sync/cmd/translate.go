package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/bobadilla-tech/locale-sync/internal/locale"
	"github.com/bobadilla-tech/locale-sync/internal/translate"
	"github.com/spf13/cobra"
)

var (
	translateTo        []string
	translateFrom      string
	translateDryRun    bool
	translateForce     bool
	translateProvider  string
	translateCacheOnly bool
	translateNoCache   bool
	translateCacheFile string
	translateQuota     int
	translateBatchSize int
)

var translateCmd = &cobra.Command{
	Use:   "translate [lang...]",
	Short: "Translate missing locale keys using Google Cloud Translation API",
	Long: `translate finds keys present in the source locale (default: en) that are
missing in one or more target locale files, then fills them using the
Google Cloud Translation API.

The command is idempotent: it never overwrites existing translations.
Run with --dry-run to preview what would be translated without calling the API.

Authentication:
  Set GOOGLE_APPLICATION_CREDENTIALS to a service account JSON key file.

Examples:
  locale-sync translate --to=es,fr,de --root=../../apps/dashboard
  locale-sync translate es fr de
  locale-sync translate --to=all --dry-run
  locale-sync translate --to=es --cache-only`,
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
		"Stop translating when this many characters have been sent to the API")
	translateCmd.Flags().IntVar(&translateBatchSize, "batch-size", 100,
		"Number of strings to send per API call")
	rootCmd.AddCommand(translateCmd)
}

func runTranslate(cmd *cobra.Command, args []string) error {
	// 1. Resolve target languages.
	targets, err := resolveTargetLangs(translateTo, args, translateFrom)
	if err != nil {
		return err
	}
	if len(targets) == 0 {
		return fmt.Errorf("no target languages specified; use --to=es,fr or positional args")
	}

	// 2. Load source entries.
	fmt.Printf("Loading %s locale entries from %s...\n", translateFrom, rootPath)
	sourceEntries, err := locale.Scan(rootPath, translateFrom)
	if err != nil {
		return fmt.Errorf("scan %s locales: %w", translateFrom, err)
	}
	fmt.Printf("  %d source entries found.\n\n", len(sourceEntries))

	// 3. For each target lang, find missing entries.
	var allLangs []langMissing
	totalCharsNeeded := 0

	for _, tLang := range targets {
		targetEntries, err := locale.Scan(rootPath, tLang)
		if err != nil {
			// Target may not exist yet — treat as zero entries.
			targetEntries = nil
		}
		missing := translate.FindMissing(sourceEntries, targetEntries, translateFrom, tLang, rootPath)
		chars := translate.CharCount(missing)
		allLangs = append(allLangs, langMissing{lang: tLang, missing: missing})
		totalCharsNeeded += chars
	}

	// 4. Print summary.
	fmt.Println("=== Translation Plan ===")
	for _, lm := range allLangs {
		chars := translate.CharCount(lm.missing)
		fmt.Printf("  %s: %d missing keys, ~%s characters\n",
			lm.lang, len(lm.missing), formatInt(chars))
	}
	fmt.Printf("Total: ~%s characters (Google free tier: ~500,000/month)\n\n",
		formatInt(totalCharsNeeded))

	// 5. Dry-run: just show entries and exit.
	if translateDryRun {
		printDryRunPlan(allLangs)
		fmt.Println("No files written (--dry-run).")
		return nil
	}

	// Check if anything to do.
	if totalCharsNeeded == 0 {
		fmt.Println("All target locales are up to date. Nothing to translate.")
		return nil
	}

	// 6. Double confirmation (unless --force or --cache-only).
	if !translateForce && !translateCacheOnly {
		if !confirm("WARNING: This will call the Google Cloud Translation API.\nAre you sure?", "yes") {
			fmt.Println("Aborted.")
			return nil
		}
		fmt.Println()
		if !confirm(fmt.Sprintf("This will consume your Google quota (free tier: ~500,000 chars/month).\nYou have ~%s characters to translate.\nType YES to continue", formatInt(totalCharsNeeded)), "YES") {
			fmt.Println("Aborted.")
			return nil
		}
		fmt.Println()
	}

	// 7. Load cache.
	var cache *translate.Cache
	if !translateNoCache {
		cache, err = translate.LoadCache(translateCacheFile)
		if err != nil {
			return fmt.Errorf("load cache: %w", err)
		}
	}

	// 8. Create API client (unless cache-only).
	var client *translate.Client
	if !translateCacheOnly {
		ctx := context.Background()
		client, err = translate.NewClient(ctx)
		if err != nil {
			return fmt.Errorf("init translation client: %w\n\nEnsure GOOGLE_APPLICATION_CREDENTIALS is set to a valid service account JSON file.", err)
		}
		defer client.Close()
	}

	// 9. Translate per language.
	totalCharsUsed := 0
	quotaExhausted := false

	for _, lm := range allLangs {
		if len(lm.missing) == 0 {
			fmt.Printf("%s: already up to date.\n", lm.lang)
			continue
		}

		fmt.Printf("Translating to %s (%d keys)...\n", lm.lang, len(lm.missing))

		var translated []translate.TranslatedEntry
		var cacheHits, apiCalls, cacheSkipped int

		// Separate cache hits from API-needed.
		var needsAPI []translate.MissingEntry
		for _, m := range lm.missing {
			if !translateNoCache && cache != nil {
				if cached, ok := cache.Get(m.Value, lm.lang); ok {
					p := translate.Protect(m.Value)
					_ = p // cached value already has correct placeholders from original translation
					translated = append(translated, translate.TranslatedEntry{
						MissingEntry:   m,
						TranslatedText: cached,
						FromCache:      true,
					})
					cacheHits++
					continue
				}
			}
			if translateCacheOnly {
				cacheSkipped++
				continue
			}
			needsAPI = append(needsAPI, m)
		}

		// API translation in batches.
		if !quotaExhausted && len(needsAPI) > 0 && client != nil {
			ctx := context.Background()
			for i := 0; i < len(needsAPI); i += translateBatchSize {
				end := i + translateBatchSize
				if end > len(needsAPI) {
					end = len(needsAPI)
				}
				batch := needsAPI[i:end]

				// Protect placeholders, build texts for API.
				protecteds := make([]translate.Protected, len(batch))
				apiTexts := make([]string, len(batch))
				batchChars := 0
				for j, m := range batch {
					p := translate.Protect(m.Value)
					protecteds[j] = p
					apiTexts[j] = p.Text
					batchChars += len([]rune(p.Text))
				}

				// Quota check.
				if totalCharsUsed+batchChars > translateQuota {
					fmt.Fprintf(os.Stderr, "  Warning: approaching quota limit (%s chars). Stopping API calls.\n",
						formatInt(translateQuota))
					quotaExhausted = true
					break
				}

				results, err := client.Translate(ctx, apiTexts, lm.lang, translateFrom)
				if err != nil {
					return fmt.Errorf("translate batch (lang=%s, batch starting at %d): %w", lm.lang, i, err)
				}

				totalCharsUsed += batchChars
				apiCalls++

				for j, m := range batch {
					restored := protecteds[j].Restore(results[j])
					translated = append(translated, translate.TranslatedEntry{
						MissingEntry:   m,
						TranslatedText: restored,
						FromCache:      false,
					})
					// Store in cache.
					if !translateNoCache && cache != nil {
						cache.Set(m.Value, lm.lang, restored)
					}
				}

				// Save cache after each batch.
				if !translateNoCache && cache != nil {
					if err := cache.Save(); err != nil {
						fmt.Fprintf(os.Stderr, "  Warning: could not save cache: %v\n", err)
					}
				}
			}
		}

		// Print per-entry results.
		for _, e := range translated {
			src := truncate(e.Value, 40)
			dst := truncate(e.TranslatedText, 40)
			tag := "[api]  "
			if e.FromCache {
				tag = "[cache]"
			}
			fmt.Printf("  %s %s\n", tag, dotAfterLang(e.TargetKey))
			fmt.Printf("         %q → %q\n", src, dst)
		}
		if cacheSkipped > 0 {
			fmt.Printf("  (skipped %d entries — not in cache, --cache-only set)\n", cacheSkipped)
		}

		// Write results.
		if len(translated) > 0 {
			written, err := translate.WriteTranslations(translated)
			if err != nil {
				return fmt.Errorf("write %s translations: %w", lm.lang, err)
			}
			fmt.Printf("  Written %d file(s): %s\n", len(written), strings.Join(written, ", "))
		}

		totalCached := cacheHits
		totalAPI := apiCalls * translateBatchSize
		if totalAPI > len(needsAPI) {
			totalAPI = len(needsAPI)
		}
		fmt.Printf("  Summary: %d from cache, %d via API\n\n", totalCached, len(translated)-totalCached)
	}

	fmt.Printf("Total API characters used this run: ~%s\n", formatInt(totalCharsUsed))
	if quotaExhausted {
		fmt.Fprintf(os.Stderr, "Warning: quota limit reached. Run again to continue remaining translations.\n")
	}
	return nil
}

// resolveTargetLangs merges --to flag values and positional args into a deduplicated
// list of target language codes. "--to=all" triggers auto-discovery.
func resolveTargetLangs(flagTo []string, args []string, sourceLang string) ([]string, error) {
	raw := append(append([]string{}, flagTo...), args...)
	// Expand comma-separated values (cobra does this for StringSlice but be safe).
	var expanded []string
	for _, v := range raw {
		for _, part := range strings.Split(v, ",") {
			part = strings.TrimSpace(part)
			if part != "" {
				expanded = append(expanded, part)
			}
		}
	}

	// Handle --to=all.
	for _, v := range expanded {
		if v == "all" {
			return translate.DiscoverLangs(rootPath, sourceLang)
		}
	}

	// Deduplicate, exclude source lang.
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

// confirm prompts the user with a message and returns true if they type the
// expected confirmation string.
func confirm(message, expected string) bool {
	fmt.Printf("%s (%s/no): ", message, expected)
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return false
	}
	return strings.TrimSpace(scanner.Text()) == expected
}

type langMissing struct {
	lang    string
	missing []translate.MissingEntry
}

// printDryRunPlan prints a preview of what would be translated.
func printDryRunPlan(allLangs []langMissing) {
	fmt.Println("=== Dry Run: Translation Plan ===")
	for _, lm := range allLangs {
		chars := translate.CharCount(lm.missing)
		fmt.Printf("\n  %s — %d missing keys (~%s chars):\n", lm.lang, len(lm.missing), formatInt(chars))
		for _, m := range lm.missing {
			key := dotAfterLang(m.TargetKey)
			val := truncate(m.Value, 60)
			fmt.Printf("    [api] %s: %q → (would translate)\n", key, val)
		}
	}
	fmt.Println()
}

func formatInt(n int) string {
	// Simple thousands-separator formatting.
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
