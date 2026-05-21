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
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/profanity", strings.NewReader(`{}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestProfanity_MissingBody(t *testing.T) {
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/profanity", http.NoBody)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestProfanity_Batch_HappyPath(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"texts":["hello","fuck"]}`
	req := httptest.NewRequest(http.MethodPost, "/profanity/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[httpx.BatchResponse[BatchResult]]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Equal(t, 2, resp.Data.Total)
	require.Len(t, resp.Data.Results, 2)
}

func TestProfanity_Batch_EmptyArray(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/profanity/batch", strings.NewReader(`{"texts":[]}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestProfanity_Batch_ExceedsLimit(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	items := make([]string, 51)
	for i := range items {
		items[i] = `"text"`
	}
	body := `{"texts":[` + strings.Join(items, ",") + `]}`
	req := httptest.NewRequest(http.MethodPost, "/profanity/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestProfanity_Batch_OrderPreserved(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"texts":["alpha","fuck","beta"]}`
	req := httptest.NewRequest(http.MethodPost, "/profanity/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[httpx.BatchResponse[BatchResult]]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	require.Len(t, resp.Data.Results, 3)
	assert.Equal(t, "alpha", resp.Data.Results[0].Text)
	assert.Equal(t, "fuck", resp.Data.Results[1].Text)
	assert.Equal(t, "beta", resp.Data.Results[2].Text)
	assert.False(t, resp.Data.Results[0].Result.HasProfanity)
	assert.True(t, resp.Data.Results[1].Result.HasProfanity)
	assert.False(t, resp.Data.Results[2].Result.HasProfanity)
}
