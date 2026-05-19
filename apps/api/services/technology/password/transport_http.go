package password

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request holds the optional query parameters for the password endpoint.
type Request struct {
	Length    int  `query:"length"    validate:"min=8,max=128"`
	Uppercase bool `query:"uppercase"`
	Numbers   bool `query:"numbers"`
	Symbols   bool `query:"symbols"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/password", func(w http.ResponseWriter, r *http.Request) {
		req := Request{Length: 16}

		if err := httpx.BindQuery(r, &req); err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		result, err := svc.Generate(req.Length, req.Uppercase, req.Numbers, req.Symbols)
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "failed to generate password")
			return
		}

		httpx.JSON(w, http.StatusOK, result)
	})
}
