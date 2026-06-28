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

// LintValues inspects locale entry values and returns findings for values that
// look like technical identifiers rather than translatable UI copy:
//
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
		if reason := technicalValueReason(v); reason != "" {
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

func technicalValueReason(v string) string {
	if strings.Contains(v, "@") {
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
