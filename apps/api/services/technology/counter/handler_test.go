package counter

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
)

var errRedisDown = errors.New("connection refused")

type mockService struct {
	incrementFn func(ctx context.Context, namespace string) (int64, error)
	getFn       func(ctx context.Context, namespace string) (int64, error)
}

func (m *mockService) Increment(ctx context.Context, namespace string) (int64, error) {
	return m.incrementFn(ctx, namespace)
}

func (m *mockService) Get(ctx context.Context, namespace string) (int64, error) {
	return m.getFn(ctx, namespace)
}

func newTestRouter(svc Service) http.Handler {
	r := chi.NewRouter()
	RegisterRoutes(r, svc)
	return r
}

func TestIncrementHandler(t *testing.T) {
	t.Parallel()
	t.Run("returns updated counter value", func(t *testing.T) {
		t.Parallel()
		svc := &mockService{
			incrementFn: func(_ context.Context, ns string) (int64, error) {
				return 5, nil
			},
		}

		req := httptest.NewRequest(http.MethodPost, "/counter/hits", http.NoBody)
		w := httptest.NewRecorder()

		newTestRouter(svc).ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)

		var resp httpx.Response[Counter]

		err := json.NewDecoder(w.Body).Decode(&resp)
		require.NoError(t, err)
		got := resp.Data

		assert.Equal(t, "hits", got.Namespace)
		assert.Equal(t, int64(5), got.Value)
	})

	t.Run("returns 422 for invalid namespace from URL param validation", func(t *testing.T) {
		t.Parallel()
		svc := &mockService{
			incrementFn: func(_ context.Context, ns string) (int64, error) {
				return 1, nil
			},
		}

		req := httptest.NewRequest(http.MethodPost, "/counter/!!!invalid!!!", http.NoBody)
		w := httptest.NewRecorder()

		newTestRouter(svc).ServeHTTP(w, req)

		require.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("returns 500 for internal server error", func(t *testing.T) {
		t.Parallel()
		svc := &mockService{
			incrementFn: func(_ context.Context, ns string) (int64, error) {
				return 0, errRedisDown
			},
		}

		req := httptest.NewRequest(http.MethodPost, "/counter/hits", http.NoBody)
		w := httptest.NewRecorder()

		newTestRouter(svc).ServeHTTP(w, req)

		require.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestGetHandler(t *testing.T) {
	t.Parallel()
	t.Run("returns counter value", func(t *testing.T) {
		t.Parallel()
		svc := &mockService{
			getFn: func(_ context.Context, ns string) (int64, error) {
				return 42, nil
			},
		}

		req := httptest.NewRequest(http.MethodGet, "/counter/page-views", http.NoBody)
		w := httptest.NewRecorder()

		newTestRouter(svc).ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)

		var resp httpx.Response[Counter]
		err := json.NewDecoder(w.Body).Decode(&resp)
		require.NoError(t, err)
		got := resp.Data

		assert.Equal(t, "page-views", got.Namespace)
		assert.Equal(t, int64(42), got.Value)
	})

	t.Run("returns 422 for invalid namespace from URL param validation", func(t *testing.T) {
		t.Parallel()
		svc := &mockService{
			getFn: func(_ context.Context, ns string) (int64, error) {
				return 42, nil
			},
		}

		req := httptest.NewRequest(http.MethodGet, "/counter/!!!invalid!!!", http.NoBody)
		w := httptest.NewRecorder()

		newTestRouter(svc).ServeHTTP(w, req)

		require.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("returns 500 for internal server error", func(t *testing.T) {
		t.Parallel()
		svc := &mockService{
			getFn: func(_ context.Context, ns string) (int64, error) {
				return 0, errRedisDown
			},
		}

		req := httptest.NewRequest(http.MethodGet, "/counter/page-views", http.NoBody)
		w := httptest.NewRecorder()

		newTestRouter(svc).ServeHTTP(w, req)

		require.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
