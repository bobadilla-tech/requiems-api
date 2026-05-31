package advice

import (
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

func newRouterWithRows(rows *mockRows) chi.Router {
	svc := &Service{db: &mockDB{
		row:  &mockRow{scanFn: func(_ ...any) error { return errors.New("unused") }},
		rows: rows,
	}}
	r := chi.NewRouter()
	RegisterRoutes(r, svc)
	return r
}

func TestAdviceBatch_HappyPath(t *testing.T) {
	t.Parallel()

	r := newRouterWithRows(&mockRows{
		items: [][]any{
			{1, "Stay consistent."},
			{2, "Stay consistent."},
		},
	})

	req := httptest.NewRequest(
		http.MethodPost,
		"/advice/batch",
		strings.NewReader(`{"count":2}`),
	)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var got struct {
		Data httpx.BatchResponse[Advice] `json:"data"`
	}
	err := json.NewDecoder(rec.Body).Decode(&got)
	require.NoError(t, err)

	require.Len(t, got.Data.Results, 2)
	assert.Equal(t, 1, got.Data.Results[0].ID)
	assert.Equal(t, "Stay consistent.", got.Data.Results[0].Text)
	assert.Equal(t, 2, got.Data.Results[1].ID)
	assert.Equal(t, "Stay consistent.", got.Data.Results[1].Text)
}

func TestAdviceBatch_InvalidJSON(t *testing.T) {
	t.Parallel()

	r := newRouterWithRows(&mockRows{})
	req := httptest.NewRequest(
		http.MethodPost,
		"/advice/batch",
		strings.NewReader(`{invalid-json}`),
	)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestAdviceBatch_InvalidCount(t *testing.T) {
	t.Parallel()

	r := newRouterWithRows(&mockRows{})
	req := httptest.NewRequest(
		http.MethodPost,
		"/advice/batch",
		strings.NewReader(`{"count":0}`),
	)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
}

func TestAdviceBatch_ServiceError(t *testing.T) {
	t.Parallel()

	r := newRouterWithRows(&mockRows{err: errors.New("db unavailable")})
	req := httptest.NewRequest(
		http.MethodPost,
		"/advice/batch",
		strings.NewReader(`{"count":2}`),
	)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusServiceUnavailable, rec.Code)
}
