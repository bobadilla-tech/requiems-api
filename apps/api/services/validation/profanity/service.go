package profanity

import (
	"context"
	"strings"

	goaway "github.com/TwiN/go-away"
)

// detector is the default profanity detector shared across all requests.
var detector = goaway.NewProfanityDetector()

// Result is the response payload for the profanity check endpoint.
type Result struct {
	HasProfanity bool     `json:"has_profanity"`
	Censored     string   `json:"censored"`
	FlaggedWords []string `json:"flagged_words"`
}


// Service performs profanity detection and censoring.
type Service struct{}

// NewService returns a new profanity Service.
func NewService() *Service { return &Service{} }

// Check inspects text for profanity, returning a censored copy of the text
// and the deduplicated list of flagged words found.
func (s *Service) Check(ctx context.Context, text string) Result {
	censored := detector.Censor(text)
	hasProfanity := detector.IsProfane(text)

	flaggedSet := map[string]bool{}
	flaggedWords := make([]string, 0)

	if hasProfanity {
		remaining := text
		for {
			word := goaway.ExtractProfanity(remaining)
			if word == "" {
				break
			}
			if !flaggedSet[word] {
				flaggedSet[word] = true
				flaggedWords = append(flaggedWords, word)
			}
			// Advance past the first occurrence of the canonical word so
			// subsequent iterations can find additional distinct words.
			// go-away sanitises the text before matching, so the word may
			// appear as a substring (e.g. "shit" inside "bullshit").
			idx := strings.Index(strings.ToLower(remaining), word)
			if idx == -1 {
				// Leet-speak or special-char obfuscation: the canonical word
				// doesn't appear literally — stop to avoid an infinite loop.
				break
			}
			remaining = remaining[idx+len(word):]
		}
	}

	return Result{
		HasProfanity: hasProfanity,
		Censored:     censored,
		FlaggedWords: flaggedWords,
	}
}
