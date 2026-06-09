package detectlanguage

import (
	"strings"

	lingua "github.com/pemistahl/lingua-go"
)

// detector is the default language detector shared across all requests.
var detector = lingua.NewLanguageDetectorBuilder().
	FromAllLanguages().
	Build()

// Result is the response payload for the detect-language endpoint.
type Result struct {
	Language   string  `json:"language"`
	Code       string  `json:"code"`
	Confidence float64 `json:"confidence"`
}

// Service performs language detection.
type Service struct{}

// NewService returns a new language detection Service.
func NewService() *Service { return &Service{} }

// Detect identifies the language of the given text.
func (s *Service) Detect(text string) Result {
	language, reliable := detector.DetectLanguageOf(text)
	if !reliable {
		return Result{
			Language:   "Unknown",
			Code:       "",
			Confidence: 0,
		}
	}

	confidence := detector.ComputeLanguageConfidence(text, language)

	return Result{
		Language:   language.String(),
		Code:       strings.ToLower(language.IsoCode639_1().String()),
		Confidence: confidence,
	}
}

// BatchDetectItem is the result for a single item in a batch language detection request.
type BatchDetectItem struct {
	Text   string  `json:"text"`
	Result *Result `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

// DetectBatch detects the language for each text in the slice.
// Empty texts return an in-band error; all other items are processed.
func (s *Service) DetectBatch(texts []string) []BatchDetectItem {
	results := make([]BatchDetectItem, len(texts))
	for i, text := range texts {
		if strings.TrimSpace(text) == "" {
			results[i] = BatchDetectItem{Text: text, Error: "text is required"}
			continue
		}
		r := s.Detect(text)
		results[i] = BatchDetectItem{Text: text, Result: &r}
	}
	return results
}
