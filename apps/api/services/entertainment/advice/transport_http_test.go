package advice

import (
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

func TestAdviceBatch_HappyPath(t *testing.T) {
	t.Parallel()

	svc := &Service{
		db: &mockQuerier{
			row: &mockRow{
				scanFn: func(dest ...any) error {
					*dest[0].(*int) = 1
					*dest[1].(*string) = "Stay consistent."
					return nil
				},
			},
		},
	}

	r := chi.NewRouter()
	RegisterRoutes(r, svc)

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
		Data BatchResponse[Advice] `json:"data"`
	}

	err := json.NewDecoder(rec.Body).Decode(&got)
	require.NoError(t, err)

	require.Len(t, got.Data.Results, 2)

	assert.Equal(t, 1, got.Data.Results[0].ID)
	assert.Equal(t, "Stay consistent.", got.Data.Results[0].Text)

	assert.Equal(t, 1, got.Data.Results[1].ID)
	assert.Equal(t, "Stay consistent.", got.Data.Results[1].Text)
}

func TestAdviceBatch_InvalidJSON(t *testing.T) {
	t.Parallel()

	svc := &Service{}

	r := chi.NewRouter()
	RegisterRoutes(r, svc)

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

	svc := &Service{}

	r := chi.NewRouter()
	RegisterRoutes(r, svc)

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

	svc := &Service{
		db: &mockQuerier{
			row: &mockRow{
				scanFn: func(_ ...any) error {
					return pgx.ErrNoRows
				},
			},
		},
	}

	r := chi.NewRouter()
	RegisterRoutes(r, svc)

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
