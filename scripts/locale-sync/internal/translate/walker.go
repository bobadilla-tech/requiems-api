package translate

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/bobadilla-tech/locale-sync/internal/locale"
)

// MissingEntry describes one source-language locale entry that is absent in
// the target language.
type MissingEntry struct {
	locale.Entry           // source entry: Key, Value, ShortPath, etc.
	TargetKey  string      // full dot-notation key with target lang prefix: "es.tools.foo"
	TargetFile string      // absolute path to the target locale file
}

// FindMissing returns source entries whose keys are not present in target entries.
// Empty-valued source entries (skeleton placeholders) are skipped.
func FindMissing(
	sourceEntries []locale.Entry,
	targetEntries []locale.Entry,
	sourceLang, targetLang, root string,
) []MissingEntry {
	targetSet := make(map[string]bool, len(targetEntries))
	for _, e := range targetEntries {
		// Strip leading lang prefix for comparison.
		targetSet[stripLang(e.Key, targetLang)] = true
	}

	var missing []MissingEntry
	for _, e := range sourceEntries {
		if e.Value == "" {
			continue
		}
		keyWithoutLang := stripLang(e.Key, sourceLang)
		if targetSet[keyWithoutLang] {
			continue
		}
		targetKey := targetLang + "." + keyWithoutLang
		targetFile := targetFilePath(root, e.ShortPath, sourceLang, targetLang)
		missing = append(missing, MissingEntry{
			Entry:      e,
			TargetKey:  targetKey,
			TargetFile: targetFile,
		})
	}
	return missing
}

// DiscoverLangs returns all language codes that have a subdirectory under
// {root}/config/locales/, excluding the source language.
func DiscoverLangs(root, sourceLang string) ([]string, error) {
	localesDir := filepath.Join(root, "config", "locales")
	entries, err := filepath.Glob(filepath.Join(localesDir, "*"))
	if err != nil {
		return nil, fmt.Errorf("glob locales: %w", err)
	}
	var langs []string
	for _, e := range entries {
		base := filepath.Base(e)
		// Only directories, not files; exclude source lang.
		if base == sourceLang {
			continue
		}
		// Check it's a directory (language subdirs, not yml files at root level).
		if !strings.Contains(base, ".") {
			langs = append(langs, base)
		}
	}
	return langs, nil
}

// CharCount returns the total UTF-8 character count across all entry values.
func CharCount(entries []MissingEntry) int {
	n := 0
	for _, e := range entries {
		n += len([]rune(e.Value))
	}
	return n
}

func stripLang(key, lang string) string {
	prefix := lang + "."
	if strings.HasPrefix(key, prefix) {
		return key[len(prefix):]
	}
	return key
}

// targetFilePath converts a source ShortPath to an absolute target file path.
// "en/tools.en.yml" → "{root}/config/locales/es/tools.es.yml"
func targetFilePath(root, shortPath, sourceLang, targetLang string) string {
	// shortPath is relative to config/locales/, e.g. "en/tools.en.yml"
	base := filepath.Base(shortPath)
	// Replace ".{sourceLang}.yml" suffix with ".{targetLang}.yml"
	base = strings.TrimSuffix(base, "."+sourceLang+".yml") + "." + targetLang + ".yml"
	return filepath.Join(root, "config", "locales", targetLang, base)
}
