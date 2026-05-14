package quotes

import "requiems-api/platform/httpx"

// BatchRandomRequest is the request body for fetching multiple random quotes at once.
type BatchRandomRequest struct {
	Count int `json:"count" validate:"required,min=1,max=50"`
}

// BatchRandomResponse is the canonical batch response for the quotes service.
// It is a type alias for httpx.BatchResponse[Quote]; Total and X-Usage-Count
// are populated automatically by httpx.HandleBatch from len(Results).
type BatchRandomResponse = httpx.BatchResponse[Quote]
