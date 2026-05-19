package emoji

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// SearchRequest holds the query parameter for emoji search.
type SearchRequest struct {
	Query string `query:"q" validate:"required,min=1,max=100"`
}

// RegisterRoutes mounts emoji handlers on the given router.
// Paths are relative to the parent mount point (/v1/entertainment).
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/emoji/random", func(w http.ResponseWriter, r *http.Request) {
		httpx.JSON(w, http.StatusOK, svc.Random())
	})

	r.Get("/emoji/search", httpx.HandleGet(func(ctx context.Context, req SearchRequest) (List, error) {
		return svc.Search(req.Query), nil
	}))

	r.Get("/emoji/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := strings.ToLower(chi.URLParam(r, "name"))
		e, ok := svc.GetByName(name)
		if !ok {
			httpx.Error(w, http.StatusNotFound, "not_found", "emoji not found")
			return
		}

		httpx.JSON(w, http.StatusOK, e)
	})
}
