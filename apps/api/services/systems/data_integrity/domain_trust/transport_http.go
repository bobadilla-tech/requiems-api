package domaintrust

import (
	"context"
	"net/http"
	"regexp"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
	"requiems-api/platform/middleware"
)

type EvaluateService interface {
	Evaluate(ctx context.Context, domain string) Response
}

// domainRe accepts standard hostnames such as "example.com" or "sub.example.co.uk".
// Each label is 1–63 chars (alphanumeric or hyphens, not starting/ending with a
// hyphen), and there must be at least one dot separating the labels.
var domainRe = regexp.MustCompile(`^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`)

func RegisterRoutes(r chi.Router, svc EvaluateService) {
	r.Group(func(validated chi.Router) {
		validated.Use(middleware.ValidateURLParam("domain", domainRe, "invalid domain: must be a valid hostname such as example.com"))

		validated.Get("/domain/trust/{domain}", func(w http.ResponseWriter, r *http.Request) {
			d := chi.URLParam(r, "domain")
			httpx.JSON(w, http.StatusOK, svc.Evaluate(r.Context(), d))
		})
	})
}
