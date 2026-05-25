package facts

// Request holds optional query parameters for the facts endpoint.
type Request struct {
	Category string `query:"category"`
}

// Fact is the response payload for a single random fact.
type Fact struct {
	Fact     string `json:"fact"`
	Category string `json:"category"`
	Source   string `json:"source"`
}

func (Fact) IsData() {}

// BatchRequest holds the request payload for the batch facts endpoint.
type BatchRequest struct {
	Categories []string `json:"categories" validate:"required,min=1,max=50,dive"`
}

// BatchItem is a single result in a batch response.
// If a category is invalid or has no facts, Error is set and Fact is zero.
type BatchItem struct {
	Category string `json:"category"`
	Fact     string `json:"fact,omitempty"`
	Source   string `json:"source,omitempty"`
	Error    string `json:"error,omitempty"`
}