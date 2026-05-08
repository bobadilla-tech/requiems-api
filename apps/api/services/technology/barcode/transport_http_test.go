package barcode

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

// ── GET /barcode (PNG) ─────────────────────────────────────────────────────

func TestBarcode_PNG_Code128(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/barcode?data=123456789&type=code128", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "image/png", w.Header().Get("Content-Type"))

	body := w.Body.Bytes()
	assert.NotEmpty(t, body)
	assert.Equal(t, "\x89PNG", string(body[:4]))
}

func TestBarcode_PNG_AllTypes(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	tests := []struct {
		name  string
		query string
	}{
		{"code128", "/barcode?data=HELLO123&type=code128"},
		{"code93", "/barcode?data=HELLO&type=code93"},
		{"code39", "/barcode?data=HELLO&type=code39"},
		{"ean8", "/barcode?data=1234567&type=ean8"},
		{"ean13", "/barcode?data=123456789012&type=ean13"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequest(http.MethodGet, tc.query, http.NoBody)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code, "type=%q: expected status 200, got %d", tc.name, w.Code)
		})
	}
}

func TestBarcode_PNG_MissingData(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/barcode?type=code128", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestBarcode_PNG_MissingType(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/barcode?data=123456789", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestBarcode_PNG_InvalidType(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/barcode?data=123456789&type=invalid", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestBarcode_PNG_InvalidEAN8Data(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/barcode?data=123&type=ean8", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

// ── GET /barcode/base64 (JSON) ─────────────────────────────────────────────

func TestBarcode_Base64_Code128(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/barcode/base64?data=123456789&type=code128", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var resp httpx.Response[Base64Response]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.NotEmpty(t, resp.Data.Image)
	assert.Equal(t, "code128", resp.Data.Type)
	assert.Equal(t, defaultWidth, resp.Data.Width)
	assert.Equal(t, defaultHeight, resp.Data.Height)

	// Verify the base64 string decodes to valid PNG bytes.
	decoded, err := base64.StdEncoding.DecodeString(resp.Data.Image)
	require.NoError(t, err)

	assert.Equal(t, "\x89PNG", string(decoded[:4]))
}

func TestBarcode_Base64_MissingData(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/barcode/base64?type=code128", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestBarcode_Base64_InvalidType(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/barcode/base64?data=123456789&type=invalid", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
