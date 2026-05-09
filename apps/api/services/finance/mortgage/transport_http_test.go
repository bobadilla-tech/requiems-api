package mortgage

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
)

// stubCalculator implements Calculator for transport tests.
type stubCalculator struct {
	result Response
}

func (s *stubCalculator) Calculate(principal, annualRate float64, years int) Response {
	r := s.result
	r.Principal = principal
	r.Rate = annualRate
	r.Years = years
	return r
}

func setupRouter(c Calculator) chi.Router {
	r := chi.NewRouter()
	registerMortgageRoutes(r, c)
	return r
}

func decodeResponse(t *testing.T, w *httptest.ResponseRecorder) httpx.Response[Response] {
	t.Helper()
	var resp httpx.Response[Response]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	return resp
}

func TestMortgage_HappyPath_Returns200(t *testing.T) {
	t.Parallel()
	svc := &stubCalculator{result: Response{MonthlyPayment: 1896.20}}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/mortgage?principal=300000&rate=6.5&years=30", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	resp := decodeResponse(t, w)
	assert.Equal(t, float64(300000), resp.Data.Principal)
	assert.Equal(t, 6.5, resp.Data.Rate)
	assert.Equal(t, 30, resp.Data.Years)
}

func TestMortgage_ResponseEnvelope(t *testing.T) {
	t.Parallel()
	svc := &stubCalculator{}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/mortgage?principal=100000&rate=5.0&years=15", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var raw map[string]json.RawMessage
	err := json.NewDecoder(w.Body).Decode(&raw)
	require.NoError(t, err)
	_, ok := raw["data"]
	assert.True(t, ok, "response must have a 'data' key")
	_, ok = raw["metadata"]
	assert.True(t, ok, "response must have a 'metadata' key")
}

func TestMortgage_MissingPrincipal_Returns400(t *testing.T) {
	t.Parallel()
	svc := &stubCalculator{}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/mortgage?rate=6.5&years=30", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestMortgage_MissingRate_Returns400(t *testing.T) {
	t.Parallel()
	svc := &stubCalculator{}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/mortgage?principal=300000&years=30", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestMortgage_MissingYears_Returns400(t *testing.T) {
	t.Parallel()
	svc := &stubCalculator{}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/mortgage?principal=300000&rate=6.5", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestMortgage_YearsZero_Returns400(t *testing.T) {
	t.Parallel()
	svc := &stubCalculator{}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/mortgage?principal=300000&rate=6.5&years=0", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestMortgage_YearsExceedsMax_Returns400(t *testing.T) {
	t.Parallel()
	svc := &stubCalculator{}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/mortgage?principal=300000&rate=6.5&years=51", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestMortgage_MetadataTimestampSet(t *testing.T) {
	t.Parallel()
	svc := &stubCalculator{}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/mortgage?principal=200000&rate=4.5&years=20", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := decodeResponse(t, w)
	assert.NotEmpty(t, resp.Metadata.Timestamp)
}
