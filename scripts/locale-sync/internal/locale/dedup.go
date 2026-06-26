package locale

import (
	"sort"
	"strings"
)

// KeyDupGroup holds entries sharing the same leaf key name across different paths/files.
type KeyDupGroup struct {
	KeyName string
	Entries []Entry
	// SameValue is true when all entries have identical text — safe to consolidate automatically.
	SameValue bool
}

// ValueDupGroup holds entries that share the exact same string value under different keys.
type ValueDupGroup struct {
	Value   string
	Entries []Entry
}

// Duplicates is the complete duplication report for a set of locale entries.
type Duplicates struct {
	// KeyDups: same leaf key name (e.g. error_rate_limit) in 2+ places.
	KeyDups []KeyDupGroup
	// ValueDups: same text value under 2+ different full key paths.
	ValueDups []ValueDupGroup
}

// FindDuplicates scans entries and returns all duplication groups with at least
// minOccurrences members. Entries in the default shared file are excluded from
// being flagged as duplicates (they're already shared).
func FindDuplicates(entries []Entry, minOccurrences int) Duplicates {
	// --- key-name duplicates ---
	byKeyName := make(map[string][]Entry)
	for _, e := range entries {
		if isSharedFile(e.ShortPath) {
			continue
		}
		byKeyName[e.KeyName] = append(byKeyName[e.KeyName], e)
	}

	var keyDups []KeyDupGroup
	for name, group := range byKeyName {
		if len(group) < minOccurrences {
			continue
		}
		sameVal := allSameValue(group)
		keyDups = append(keyDups, KeyDupGroup{
			KeyName:   name,
			Entries:   group,
			SameValue: sameVal,
		})
	}
	sort.Slice(keyDups, func(i, j int) bool {
		if len(keyDups[i].Entries) != len(keyDups[j].Entries) {
			return len(keyDups[i].Entries) > len(keyDups[j].Entries)
		}
		return keyDups[i].KeyName < keyDups[j].KeyName
	})

	// --- value duplicates ---
	byValue := make(map[string][]Entry)
	for _, e := range entries {
		if isSharedFile(e.ShortPath) || e.Value == "" || looksLikePlaceholder(e.Value) {
			continue
		}
		byValue[e.Value] = append(byValue[e.Value], e)
	}

	var valueDups []ValueDupGroup
	for val, group := range byValue {
		if len(group) < minOccurrences {
			continue
		}
		// De-duplicate entries with same key path (can't be in two places)
		group = uniqueByKey(group)
		if len(group) < minOccurrences {
			continue
		}
		valueDups = append(valueDups, ValueDupGroup{
			Value:   val,
			Entries: group,
		})
	}
	sort.Slice(valueDups, func(i, j int) bool {
		if len(valueDups[i].Entries) != len(valueDups[j].Entries) {
			return len(valueDups[i].Entries) > len(valueDups[j].Entries)
		}
		return valueDups[i].Value < valueDups[j].Value
	})

	return Duplicates{
		KeyDups:   keyDups,
		ValueDups: valueDups,
	}
}

// sharedRoutes maps key-name prefixes/exact matches to their semantic shared namespace.
// First match wins. Order matters — longer/more-specific prefixes first.
var sharedRoutes = []struct {
	prefixes []string
	target   string
}{
	{[]string{"error_", "err_"}, "errors"},
	{[]string{"badge_"}, "badges"},
	{[]string{"active", "suspended", "reported", "revoked", "pending", "status"}, "status"},
	{[]string{"copy", "cancel", "close", "dismiss", "save", "back", "delete", "edit",
		"submit", "confirm", "get_api_key", "read_docs", "contact_support",
		"browse_apis", "go_to_dashboard", "try_live", "explore_system",
		"apply_filters", "change_password"}, "buttons"},
}

// SuggestSharedKey proposes a semantic shared key path for a duplicate group.
// e.g. "error_rate_limit" → "en.shared.errors.rate_limit"
// e.g. "copy_btn"        → "en.shared.buttons.copy_btn"
// e.g. "badge_valid"     → "en.shared.badges.valid"
func SuggestSharedKey(lang, keyName string) string {
	lower := strings.ToLower(keyName)
	for _, route := range sharedRoutes {
		for _, prefix := range route.prefixes {
			if strings.HasPrefix(lower, prefix) || lower == strings.TrimSuffix(prefix, "_") {
				// Strip well-known prefixes so badge_valid → badges.valid (not badges.badge_valid).
				clean := strings.TrimPrefix(lower, "badge_")
				clean = strings.TrimPrefix(clean, "error_")
				clean = strings.TrimPrefix(clean, "err_")
				// For exact matches (e.g. "copy", "active") keep original casing from keyName.
				if clean == lower {
					clean = keyName
				}
				return lang + ".shared." + route.target + "." + clean
			}
		}
	}
	return lang + ".shared.common." + keyName
}

func isErrorKey(k string) bool {
	return strings.HasPrefix(k, "error_") || strings.HasPrefix(k, "err_")
}

func isSharedFile(short string) bool {
	return strings.Contains(short, "shared.")
}

func looksLikePlaceholder(s string) bool {
	return strings.Contains(s, "%{") ||
		s == "true" || s == "false" ||
		len(s) < 2
}

func allSameValue(entries []Entry) bool {
	if len(entries) == 0 {
		return true
	}
	v := entries[0].Value
	for _, e := range entries[1:] {
		if e.Value != v {
			return false
		}
	}
	return true
}

func uniqueByKey(entries []Entry) []Entry {
	seen := make(map[string]bool)
	out := entries[:0]
	for _, e := range entries {
		if !seen[e.Key] {
			seen[e.Key] = true
			out = append(out, e)
		}
	}
	return out
}
