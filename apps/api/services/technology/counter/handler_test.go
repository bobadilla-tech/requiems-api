package counter

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
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
	batchFn     func(ctx context.Context, namespaces []string) []BatchCounterItem
}

func (m *mockService) Increment(ctx context.Context, namespace string) (int64, error) {
	return m.incrementFn(ctx, namespace)
}

func (m *mockService) Get(ctx context.Context, namespace string) (int64, error) {
	return m.getFn(ctx, namespace)
}

func (m *mockService) IncrementBatch(ctx context.Context, namespaces []string) []BatchCounterItem {
	return m.batchFn(ctx, namespaces)
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

func TestIncrementBatchHandler(t *testing.T) {
	t.Parallel()
	t.Run("happy path", func(t *testing.T) {
		t.Parallel()
		svc := &mockService{
			batchFn: func(_ context.Context, _ []string) []BatchCounterItem {
				return []BatchCounterItem{
					{Namespace: "hits", Value: 10},
					{Namespace: "page-views", Value: 5},
				}
			},
		}
		req := httptest.NewRequest(http.MethodPost, "/counter/batch", strings.NewReader(`{"namespaces":["hits","page-views"]}`))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		newTestRouter(svc).ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)
		var resp httpx.Response[httpx.BatchResponse[BatchCounterItem]]
		err := json.NewDecoder(w.Body).Decode(&resp)
		require.NoError(t, err)
		assert.Len(t, resp.Data.Results, 2)
		assert.Equal(t, 2, resp.Data.Total)
		assert.Equal(t, "hits", resp.Data.Results[0].Namespace)
		assert.Equal(t, int64(10), resp.Data.Results[0].Value)
		assert.Empty(t, resp.Data.Results[0].Error)
	})

	t.Run("invalid namespace in-band error", func(t *testing.T) {
		t.Parallel()
		svc := &mockService{
			batchFn: func(_ context.Context, _ []string) []BatchCounterItem {
				return []BatchCounterItem{
					{Namespace: "hits", Value: 1},
					{Namespace: "!!!bad!!!", Error: "invalid namespace"},
				}
			},
		}
		req := httptest.NewRequest(http.MethodPost, "/counter/batch", strings.NewReader(`{"namespaces":["hits","!!!bad!!!"]}`))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		newTestRouter(svc).ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)
		var resp httpx.Response[httpx.BatchResponse[BatchCounterItem]]
		err := json.NewDecoder(w.Body).Decode(&resp)
		require.NoError(t, err)
		assert.Len(t, resp.Data.Results, 2)
		assert.Empty(t, resp.Data.Results[0].Error, "valid ns should have no error")
		assert.NotEmpty(t, resp.Data.Results[1].Error, "invalid ns should have in-band error")
	})

	t.Run("empty namespaces returns 422", func(t *testing.T) {
		t.Parallel()
		req := httptest.NewRequest(http.MethodPost, "/counter/batch", strings.NewReader(`{"namespaces":[]}`))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		newTestRouter(&mockService{}).ServeHTTP(w, req)
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
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
