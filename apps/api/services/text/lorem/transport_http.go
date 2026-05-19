package lorem

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

type Generator interface {
	Generate(paragraphs, sentences int) Lorem
}

// Request holds the optional query parameters for the lorem endpoint.
type Request struct {
	Paragraphs int `query:"paragraphs" validate:"min=1,max=20"`
	Sentences  int `query:"sentences"  validate:"min=1,max=20"`
}

// BatchItemRequest defines the JSON parameters for a single item within the batch.
type BatchItemRequest struct {
	Paragraphs int `json:"paragraphs" validate:"omitempty,min=1,max=20"`
	Sentences  int `json:"sentences"  validate:"omitempty,min=1,max=20"`
}

// BatchRequest represents the incoming JSON body.
// The dive validation ensures that internal items are also validated.
type BatchRequest struct {
	Items []BatchItemRequest `json:"items" validate:"required,min=1,max=50,dive"`
}

// BatchItemResponse represents the response for a single item.
// It uses the "partial success" structure requested by the RFC.
type BatchItemResponse struct {
	Data  *Lorem  `json:"data,omitempty"`
	Error *string `json:"error,omitempty"` // Used if this specific item fails
}

func RegisterRoutes(r chi.Router, svc Generator) {
	// Original GET endpoint (unmodified)
	r.Get("/lorem", func(w http.ResponseWriter, r *http.Request) {
		// Set defaults before binding so unset params keep their default value.
		req := Request{Paragraphs: 1, Sentences: 5}

		if err := httpx.BindQuery(r, &req); err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		httpx.JSON(w, http.StatusOK, svc.Generate(req.Paragraphs, req.Sentences))
	})

	// New POST endpoint to process batches using the correct generic signature
	r.Post("/lorem/batch", httpx.HandleBatch(func(ctx context.Context, req BatchRequest) (httpx.BatchResponse[BatchItemResponse], error) {
		var results []BatchItemResponse

		for _, item := range req.Items {
			// Set defaults in case the client sends 0 or omits the field
			p := item.Paragraphs
			if p == 0 {
				p = 1
			}
			s := item.Sentences
			if s == 0 {
				s = 5
			}

			// Generate the text.
			lorem := svc.Generate(p, s)
			
			results = append(results, BatchItemResponse{
				Data: &lorem,
			})
		}

		// Return the batch response struct wrapped by httpx
		return httpx.BatchResponse[BatchItemResponse]{
			Results: results,
		}, nil
	}))
}