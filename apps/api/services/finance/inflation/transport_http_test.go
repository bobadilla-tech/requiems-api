package inflation

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
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
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	return resp
}

func decodeBatchResponse(t *testing.T, w *httptest.ResponseRecorder) httpx.Response[BatchResponse] {
	t.Helper()
	var resp httpx.Response[BatchResponse]
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode batch response: %v", err)
	}
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
	svc := &stubGetter{result: Response{
		Rate:       3.2,
		Period:     "2024",
		Historical: []HistoricalRate{{Period: "2023", Rate: 4.1}},
	}}

	r := setupRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/inflation?country=US", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}

	resp := decodeResponse(t, w)
	if resp.Data.Country != "US" {
		t.Errorf("expected country US, got %q", resp.Data.Country)
	}
	if resp.Metadata.Timestamp == "" {
		t.Error("expected metadata.timestamp to be set")
	}
}

func TestInflation_ResponseEnvelope(t *testing.T) {
	svc := &stubGetter{result: Response{Rate: 2.5, Period: "2024"}}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/inflation?country=DE", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var raw map[string]json.RawMessage
	if err := json.NewDecoder(w.Body).Decode(&raw); err != nil {
		t.Fatalf("failed to decode: %v", err)
	}
	if _, ok := raw["data"]; !ok {
		t.Error("response must have a 'data' key")
	}
	if _, ok := raw["metadata"]; !ok {
		t.Error("response must have a 'metadata' key")
	}
}

func TestInflation_UnknownCountry_Returns404(t *testing.T) {
	svc := &stubGetter{err: &httpx.AppError{
		Status:  http.StatusNotFound,
		Code:    "not_found",
		Message: "no inflation data found for country",
	}}

	r := setupRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/inflation?country=XK", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d: %s", w.Code, w.Body.String())
	}
}

func TestInflation_MissingCountryParam_Returns400(t *testing.T) {
	svc := &stubGetter{}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/inflation", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d: %s", w.Code, w.Body.String())
	}
}

func TestInflation_InvalidCountryCode_Returns400(t *testing.T) {
	svc := &stubGetter{}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/inflation?country=ZZZ", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d: %s", w.Code, w.Body.String())
	}
}

func TestInflation_LowercaseCountryAccepted(t *testing.T) {
	svc := &stubGetter{result: Response{Rate: 1.5, Period: "2024"}}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/inflation?country=us", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 for lowercase country, got %d: %s", w.Code, w.Body.String())
	}

	resp := decodeResponse(t, w)
	if resp.Data.Country != "US" {
		t.Errorf("expected country US (uppercased), got %q", resp.Data.Country)
	}
}

func TestInflation_HistoricalFieldPresent(t *testing.T) {
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
	if len(resp.Data.Historical) != 2 {
		t.Errorf("expected 2 historical entries, got %d", len(resp.Data.Historical))
	}
}

// ---- batch endpoint tests ----

func TestBatch_HappyPath_Returns200(t *testing.T) {
	// All three countries exist — expect 200 and three found: true items.
	svc := &stubGetter{result: Response{Rate: 3.2, Period: "2024"}}
	r := setupRouter(svc)

	w := postBatch(r, `{"countries": ["US", "AR", "DE"]}`)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}

	resp := decodeBatchResponse(t, w)
	if resp.Data.Total != 3 {
		t.Errorf("expected total 3, got %d", resp.Data.Total)
	}
	for _, item := range resp.Data.Results {
		if !item.Found {
			t.Errorf("expected found: true for %s, got false", item.Country)
		}
	}
}

func TestBatch_PartialFailure_NotFoundItemIsInBand(t *testing.T) {
	// Stub returns error for every call — all items should be found: false, but status is still 200.
	svc := &stubGetter{err: &httpx.AppError{
		Status:  http.StatusNotFound,
		Code:    "not_found",
		Message: "no inflation data found for country",
	}}
	r := setupRouter(svc)

	w := postBatch(r, `{"countries": ["US", "AR"]}`)

	// Batch never returns 404 — not_found is handled per item.
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 even when countries are not found, got %d: %s", w.Code, w.Body.String())
	}

	resp := decodeBatchResponse(t, w)
	for _, item := range resp.Data.Results {
		if item.Found {
			t.Errorf("expected found: false for %s, got true", item.Country)
		}
	}
}

func TestBatch_OrderPreserved(t *testing.T) {
	// Results must come back in the same order as the input array.
	svc := &stubGetter{result: Response{Rate: 1.0, Period: "2024"}}
	r := setupRouter(svc)

	w := postBatch(r, `{"countries": ["DE", "AR", "US"]}`)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	resp := decodeBatchResponse(t, w)
	expected := []string{"DE", "AR", "US"}
	for i, item := range resp.Data.Results {
		if item.Country != expected[i] {
			t.Errorf("position %d: expected %s, got %s", i, expected[i], item.Country)
		}
	}
}

func TestBatch_EmptyArray_Returns422(t *testing.T) {
	// An empty countries array must be rejected before hitting the service.
	svc := &stubGetter{}
	r := setupRouter(svc)

	w := postBatch(r, `{"countries": []}`)

	if w.Code != http.StatusUnprocessableEntity {
		t.Fatalf("expected 422, got %d: %s", w.Code, w.Body.String())
	}
}

func TestBatch_OverLimit_Returns422(t *testing.T) {
	// 51 countries exceeds the max of 50 — must be rejected.
	countries := make([]string, 51)
	for i := range countries {
		countries[i] = `"US"`
	}
	body := `{"countries": [` + strings.Join(countries, ",") + `]}`
	svc := &stubGetter{}
	r := setupRouter(svc)

	w := postBatch(r, body)

	if w.Code != http.StatusUnprocessableEntity {
		t.Fatalf("expected 422 for 51 countries, got %d: %s", w.Code, w.Body.String())
	}
}

func TestBatch_InvalidCountryCode_Returns422(t *testing.T) {
	// ZZZ is not a valid iso3166_1_alpha2 code — must be rejected.
	svc := &stubGetter{}
	r := setupRouter(svc)

	w := postBatch(r, `{"countries": ["US", "ZZZ"]}`)

	if w.Code != http.StatusUnprocessableEntity {
		t.Fatalf("expected 422 for invalid country code, got %d: %s", w.Code, w.Body.String())
	}
}

func TestBatch_TotalMatchesInput(t *testing.T) {
	// total in the response must always equal the number of countries sent.
	svc := &stubGetter{result: Response{Rate: 2.0, Period: "2024"}}
	r := setupRouter(svc)

	w := postBatch(r, `{"countries": ["US", "DE"]}`)

	resp := decodeBatchResponse(t, w)
	if resp.Data.Total != 2 {
		t.Errorf("expected total 2, got %d", resp.Data.Total)
	}
}
