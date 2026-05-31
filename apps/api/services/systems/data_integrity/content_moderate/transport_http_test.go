package contentmoderate

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
	"requiems-api/services/text/detectlanguage"
	"requiems-api/services/text/sentiment"
	"requiems-api/services/validation/profanity"
)

func setupRouter() chi.Router {
	r := chi.NewRouter()
	svc := NewService(
		detectlanguage.NewService(),
		profanity.NewService(),
		sentiment.NewService(),
	)
	RegisterRoutes(r, svc)
	return r
}

func post(t *testing.T, r chi.Router, body string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodPost, "/content/moderate", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestModerate_HappyPath(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	w := post(t, r, `{"text": "i love you", "language": "en"}`)
	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Response]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))

	require.NotNil(t, resp.Data.LanguageConfidence)
	assert.Equal(t, 1.0, *resp.Data.LanguageConfidence)
	assert.True(t, resp.Data.IsSafe)
	assert.False(t, resp.Data.Categories.Profanity)
	assert.Empty(t, resp.Data.Flags)
}

func TestModerate_EmptyText(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	w := post(t, r, `{"text": "", "language": "en"}`)
	require.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestModerate_LanguageAutoDetected(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	w := post(t, r, `{"text": "i love you"}`)
	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Response]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))

	require.NotNil(t, resp.Data.LanguageConfidence)
	assert.Greater(t, *resp.Data.LanguageConfidence, 0.0)
	assert.LessOrEqual(t, *resp.Data.LanguageConfidence, 1.0)

	require.NotNil(t, resp.Data.Language)
	assert.Equal(t, "en", *resp.Data.Language)
}

func TestModerate_TextWithProfanity(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	w := post(t, r, `{"text": "fuck you"}`)
	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Response]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))

	assert.False(t, resp.Data.IsSafe)
	assert.True(t, resp.Data.Categories.Profanity)
	assert.Contains(t, resp.Data.Flags, "text_profanity")
}

func TestModerate_EmptyBody(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/content/moderate", http.NoBody)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
