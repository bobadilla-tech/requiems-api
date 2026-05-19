package thesaurus

import (
	"fmt"
	"strings"
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

	e, ok := thesaurusData[normalized]
	if !ok {
		return Result{}, fmt.Errorf("word not found: %s", normalized)
	}

	synonyms := e.synonyms
	if synonyms == nil {
		synonyms = []string{}
	}

	antonyms := e.antonyms
	if antonyms == nil {
		antonyms = []string{}
	}

	return Result{
		Word:     normalized,
		Synonyms: synonyms,
		Antonyms: antonyms,
	}, nil
}
