package useragent

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

func TestUserAgent_HappyPath(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/useragent?ua=Mozilla%2F5.0+%28Windows+NT+10.0%3B+Win64%3B+x64%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Chrome%2F120.0.0.0+Safari%2F537.36", http.NoBody)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Result]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "Chrome", resp.Data.Browser)
	assert.Equal(t, "desktop", resp.Data.Device)
	assert.False(t, resp.Data.IsBot)
}

func TestUserAgent_MissingUA(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/useragent", http.NoBody)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var resp httpx.ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Equal(t, "bad_request", resp.Error)
}

func TestUserAgent_BotDetection(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/useragent?ua=Mozilla%2F5.0+%28compatible%3B+Googlebot%2F2.1%3B+%2Bhttp%3A%2F%2Fwww.google.com%2Fbot.html%29", http.NoBody)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Result]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.True(t, resp.Data.IsBot)
	assert.Equal(t, "bot", resp.Data.Device)
}

func TestUserAgent_BatchParse_HappyPath(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"user_agents":[
		"Mozilla/5.0 (Linux; Android 14; Pixel 8) AppleWebKit/537.36",
        "curl/8.4.0"
	]}`

	req := httptest.NewRequest(http.MethodPost, "/useragent/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	require.Contains(t, w.Header().Get("Content-Type"), "application/json")

	var resp httpx.Response[httpx.BatchResponse[BatchParseItem]]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	require.Equal(t, 2, resp.Data.Total)
	require.Len(t, resp.Data.Results, 2)

	assert.Equal(t, "Mozilla/5.0 (Linux; Android 14; Pixel 8) AppleWebKit/537.36", resp.Data.Results[0].UserAgent)
	assert.Equal(t, "mobile", resp.Data.Results[0].Data.Device)
	assert.False(t, resp.Data.Results[0].Data.IsBot)

	assert.Equal(t, "curl/8.4.0", resp.Data.Results[1].UserAgent)
	assert.True(t, resp.Data.Results[1].Data.IsBot)
}

func TestUserAgent_BatchParse_EmptyUserAgents(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"user_agents":[]}`
	req := httptest.NewRequest(http.MethodPost, "/useragent/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestUserAgent_BatchParse_ExceedsLimit(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	userAgents := make([]string, 51)

	for i := range userAgents {
		userAgents[i] = "curl/8.4.0"
	}

	body, _ := json.Marshal(BatchParseRequest{
		UserAgents: userAgents,
	})

	req := httptest.NewRequest(http.MethodPost, "/useragent/batch", strings.NewReader(string(body)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusUnprocessableEntity, w.Code)

}

func TestUserAgent_BatchParse_SetUsageCountHeader(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"user_agents":["curl/8.4.0","curl/8.4.0","curl/8.4.0"]}`
	req := httptest.NewRequest(http.MethodPost, "/useragent/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	require.Equal(t, "3", w.Header().Get("X-Usage-Count"))
}
