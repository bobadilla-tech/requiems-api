package workingdays

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newBatchRouter() chi.Router {
	r := chi.NewRouter()
	RegisterRoutes(r, NewService())
	return r
}

func mustMarshal(t *testing.T, v any) *bytes.Buffer {
	t.Helper()

	b, err := json.Marshal(v)
	require.NoError(t, err)

	return bytes.NewBuffer(b)
}

func TestWorkingDaysBatch_HappyPath(t *testing.T) {
	t.Parallel()

	body := BatchRequest{
		Items: []BatchItem{
			{
				From: Date{
					Time: mustParseDate(t, "2026-01-05"),
				},
				To: Date{
					Time: mustParseDate(t, "2026-01-09"),
				},
			},
			{
				From: Date{
					Time: mustParseDate(t, "2026-01-05"),
				},
				To: Date{
					Time: mustParseDate(t, "2026-01-09"),
				},
				Country: "US",
			},
		},
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/working-days/batch",
		mustMarshal(t, body),
	)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	newBatchRouter().ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var got struct {
		Data BatchResponse `json:"data"`
	}

	require.NoError(t, json.NewDecoder(rec.Body).Decode(&got))

	require.Len(t, got.Data.Results, 2)

	assert.Equal(t, "2026-01-05", got.Data.Results[0].From)
	assert.Equal(t, "2026-01-09", got.Data.Results[0].To)
	assert.Equal(t, "", got.Data.Results[0].Country)

	assert.Equal(t, "US", got.Data.Results[1].Country)
}

func TestWorkingDaysBatch_InvalidJSON(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(
		http.MethodPost,
		"/working-days/batch",
		bytes.NewBufferString(`{invalid}`),
	)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	newBatchRouter().ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestWorkingDaysBatch_EmptyItems(t *testing.T) {
	t.Parallel()

	body := BatchRequest{
		Items: []BatchItem{},
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/working-days/batch",
		mustMarshal(t, body),
	)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	newBatchRouter().ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
}

func TestWorkingDaysBatch_NilItems(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(
		http.MethodPost,
		"/working-days/batch",
		bytes.NewBufferString(`{}`),
	)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	newBatchRouter().ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
}

func TestWorkingDaysBatch_ExceedsLimit(t *testing.T) {
	t.Parallel()

	items := make([]BatchItem, 51)

	for i := range items {
		items[i] = BatchItem{
			From: Date{
				Time: mustParseDate(t, "2026-01-05"),
			},
			To: Date{
				Time: mustParseDate(t, "2026-01-09"),
			},
		}
	}

	body := BatchRequest{
		Items: items,
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/working-days/batch",
		mustMarshal(t, body),
	)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	newBatchRouter().ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
}

func TestWorkingDaysBatch_MaxItems(t *testing.T) {
	t.Parallel()

	items := make([]BatchItem, 50)

	for i := range items {
		items[i] = BatchItem{
			From: Date{
				Time: mustParseDate(t, "2026-01-05"),
			},
			To: Date{
				Time: mustParseDate(t, "2026-01-09"),
			},
		}
	}

	body := BatchRequest{
		Items: items,
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/working-days/batch",
		mustMarshal(t, body),
	)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	newBatchRouter().ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var got struct {
		Data BatchResponse `json:"data"`
	}

	require.NoError(t, json.NewDecoder(rec.Body).Decode(&got))

	assert.Len(t, got.Data.Results, 50)
}

func mustParseDate(t *testing.T, s string) time.Time {
	t.Helper()

	parsed, err := time.Parse(time.DateOnly, s)
	require.NoError(t, err)

	return parsed
}
