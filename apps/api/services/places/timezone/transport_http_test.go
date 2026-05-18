package timezone

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
	"strings"


	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
)

func setupRouter(t *testing.T) chi.Router {
	t.Helper()
	svc, err := NewService()
	require.NoError(t, err)
	r := chi.NewRouter()
	RegisterRoutes(r, svc)
	return r
}

func TestTimezone_NilService_Returns500(t *testing.T) {
	t.Parallel()
	r := chi.NewRouter()
	RegisterRoutes(r, nil)

	for _, path := range []string{"/timezone?lat=51.5&lon=-0.1", "/time/UTC"} {
		req := httptest.NewRequest(http.MethodGet, path, http.NoBody)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusInternalServerError {
			t.Errorf("path %q: expected 500, got %d", path, w.Code)
		}

		var resp httpx.ErrorResponse
		err := json.NewDecoder(w.Body).Decode(&resp)
		if err != nil {
			t.Errorf("path %q: failed to decode response: %v", path, err)
			continue
		}

		if resp.Error != "internal_error" {
			t.Errorf("path %q: expected error code %q, got %q", path, "internal_error", resp.Error)
		}
	}
}

func TestTimezone_ByCoords(t *testing.T) {
	t.Parallel()
	r := setupRouter(t)

	// London coordinates
	req := httptest.NewRequest(http.MethodGet, "/timezone?lat=51.5&lon=-0.1", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Info]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "Europe/London", resp.Data.Timezone)
	assert.NotEmpty(t, resp.Data.CurrentTime)
	assert.NotEmpty(t, resp.Data.Offset)
}

func TestTimezone_ByCity(t *testing.T) {
	t.Parallel()
	r := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/timezone?city=Tokyo", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Info]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "Asia/Tokyo", resp.Data.Timezone)
}

func TestTimezone_CityNotFound(t *testing.T) {
	t.Parallel()
	r := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/timezone?city=Atlantis", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestTimezone_MissingParams(t *testing.T) {
	t.Parallel()
	r := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/timezone", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestTimezone_MissingLon(t *testing.T) {
	t.Parallel()
	r := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/timezone?lat=51.5", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestTimezone_MissingLat(t *testing.T) {
	t.Parallel()
	r := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/timezone?lon=-0.1", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestTimezone_InvalidLatRange(t *testing.T) {
	t.Parallel()
	r := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/timezone?lat=200&lon=0", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestTimezone_NewYork(t *testing.T) {
	t.Parallel()
	r := setupRouter(t)

	// New York City coordinates
	req := httptest.NewRequest(http.MethodGet, "/timezone?lat=40.7&lon=-74.0", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Info]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "America/New_York", resp.Data.Timezone)
}

func TestTimezone_CityLookup_CaseInsensitive(t *testing.T) {
	t.Parallel()
	svc, err := NewService()
	require.NoError(t, err)

	info, err := svc.GetTimezoneByCity("TOKYO")
	require.NoError(t, err)

	assert.Equal(t, "Asia/Tokyo", info.Timezone)
}

func TestWorldTime_ValidTimezone(t *testing.T) {
	t.Parallel()
	r := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/time/America/New_York", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Info]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "America/New_York", resp.Data.Timezone)
	assert.NotEmpty(t, resp.Data.CurrentTime)
	assert.NotEmpty(t, resp.Data.Offset)
}

func TestWorldTime_UTCTimezone(t *testing.T) {
	t.Parallel()
	r := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/time/UTC", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Info]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "UTC", resp.Data.Timezone)
	assert.Equal(t, "+00:00", resp.Data.Offset)
}

func TestWorldTime_InvalidTimezone(t *testing.T) {
	t.Parallel()
	r := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/time/Fake/Timezone", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestWorldTime_AsiaKolkata(t *testing.T) {
	t.Parallel()
	r := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/time/Asia/Kolkata", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Info]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "Asia/Kolkata", resp.Data.Timezone)
	assert.Equal(t, "+05:30", resp.Data.Offset)
}

func TestTimezone_OffsetFormat(t *testing.T) {
	t.Parallel()
	tests := []struct {
		offsetSecs int
		expected   string
	}{
		{0, "+00:00"},
		{3600, "+01:00"},
		{-18000, "-05:00"},
		{19800, "+05:30"},  // India
		{20700, "+05:45"},  // Nepal
		{-34200, "-09:30"}, // Marquesas Islands
	}

	for _, tc := range tests {
		got := formatOffset(tc.offsetSecs)
		if got != tc.expected {
			t.Errorf("formatOffset(%d) = %q, want %q", tc.offsetSecs, got, tc.expected)
		}
	}
}

func TestTimezone_Batch_HappyPath(t *testing.T) {
	t.Parallel()

	r := setupRouter(t)

	body := `{
		"cities": [
			"Tokyo",
			"Lima",
			"New York"
		]
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/timezone/batch",
		strings.NewReader(body),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[BatchResponse]

	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	require.Len(t, resp.Data.Results, 3)

	assert.Equal(t, "Tokyo", resp.Data.Results[0].City)
	assert.Equal(t, "Lima", resp.Data.Results[1].City)
	assert.Equal(t, "New York", resp.Data.Results[2].City)

	require.NotNil(t, resp.Data.Results[0].Info)
	require.NotNil(t, resp.Data.Results[1].Info)
	require.NotNil(t, resp.Data.Results[2].Info)
}

func TestTimezone_Batch_InvalidJSON(t *testing.T) {
	t.Parallel()

	r := setupRouter(t)

	req := httptest.NewRequest(
		http.MethodPost,
		"/timezone/batch",
		strings.NewReader(`{invalid-json}`),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var resp httpx.ErrorResponse

	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "bad_request", resp.Error)
}

func TestTimezone_Batch_EmptyCities(t *testing.T) {
	t.Parallel()

	r := setupRouter(t)

	body := `{
		"cities": []
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/timezone/batch",
		strings.NewReader(body),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var resp httpx.ErrorResponse

	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "bad_request", resp.Error)
}

func TestTimezone_Batch_InvalidCities(t *testing.T) {
	t.Parallel()

	r := setupRouter(t)

	body := `{
		"cities": [
			"Tokyo",
			"InvalidCity",
			"",
			"Atlantis"
		]
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/timezone/batch",
		strings.NewReader(body),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[BatchResponse]

	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	require.Len(t, resp.Data.Results, 4)

	assert.Equal(t, "Tokyo", resp.Data.Results[0].City)
	assert.Equal(t, "InvalidCity", resp.Data.Results[1].City)
	assert.Equal(t, "", resp.Data.Results[2].City)
	assert.Equal(t, "Atlantis", resp.Data.Results[3].City)

	require.NotNil(t, resp.Data.Results[0].Info)

	assert.Nil(t, resp.Data.Results[1].Info)
	assert.Nil(t, resp.Data.Results[2].Info)
	assert.Nil(t, resp.Data.Results[3].Info)
}

func TestTimezone_Batch_PreservesOrder(t *testing.T) {
	t.Parallel()

	r := setupRouter(t)

	body := `{
		"cities": [
			"Lima",
			"Tokyo",
			"Paris",
			"New York"
		]
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/timezone/batch",
		strings.NewReader(body),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[BatchResponse]

	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	expected := []string{
		"Lima",
		"Tokyo",
		"Paris",
		"New York",
	}

	require.Len(t, resp.Data.Results, len(expected))

	for i := range expected {
		assert.Equal(t, expected[i], resp.Data.Results[i].City)
	}
}

func TestTimezone_Batch_TooManyCities(t *testing.T) {
	t.Parallel()

	r := setupRouter(t)

	cities := make([]string, 0, 51)

	for i := 0; i < 51; i++ {
		cities = append(cities, "Tokyo")
	}

	reqBody := BatchRequest{
		Cities: cities,
	}

	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(
		http.MethodPost,
		"/timezone/batch",
		bytes.NewReader(body),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var resp httpx.ErrorResponse

	err = json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "bad_request", resp.Error)
}
