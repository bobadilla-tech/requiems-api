package translate

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
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

// Protector protects Rails interpolation placeholders and optionally a custom
// dictionary of words/phrases from being altered during translation.
//
// Create one Protector per command invocation and reuse it across all strings.
type Protector struct {
	// customWords is sorted longest-first so longer phrases take precedence over
	// shorter sub-phrases when both appear in the same string.
	customWords []string
}

// NewProtector returns a Protector that will shield the given words from
// translation in addition to the standard placeholder patterns.
// Words are matched case-sensitively and protected as whole tokens.
func NewProtector(customWords []string) *Protector {
	words := make([]string, len(customWords))
	copy(words, customWords)
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j]) // longest first
	})
	return &Protector{customWords: words}
}

// Protected holds a string with its placeholders and protected words replaced
// by opaque tokens, plus the reverse mapping to restore them after translation.
type Protected struct {
	Text   string            // text sent to translation API
	tokens map[string]string // token → original text
}

// Protect replaces all placeholder patterns and custom dictionary words in s
// with tokens like XTOKEN0X that the API will leave untouched.
func (p *Protector) Protect(s string) Protected {
	tokens := make(map[string]string)
	i := 0

	nextToken := func(original string) string {
		tok := fmt.Sprintf("XTOKEN%dX", i)
		i++
		tokens[tok] = original
		return tok
	}

	text := s

	// 1. Protect custom dictionary words first using a single-pass regex alternation.
	//    A sequential strings.ReplaceAll loop would corrupt already-inserted tokens
	//    when a short custom word (e.g. "X") matches characters inside "XTOKEN0X".
	//    Words are already sorted longest-first by NewProtector.
	if len(p.customWords) > 0 {
		parts := make([]string, 0, len(p.customWords))
		for _, w := range p.customWords {
			if w != "" {
				parts = append(parts, regexp.QuoteMeta(w))
			}
		}
		if len(parts) > 0 {
			customRe := regexp.MustCompile(strings.Join(parts, "|"))
			text = customRe.ReplaceAllStringFunc(text, func(match string) string {
				return nextToken(match)
			})
		}
	}

	// 2. Protect standard interpolation placeholders.
	text = placeholderRe.ReplaceAllStringFunc(text, func(match string) string {
		return nextToken(match)
	})

	return Protected{Text: text, tokens: tokens}
}

// Restore replaces all tokens in translated back with their original strings.
func (pr Protected) Restore(translated string) string {
	result := translated
	for tok, orig := range pr.tokens {
		result = strings.ReplaceAll(result, tok, orig)
	}
	return result
}

// LoadDictionaryFile reads a protected-words file: one word/phrase per line,
// blank lines and lines starting with # are ignored.
func LoadDictionaryFile(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var words []string
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		words = append(words, line)
	}
	return words, nil
}
