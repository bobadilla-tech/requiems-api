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
