package jokes

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
	svc := NewService()
	RegisterRoutes(r, svc)
	return r
}

func TestDadJoke_Random(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/jokes/dad", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[DadJoke]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	j := resp.Data
	assert.NotEmpty(t, j.Joke)
	assert.NotEmpty(t, j.ID)
	assert.True(t, strings.HasPrefix(j.ID, "joke_"), "expected id to start with 'joke_', got %q", j.ID)
}

func TestDadJoke_Random_MultipleCallsReturnValidJokes(t *testing.T) {
	t.Parallel()
	svc := NewService()

	for range 10 {
		j := svc.Random()
		assert.NotEmpty(t, j.Joke)
		assert.NotEmpty(t, j.ID)
		assert.True(t, strings.HasPrefix(j.ID, "joke_"), "expected id to start with 'joke_', got %q", j.ID)
	}
}

func TestDadJoke_Batch_ReturnsCount(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/jokes/dad/batch", strings.NewReader(`{"count":5}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[httpx.BatchResponse[DadJoke]]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Len(t, resp.Data.Results, 5)
	assert.Equal(t, 5, resp.Data.Total)
	for _, j := range resp.Data.Results {
		assert.NotEmpty(t, j.Joke)
		assert.True(t, strings.HasPrefix(j.ID, "joke_"), "expected ID prefix 'joke_', got %q", j.ID)
	}
}

func TestDadJoke_Batch_ValidationError(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/jokes/dad/batch", strings.NewReader(`{"count":0}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
