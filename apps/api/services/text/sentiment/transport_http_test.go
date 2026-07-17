package sentiment

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
)

var testSvc *Service

func TestMain(m *testing.M) {
	svc, err := NewService()
	if err != nil {
		log.Fatalf("failed to init service for tests: %v", err)
	}
	testSvc = svc

	os.Exit(m.Run())
}

func setupRouter() chi.Router {
	r := chi.NewRouter()
	RegisterRoutes(r, testSvc)
	return r
}

func TestSentimentHandler_Positive(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"text":"I love this product! It's amazing."}`
	req := httptest.NewRequest(http.MethodPost, "/sentiment", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Result]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "positive", resp.Data.Sentiment)
	assert.True(t, resp.Data.Score > 0 && resp.Data.Score <= 1, "score out of range [0,1]: %.2f", resp.Data.Score)
}

func TestSentimentHandler_Negative(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"text":"This is terrible and I hate it."}`
	req := httptest.NewRequest(http.MethodPost, "/sentiment", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Result]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "negative", resp.Data.Sentiment)
}

func TestSentimentHandler_MissingTextField(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/sentiment", strings.NewReader(`{}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestSentimentHandler_MissingBody(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/sentiment", http.NoBody)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSentimentBatchHandler_OK(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"texts":["I love this!","This is terrible.","The table is there."]}`
	req := httptest.NewRequest(http.MethodPost, "/sentiment/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[httpx.BatchResponse[Result]]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, 3, resp.Data.Total)
	assert.Len(t, resp.Data.Results, 3)
	assert.Equal(t, "positive", resp.Data.Results[0].Sentiment)
	assert.Equal(t, "negative", resp.Data.Results[1].Sentiment)
}

func TestSentimentBatchHandler_EmptyTexts(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/sentiment/batch", strings.NewReader(`{"texts":[]}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestSentimentBatchHandler_MissingBody(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/sentiment/batch", http.NoBody)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSentimentBatchHandler_OversizeTexts(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	texts := make([]string, 51)
	for i := range texts {
		texts[i] = "some text"
	}
	bodyBytes, _ := json.Marshal(map[string][]string{"texts": texts})
	req := httptest.NewRequest(http.MethodPost, "/sentiment/batch", strings.NewReader(string(bodyBytes)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
