package profanity

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

func TestProfanity_CleanText(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"text":"Hello, world!"}`
	req := httptest.NewRequest(http.MethodPost, "/profanity", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Result]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.False(t, resp.Data.HasProfanity)
	assert.Equal(t, "Hello, world!", resp.Data.Censored)
	assert.Empty(t, resp.Data.FlaggedWords)
}

func TestProfanity_ProfaneText(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"text":"What the fuck is this shit"}`
	req := httptest.NewRequest(http.MethodPost, "/profanity", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Result]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.True(t, resp.Data.HasProfanity)
	assert.Equal(t, "What the **** is this ****", resp.Data.Censored)
	assert.Len(t, resp.Data.FlaggedWords, 2)
	// Verify the specific words detected
	found := map[string]bool{}
	for _, w := range resp.Data.FlaggedWords {
		found[w] = true
	}
	if !found["fuck"] || !found["shit"] {
		t.Errorf("expected flagged words [\"fuck\", \"shit\"], got %v", resp.Data.FlaggedWords)
	}
}

func TestProfanity_MissingTextField(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/profanity", strings.NewReader(`{}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestProfanity_MissingBody(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/profanity", http.NoBody)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
