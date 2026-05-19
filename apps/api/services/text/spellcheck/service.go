package spellcheck

import (
	"bufio"
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/sajari/fuzzy"
)

//go:embed data/words.txt
var wordListData []byte

// wordRe matches sequences of ASCII letters (words).
var wordRe = regexp.MustCompile(`[a-zA-Z]+`)

// Correction describes a single spelling mistake and its suggested fix.
type Correction struct {
	Original  string `json:"original"`
	Suggested string `json:"suggested"`
	Position  int    `json:"position"`
}

// Result is the response payload for the spell check endpoint.
type Result struct {
	Corrected   string       `json:"corrected"`
	Corrections []Correction `json:"corrections"`
}


// Service performs spell checking and correction.
type Service struct {
	model   *fuzzy.Model
	initErr error
}

// NewService returns a new spellcheck Service. If the embedded word list
// cannot be loaded the service is still valid but Check will return an error.
func NewService() *Service {
	m := fuzzy.NewModel()
	m.SetThreshold(1)

	var words []string
	scanner := bufio.NewScanner(bytes.NewReader(wordListData))
	for scanner.Scan() {
		if w := strings.TrimSpace(scanner.Text()); w != "" {
			words = append(words, w)
		}
	}
	if err := scanner.Err(); err != nil {
		return &Service{initErr: fmt.Errorf("spellcheck: word list unavailable: %w", err)}
	}
	m.Train(words)
	return &Service{model: m}
}

// errUnavailable is returned by Check when the service failed to initialise.
var errUnavailable = errors.New("spellcheck: service unavailable")

// Check inspects text for spelling mistakes and returns the corrected text
// together with a list of individual corrections. It returns an error if the
// service failed to initialise.
func (s *Service) Check(text string) (Result, error) {
	if s.initErr != nil {
		return Result{}, errUnavailable
	}

	var corrections []Correction

	// Find all word positions so that we can replace in one pass.
	locs := wordRe.FindAllStringIndex(text, -1)

	// Build the corrected string by walking through every word location.
	var builder strings.Builder
	prev := 0
	for _, loc := range locs {
		start, end := loc[0], loc[1]
		word := text[start:end]
		lower := strings.ToLower(word)

		suggestion := bestSuggestion(lower, s.model)

		// bestSuggestion returns the input unchanged when the word is already
		// in the model's vocabulary.
		if suggestion != lower && suggestion != "" {
			display := matchCase(word, suggestion)

			corrections = append(corrections, Correction{
				Original:  word,
				Suggested: display,
				Position:  len([]rune(text[:start])),
			})

			// Copy verbatim text before this word, then the correction.
			builder.WriteString(text[prev:start])
			builder.WriteString(display)
		} else {
			builder.WriteString(text[prev:end])
		}
		prev = end
	}
	// Append any trailing non-word characters.
	builder.WriteString(text[prev:])

	if corrections == nil {
		corrections = []Correction{}
	}

	return Result{
		Corrected:   builder.String(),
		Corrections: corrections,
	}, nil
}

// bestSuggestion picks the most likely correction from the model's exhaustive
// potentials by preferring (in order): lowest Levenshtein distance, then same
// first letter as the input, then highest corpus score.
func bestSuggestion(input string, m *fuzzy.Model) string {
	potentials := m.Potentials(input, true)
	if len(potentials) == 0 {
		return input
	}

	// Return the word unchanged when it is already in the vocabulary.
	if p, ok := potentials[input]; ok && p.Leven == 0 {
		return input
	}

	type candidate struct {
		term  string
		lev   int
		score int
	}

	var best candidate
	first := true
	for _, p := range potentials {
		if p.Leven == 0 {
			// Exact match – the word is correct, no suggestion needed.
			return input
		}
		bonus := 0
		if p.Term != "" && input != "" && p.Term[0] == input[0] {
			bonus = 100
		}
		effectiveScore := p.Score + bonus
		if first || p.Leven < best.lev || (p.Leven == best.lev && effectiveScore > best.score) {
			best = candidate{term: p.Term, lev: p.Leven, score: effectiveScore}
			first = false
		}
	}

	return best.term
}

// matchCase applies the capitalisation pattern of src to dst so that, for
// example, "Ths" → "This" and "THs" → "THIS".
func matchCase(src, dst string) string {
	if src == "" || dst == "" {
		return dst
	}

	srcRunes := []rune(src)
	dstRunes := []rune(dst)

	// All-uppercase source → uppercase suggestion.
	allUpper := true
	for _, r := range srcRunes {
		if unicode.IsLetter(r) && !unicode.IsUpper(r) {
			allUpper = false
			break
		}
	}
	if allUpper {
		return strings.ToUpper(dst)
	}

	// First letter capitalised → capitalise first letter of suggestion.
	if unicode.IsUpper(srcRunes[0]) {
		dstRunes[0] = unicode.ToUpper(dstRunes[0])
		return string(dstRunes)
	}

	return dst
}
