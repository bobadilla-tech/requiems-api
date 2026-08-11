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

// Request defines the parameters for generating lorem ipsum text.
// It is used both for the GET query parameters and the POST batch JSON items.
type Request struct {
	Paragraphs int `query:"paragraphs" json:"paragraphs" validate:"omitempty,min=1,max=20"`
	Sentences  int `query:"sentences" json:"sentences" validate:"omitempty,min=1,max=20"`
}

// BatchRequest represents the incoming JSON body.
type BatchRequest struct {
	Items []Request `json:"items" validate:"required,min=1,max=50,dive"`
}

// BatchItemResponse represents the response for a single item.
type BatchItemResponse struct {
	Data  *Lorem  `json:"data,omitempty"`
	Error *string `json:"error,omitempty"`
}

func RegisterRoutes(r chi.Router, svc Generator) {
	r.Get("/lorem", handleLoremIpsum(svc))
	r.Post("/lorem/batch", handleLoremIpsumBatch(svc))
}

// handleLoremIpsum godoc
//
//	@Summary		Generate Lorem Ipsum
//	@Description	Generate Lorem Ipsum placeholder text with customizable length and format.
//	@Tags			lorem-ipsum
//	@Produce		json
//	@Param			paragraphs	query		integer	false	"Number of paragraphs (1–20, default 1)"
//	@Param			sentences	query		integer	false	"Number of sentences per paragraph (1–20, default 5)"
//	@Success		200			{object}	httpx.Response[Lorem]
//	@Failure		400			{object}	httpx.ErrorResponse
//	@Router			/v1/text/lorem [get]
func handleLoremIpsum(svc Generator) http.HandlerFunc {
	return httpx.HandleGet(func(ctx context.Context, req Request) (Lorem, error) {
		return svc.Generate(req.Paragraphs, req.Sentences), nil
	}, Request{Paragraphs: 1, Sentences: 5})
}

// handleLoremIpsumBatch godoc
//
//	@Summary		Generate Lorem Ipsum (Batch)
//	@Description	Generate multiple Lorem Ipsum texts; partial successes returned.
//	@Tags			lorem-ipsum
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchRequest	true	"List of generation options"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchItemResponse]]
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/text/lorem/batch [post]
func handleLoremIpsumBatch(svc Generator) http.HandlerFunc {
	return httpx.HandleBatch(func(ctx context.Context, req BatchRequest) (httpx.BatchResponse[BatchItemResponse], error) {
		var results []BatchItemResponse

		for _, item := range req.Items {
			p := item.Paragraphs
			if p == 0 {
				p = 1
			}
			s := item.Sentences
			if s == 0 {
				s = 5
			}

			lorem := svc.Generate(p, s)

			results = append(results, BatchItemResponse{
				Data: &lorem,
			})
		}

		return httpx.BatchResponse[BatchItemResponse]{
			Results: results,
		}, nil
	})
}
