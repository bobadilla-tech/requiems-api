package lorem

import (
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

func RegisterRoutes(r chi.Router, svc Generator) {
	r.Get("/lorem", func(w http.ResponseWriter, r *http.Request) {
		// Set defaults before binding so unset params keep their default value.
		req := Request{Paragraphs: 1, Sentences: 5}

		if err := httpx.BindQuery(r, &req); err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		httpx.JSON(w, http.StatusOK, svc.Generate(req.Paragraphs, req.Sentences))
	})
}
