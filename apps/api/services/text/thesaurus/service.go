package thesaurus

import (
	"fmt"
	"strings"

	"github.com/bobadilla-tech/thesaurus-go/pkg/thesaurus"
)

// Result is the response payload for the thesaurus endpoint.
type Result struct {
	Word     string   `json:"word"`
	Synonyms []string `json:"synonyms"`
	Antonyms []string `json:"antonyms"`
}

// Service looks up synonyms and antonyms for a given word.
type Service struct{}

// NewService returns a new thesaurus Service.
func NewService() *Service { return &Service{} }

// Lookup returns synonyms and antonyms for the given word.
// Returns an error if the word is not found in the dataset.
func (s *Service) Lookup(word string) (Result, error) {
	normalized := strings.ToLower(strings.TrimSpace(word))

	e, ok := thesaurus.Lookup(word)
	if !ok {
		return Result{}, fmt.Errorf("word not found: %s", normalized)
	}

	synonyms := e.Synonyms
	if synonyms == nil {
		synonyms = []string{}
	}

	antonyms := e.Antonyms
	if antonyms == nil {
		antonyms = []string{}
	}

	return Result{
		Word:     normalized,
		Synonyms: synonyms,
		Antonyms: antonyms,
	}, nil
}

// BatchThesaurusItem is the result for a single item in a batch thesaurus request.
type BatchThesaurusItem struct {
	Word   string  `json:"word"`
	Result *Result `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

// LookupBatch returns synonyms and antonyms for each word in the slice.
// Words not found in the dataset return an in-band error.
func (s *Service) LookupBatch(words []string) []BatchThesaurusItem {
	results := make([]BatchThesaurusItem, len(words))
	for i, word := range words {
		r, err := s.Lookup(word)
		if err != nil {
			results[i] = BatchThesaurusItem{Word: word, Error: "word not found"}
		} else {
			results[i] = BatchThesaurusItem{Word: word, Result: &r}
		}
	}
	return results
}
