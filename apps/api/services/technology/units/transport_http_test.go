package units

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
)

func setupRouter() chi.Router {
	r := chi.NewRouter()
	RegisterRoutes(r, NewService())
	return r
}

func TestUnits_BatchConvert_HappyPath(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"operations":[{"from": "miles", "to": "km", "value": 10}]}`
	req := httptest.NewRequest(http.MethodPost, "/convert/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[httpx.BatchResponse[BatchResponse]]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	require.Contains(t, w.Header().Get("Content-Type"), "application/json")

	require.Equal(t, 1, resp.Data.Total)
	require.Len(t, resp.Data.Results, 1)

	assert.True(t, resp.Data.Results[0].Success)
	assert.Empty(t, resp.Data.Results[0].Error)

	assert.InDelta(
		t,
		16.0934,
		resp.Data.Results[0].Data.Result,
		0.001,
	)
}

// Test EmptyBody
func TestUnits_BatchConvert_EmptyOperations(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"operations":[]}`
	req := httptest.NewRequest(http.MethodPost, "/convert/batch", strings.NewReader(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestUnits_BatchConvert_ExceedsLimit(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	operations := make([]BatchItem, 51)

	for i := range operations {
		operations[i] = BatchItem{From: "miles", To: "km", Value: new(float64(10))}
	}

	body, _ := json.Marshal(BatchRequest{
		Operations: operations,
	})

	req := httptest.NewRequest(http.MethodPost, "/convert/batch", strings.NewReader(string(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestUnits_BatchConvert_SetUsageCountHeader(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"operations":[{"from": "miles", "to": "km", "value": 10}]}`
	req := httptest.NewRequest(http.MethodPost, "/convert/batch", strings.NewReader(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	require.Equal(t, "1", w.Header().Get("X-Usage-Count"))
}
