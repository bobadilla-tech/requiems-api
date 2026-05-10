package inflation

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
	"requiems-api/platform/svcerr"
)

// stubGetter implements Getter for transport tests. It returns a fixed result
// or a fixed error on every call, keeping tests DB-free and fast.
type stubGetter struct {
	result Response
	err    error
}

func (s *stubGetter) GetInflation(_ context.Context, countryCode string) (Response, error) {
	if s.err != nil {
		return Response{}, s.err
	}
	r := s.result
	r.Country = countryCode
	return r, nil
}

// GetInflationBatch delegates to GetInflation per item, matching real service behaviour.
func (s *stubGetter) GetInflationBatch(ctx context.Context, countries []string) BatchResponse {
	results := make([]BatchItem, len(countries))
	for i, c := range countries {
		resp, err := s.GetInflation(ctx, c)
		if err != nil {
			results[i] = BatchItem{Country: strings.ToUpper(c), Found: false}
		} else {
			results[i] = BatchItem{
				Country:    resp.Country,
				Found:      true,
				Rate:       resp.Rate,
				Period:     resp.Period,
				Historical: resp.Historical,
			}
		}
	}
	return BatchResponse{Results: results, Total: len(results)}
}

// setupRouter wires up a stub getter into a chi router for handler testing.
func setupRouter(g Getter) chi.Router {
	r := chi.NewRouter()
	registerInflationRoutes(r, g)
	return r
}

// ---- helpers ----

func decodeResponse(t *testing.T, w *httptest.ResponseRecorder) httpx.Response[Response] {
	t.Helper()
	var resp httpx.Response[Response]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	return resp
}

func decodeBatchResponse(t *testing.T, w *httptest.ResponseRecorder) httpx.Response[BatchResponse] {
	t.Helper()
	var resp httpx.Response[BatchResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	return resp
}

func postBatch(r chi.Router, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/inflation/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// ---- single endpoint tests (unchanged) ----

func TestInflation_KnownCountry_Returns200(t *testing.T) {
	t.Parallel()
	svc := &stubGetter{result: Response{
		Rate:       3.2,
		Period:     "2024",
		Historical: []HistoricalRate{{Period: "2023", Rate: 4.1}},
	}}

	r := setupRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/inflation?country=US", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	resp := decodeResponse(t, w)
	assert.Equal(t, "US", resp.Data.Country)
	assert.NotEmpty(t, resp.Metadata.Timestamp)
}

func TestInflation_ResponseEnvelope(t *testing.T) {
	t.Parallel()
	svc := &stubGetter{result: Response{Rate: 2.5, Period: "2024"}}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/inflation?country=DE", http.NoBody)
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

func TestInflation_UnknownCountry_Returns404(t *testing.T) {
	t.Parallel()
	svc := &stubGetter{err: svcerr.NotFound("not_found", "no inflation data found for country")}

	r := setupRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/inflation?country=XK", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
}

func TestInflation_MissingCountryParam_Returns400(t *testing.T) {
	t.Parallel()
	svc := &stubGetter{}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/inflation", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestInflation_InvalidCountryCode_Returns400(t *testing.T) {
	t.Parallel()
	svc := &stubGetter{}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/inflation?country=ZZZ", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestInflation_LowercaseCountryAccepted(t *testing.T) {
	t.Parallel()
	svc := &stubGetter{result: Response{Rate: 1.5, Period: "2024"}}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/inflation?country=us", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	resp := decodeResponse(t, w)
	assert.Equal(t, "US", resp.Data.Country)
}

func TestInflation_HistoricalFieldPresent(t *testing.T) {
	t.Parallel()
	historical := []HistoricalRate{
		{Period: "2023", Rate: 4.1},
		{Period: "2022", Rate: 8.0},
	}
	svc := &stubGetter{result: Response{
		Rate:       3.2,
		Period:     "2024",
		Historical: historical,
	}}

	r := setupRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/inflation?country=US", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := decodeResponse(t, w)
	assert.Len(t, resp.Data.Historical, 2)
}

// ---- batch endpoint tests ----

func TestBatch_HappyPath_Returns200(t *testing.T) {
	t.Parallel()
	// All three countries exist — expect 200 and three found: true items.
	svc := &stubGetter{result: Response{Rate: 3.2, Period: "2024"}}
	r := setupRouter(svc)

	w := postBatch(r, `{"countries": ["US", "AR", "DE"]}`)

	require.Equal(t, http.StatusOK, w.Code)

	resp := decodeBatchResponse(t, w)
	assert.Equal(t, 3, resp.Data.Total)
	for _, item := range resp.Data.Results {
		assert.True(t, item.Found, "expected found: true for %s, got false", item.Country)
	}
}

func TestBatch_PartialFailure_NotFoundItemIsInBand(t *testing.T) {
	t.Parallel()
	// Stub returns error for every call — all items should be found: false, but status is still 200.
	svc := &stubGetter{err: svcerr.NotFound("not_found", "no inflation data found for country")}
	r := setupRouter(svc)

	w := postBatch(r, `{"countries": ["US", "AR"]}`)

	// Batch never returns 404 — not_found is handled per item.
	require.Equal(t, http.StatusOK, w.Code)

	resp := decodeBatchResponse(t, w)
	for _, item := range resp.Data.Results {
		assert.False(t, item.Found, "expected found: false for %s, got true", item.Country)
	}
}

func TestBatch_OrderPreserved(t *testing.T) {
	t.Parallel()
	// Results must come back in the same order as the input array.
	svc := &stubGetter{result: Response{Rate: 1.0, Period: "2024"}}
	r := setupRouter(svc)

	w := postBatch(r, `{"countries": ["DE", "AR", "US"]}`)

	require.Equal(t, http.StatusOK, w.Code)

	resp := decodeBatchResponse(t, w)
	expected := []string{"DE", "AR", "US"}
	for i, item := range resp.Data.Results {
		assert.Equal(t, expected[i], item.Country)
	}
}

func TestBatch_EmptyArray_Returns422(t *testing.T) {
	t.Parallel()
	// An empty countries array must be rejected before hitting the service.
	svc := &stubGetter{}
	r := setupRouter(svc)

	w := postBatch(r, `{"countries": []}`)

	require.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestBatch_OverLimit_Returns422(t *testing.T) {
	t.Parallel()
	// 51 countries exceeds the max of 50 — must be rejected.
	countries := make([]string, 51)
	for i := range countries {
		countries[i] = `"US"`
	}
	body := `{"countries": [` + strings.Join(countries, ",") + `]}`
	svc := &stubGetter{}
	r := setupRouter(svc)

	w := postBatch(r, body)

	require.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestBatch_InvalidCountryCode_Returns422(t *testing.T) {
	t.Parallel()
	// ZZZ is not a valid iso3166_1_alpha2 code — must be rejected.
	svc := &stubGetter{}
	r := setupRouter(svc)

	w := postBatch(r, `{"countries": ["US", "ZZZ"]}`)

	require.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestBatch_TotalMatchesInput(t *testing.T) {
	t.Parallel()
	// total in the response must always equal the number of countries sent.
	svc := &stubGetter{result: Response{Rate: 2.0, Period: "2024"}}
	r := setupRouter(svc)

	w := postBatch(r, `{"countries": ["US", "DE"]}`)

	resp := decodeBatchResponse(t, w)
	assert.Equal(t, 2, resp.Data.Total)
}
