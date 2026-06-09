package color_test

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
	"requiems-api/services/technology/color"
)

func setupRouter() chi.Router {
	r := chi.NewRouter()
	color.RegisterRoutes(r, color.NewService())
	return r
}

func TestColor_HappyPath_HexToRGB(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(
		http.MethodGet,
		"/color?from=hex&to=rgb&value=%23ffffff",
		http.NoBody,
	)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	assert.True(t, strings.Contains(w.Header().Get("Content-Type"), "application/json"), "expected application/json, got %s", w.Header().Get("Content-Type"))

	var res httpx.Response[color.Response]
	err := json.NewDecoder(w.Body).Decode(&res)
	require.NoError(t, err)

	assert.Equal(t, "#ffffff", res.Data.Input)
	assert.True(t, strings.Contains(res.Data.Result, "rgb"), "expected RGB result, got %s", res.Data.Result)
}

func TestColor_MissingParam(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(
		http.MethodGet,
		"/color?from=hex&to=rgb",
		http.NoBody,
	)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestColor_InvalidValue(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(
		http.MethodGet,
		"/color?from=hex&to=rgb&value=invalid",
		http.NoBody,
	)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestColor_ServiceError(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(
		http.MethodGet,
		"/color?from=hex&to=rgb&value=notahex",
		http.NoBody,
	)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
	assert.True(t, strings.Contains(w.Header().Get("Content-Type"), "application/json"), "expected application/json, got %s", w.Header().Get("Content-Type"))

	var resp httpx.ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "invalid_color", resp.Error)
}

func TestColor_Batch_HappyPath(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"items":[{"from":"hex","to":"rgb","value":"#ffffff"},{"from":"rgb","to":"hex","value":"0,0,0"}]}`
	req := httptest.NewRequest(http.MethodPost, "/color/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[httpx.BatchResponse[color.BatchColorItem]]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Len(t, resp.Data.Results, 2)
	assert.Equal(t, 2, resp.Data.Total)
	assert.NotNil(t, resp.Data.Results[0].Result)
	assert.Empty(t, resp.Data.Results[0].Error)
}

func TestColor_Batch_InvalidItem_InBandError(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"items":[{"from":"hex","to":"rgb","value":"#ffffff"},{"from":"hex","to":"rgb","value":"notahex"}]}`
	req := httptest.NewRequest(http.MethodPost, "/color/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[httpx.BatchResponse[color.BatchColorItem]]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Len(t, resp.Data.Results, 2)
	assert.NotNil(t, resp.Data.Results[0].Result, "valid item should have result")
	assert.NotEmpty(t, resp.Data.Results[1].Error, "invalid item should have in-band error")
}

func TestColor_Batch_EmptyItems422(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/color/batch", strings.NewReader(`{"items":[]}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
