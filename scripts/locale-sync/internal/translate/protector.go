package translate

import (
	"fmt"
	"regexp"
)

// placeholderRe matches Rails i18n interpolation patterns and printf-style verbs
// that must survive translation unchanged.
//
// Matches (in order):
//   - %{name}       — Rails named interpolation
//   - {{name}}      — Handlebars/Liquid-style
//   - %<name>s      — printf named verb
//   - %08.2f etc.   — printf positional verbs (%s %d %i %f %g with optional flags)
var placeholderRe = regexp.MustCompile(`%\{[^}]+\}|\{\{[^}]+\}\}|%<[^>]+>[sdifg]|%[-+0 #]*\d*(?:\.\d+)?[sdifgv]`)

// Protected holds a string with its placeholders replaced by opaque tokens,
// plus the reverse mapping to restore them after translation.
type Protected struct {
	Text   string            // text sent to translation API
	tokens map[string]string // token → original placeholder
}

// Protect replaces all placeholder patterns in s with tokens like XTOKEN0X.
func Protect(s string) Protected {
	tokens := make(map[string]string)
	i := 0
	text := placeholderRe.ReplaceAllStringFunc(s, func(match string) string {
		tok := fmt.Sprintf("XTOKEN%dX", i)
		i++
		tokens[tok] = match
		return tok
	})
	return Protected{Text: text, tokens: tokens}
}

// Restore replaces tokens in translated back with their original placeholders.
func (p Protected) Restore(translated string) string {
	result := translated
	for tok, orig := range p.tokens {
		// Simple string replacement; tokens are unique per Protected instance.
		result = replaceAll(result, tok, orig)
	}
	return result
}

// replaceAll replaces all occurrences of old with new in s.
func replaceAll(s, old, new string) string {
	if old == "" {
		return s
	}
	out := make([]byte, 0, len(s))
	olen := len(old)
	for i := 0; i < len(s); {
		if i+olen <= len(s) && s[i:i+olen] == old {
			out = append(out, new...)
			i += olen
		} else {
			out = append(out, s[i])
			i++
		}
	}
	return string(out)
}
