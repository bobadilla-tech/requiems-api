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
		{"has space passes", "Submit form now", true},
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
		// Stimulus controller#action descriptors must never be flagged as user-facing text.
		{"stimulus bare action", "modal#close", true},
		{"stimulus with event", "click->modal#close", true},
		{"stimulus with garbage suffix", `modal#close">`, true},
		{"stimulus dismiss", "modal#dismiss", true},
		{"stimulus multi-word controller", "dropdown-menu#toggle", true},
		{"normal word with hash", "C# programming", false}, // C# is not a Stimulus action
		// Single-word strings with 2+ hyphens are HTTP headers / technical tokens.
		{"http header two hyphens", "X-Backend-Secret", true},
		{"kebab three segments", "data-turbo-frame", true},
		{"one hyphen no space not caught", "two-factor", false},
		// Multi-word string: has space → new single-word rule does not apply.
		{"hyphenated phrase with space", "sign-in and register now", false},
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

// ── i18nCallRe — guards the early-return that prevents false positives ────────

func TestI18nCallRe_MatchesI18nT(t *testing.T) {
	// I18n.t() in model/service context must be recognised so the line is skipped.
	if !i18nCallRe.MatchString(`I18n.t("api_key.failed_to_generate")`) {
		t.Error("i18nCallRe must match I18n.t()")
	}
}

func TestI18nCallRe_MatchesBareT(t *testing.T) {
	if !i18nCallRe.MatchString(`t('admin.sidebar.analytics')`) {
		t.Error("i18nCallRe must match bare t()")
	}
}

func TestI18nCallRe_MatchesTranslate(t *testing.T) {
	if !i18nCallRe.MatchString(`translate('some.key')`) {
		t.Error("i18nCallRe must match translate()")
	}
}

func TestI18nCallRe_NoFalsePositiveOnStringLiteral(t *testing.T) {
	// A plain hardcoded string must NOT match i18nCallRe (or it would be silently skipped).
	if i18nCallRe.MatchString(`title_text = "My Title"`) {
		t.Error("i18nCallRe must not match variable assignments with string literals")
	}
}

// ── scanAttrPatterns ──────────────────────────────────────────────────────────

func TestScanAttrPatterns_Placeholder(t *testing.T) {
	line := `<input placeholder="Email address">`
	got := scanAttrPatterns(line, line, "test.erb", 1)
	if len(got) != 1 {
		t.Fatalf("expected 1 result for placeholder attr, got %d: %v", len(got), got)
	}
	if got[0].Text != "Email address" {
		t.Errorf("Text = %q, want %q", got[0].Text, "Email address")
	}
}

func TestScanAttrPatterns_AriaLabel(t *testing.T) {
	line := `<button aria-label="Close dialog">`
	got := scanAttrPatterns(line, line, "test.erb", 1)
	if len(got) != 1 {
		t.Fatalf("expected 1 result for aria-label, got %d: %v", len(got), got)
	}
	if got[0].Text != "Close dialog" {
		t.Errorf("Text = %q, want %q", got[0].Text, "Close dialog")
	}
}

func TestScanAttrPatterns_AlreadyTranslated(t *testing.T) {
	// After ERB tag stripping, the attribute value is empty — pattern requires 4+ chars.
	line := `<input placeholder="<%= t('key') %>">`
	got := scanAttrPatterns(line, line, "test.erb", 1)
	if len(got) != 0 {
		t.Errorf("expected 0 results for already-translated placeholder, got %d: %v", len(got), got)
	}
}

func TestScanAttrPatterns_DecorativeAlt(t *testing.T) {
	// Alt text describing an icon/logo is decorative — not user-facing UI copy.
	line := `<img alt="Company logo">`
	got := scanAttrPatterns(line, line, "test.erb", 1)
	if len(got) != 0 {
		t.Errorf("expected 0 results for decorative alt text, got %d: %v", len(got), got)
	}
}

// ── scanTagContent ────────────────────────────────────────────────────────────

func TestScanTagContent_BareText(t *testing.T) {
	line := `<p>Hello world today</p>`
	got := scanTagContent(line, line, "test.erb", 1)
	if len(got) != 1 {
		t.Fatalf("expected 1 result for bare tag content, got %d: %v", len(got), got)
	}
	if got[0].Text != "Hello world today" {
		t.Errorf("Text = %q, want %q", got[0].Text, "Hello world today")
	}
}

func TestScanTagContent_I18nCallSkipped(t *testing.T) {
	// Line has i18nCallRe match → scanTagContent returns nil immediately.
	line := `<p><%= t('key') %></p>`
	got := scanTagContent(line, line, "test.erb", 1)
	if len(got) != 0 {
		t.Errorf("expected 0 results when i18n call present, got %d: %v", len(got), got)
	}
}

func TestScanTagContent_ErbOutputSkipped(t *testing.T) {
	// Any line with <%= is skipped to avoid orphaned fragments.
	line := `<p>Welcome <%= @email %>!</p>`
	got := scanTagContent(line, line, "test.erb", 1)
	if len(got) != 0 {
		t.Errorf("expected 0 results for line with ERB output, got %d: %v", len(got), got)
	}
}

// ── Stimulus action regression ────────────────────────────────────────────────
// tagContentRe `>([A-Za-z\d][^<\n]{3,})<` can match the `>` inside a Stimulus
// action arrow `->` and capture `controller#action">` as "tag content".
// These tests lock down that no Stimulus descriptor leaks through as user-facing text.

func TestScanTagContent_StimulusActionNotDetected(t *testing.T) {
	// Exact line that caused the regression: data-action="click->modal#close"></div>
	// tagContentRe sees `->modal#close">` and captures `modal#close">` as content.
	// looksLikeCode must kill it before it becomes a finding.
	line := `  <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" data-action="click->modal#close"></div>`
	got := scanTagContent(line, line, "partials/shared/_modal.html.erb", 1)
	for _, h := range got {
		if strings.Contains(h.Text, "modal#close") || strings.Contains(h.Text, "#close") {
			t.Errorf("Stimulus action leaked as tag content: %q", h.Text)
		}
	}
}

func TestScanAttrPatterns_StimulusActionNotDetected(t *testing.T) {
	// data-action is not in attrPatterns, but guard: no attr scanner should emit Stimulus descriptors.
	line := `  <div data-action="click->modal#close" data-controller="modal"></div>`
	got := scanAttrPatterns(line, line, "test.erb", 1)
	for _, h := range got {
		if strings.Contains(h.Text, "modal#close") || strings.Contains(h.Text, "#close") {
			t.Errorf("Stimulus action leaked via attr pattern: %q", h.Text)
		}
	}
}

func TestLooksLikeCode_StimulusWithTrailingHTML(t *testing.T) {
	// The exact garbage string tagContentRe captured: `modal#close">`
	// Must be identified as code, not user-facing text.
	if !looksLikeCode(`modal#close">`) {
		t.Error(`looksLikeCode("modal#close\">") must return true — Stimulus action with trailing HTML`)
	}
}

// ── scanHelperKeywordArgs ─────────────────────────────────────────────────────

func TestScanHelperKeywordArgs_DisableWith(t *testing.T) {
	line := `<%= f.submit "Submit", disable_with: "Sending..." %>`
	got := scanHelperKeywordArgs(line, line, "test.erb", 1)
	if len(got) != 1 {
		t.Fatalf("expected 1 result for disable_with, got %d: %v", len(got), got)
	}
	if got[0].Text != "Sending..." {
		t.Errorf("Text = %q, want %q", got[0].Text, "Sending...")
	}
}

func TestScanHelperKeywordArgs_SubmitTag(t *testing.T) {
	line := `<%= submit_tag "Save Changes" %>`
	got := scanHelperKeywordArgs(line, line, "test.erb", 1)
	if len(got) != 1 {
		t.Fatalf("expected 1 result for submit_tag, got %d: %v", len(got), got)
	}
	if got[0].Text != "Save Changes" {
		t.Errorf("Text = %q, want %q", got[0].Text, "Save Changes")
	}
}

func TestScanHelperKeywordArgs_TextKeyword(t *testing.T) {
	line := `<%= link_to root_path, text: "Convert" %>`
	got := scanHelperKeywordArgs(line, line, "test.erb", 1)
	if len(got) != 1 {
		t.Fatalf("expected 1 result for text: keyword, got %d: %v", len(got), got)
	}
	if got[0].Text != "Convert" {
		t.Errorf("Text = %q, want %q", got[0].Text, "Convert")
	}
}

// ── Regression: HTTP header literals must not be extracted as hardcoded UI strings ──
//
// X-Backend-Secret (and similar kebab-identifier tokens) appeared in views as
// literal strings and were incorrectly flagged by the extractor because
// looksLikeCode did not handle single-word strings with ≥ 2 hyphens.
// The fix added that rule to looksLikeCode; these tests lock it in.

func TestRegression_HttpHeaderNotFlaggedByTableHeaders(t *testing.T) {
	line := `<th>X-Backend-Secret</th>`
	got := scanTableHeaders(line, line, "test.erb", 1)
	if len(got) != 0 {
		t.Errorf("X-Backend-Secret in <th> must not be extracted as hardcoded string; got %d findings: %v", len(got), got)
	}
}

func TestRegression_HttpHeaderNotFlaggedByTagContent(t *testing.T) {
	line := `<dt>X-Backend-Secret</dt>`
	got := scanTagContent(line, line, "test.erb", 1)
	if len(got) != 0 {
		t.Errorf("X-Backend-Secret in <dt> must not be extracted as hardcoded string; got %d findings: %v", len(got), got)
	}
}

func TestRegression_HttpHeaderNotFlaggedByHelperKeywordArgs(t *testing.T) {
	line := `<%= f.label :secret, text: "X-Backend-Secret" %>`
	got := scanHelperKeywordArgs(line, line, "test.erb", 1)
	if len(got) != 0 {
		t.Errorf("X-Backend-Secret in text: arg must not be extracted as hardcoded string; got %d findings: %v", len(got), got)
	}
}

func TestRegression_MetricCodeNotFlaggedByTagContent(t *testing.T) {
	// p50 / p99 are single lowercase tokens — looksUserFacingLabel requires a
	// letter start and min length 2, but looksLikeCode does not catch them.
	// They are filtered upstream by looksUserFacingLabel's length and start-char rules.
	// This test confirms they do not produce false-positive findings.
	for _, token := range []string{"p50", "p99"} {
		line := `<span>` + token + `</span>`
		got := scanTagContent(line, line, "test.erb", 1)
		if len(got) != 0 {
			t.Errorf("metric code %q in <span> must not be extracted; got %d findings: %v", token, len(got), got)
		}
	}
}

func TestScanHelperKeywordArgs_AlreadyTranslated(t *testing.T) {
	// i18nCallRe matches t() → scanHelperKeywordArgs returns nil.
	line := `<%= f.submit t('shared.submit'), disable_with: t('shared.sending') %>`
	got := scanHelperKeywordArgs(line, line, "test.erb", 1)
	if len(got) != 0 {
		t.Errorf("expected 0 results when t() used for values, got %d: %v", len(got), got)
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
