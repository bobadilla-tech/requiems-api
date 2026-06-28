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

// ── HTML-in-locale tests ──────────────────────────────────────────────────────

func TestLintValues_HtmlWithCssClasses(t *testing.T) {
	// _html key with Tailwind class attribute — should be flagged.
	e := makeEntry("en.admin.analytics.revenue.low_churn_html",
		`<strong class="text-green-600">Low churn rate:</strong> Customers are satisfied`)
	got := LintValues([]Entry{e})
	if len(got) != 1 {
		t.Fatalf("expected 1 finding for _html key with CSS class, got %d: %v", len(got), got)
	}
}

func TestLintValues_HtmlInNonHtmlKey(t *testing.T) {
	// Non-_html key with <code> — Rails won't html_safe this; should be flagged.
	e := makeEntry("en.home.developer_docs.examples.disposable_email.example",
		`Example: <code class="bg-gray-100 px-2 py-1 rounded">POST /v1/networking/check</code>`)
	got := LintValues([]Entry{e})
	if len(got) != 1 {
		t.Fatalf("expected 1 finding for HTML in non-_html key, got %d: %v", len(got), got)
	}
}

func TestLintValues_SimpleInlineHtmlNotFlagged(t *testing.T) {
	// _html key with plain <strong> wrapping an interpolation variable — no CSS class,
	// acceptable Rails pattern; must NOT be flagged.
	e := makeEntry("en.shared.formats.greeting_html", "Hi <strong>%{email}</strong>,")
	got := LintValues([]Entry{e})
	if len(got) != 0 {
		t.Errorf("expected 0 findings for simple inline _html, got %d: %v", len(got), got)
	}
}

func TestLintValues_CodePlaceholderNotFlagged(t *testing.T) {
	// <key> in a migration step is a code placeholder, not an HTML tag.
	// knownHtmlTagRe must not match it.
	e := makeEntry("en.comparisons.steps.0",
		"Swap `apikey` query param for `Authorization: Bearer <key>` header")
	got := LintValues([]Entry{e})
	if len(got) != 0 {
		t.Errorf("expected 0 findings for <key> code placeholder, got %d: %v", len(got), got)
	}
}

// ── Regression: HTML-in-locale patterns found in locale-sync hygiene pass ────

// TestRegression_AdminChurnHtml guards the pattern where admin analytics insight
// keys used _html suffix to embed Tailwind colour classes directly in translations.
// Example: low_churn_html: "<strong class=\"text-green-600\">Low churn rate:</strong> ..."
func TestRegression_AdminChurnHtml(t *testing.T) {
	cases := []struct{ key, value string }{
		{"en.admin.analytics.revenue.low_churn_html", `<strong class="text-green-600">Low churn rate:</strong> Customers are satisfied`},
		{"en.admin.analytics.revenue.high_churn_html", `<strong class="text-red-600">High churn rate:</strong> Focus on customer retention`},
		{"en.admin.analytics.uptime.uptime_excellent_html", `<strong class="text-green-600">Excellent uptime.</strong> The API is highly reliable.`},
	}
	for _, tc := range cases {
		t.Run(tc.key, func(t *testing.T) {
			got := LintValues([]Entry{makeEntry(tc.key, tc.value)})
			if len(got) == 0 {
				t.Errorf("LintValues(%q) returned 0 findings — CSS in _html key must be flagged", tc.key)
			}
		})
	}
}

// TestRegression_HomeExampleWithCode guards the pattern where home.en.yml FAQ/tutorial
// `example:` keys (no _html suffix) embedded <code> tags with full Tailwind classes.
func TestRegression_HomeExampleWithCode(t *testing.T) {
	e := makeEntry("en.home.developer_docs.steps.auth.example",
		`Example: <code class="bg-gray-100 dark:bg-gray-700 px-2 py-1 rounded">Authorization: Bearer YOUR_API_KEY</code>`)
	got := LintValues([]Entry{e})
	if len(got) == 0 {
		t.Error("LintValues returned 0 findings — <code> in non-_html example key must be flagged")
	}
}

// TestLintValues_MultipleEmailsNotFlagged guards that illustrative format strings
// containing two or more @ signs (e.g. normalization examples) are not flagged
// as contact email addresses.
func TestLintValues_MultipleEmailsNotFlagged(t *testing.T) {
	v := "Ensure User.Name@Gmail.com and username@gmail.com match as the same identity."
	got := LintValues([]Entry{makeEntry("en.tools.email_normalizer.use_cases.integrations_desc", v)})
	if len(got) != 0 {
		t.Errorf("LintValues flagged illustrative multi-email string: %v", got[0].Reason)
	}
}

// TestLintValues_MultilineJsonNotFlagged guards that multi-line JSON code blocks
// containing an email address (e.g. API example payloads) are not flagged.
func TestLintValues_MultilineJsonNotFlagged(t *testing.T) {
	v := "{\n  \"email\": \"user@tempmail.io\",\n  \"ip_address\": \"45.33.32.156\"\n}\n"
	got := LintValues([]Entry{makeEntry("en.home.index.engine_spotlight.request", v)})
	if len(got) != 0 {
		t.Errorf("LintValues flagged multiline JSON payload: %v", got[0].Reason)
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
