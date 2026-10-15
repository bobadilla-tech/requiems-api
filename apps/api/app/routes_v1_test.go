package app

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/config"
	"requiems-api/platform/middleware"
)

func TestServiceEnabled(t *testing.T) {
	t.Parallel()

	t.Run("returns true when enabled services is blank", func(t *testing.T) {
		t.Parallel()

		assert.True(t, serviceEnabled(config.Config{}, "text"), "expected service to be enabled when config is blank")
	})

	t.Run("matches trimmed service names from config", func(t *testing.T) {
		t.Parallel()

		cfg := config.Config{EnabledServices: "text, validation,networking"}

		assert.True(t, serviceEnabled(cfg, "text"), "expected text service to be enabled")
		assert.False(t, serviceEnabled(cfg, "finance"), "expected finance service to be disabled")
	})
}

func TestRegisterV1Routes(t *testing.T) {
	t.Parallel()

	t.Run("mounts all services when enabled services is blank", func(t *testing.T) {
		t.Parallel()

		r := chi.NewRouter()

		registerV1Routes(context.Background(), r, nil, nil, config.Config{})

		routes := walkRoutes(t, r)
		for _, prefix := range []string{
			"/entertainment", "/finance", "/health", "/networking",
			"/places", "/technology", "/text", "/validation",
		} {
			require.True(t, hasRoutePrefix(routes, prefix), "expected mounted routes to include prefix %s; got %v", prefix, routes)
		}
	})

	t.Run("mounts only explicitly enabled services", func(t *testing.T) {
		t.Parallel()

		r := chi.NewRouter()

		registerV1Routes(context.Background(), r, nil, nil, config.Config{EnabledServices: "validation,text"})

		routes := walkRoutes(t, r)
		require.True(t, hasRoutePrefix(routes, "/validation"), "expected validation routes to be mounted; got %v", routes)
		require.True(t, hasRoutePrefix(routes, "/text"), "expected text routes to be mounted; got %v", routes)
		assert.False(t, hasRoutePrefix(routes, "/finance"), "expected finance routes to be absent; got %v", routes)
		assert.False(t, hasRoutePrefix(routes, "/technology"), "expected technology routes to be absent; got %v", routes)
	})
}

func walkRoutes(t *testing.T, r chi.Router) []string {
	t.Helper()

	var routes []string
	if err := chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		routes = append(routes, route)
		return nil
	}); err != nil {
		t.Fatalf("walk routes: %v", err)
	}

	return routes
}

func hasRoutePrefix(routes []string, prefix string) bool {
	for _, route := range routes {
		if strings.HasPrefix(route, prefix) {
			return true
		}
	}
	return false
}

// serveTracked runs the request through r and reports the status and whether a
// NotFound/MethodNotAllowed handler flagged it as a route miss.
func serveTracked(r http.Handler, method, target string) (status int, missed bool) {
	req := httptest.NewRequest(method, target, http.NoBody)
	ctx, missedFn := middleware.TrackRouteMiss(req.Context())

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req.WithContext(ctx))

	return w.Code, missedFn()
}

func TestRegisterV1Routes_MissesAreFlaggedOnEveryMount(t *testing.T) {
	t.Parallel()

	root := chi.NewRouter()
	root.Route("/v1", func(v1 chi.Router) {
		registerV1Routes(context.Background(), v1, nil, nil, config.Config{})
	})

	t.Run("unknown top-level route", func(t *testing.T) {
		t.Parallel()

		status, missed := serveTracked(root, http.MethodGet, "/v1/x")
		assert.Equal(t, http.StatusNotFound, status)
		assert.True(t, missed)
	})

	for _, prefix := range []string{
		"entertainment", "finance", "health", "networking",
		"places", "technology", "text", "validation", "systems",
	} {
		t.Run("unknown route under "+prefix, func(t *testing.T) {
			t.Parallel()

			status, missed := serveTracked(root, http.MethodGet, "/v1/"+prefix+"/__definitely_not_a_route__")
			assert.Equal(t, http.StatusNotFound, status)
			assert.True(t, missed, "the %s router must carry the NotFound flagging handler", prefix)
		})
	}

	t.Run("services that are not enabled are plain misses", func(t *testing.T) {
		t.Parallel()

		only := chi.NewRouter()
		only.Route("/v1", func(v1 chi.Router) {
			registerV1Routes(context.Background(), v1, nil, nil, config.Config{EnabledServices: "text"})
		})

		status, missed := serveTracked(only, http.MethodGet, "/v1/finance/anything")
		assert.Equal(t, http.StatusNotFound, status)
		assert.True(t, missed)
	})
}

func TestNewServiceRouter_NestedRoutesInheritTheMissHandlers(t *testing.T) {
	t.Parallel()

	service := newServiceRouter()
	service.Get("/top", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusOK) })
	service.Route("/nested", func(n chi.Router) {
		n.Get("/ok", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusOK) })
	})

	root := chi.NewRouter()
	root.Mount("/v1/svc", service)

	t.Run("a real route is not a miss", func(t *testing.T) {
		t.Parallel()

		for _, target := range []string{"/v1/svc/top", "/v1/svc/nested/ok"} {
			status, missed := serveTracked(root, http.MethodGet, target)
			assert.Equal(t, http.StatusOK, status, target)
			assert.False(t, missed, target)
		}
	})

	t.Run("an unknown route in a nested Route is a 404 miss", func(t *testing.T) {
		t.Parallel()

		status, missed := serveTracked(root, http.MethodGet, "/v1/svc/nested/zzz")
		assert.Equal(t, http.StatusNotFound, status)
		assert.True(t, missed)
	})

	t.Run("a wrong method in a nested Route is a 405 miss", func(t *testing.T) {
		t.Parallel()

		status, missed := serveTracked(root, http.MethodPost, "/v1/svc/nested/ok")
		assert.Equal(t, http.StatusMethodNotAllowed, status)
		assert.True(t, missed)
	})
}

func TestRootRouterMatchSeesThroughMountsAndEncodedPaths(t *testing.T) {
	t.Parallel()

	service := newServiceRouter()
	service.Get("/echo/{word}", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusOK) })

	root := chi.NewRouter()
	root.Group(func(g chi.Router) {
		g.Route("/v1", func(v1 chi.Router) { v1.Mount("/svc", service) })
	})

	match := func(method, path string) bool { return root.Match(chi.NewRouteContext(), method, path) }

	assert.True(t, match(http.MethodGet, "/v1/svc/echo/hello"))
	assert.True(t, match(http.MethodGet, "/v1/svc/echo/a%20b"), "an encoded parameter still matches the route")
	assert.False(t, match(http.MethodGet, "/v1/svc/echo"), "missing parameter")
	assert.False(t, match(http.MethodPost, "/v1/svc/echo/hello"), "wrong method")
	assert.False(t, match(http.MethodGet, "/v1/svc/nope"))
	assert.False(t, match(http.MethodGet, "/v1/nope"))
}
