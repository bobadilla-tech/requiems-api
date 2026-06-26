package source

import (
	"testing"
	"strings"
)

// ── looksUserFacing ───────────────────────────────────────────────────────────

func TestLooksUserFacing(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  bool
	}{
		{"too short 3 chars", "Hi!", false},
		{"exactly 4 chars", "Help", false}, // no space → false
		{"normal sentence", "Submit your form", true},
		{"no space single word", "Dashboard", false},
		{"has space passes", "API Key", true},
		{"too long 121 chars", strings.Repeat("a", 121), false},
		{"ruby interpolation", "Hello #{user.name}", false},
		{"contains class=", `class="foo bar"`, false},
		{"contains data-", "data-controller value", false},
		{"contains https://", "Visit https://example.com", false},
		{"contains <%", "<% foo %>", false},
		{"starts with digit", "42 items found", true},
		{"brand name exact", "Requiems API", false}, // brandNames skip
		{"brand name in sentence", "Use Requiems API today", true}, // not exact match
		{"contains select sql", "select from table", false},
		{"css asset ref .css", "app.css file", false},
		{"contains application/json", "content-type application/json", false},
		{"valid spanish-like", "Hola mundo amigos", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := looksUserFacing(tc.input)
			if got != tc.want {
				t.Errorf("looksUserFacing(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

// ── looksLikeCode ─────────────────────────────────────────────────────────────

func TestLooksLikeCode(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  bool
	}{
		{"css class string", "flex items-center gap-3", true},
		{"high hyphen density", "text-gray-900 dark:bg-gray-800", true},
		{"tailwind combo", "px-4 py-2 bg-blue-600 rounded", true},
		{"normal sentence", "Submit your form today", false},
		{"two css keywords", "px-4 flex items", true},
		{"one css keyword", "flex layout", false}, // only 1 match
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := looksLikeCode(tc.input)
			if got != tc.want {
				t.Errorf("looksLikeCode(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

// ── scanTableHeaders ──────────────────────────────────────────────────────────

func TestScanTableHeaders(t *testing.T) {
	cases := []struct {
		name  string
		line  string
		wantN int
		wantT string // first detected text if wantN > 0
	}{
		{
			"plain th",
			`<th>Name</th>`,
			1, "Name",
		},
		{
			"th with class",
			`<th class="px-4 py-3 uppercase">Type</th>`,
			1, "Type",
		},
		{
			"th already i18n'd",
			`<th><%= t('admin.name') %></th>`,
			0, "",
		},
		{
			"th single letter skipped",
			`<th>A</th>`,
			0, "",
		},
		{
			"th with ERB output skipped",
			`<th><%= @value %></th>`,
			0, "",
		},
		{
			"multiple th in one line",
			`<th>Name</th><th>Status</th>`,
			2, "Name",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := scanTableHeaders(tc.line, tc.line, "test.erb", 1)
			if len(got) != tc.wantN {
				t.Errorf("scanTableHeaders(%q): got %d results, want %d; results: %v",
					tc.line, len(got), tc.wantN, got)
				return
			}
			if tc.wantN > 0 && got[0].Text != tc.wantT {
				t.Errorf("Text = %q, want %q", got[0].Text, tc.wantT)
			}
		})
	}
}

// ── scanBareTextSuffixes — punctuation preservation (fix 5b regression) ───────

func TestScanBareTextSuffixes_PreservesPunctuation(t *testing.T) {
	// After fix 5b: Text should include trailing "." (stored as suffix, not cleaned).
	line := `<strong><%= t('admin.analytics.revenue.high_churn_rate') %></strong> Focus on customer retention.`
	got := scanBareTextSuffixes(line, line, "test.erb", 1)
	if len(got) == 0 {
		t.Fatal("expected at least 1 result, got 0")
	}
	if got[0].Text != "Focus on customer retention." {
		t.Errorf("Text = %q, want %q (trailing dot must be preserved)", got[0].Text, "Focus on customer retention.")
	}
}

func TestScanBareTextSuffixes_NoERBOutput(t *testing.T) {
	// Must have ERB output tag to trigger.
	got := scanBareTextSuffixes("Plain text here no erb", "Plain text here no erb", "test.erb", 1)
	if len(got) != 0 {
		t.Errorf("expected 0 results for line without ERB output, got %d", len(got))
	}
}

func TestScanBareTextSuffixes_MixedWithI18nCall(t *testing.T) {
	// Fix 5a: i18nCallRe no longer causes early return — suffix after t() closer still detected.
	line := `<%= t('label.key') %> Some additional text here`
	got := scanBareTextSuffixes(line, line, "test.erb", 1)
	if len(got) == 0 {
		t.Error("expected suffix text to be detected even when i18n call appears on same line")
	}
}

func TestScanBareTextSuffixes_NoPunctuationPreserved(t *testing.T) {
	// When no trailing punctuation exists, Text == suffix unchanged.
	line := `<strong><%= t('foo') %></strong> Keep testing right away`
	got := scanBareTextSuffixes(line, line, "test.erb", 1)
	if len(got) == 0 {
		t.Fatal("expected 1 result")
	}
	if got[0].Text != "Keep testing right away" {
		t.Errorf("Text = %q, want %q", got[0].Text, "Keep testing right away")
	}
}

// ── scanErbTernary — i18nCallRe early-return removed (fix 5a regression) ─────

func TestScanErbTernary_DetectsHardcodedBranchAlongsideI18n(t *testing.T) {
	// Before fix 5a: the whole line was skipped because t() appears in one branch.
	// After fix 5a: the hardcoded branch is still detected.
	line := `<%= user.active? ? t('users.status.active') : 'Account suspended' %>`
	got := scanErbTernary(line, line, "test.erb", 1)
	if len(got) == 0 {
		t.Fatal("expected hardcoded ternary branch to be detected even when other branch uses t()")
	}
	found := false
	for _, h := range got {
		if h.Text == "Account suspended" {
			found = true
		}
	}
	if !found {
		t.Errorf("expected 'Account suspended' in results; got: %v", got)
	}
}

func TestScanErbTernary_BothHardcoded(t *testing.T) {
	line := `<%= active ? 'User is active now' : 'User is not active' %>`
	got := scanErbTernary(line, line, "test.erb", 1)
	if len(got) < 2 {
		t.Errorf("expected 2 hardcoded branches, got %d: %v", len(got), got)
	}
}

func TestScanErbTernary_NoQuestionMark(t *testing.T) {
	got := scanErbTernary(`<%= t('foo.bar') %>`, `<%= t('foo.bar') %>`, "test.erb", 1)
	if len(got) != 0 {
		t.Errorf("expected 0 results for line without ?, got %d", len(got))
	}
}

func TestScanErbTernary_ShortStringIgnored(t *testing.T) {
	// "No" is only 2 chars — ternaryBranchRe requires {2,} inside quotes (min 3 chars total).
	line := `<%= x ? 'Yes' : 'No' %>`
	got := scanErbTernary(line, line, "test.erb", 1)
	// "Yes" is 3 chars → matches {2,}; "No" is 2 chars → does not match. looksUserFacing("Yes") → no space → false.
	// Expect 0 valid user-facing results.
	for _, h := range got {
		if h.Text == "No" {
			t.Errorf("'No' (2 chars) should not be detected")
		}
	}
}
