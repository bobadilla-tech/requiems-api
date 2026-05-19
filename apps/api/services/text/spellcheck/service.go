package spellcheck

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const defaultLanguageToolURL = "http://localhost:8010"

// ltResponse is the subset of the LanguageTool /v2/check response we need.
type ltResponse struct {
	Matches []ltMatch `json:"matches"`
}

type ltMatch struct {
	Offset       int             `json:"offset"`
	Length       int             `json:"length"`
	Replacements []ltReplacement `json:"replacements"`
}

type ltReplacement struct {
	Value string `json:"value"`
}

// Service calls LanguageTool for spell checking.
type Service struct {
	client  *http.Client
	baseURL string
}

// NewService returns a new spellcheck Service backed by LanguageTool.
// baseURL defaults to http://localhost:8010 if empty.
func NewService(baseURL string) *Service {
	if baseURL == "" {
		baseURL = defaultLanguageToolURL
	}
	return &Service{
		client:  &http.Client{Timeout: 10 * time.Second},
		baseURL: strings.TrimRight(baseURL, "/"),
	}
}

// Check inspects text for spelling mistakes using LanguageTool and returns
// the corrected text together with a list of individual corrections.
func (s *Service) Check(text string) (Result, error) {
	form := url.Values{}
	form.Set("text", text)
	form.Set("language", "en-US")

	resp, err := s.client.Post(
		s.baseURL+"/v2/check",
		"application/x-www-form-urlencoded",
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		return Result{}, fmt.Errorf("spellcheck: languagetool unreachable: %w", err)
	}
	defer resp.Body.Close()

	var ltResp ltResponse
	if err := json.NewDecoder(resp.Body).Decode(&ltResp); err != nil {
		return Result{}, fmt.Errorf("spellcheck: failed to decode response: %w", err)
	}

	return buildResult(text, ltResp.Matches), nil
}

// CheckBatch processes multiple texts and returns spell-check results for each.
// It returns an error only if LanguageTool is unreachable; individual texts do
// not produce per-item errors.
func (s *Service) CheckBatch(texts []string) ([]Result, error) {
	results := make([]Result, 0, len(texts))
	for _, text := range texts {
		result, err := s.Check(text)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

// buildResult converts LanguageTool matches into our Result type.
// Replacements are applied in reverse order to keep offsets valid.
func buildResult(text string, matches []ltMatch) Result {
	runes := []rune(text)
	corrections := make([]Correction, 0, len(matches))

	for _, m := range matches {
		if len(m.Replacements) == 0 || m.Offset+m.Length > len(runes) {
			continue
		}
		// Collect up to 3 suggestions so callers can present a picker.
		suggestions := make([]string, 0, 3)
		for i, r := range m.Replacements {
			if i >= 3 {
				break
			}
			suggestions = append(suggestions, r.Value)
		}
		corrections = append(corrections, Correction{
			Original:    string(runes[m.Offset : m.Offset+m.Length]),
			Suggested:   m.Replacements[0].Value,
			Suggestions: suggestions,
			Position:    m.Offset,
		})
	}

	// Apply replacements in reverse order to preserve offsets.
	corrected := make([]rune, len(runes))
	copy(corrected, runes)
	for i := len(matches) - 1; i >= 0; i-- {
		m := matches[i]
		if len(m.Replacements) == 0 || m.Offset+m.Length > len(corrected) {
			continue
		}
		repl := []rune(m.Replacements[0].Value)
		corrected = append(corrected[:m.Offset], append(repl, corrected[m.Offset+m.Length:]...)...)
	}

	if corrections == nil {
		corrections = []Correction{}
	}

	return Result{
		Corrected:   string(corrected),
		Corrections: corrections,
	}
}
