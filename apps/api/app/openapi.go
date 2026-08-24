package app

import (
	_ "embed"
	"net/http"
)

//go:embed openapi.json
var openapiSpec []byte

// OpenAPISpec serves the build-time-embedded OpenAPI spec, matching what the
// retired Cloudflare Worker used to serve at the same path.
func OpenAPISpec() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(openapiSpec)
	}
}
