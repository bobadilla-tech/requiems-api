package middleware

import (
	"net/http"
	"regexp"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// ValidateURLParam validates a URL parameter against a regex pattern.
// If validation fails, responds with 400 Bad Request and stops handler execution.
func ValidateURLParam(paramName string, pattern *regexp.Regexp, errorMsg string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			paramVal := chi.URLParam(r, paramName)

			if !pattern.MatchString(paramVal) {
				httpx.Error(w, http.StatusUnprocessableEntity, "validation_error", errorMsg)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
