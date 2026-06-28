package translate

import (
	"fmt"
	"path/filepath"
	"regexp"
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

// machinePlaceholderRe matches machine-generated placeholder strings like
// "TODO: translate", "TODO:translate", "FIXME: review", "FIXME:check".
// It requires the word after the colon to be lowercase — human-written prose
// that begins with "TODO:" (e.g. Spanish "TODO: Necesitamos...") uses
// uppercase after the colon and is NOT treated as a placeholder.
var machinePlaceholderRe = regexp.MustCompile(`^(?:TODO|FIXME):\s*[a-z]`)

// IsPlaceholder reports whether a locale value is an untranslated placeholder
// that should be overwritten by a real translation.
// extraPlaceholders are additional values to treat as placeholders (e.g. the
// value used by `generate --placeholder`).
func IsPlaceholder(v string, extraPlaceholders ...string) bool {
	v = strings.TrimSpace(v)
	if v == "" || machinePlaceholderRe.MatchString(v) {
		return true
	}
	for _, p := range extraPlaceholders {
		if p != "" && v == p {
			return true
		}
	}
	return false
}

// FindMissing returns source entries whose keys are absent in target entries
// or whose target value is a known placeholder (e.g. "TODO: translate").
// skeletonPlaceholder is the custom placeholder value used by `generate
// --placeholder`; empty string is always treated as a placeholder regardless.
func FindMissing(
	sourceEntries []locale.Entry,
	targetEntries []locale.Entry,
	sourceLang, targetLang, root string,
	skeletonPlaceholder string,
) []MissingEntry {
	targetSet := make(map[string]bool, len(targetEntries))
	for _, e := range targetEntries {
		if IsPlaceholder(e.Value, skeletonPlaceholder) {
			continue // treat placeholder as absent so it gets translated
		}
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

// pluralForms are the Rails i18n pluralization key suffixes.
var pluralForms = map[string]bool{
	"zero": true, "one": true, "two": true,
	"few": true, "many": true, "other": true,
}

// pluralSep is a protected token injected between plural sibling values before
// sending to the translation API. It is added to the Protector's custom-words
// list so it survives translation unchanged.
const PluralSep = "XPLURALSEPX"

// PluralGroup holds sibling pluralization entries that share a parent key
// (e.g. "en.items.one" and "en.items.other" → parent "items").
type PluralGroup struct {
	Entries []MissingEntry // one per plural form, sorted by leaf name
	Parent  string        // shared key prefix, target-lang qualified
}

// GroupPlurals partitions missing entries into singletons and groups of plural
// siblings. Plural siblings are entries whose leaf key is a Rails i18n plural
// form ("one", "other", "zero", etc.) AND share the same parent key. Groups
// with only one plural form found (the sibling is already translated) are kept
// as singletons so they are still translated individually.
func GroupPlurals(missing []MissingEntry) (singles []MissingEntry, groups []PluralGroup) {
	parentMap := map[string][]MissingEntry{}
	var nonPlural []MissingEntry

	for _, m := range missing {
		parts := strings.Split(m.TargetKey, ".")
		if len(parts) >= 2 && pluralForms[parts[len(parts)-1]] {
			parent := strings.Join(parts[:len(parts)-1], ".")
			parentMap[parent] = append(parentMap[parent], m)
		} else {
			nonPlural = append(nonPlural, m)
		}
	}

	for parent, entries := range parentMap {
		if len(entries) >= 2 {
			groups = append(groups, PluralGroup{Entries: entries, Parent: parent})
		} else {
			singles = append(singles, entries...)
		}
	}
	singles = append(singles, nonPlural...)
	return
}

// CombinePluralGroup serializes plural sibling values into one string for the
// translation API, separated by PluralSep. Returns the combined value and a
// slice of leaf-key names in the same order.
func CombinePluralGroup(g PluralGroup) (combined string, order []string) {
	parts := make([]string, len(g.Entries))
	order = make([]string, len(g.Entries))
	for i, e := range g.Entries {
		kParts := strings.Split(e.TargetKey, ".")
		order[i] = kParts[len(kParts)-1]
		parts[i] = e.Entry.Value
	}
	return strings.Join(parts, " "+PluralSep+" "), order
}

// SplitPluralTranslation splits a translated combined plural string back into
// per-form values using order returned by CombinePluralGroup. Falls back to
// repeating the full string for all forms if the separator is missing (e.g.
// translation API stripped it).
func SplitPluralTranslation(translated string, order []string) map[string]string {
	parts := strings.Split(translated, " "+PluralSep+" ")
	result := make(map[string]string, len(order))
	for i, form := range order {
		if i < len(parts) {
			result[form] = strings.TrimSpace(parts[i])
		} else {
			result[form] = strings.TrimSpace(parts[0]) // fallback
		}
	}
	return result
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
