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

// keyAliases normalizes legacy key names to their canonical clean names.
// Prevents generate from writing copy_btn alongside the already-existing copy, etc.
var keyAliases = map[string]string{
	"copy_btn":               "copy",
	"copied_btn":             "copied",
	"status_col":             "col_status",
	"badge_json":             "json_response",
	"badge_rest":             "rest_api",
	"badge_uptime":           "uptime",
	"badge_latency":          "latency",
	"badge_yes":              "yes",
	"badge_no":               "no",
	"badge_valid":            "valid",
	"badge_invalid":          "invalid",
	"popular_badge":          "popular",
	"read_the_docs":          "read_docs",
	"support_link":           "contact_support",
	"apis_catalog_cta":       "browse_apis",
	"cta_btn":                "go_to_dashboard",
	"explore_cta":            "explore_system",
	"change_my_password":     "change_password",
	"utilities_cta":          "utilities",
	"examples_title":         "examples",
	"accuracy_heading":       "accuracy",
	"simplicity_heading":     "simplicity",
	"apis_catalog_title":     "apis_catalog",
	"use_cases_heading":      "use_cases_subheading",
	"integrations_context":   "integrations",
	"auth_context":           "auth",
	"onboarding_context":     "onboarding",
	"fraud_context":          "fraud",
	"label_normalized":       "normalized",
	"label_benefit":          "benefit",
	"label_outcome":          "outcome",
	"label_domain":           "domain",
	"we_usually_reply_within_24_hours": "we_usually_reply",
	"try_adjusting_your_search_or_filters": "try_search",
	"email_input_placeholder": "email_placeholder",
	"are_you_sure_you_want_to_suspend_this": "are_you_sure_suspend",
}

// keyNamespace overrides the namespace for keys that don't match route prefixes cleanly.
var keyNamespace = map[string]string{
	"col_status":          "labels",
	"status_col":          "labels",
	"admin":               "labels",
	"admin_actions":       "labels",
	"environment":         "labels",
	"joined":              "labels",
	"monthly":             "labels",
	"yearly":              "labels",
	"this_month":          "labels",
	"key_prefix":          "labels",
	"system_health":       "labels",
	"api_reference":       "labels",
	"are_you_sure_suspend": "labels",
	"no_account":          "info",
	"no_account_demo":     "info",
	"coverage":            "info",
	"support_prompt":      "info",
	"eu_title":            "info",
	"soc2_desc":           "info",
	"we_usually_reply":    "info",
	"time_ago":            "formats",
	"of_requests":         "formats",
	"greeting_html":       "formats",
	"no_data":             "empty_states",
	"try_search":          "empty_states",
	"examples":            "headings",
	"use_cases":           "headings",
	"use_cases_subheading": "headings",
	"accuracy":            "headings",
	"simplicity":          "headings",
	"apis_catalog":        "headings",
	"integrations":        "contexts",
	"auth":                "contexts",
	"onboarding":          "contexts",
	"fraud":               "contexts",
	"email_placeholder":   "forms",
	"text_to_analyze":     "forms",
	"password":            "forms",
	"missing_fields":      "forms",
	"domain":              "forms",
	"normalized":          "result_labels",
	"benefit":             "result_labels",
	"outcome":             "result_labels",
}

// SuggestSharedKey proposes a semantic shared key path for a duplicate group.
// e.g. "error_rate_limit" → "en.shared.errors.rate_limit"
// e.g. "copy_btn"        → "en.shared.buttons.copy"   (alias normalised)
// e.g. "badge_valid"     → "en.shared.badges.valid"
func SuggestSharedKey(lang, keyName string) string {
	// Normalise legacy names first.
	if canonical, ok := keyAliases[strings.ToLower(keyName)]; ok {
		keyName = canonical
	}

	lower := strings.ToLower(keyName)

	// Namespace override table wins before prefix routing.
	if ns, ok := keyNamespace[lower]; ok {
		return lang + ".shared." + ns + "." + lower
	}

	for _, route := range sharedRoutes {
		for _, prefix := range route.prefixes {
			if strings.HasPrefix(lower, prefix) || lower == strings.TrimSuffix(prefix, "_") {
				// Strip well-known prefixes so badge_valid → badges.valid.
				clean := strings.TrimPrefix(lower, "badge_")
				clean = strings.TrimPrefix(clean, "error_")
				clean = strings.TrimPrefix(clean, "err_")
				if clean == lower {
					clean = keyName
				}
				return lang + ".shared." + route.target + "." + clean
			}
		}
	}
	return lang + ".shared.common." + keyName
}

// FilterExistingValueCandidates drops candidates already covered by shared:
// either their value exists in shared (under any key) or their SuggestedKey is
// already defined. The key check catches interpolated values (%{…}) that
// looksLikePlaceholder excludes from the value dedup set.
func FilterExistingValueCandidates(candidates []MergeCandidate, allEntries []Entry) []MergeCandidate {
	sharedValues := make(map[string]bool)
	sharedKeys := make(map[string]bool)
	for _, e := range allEntries {
		if isSharedFile(e.ShortPath) {
			sharedKeys[e.Key] = true
			if e.Value != "" && !looksLikePlaceholder(e.Value) {
				sharedValues[e.Value] = true
			}
		}
	}
	out := candidates[:0]
	for _, c := range candidates {
		if !sharedValues[c.Value] && !sharedKeys[c.SuggestedKey] {
			out = append(out, c)
		}
	}
	return out
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
