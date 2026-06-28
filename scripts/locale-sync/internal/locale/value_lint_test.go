package locale

import "testing"

func makeEntry(key, value string) Entry {
	return Entry{Key: key, KeyName: key, Value: value, File: "en/shared.en.yml", Line: 1}
}

func TestLintValues_HttpHeader(t *testing.T) {
	got := LintValues([]Entry{makeEntry("en.shared.common.x_backend_secret", "X-Backend-Secret")})
	if len(got) != 1 {
		t.Fatalf("expected 1 finding for HTTP header, got %d: %v", len(got), got)
	}
	if got[0].Key != "en.shared.common.x_backend_secret" {
		t.Errorf("Key = %q, want %q", got[0].Key, "en.shared.common.x_backend_secret")
	}
}

func TestLintValues_Email(t *testing.T) {
	got := LintValues([]Entry{makeEntry("en.shared.common.eliaz_bobadilla_tech", "eliaz@bobadilla.tech")})
	if len(got) != 1 {
		t.Fatalf("expected 1 finding for email, got %d: %v", len(got), got)
	}
}

func TestLintValues_MetricCode(t *testing.T) {
	entries := []Entry{
		makeEntry("en.shared.common.p50", "p50"),
		makeEntry("en.shared.common.p99", "p99"),
	}
	got := LintValues(entries)
	if len(got) != 2 {
		t.Fatalf("expected 2 findings for metric codes, got %d: %v", len(got), got)
	}
}

func TestLintValues_CleanValues(t *testing.T) {
	entries := []Entry{
		makeEntry("en.shared.buttons.cancel", "Cancel"),
		makeEntry("en.shared.labels.view", "View"),
		makeEntry("en.shared.status.active", "Active"),
	}
	got := LintValues(entries)
	if len(got) != 0 {
		t.Errorf("expected 0 findings for clean values, got %d: %v", len(got), got)
	}
}

func TestLintValues_MultiHyphenWithSpace(t *testing.T) {
	// "well-known key" has 2 hyphens but also a space → not an HTTP header pattern.
	got := LintValues([]Entry{makeEntry("en.foo.bar", "well-known key")})
	if len(got) != 0 {
		t.Errorf("expected 0 findings for hyphenated phrase with space, got %d: %v", len(got), got)
	}
}

func TestLintValues_SkipsEmpty(t *testing.T) {
	got := LintValues([]Entry{makeEntry("en.foo", ""), makeEntry("en.bar", "[array]")})
	if len(got) != 0 {
		t.Errorf("expected 0 findings for empty/array entries, got %d: %v", len(got), got)
	}
}

// TestRegression_SharedCommonBadEntries guards against the specific bad values that
// were added to en/fr/es shared.common and removed in the locale-sync hygiene pass.
// If any of these values reappear in a locale file, LintValues must flag them.
func TestRegression_SharedCommonBadEntries(t *testing.T) {
	cases := []struct {
		key   string
		value string
	}{
		// HTTP header name — technical identifier, not translatable UI copy.
		{"en.shared.common.x_backend_secret", "X-Backend-Secret"},
		// Personal email address — must never live in a locale file.
		{"en.shared.common.eliaz_bobadilla_tech", "eliaz@bobadilla.tech"},
		// Metric abbreviations — language-neutral codes, no translation needed.
		{"en.shared.common.p50", "p50"},
		{"en.shared.common.p99", "p99"},
	}
	for _, tc := range cases {
		t.Run(tc.key, func(t *testing.T) {
			got := LintValues([]Entry{makeEntry(tc.key, tc.value)})
			if len(got) == 0 {
				t.Errorf("LintValues(%q = %q) returned 0 findings — this value should be flagged as technical",
					tc.key, tc.value)
			}
		})
	}
}
