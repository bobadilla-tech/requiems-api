package locale

import (
	"regexp"
	"strings"
)

// LintFinding is a locale value that looks technically incorrect or untranslatable.
type LintFinding struct {
	Key    string
	Value  string
	File   string
	Line   int
	Reason string
}

// metricCodeRe matches short metric/version codes like p50, p99, s3, v2.
var metricCodeRe = regexp.MustCompile(`^[a-z]\d+$`)

// htmlWithClassRe matches HTML tags that carry CSS or style attributes —
// presentation logic embedded in translation strings.
var htmlWithClassRe = regexp.MustCompile(`<[a-zA-Z][^>]* (?:class|style)=`)

// knownHtmlTagRe matches real HTML element names. Excludes angle-bracket
// placeholders like <key> or <url> that appear in code examples.
var knownHtmlTagRe = regexp.MustCompile(`</?\b(?:strong|em|a|code|span|b|i|br|s|u|mark|p|div|h[1-6]|ul|ol|li)\b`)

// LintValues inspects locale entry values and returns findings for values that
// look like technical identifiers or contain misplaced HTML:
//
//   - _html keys whose values embed CSS/style attributes (Tailwind in translations)
//   - Known HTML tags in non-_html keys (won't be auto-marked safe by Rails)
//   - Single-word strings with ≥ 2 hyphens (HTTP headers, kebab identifiers)
//   - Strings containing "@" (email addresses)
//   - Short metric/version codes matching [a-z]\d+ (p50, p99, s3)
func LintValues(entries []Entry) []LintFinding {
	var out []LintFinding
	for _, e := range entries {
		v := e.Value
		if v == "" || v == "[array]" {
			continue
		}
		if reason := entryReason(e.KeyName, v); reason != "" {
			out = append(out, LintFinding{
				Key:    e.Key,
				Value:  v,
				File:   e.ShortPath,
				Line:   e.Line,
				Reason: reason,
			})
		}
	}
	return out
}

func entryReason(keyName, v string) string {
	// HTML checks first — these indicate structural/presentation problems.
	if strings.HasSuffix(keyName, "_html") && htmlWithClassRe.MatchString(v) {
		return "HTML with CSS/style classes in _html key — move styling to view, use %{variable} interpolation for formatting"
	}
	if !strings.HasSuffix(keyName, "_html") && knownHtmlTagRe.MatchString(v) {
		return "HTML tag in non-_html key — Rails won't html_safe this automatically; move markup to view"
	}

	// Technical identifier checks.
	// Skip multi-line values (JSON/code blocks) and values with multiple @ signs
	// (illustrative format examples); only flag a single embedded contact email.
	if strings.Count(v, "@") == 1 && !strings.Contains(v, "\n") {
		return "email address — hardcode the address in source, not in a locale file"
	}
	if !strings.Contains(v, " ") && strings.Count(v, "-") >= 2 {
		return "HTTP header or kebab identifier — technical token, does not need translation"
	}
	if metricCodeRe.MatchString(v) {
		return "metric/version code (e.g. p50, p99) — technical abbreviation, does not need translation"
	}
	return ""
}
