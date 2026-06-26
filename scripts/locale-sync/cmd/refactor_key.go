package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/bobadilla-tech/locale-sync/internal/locale"
	"github.com/spf13/cobra"
)

var (
	refactorOld     string
	refactorNew     string
	refactorDryRun  bool
	refactorAllLang bool
)

var refactorKeyCmd = &cobra.Command{
	Use:   "refactor-key",
	Short: "Rename a locale key and update all t() callers in source files",
	Long: `refactor-key renames a locale key across YAML and source files.

It:
  1. Looks up the value of --old in the locale YAML files
  2. Writes the value under --new in the same YAML file (or shared if --shared)
  3. Rewrites every t('old.key') / t("old.key") reference in app/ source to t('new.key')

The old key is NOT deleted — run ` + "`locale-sync prune`" + ` afterwards to clean it up
once you've verified the app works correctly.

Examples:
  locale-sync refactor-key --old shared.common.copy_btn --new shared.buttons.copy \
    --root ../../apps/dashboard --lang en
  locale-sync refactor-key --old shared.close --new shared.buttons.close \
    --root ../../apps/dashboard --lang en --dry-run`,
	RunE: runRefactorKey,
}

func init() {
	refactorKeyCmd.Flags().StringVarP(&rootPath, "root", "r", ".", "Rails app root directory")
	refactorKeyCmd.Flags().StringVar(&refactorOld, "old", "", "Current dot-notation key path (required)")
	refactorKeyCmd.Flags().StringVar(&refactorNew, "new", "", "New dot-notation key path (required)")
	refactorKeyCmd.Flags().StringVarP(&pruneLang, "lang", "l", "en", "Source language code")
	refactorKeyCmd.Flags().BoolVar(&refactorDryRun, "dry-run", false, "Preview changes without writing")
	refactorKeyCmd.Flags().BoolVar(&refactorAllLang, "all-lang", true,
		"Also rename the key in all other language files (default true)")
	if err := refactorKeyCmd.MarkFlagRequired("old"); err != nil {
		panic(err)
	}
	if err := refactorKeyCmd.MarkFlagRequired("new"); err != nil {
		panic(err)
	}
	rootCmd.AddCommand(refactorKeyCmd)
}

func runRefactorKey(_ *cobra.Command, _ []string) error {
	if refactorOld == refactorNew {
		return fmt.Errorf("--old and --new are the same key: %q", refactorOld)
	}

	// ── 1. Discover locale files that contain --old ────────────────────────
	entries, err := locale.Scan(rootPath, pruneLang)
	if err != nil {
		return fmt.Errorf("scan locales: %w", err)
	}

	// Find every locale file that holds the old key (lang prefix is optional in --old).
	type fileOccurrence struct {
		file  string
		value string
		key   string // full key as stored in YAML
	}
	var occurrences []fileOccurrence
	for _, e := range entries {
		// Match with or without leading lang prefix.
		keyWithout := stripLangPrefixStr(e.Key, pruneLang)
		if keyWithout == refactorOld || e.Key == refactorOld {
			occurrences = append(occurrences, fileOccurrence{
				file:  e.File,
				value: e.Value,
				key:   e.Key,
			})
		}
	}

	if len(occurrences) == 0 {
		return fmt.Errorf("key %q not found in %s locale files", refactorOld, pruneLang)
	}

	fmt.Printf("Found %d occurrence(s) of %q:\n", len(occurrences), refactorOld)
	for _, o := range occurrences {
		fmt.Printf("  %s  →  %q\n", shortPath(o.file, rootPath), o.value)
	}
	fmt.Println()

	// ── 2. Discover source-file callers ───────────────────────────────────
	callers, err := findTCallers(rootPath, refactorOld)
	if err != nil {
		return fmt.Errorf("scan source files: %w", err)
	}

	fmt.Printf("Found %d t() caller(s) of %q:\n", len(callers), refactorOld)
	for _, c := range callers {
		fmt.Printf("  %s:%d\n", c.relFile, c.line)
	}
	fmt.Println()

	if refactorDryRun {
		fmt.Printf("YAML: add %q → %q in %d file(s)\n", refactorNew, occurrences[0].value, len(occurrences))
		fmt.Printf("Source: rewrite %d caller(s) to t('%s')\n", len(callers), refactorNew)
		fmt.Printf("--dry-run: no files modified.\n")
		return nil
	}

	// ── 3. Write new key to YAML ──────────────────────────────────────────
	for _, o := range occurrences {
		// Determine lang for this file by reading its lang prefix.
		fileLang := extractLangFromKey(o.key)
		// Build new full key: replace lang prefix + old path with lang prefix + new path.
		newFullKey := fileLang + "." + stripLangPrefixStr(refactorNew, fileLang)
		candidate := locale.MergeCandidate{
			KeyName:      lastSegment(refactorNew),
			Value:        o.value,
			SuggestedKey: newFullKey,
		}
		_, changed, werr := locale.UpsertTopicFile(rootPath, fileLang, topicFromFile(o.file, fileLang), []locale.MergeCandidate{candidate})
		if werr != nil {
			fmt.Fprintf(os.Stderr, "warning: could not write new key to %s: %v\n", shortPath(o.file, rootPath), werr)
			continue
		}
		if changed {
			fmt.Printf("YAML: added %q to %s\n", newFullKey, shortPath(o.file, rootPath))
		} else {
			fmt.Printf("YAML: %q already present in %s (skipped)\n", newFullKey, shortPath(o.file, rootPath))
		}
	}

	// ── 4. Rewrite t() callers in source ──────────────────────────────────
	if err := applyCallerRewrites(rootPath, callers, refactorOld, refactorNew); err != nil {
		return fmt.Errorf("rewrite source files: %w", err)
	}

	fmt.Printf("\nDone. Old key %q preserved — run `locale-sync prune` to remove it after verification.\n", refactorOld)
	return nil
}

// ── helpers ───────────────────────────────────────────────────────────────────

type tCaller struct {
	relFile string
	absFile string
	line    int
	raw     string // original line text
}

// tCallRe matches t('key') and t("key") with or without I18n. prefix.
// Capture group 1 = quote char, capture group 2 = key path.
// RE2 has no backreferences; match either single- or double-quoted key separately.
var tCallRe = regexp.MustCompile(`(?:I18n\.)?t\((?:'([^']+)'|"([^"]+)")\)`)

// findTCallers scans all .rb/.erb/.haml/.slim files under {root}/app/ for t() calls
// referencing keyPath (with or without leading lang prefix).
func findTCallers(root, keyPath string) ([]tCaller, error) {
	appDir := filepath.Join(root, "app")
	if _, err := os.Stat(appDir); os.IsNotExist(err) {
		appDir = root
	}

	var callers []tCaller
	sourceExts := map[string]bool{".rb": true, ".erb": true, ".haml": true, ".slim": true}

	err := filepath.WalkDir(appDir, func(path string, d fs.DirEntry, werr error) error {
		if werr != nil || d.IsDir() {
			return nil
		}
		if !sourceExts[filepath.Ext(path)] {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return nil // skip unreadable
		}
		rel, _ := filepath.Rel(root, path)
		for i, line := range strings.Split(string(data), "\n") {
			matches := tCallRe.FindAllStringSubmatch(line, -1)
			for _, m := range matches {
				// m[1] = single-quoted key, m[2] = double-quoted key
				k := m[1]
				if k == "" {
					k = m[2]
				}
				// Strip leading lang prefix if present.
				for _, lang := range []string{"en.", "es.", "fr.", "de.", "pt.", "ja.", "zh."} {
					k = strings.TrimPrefix(k, lang)
				}
				if k == keyPath {
					callers = append(callers, tCaller{
						relFile: rel,
						absFile: path,
						line:    i + 1,
						raw:     line,
					})
				}
			}
		}
		return nil
	})
	return callers, err
}

// applyCallerRewrites replaces t('oldKey') with t('newKey') in all caller files.
func applyCallerRewrites(root string, callers []tCaller, oldKey, newKey string) error {
	byFile := map[string][]tCaller{}
	for _, c := range callers {
		byFile[c.absFile] = append(byFile[c.absFile], c)
	}

	esc := regexp.QuoteMeta(oldKey)
	// Two separate patterns — one per quote style — avoids RE2's lack of backreferences.
	singleQ := regexp.MustCompile(`((?:I18n\.)?t\()('` + esc + `')(\))`)
	doubleQ := regexp.MustCompile(`((?:I18n\.)?t\()("` + esc + `")(\))`)

	for absFile, fileCalls := range byFile {
		data, err := os.ReadFile(absFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: cannot read %s: %v\n", absFile, err)
			continue
		}
		content := string(data)
		content = singleQ.ReplaceAllString(content, "${1}'"+newKey+"'${3}")
		content = doubleQ.ReplaceAllString(content, `${1}"`+newKey+`"${3}`)
		if string(data) == content {
			continue
		}
		if err := os.WriteFile(absFile, []byte(content), 0o644); err != nil {
			fmt.Fprintf(os.Stderr, "warning: cannot write %s: %v\n", absFile, err)
			continue
		}
		rel, _ := filepath.Rel(root, absFile)
		fmt.Printf("Source: rewrote %d caller(s) in %s\n", len(fileCalls), rel)
	}
	return nil
}

func extractLangFromKey(fullKey string) string {
	parts := strings.SplitN(fullKey, ".", 2)
	if len(parts) > 0 {
		return parts[0]
	}
	return "en"
}

func lastSegment(key string) string {
	parts := strings.Split(key, ".")
	return parts[len(parts)-1]
}

// topicFromFile derives the topic name from an absolute locale file path.
// e.g. "/…/config/locales/en/shared.en.yml" → "shared"
func topicFromFile(absFile, lang string) string {
	base := filepath.Base(absFile)
	base = strings.TrimSuffix(base, "."+lang+".yml")
	base = strings.TrimSuffix(base, ".yml")
	return base
}

func stripLangPrefixStr(key, lang string) string {
	prefix := lang + "."
	if strings.HasPrefix(key, prefix) {
		return key[len(prefix):]
	}
	return key
}
