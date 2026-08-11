package app

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/config"
)

// TestOpenAPISpec_CoversRegisteredRoutes is the drift guard: every route the
// chi router registers must be documented in the OpenAPI spec. This keeps the
// swagger annotations from silently falling behind the Go handlers.
//
// The spec is regenerated with:
//
//	swag init -g main.go --outputTypes json --parseDependency
func TestOpenAPISpec_CoversRegisteredRoutes(t *testing.T) {
	pool, err := pgxpool.New(context.Background(),
		"postgres://requiem:requiem@localhost:5432/requiem?sslmode=disable")
	require.NoError(t, err)
	defer pool.Close()

	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	defer rdb.Close()

	r := chi.NewRouter()
	r.Route("/v1", func(v1 chi.Router) {
		registerV1Routes(context.Background(), v1, pool, rdb,
			config.Config{BackendSecret: "test_secret_min_32_chars_long_for_testing_only"})
	})

	spec, err := os.ReadFile(filepath.Join("..", "docs", "swagger.json"))
	require.NoError(t, err, "swagger.json missing; run `swag init -g main.go --outputTypes json --parseDependency`")

	var doc struct {
		Paths map[string]map[string]json.RawMessage `json:"paths"`
	}
	require.NoError(t, json.Unmarshal(spec, &doc))

	// chi wildcard routes (e.g. /v1/places/time/*) are documented as {param}.
	// Match at the segment level: a `*` or `{...}` in either side accepts any
	// single path segment.
	segmentMatches := func(route, specPath string) bool {
		a := strings.Split(route, "/")
		b := strings.Split(specPath, "/")
		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if a[i] == b[i] {
				continue
			}
			isParam := func(s string) bool {
				return s == "*" || (strings.HasPrefix(s, "{") && strings.HasSuffix(s, "}"))
			}
			if isParam(a[i]) || isParam(b[i]) {
				continue
			}
			return false
		}
		return true
	}

	match := func(route string) bool {
		if _, ok := doc.Paths[route]; ok {
			return true
		}
		for specPath := range doc.Paths {
			if segmentMatches(route, specPath) {
				return true
			}
		}
		return false
	}

	var missing []string
	_ = chi.Walk(r, func(method, route string, _ http.Handler, _ ...func(http.Handler) http.Handler) error {
		if route == "/*" {
			return nil
		}
		if !match(route) {
			missing = append(missing, method+" "+route)
		}
		return nil
	})

	require.Empty(t, missing,
		"registered routes missing from the OpenAPI spec — add swagger annotations and regenerate")
}
