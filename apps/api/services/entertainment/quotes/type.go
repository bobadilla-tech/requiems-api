package quotes

// BatchRandomRequest is the request body for fetching multiple random quotes at once.
type BatchRandomRequest struct {
	Count int `json:"count" validate:"required,min=1,max=50"`
}

// BatchRandomResponse is the response for a batch random quotes request.
type BatchRandomResponse struct {
	Results []Quote `json:"results"`
	Total   int     `json:"total"`
}

func (BatchRandomResponse) IsData() {}
