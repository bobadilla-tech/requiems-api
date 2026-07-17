package sentiment

import (
	afinn "github.com/bobadilla-tech/sentiment-go"
)

// Breakdown contains the proportional score for each sentiment class.
type Breakdown struct {
	Positive float64 `json:"positive"`
	Negative float64 `json:"negative"`
	Neutral  float64 `json:"neutral"`
}

// Result is the response payload for the sentiment endpoint.
type Result struct {
	Sentiment string    `json:"sentiment"`
	Score     float64   `json:"score"`
	Breakdown Breakdown `json:"breakdown"`
}

// Service performs sentiment analysis.
type Service struct {
	sentimentSvc *afinn.Service
}

// NewService returns a new sentiment Service.
func NewService() (*Service, error) {
	svc, err := afinn.NewService()
	if err != nil {
		return nil, err
	}
	return &Service{sentimentSvc: svc}, nil
}

// Analyze scores the given text using the underlying sentiment engine.
func (s *Service) Analyze(text string) Result {
	r := s.sentimentSvc.Analyze(text)
	return Result{
		Sentiment: r.Sentiment,
		Score:     r.Score,
		Breakdown: Breakdown(r.Breakdown),
	}
}

// AnalyzeBatch analyzes a slice of texts and returns results in the same order.
func (s *Service) AnalyzeBatch(texts []string) []Result {
	results := make([]Result, len(texts))
	for i, text := range texts {
		results[i] = s.Analyze(text)
	}
	return results
}
