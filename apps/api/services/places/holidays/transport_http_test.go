package holidays

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

func setupRouter(t *testing.T) chi.Router {
	t.Helper()
	svc := NewService()
	r := chi.NewRouter()
	RegisterRoutes(r, svc)
	return r
}

func TestHolidays_ValidRequest(t *testing.T) {
	t.Parallel()
	r := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/holidays?country=US&year=2025", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[HolidayList]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "US", resp.Data.Country)
	assert.Equal(t, 2025, resp.Data.Year)
	assert.NotEmpty(t, resp.Data.Holidays)
}

func TestHolidays_UK(t *testing.T) {
	t.Parallel()
	r := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/holidays?country=GB&year=2025", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[HolidayList]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "GB", resp.Data.Country)
}

func TestHolidays_MissingCountry(t *testing.T) {
	t.Parallel()
	r := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/holidays?year=2025", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHolidays_MissingYear(t *testing.T) {
	t.Parallel()
	r := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/holidays?country=US", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHolidays_InvalidCountry(t *testing.T) {
	t.Parallel()
	r := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/holidays?country=INVALID&year=2025", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHolidays_InvalidYear(t *testing.T) {
	t.Parallel()
	r := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/holidays?country=US&year=0", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHolidays_NoParams(t *testing.T) {
	t.Parallel()
	r := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/holidays", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHolidays_LowercaseCountry(t *testing.T) {
	t.Parallel()
	r := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/holidays?country=us&year=2025", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[HolidayList]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "US", resp.Data.Country)
}

// ---- batch endpoint tests ----

func TestHolidaysBatch_ValidRequest(t *testing.T) {
	t.Parallel()
	r := setupRouter(t)

	body := `{"queries":[{"country":"US","year":2025},{"country":"GB","year":2025}]}`
	req := httptest.NewRequest(http.MethodPost, "/holidays/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[httpx.BatchResponse[BatchItem]]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, 2, resp.Data.Total)
	for _, item := range resp.Data.Results {
		if !item.Found {
			t.Errorf("expected found=true for %s %d", item.Country, item.Year)
		}
	}
}

func TestHolidaysBatch_SetsUsageCountHeader(t *testing.T) {
	t.Parallel()
	r := setupRouter(t)

	body := `{"queries":[{"country":"US","year":2025},{"country":"DE","year":2025}]}`
	req := httptest.NewRequest(http.MethodPost, "/holidays/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "2", w.Header().Get("X-Usage-Count"))
}

func TestHolidaysBatch_PartialFailure(t *testing.T) {
	t.Parallel()
	r := setupRouter(t)

	// AQ (Antarctica) is a valid ISO 3166-1 alpha-2 code but has no holiday data,
	// so it passes input validation and triggers the found:false path in the service.
	body := `{"queries":[{"country":"US","year":2025},{"country":"AQ","year":2025}]}`
	req := httptest.NewRequest(http.MethodPost, "/holidays/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[httpx.BatchResponse[BatchItem]]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.True(t, resp.Data.Results[0].Found)
	assert.False(t, resp.Data.Results[1].Found)
}

func TestHolidaysBatch_EmptyQueries(t *testing.T) {
	t.Parallel()
	r := setupRouter(t)

	body := `{"queries":[]}`
	req := httptest.NewRequest(http.MethodPost, "/holidays/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestHolidaysBatch_MissingBody(t *testing.T) {
	t.Parallel()
	r := setupRouter(t)

	req := httptest.NewRequest(http.MethodPost, "/holidays/batch", strings.NewReader(`{}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
