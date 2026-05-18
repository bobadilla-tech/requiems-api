package useragent

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// ParseRequest holds the query parameters for the user agent parse endpoint.
type ParseRequest struct {
	UA string `query:"ua" validate:"required"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/useragent", func(w http.ResponseWriter, r *http.Request) {
		var req ParseRequest

		if err := httpx.BindQuery(r, &req); err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		httpx.JSON(w, http.StatusOK, svc.Parse(req.UA))
	})
}
