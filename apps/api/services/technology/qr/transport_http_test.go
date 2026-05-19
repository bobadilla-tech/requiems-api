package qr

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
)

func setupRouter() chi.Router {
	r := chi.NewRouter()
	svc := NewService()
	RegisterRoutes(r, svc)
	return r
}

// ── GET /qr (PNG) ──────────────────────────────────────────────────────────

func TestQR_PNG_DefaultSize(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/qr?data=https://example.com", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Equal(t, "image/png", w.Header().Get("Content-Type"))

	body := w.Body.Bytes()
	assert.NotEmpty(t, body)

	if string(body[:4]) != "\x89PNG" {
		t.Error("expected valid PNG signature in response body")
	}
}

func TestQR_PNG_CustomSize(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/qr?data=https://example.com&size=200", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestQR_PNG_RecoveryLevels(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	levels := []string{"low", "medium", "high", "highest"}
	for _, level := range levels {
		req := httptest.NewRequest(http.MethodGet, "/qr?data=test&recovery="+level, http.NoBody)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("recovery=%q: expected status 200, got %d", level, w.Code)
		}
	}
}

func TestQR_PNG_MissingData(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/qr", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestQR_PNG_SizeTooSmall(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/qr?data=test&size=10", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestQR_PNG_SizeTooLarge(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/qr?data=test&size=2000", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestQR_PNG_InvalidRecovery(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/qr?data=test&recovery=ultra", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// ── GET /qr/base64 (JSON) ──────────────────────────────────────────────────

func TestQR_Base64_DefaultSize(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/qr/base64?data=https://example.com", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var resp httpx.Response[Base64Response]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.NotEmpty(t, resp.Data.Image)

	assert.Equal(t, defaultSize, resp.Data.Width)

	assert.Equal(t, defaultSize, resp.Data.Height)

	// Verify the base64 string decodes to valid PNG bytes.
	decoded, err := base64.StdEncoding.DecodeString(resp.Data.Image)
	require.NoError(t, err)

	if string(decoded[:4]) != "\x89PNG" {
		t.Error("expected valid PNG signature in decoded base64 data")
	}
}

func TestQR_Base64_CustomSize(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/qr/base64?data=hello&size=300", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Base64Response]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, 300, resp.Data.Width)

	assert.Equal(t, 300, resp.Data.Height)
}

func TestQR_Base64_RecoveryLevels(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	levels := []string{"low", "medium", "high", "highest"}
	for _, level := range levels {
		req := httptest.NewRequest(http.MethodGet, "/qr/base64?data=test&recovery="+level, http.NoBody)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("recovery=%q: expected status 200, got %d", level, w.Code)
		}
	}
}

func TestQR_Base64_MissingData(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/qr/base64", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestQR_Base64_InvalidRecovery(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/qr/base64?data=test&recovery=ultra", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
