package contentmoderate

import (
	"context"

	"requiems-api/services/text/detectlanguage"
	"requiems-api/services/text/sentiment"
	"requiems-api/services/validation/profanity"
)

// Request represents the input for the moderation endpoint.
type Request struct {
	Text     string `json:"text" validate:"required"`
	Language string `json:"language"` // optional — if empty, language is detected automatically.
}

// Categories represents the different types of harmful content detected in the text.
type Categories struct {
	Profanity  bool `json:"profanity"`
	HateSpeech bool `json:"hate_speech"`
	Spam       bool `json:"spam"`
	Violence   bool `json:"violence"`
}

// Response is the response from the moderation endpoint.
type Response struct {
	IsSafe             bool       `json:"is_safe"`
	ToxicityScore      *float64   `json:"toxicity_score"`
	Sentiment          string     `json:"sentiment"`
	Language           *string    `json:"language"`
	LanguageConfidence *float64   `json:"language_confidence"`
	Flags              []string   `json:"flags"`
	Categories         Categories `json:"categories"`
}

// DetectLanguage is the interface for language detection.
// Implemented by detectlanguage.Service in production and a stub in tests.
type DetectLanguage interface {
	Detect(text string) detectlanguage.Result
}

// Profanity is the interface for profanity checking.
// Implemented by profanity.Service in production and a stub in tests.
type Profanity interface {
	Check(ctx context.Context, text string) profanity.Result
}

// Sentiment is the interface for sentiment analysis.
// Implemented by sentiment.Service in production and a stub in tests.
type Sentiment interface {
	Analyze(text string) sentiment.Result
}

// Service holds the dependencies required to perform content moderation.
type Service struct {
	detectlanguage DetectLanguage
	profanity      Profanity
	sentiment      Sentiment
}

// NewService returns a new Service.
func NewService(
	dl DetectLanguage,
	prof Profanity,
	sent Sentiment,
) *Service {
	return &Service{
		detectlanguage: dl,
		profanity:      prof,
		sentiment:      sent,
	}
}

// Moderate runs profanity check, sentiment analysis, and language detection if needed
// and returns a unified moderation result.
func (s *Service) Moderate(ctx context.Context, text, language string) Response {

	// resolve language — use caller-provided value or detect automatically.
	// if caller provides language, confidence is set to 1.0 (certain).
	// if detectlanguage returns confidence 0.0, language and confidence stay null.
	var lang *string
	var langConfidence *float64

	if language != "" {
		confidence := 1.0
		lang = &language
		langConfidence = &confidence
	} else {
		result := s.detectlanguage.Detect(text)
		if result.Confidence != 0.0 {
			lang = &result.Code
			langConfidence = &result.Confidence
		}
	}

	isSafe := true
	// toxicityScore is null until there is a way to calculate it.
	var toxicityScore *float64
	flags := []string{}
	var categories Categories

	// run profanity check — if profane content is found, mark category,
	// add flag, and set isSafe to false
	profanityResult := s.profanity.Check(ctx, text)

	if profanityResult.HasProfanity {
		isSafe = false
		flags = append(flags, "text_profanity")
		categories.Profanity = true
	}

	// run sentiment analysis — returns positive, negative, or neutral.
	sentimentResult := s.sentiment.Analyze(text)

	return Response{
		IsSafe:             isSafe,
		ToxicityScore:      toxicityScore,
		Sentiment:          sentimentResult.Sentiment,
		Language:           lang,
		LanguageConfidence: langConfidence,
		Flags:              flags,
		Categories:         categories,
	}
}
