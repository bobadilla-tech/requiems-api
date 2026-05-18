package quotes

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type httpMockQuerier struct {
	row pgx.Row
}

func (m *httpMockQuerier) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return m.row
}

func TestTransport_HappyPath(t *testing.T) {
	t.Parallel()
	mockRow := &mockRow{
		scanFn: func(dest ...any) error {
			*dest[0].(*int) = 1
			*dest[1].(*string) = "Hello"
			*dest[2].(*string) = "Test"
			return nil
		},
	}

	svc := &Service{
		db: &httpMockQuerier{row: mockRow},
	}

	r := chi.NewRouter()
	RegisterRoutes(r, svc)

	req := httptest.NewRequest("GET", "/quotes/random", http.NoBody)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
}

func TestTransport_Error(t *testing.T) {
	t.Parallel()
	mockRow := &mockRow{
		scanFn: func(dest ...any) error {
			return pgx.ErrNoRows
		},
	}

	svc := &Service{
		db: &httpMockQuerier{row: mockRow},
	}

	r := chi.NewRouter()
	RegisterRoutes(r, svc)

	req := httptest.NewRequest("GET", "/quotes/random", http.NoBody)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusServiceUnavailable, w.Result().StatusCode)
}

func TestTransport_MethodNotAllowed(t *testing.T) {
	t.Parallel()
	svc := &Service{}

	r := chi.NewRouter()
	RegisterRoutes(r, svc)

	// POST is not registered; chi must return 405 without calling the GET handler
	// (which would panic: Service has no db in this test).
	req := httptest.NewRequest(http.MethodPost, "/quotes/random", http.NoBody)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusMethodNotAllowed, w.Result().StatusCode)
}

func TestTransportBatch_HappyPath(t *testing.T) {
	t.Parallel()
	svc := &Service{
		db: &httpMockQuerier{row: &mockRow{
			scanFn: func(dest ...any) error {
				*dest[0].(*int) = 1
				*dest[1].(*string) = "Test quote."
				*dest[2].(*string) = "Author"
				return nil
			},
		}},
	}

	r := chi.NewRouter()
	RegisterRoutes(r, svc)

	body := `{"count": 3}`
	req := httptest.NewRequest(http.MethodPost, "/quotes/random/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, "3", w.Result().Header.Get("X-Usage-Count"))

	// Verify response body structure: envelope must contain results and total.
	var resp struct {
		Data struct {
			Results []struct {
				ID     int    `json:"id"`
				Text   string `json:"text"`
				Author string `json:"author"`
			} `json:"results"`
			Total int `json:"total"`
		} `json:"data"`
	}
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Equal(t, 3, resp.Data.Total)
	assert.Len(t, resp.Data.Results, 3)
	assert.Equal(t, "Test quote.", resp.Data.Results[0].Text)
}

func TestTransportBatch_EmptyBody(t *testing.T) {
	t.Parallel()
	svc := &Service{}

	r := chi.NewRouter()
	RegisterRoutes(r, svc)

	req := httptest.NewRequest(http.MethodPost, "/quotes/random/batch", http.NoBody)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
}

func TestTransportBatch_CountZero(t *testing.T) {
	t.Parallel()
	svc := &Service{}

	r := chi.NewRouter()
	RegisterRoutes(r, svc)

	body := `{"count": 0}`
	req := httptest.NewRequest(http.MethodPost, "/quotes/random/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Result().StatusCode)
}

func TestTransportBatch_CountExceedsMax(t *testing.T) {
	t.Parallel()
	svc := &Service{}

	r := chi.NewRouter()
	RegisterRoutes(r, svc)

	body := `{"count": 51}`
	req := httptest.NewRequest(http.MethodPost, "/quotes/random/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Result().StatusCode)
}
