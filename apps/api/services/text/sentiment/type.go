package sentiment

// Request is the input for the sentiment analysis endpoint.
type Request struct {
	Text string `json:"text" validate:"required"`
}

// Breakdown contains the proportional score for each sentiment class.
// The three values always sum to 1.0.
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

func (Result) IsData() {}

// BatchAnalyzeRequest is the request body for analyzing multiple texts at once.
type BatchAnalyzeRequest struct {
	Texts []string `json:"texts" validate:"required,min=1,max=50,dive,required"`
}
