package horoscope

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

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

func TestHoroscope_ValidSign(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/horoscope/aries", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Horoscope]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	h := resp.Data

	assert.Equal(t, "aries", h.Sign)
	assert.Equal(t, time.Now().UTC().Format("2006-01-02"), h.Date)
	assert.NotEmpty(t, h.Horoscope)
	assert.True(t, h.LuckyNumber >= 1 && h.LuckyNumber <= 99, "expected lucky_number between 1 and 99, got %d", h.LuckyNumber)
	assert.NotEmpty(t, h.Mood)
}

func TestHoroscope_AllSigns(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	validSigns := []string{
		"aries", "taurus", "gemini", "cancer", "leo", "virgo",
		"libra", "scorpio", "sagittarius", "capricorn", "aquarius", "pisces",
	}

	for _, sign := range validSigns {
		t.Run(sign, func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequest(http.MethodGet, "/horoscope/"+sign, http.NoBody)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code, "expected status 200 for sign %q, got %d", sign, w.Code)
		})
	}
}

func TestHoroscope_InvalidSign(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/horoscope/invalid", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHoroscope_CaseInsensitive(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/horoscope/ARIES", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "expected status 200 for uppercase sign, got %d", w.Code)

	var resp2 httpx.Response[Horoscope]
	err := json.NewDecoder(w.Body).Decode(&resp2)
	require.NoError(t, err)

	assert.Equal(t, "aries", resp2.Data.Sign)
}

func TestHoroscope_DailyConsistency(t *testing.T) {
	t.Parallel()
	svc := NewService()

	h1, err := svc.Daily("leo")
	require.NoError(t, err)

	h2, err := svc.Daily("leo")
	require.NoError(t, err)

	assert.Equal(t, h1.Horoscope, h2.Horoscope, "expected same horoscope for same sign on same day")
	assert.Equal(t, h1.LuckyNumber, h2.LuckyNumber, "expected same lucky_number for same sign on same day")
	assert.Equal(t, h1.Mood, h2.Mood, "expected same mood for same sign on same day")
}

func TestHoroscope_DailyBatch_HappyPath(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	signs := []string{"taurus", "aries", "gemini", "cancer"}

	bodyJSON, err := json.Marshal(BatchRequest{Signs: signs})
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/horoscope/batch", strings.NewReader(string(bodyJSON)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	require.Contains(t, w.Header().Get("Content-Type"), "application/json")

	var resp httpx.Response[httpx.BatchResponse[Horoscope]]
	err = json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	require.Equal(t, 4, resp.Data.Total)
	require.Len(t, resp.Data.Results, 4)
}

func TestHoroscope_DailyBatch_EmptySigns(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	bodyJSON, err := json.Marshal(BatchRequest{Signs: []string{}})
	require.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/horoscope/batch", strings.NewReader(string(bodyJSON)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestHoroscope_DailyBatch_ExceedsLimit(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	signs := make([]string, 13)

	for i := range signs {
		signs[i] = "cancer"
	}

	signsJSON, err := json.Marshal(BatchRequest{
		Signs: signs,
	})
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/horoscope/batch", strings.NewReader(string(signsJSON)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusUnprocessableEntity, w.Code)

}

func TestHoroscope_DailyBatch_SetUsageCountHeader(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	signs := []string{"taurus", "aries", "gemini", "cancer"}

	bodyJSON, err := json.Marshal(BatchRequest{Signs: signs})
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/horoscope/batch", strings.NewReader(string(bodyJSON)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	require.Equal(t, "4", w.Header().Get("X-Usage-Count"))
}
