package similarity

import (
	"math"
	"strings"
	"unicode"
)

// Result is the response payload for the text similarity endpoint.
type Result struct {
	Similarity float64 `json:"similarity"`
	Method     string  `json:"method"`
}


// Service computes text similarity.
type Service struct{}

// NewService returns a new similarity Service.
func NewService() *Service { return &Service{} }

// Cosine compares two texts using cosine similarity on word-frequency vectors.
// The returned score is in the range [0, 1] where 1 means identical word
// distributions and 0 means no words in common.
func (s *Service) Cosine(text1, text2 string) Result {
	freq1 := wordFrequency(text1)
	freq2 := wordFrequency(text2)

	dot := 0.0
	for word, c1 := range freq1 {
		if c2, ok := freq2[word]; ok {
			dot += float64(c1) * float64(c2)
		}
	}

	mag1 := magnitude(freq1)
	mag2 := magnitude(freq2)

	var score float64
	if mag1 > 0 && mag2 > 0 {
		score = dot / (mag1 * mag2)
	}

	// Round to 4 decimal places to avoid floating-point noise.
	score = math.Round(score*1e4) / 1e4

	return Result{Similarity: score, Method: "cosine"}
}

// wordFrequency tokenises text and returns a map of lowercase word → count.
// Only alphanumeric characters are kept; punctuation and whitespace are ignored.
func wordFrequency(text string) map[string]int {
	freq := map[string]int{}
	for _, word := range tokenise(text) {
		freq[word]++
	}
	return freq
}

// tokenise splits text into lowercase alphanumeric tokens.
func tokenise(text string) []string {
	var tokens []string
	var buf strings.Builder

	for _, r := range strings.ToLower(text) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			buf.WriteRune(r)
		} else if buf.Len() > 0 {
			tokens = append(tokens, buf.String())
			buf.Reset()
		}
	}
	if buf.Len() > 0 {
		tokens = append(tokens, buf.String())
	}
	return tokens
}

// magnitude returns the Euclidean norm of a word-frequency vector.
func magnitude(freq map[string]int) float64 {
	sum := 0.0
	for _, c := range freq {
		sum += float64(c) * float64(c)
	}
	return math.Sqrt(sum)
}
