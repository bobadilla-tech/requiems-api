package mortgage

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// stubCalculator implements Calculator for transport tests.
type stubCalculator struct {
	result  Response
	results BatchResponse
}

func (s *stubCalculator) Calculate(principal, annualRate float64, years int) Response {
	r := s.result
	r.Principal = principal
	r.Rate = annualRate
	r.Years = years
	return r
}

func (s *stubCalculator) CalculateBatch(mortgages []Request) BatchResponse {
	r := s.results
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
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	return resp
}

func TestMortgage_HappyPath_Returns200(t *testing.T) {
	svc := &stubCalculator{result: Response{MonthlyPayment: 1896.20}}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/mortgage?principal=300000&rate=6.5&years=30", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}

	resp := decodeResponse(t, w)
	if resp.Data.Principal != 300000 {
		t.Errorf("expected principal 300000, got %v", resp.Data.Principal)
	}
	if resp.Data.Rate != 6.5 {
		t.Errorf("expected rate 6.5, got %v", resp.Data.Rate)
	}
	if resp.Data.Years != 30 {
		t.Errorf("expected years 30, got %v", resp.Data.Years)
	}
}

func TestMortgage_ResponseEnvelope(t *testing.T) {
	svc := &stubCalculator{}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/mortgage?principal=100000&rate=5.0&years=15", http.NoBody)
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

func TestMortgage_MissingPrincipal_Returns400(t *testing.T) {
	svc := &stubCalculator{}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/mortgage?rate=6.5&years=30", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestMortgage_MissingRate_Returns400(t *testing.T) {
	svc := &stubCalculator{}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/mortgage?principal=300000&years=30", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestMortgage_MissingYears_Returns400(t *testing.T) {
	svc := &stubCalculator{}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/mortgage?principal=300000&rate=6.5", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestMortgage_YearsZero_Returns400(t *testing.T) {
	svc := &stubCalculator{}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/mortgage?principal=300000&rate=6.5&years=0", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400 for years=0, got %d", w.Code)
	}
}

func TestMortgage_YearsExceedsMax_Returns400(t *testing.T) {
	svc := &stubCalculator{}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/mortgage?principal=300000&rate=6.5&years=51", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400 for years=51, got %d", w.Code)
	}
}

func TestMortgage_MetadataTimestampSet(t *testing.T) {
	svc := &stubCalculator{}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/mortgage?principal=200000&rate=4.5&years=20", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := decodeResponse(t, w)
	if resp.Metadata.Timestamp == "" {
		t.Error("expected metadata.timestamp to be set")
	}
}

func TestMortgage_BatchCalculate_HappyPath(t *testing.T) {
	svc := &stubCalculator{results: BatchResponse{
		Results: []Response{{MonthlyPayment: 1896.20}, {MonthlyPayment: 1896.20}},
		Total:   2,
	}}

	r := setupRouter(svc)

	body := `{"mortgages":[{"principal":150000,"rate":6.5,"years":2},{"principal":250000,"rate":6.5,"years":5}]}`
	req := httptest.NewRequest(http.MethodPost, "/mortgage/batch", strings.NewReader(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d: %s", w.Code, w.Body.String())
	}

	var resp httpx.Response[BatchResponse]
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Data.Total != 2 {
		t.Errorf("expected total 2, got %v", resp.Data.Total)
	}

	if len(resp.Data.Results) != 2 {
		t.Errorf("expected 2 results, got %d", len(resp.Data.Results))
	}

}

func TestMortgage_BatchCalculate_EmptyBody(t *testing.T) {
	svc := &stubCalculator{}

	r := setupRouter(svc)

	body := `{"mortgages": []}`
	req := httptest.NewRequest(http.MethodPost, "/mortgage/batch", strings.NewReader(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected 422, got %d: %s", w.Code, w.Body.String())
	}
}

func TestMortgage_BatchCalculate_missingBody(t *testing.T) {
	svc := &stubCalculator{}

	r := setupRouter(svc)

	body := `{}`
	req := httptest.NewRequest(http.MethodPost, "/mortgage/batch", strings.NewReader(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected 422, got %d: %s", w.Code, w.Body.String())
	}
}

func TestMortgage_BatchCalculate_ExccedsLimit(t *testing.T) {
	svc := &stubCalculator{}

	r := setupRouter(svc)

	mortgages := make([]Request, 51)

	for i := range mortgages {
		mortgages[i] = Request{Principal: 150000, Rate: 6.5, Years: 10}
	}

	body, _ := json.Marshal(map[string]any{
		"mortgages": mortgages,
	})

	req := httptest.NewRequest(http.MethodPost, "/mortgage/batch", strings.NewReader(string(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected 422, got %d: %s", w.Code, w.Body.String())
	}
}

func TestMortgage_BatchCalculate_SetsUsageCountHeader(t *testing.T) {
	svc := &stubCalculator{}

	r := setupRouter(svc)

	body := `{"mortgages":[{"principal":150000,"rate":6.5,"years":2},{"principal":250000,"rate":6.5,"years":5}]}`
	req := httptest.NewRequest(http.MethodPost, "/mortgage/batch", strings.NewReader(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	if got := w.Header().Get("X-Usage-Count"); got != "2" {
		t.Errorf("Expected X-Usage-Count: 2, got %q", got)
	}
}
