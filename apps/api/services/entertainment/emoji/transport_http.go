package emoji

import (
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

	r.Get("/emoji/search", func(w http.ResponseWriter, r *http.Request) {
		req := SearchRequest{}
		if err := httpx.BindQuery(r, &req); err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		httpx.JSON(w, http.StatusOK, svc.Search(req.Query))
	})

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
