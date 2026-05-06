package holidays

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

func setupRouter(t *testing.T) chi.Router {
	t.Helper()
	svc := NewService()
	r := chi.NewRouter()
	RegisterRoutes(r, svc)
	return r
}

func TestHolidays_ValidRequest(t *testing.T) {
	r := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/holidays?country=US&year=2025", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d: %s", w.Code, w.Body.String())
	}

	var resp httpx.Response[HolidayList]
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Data.Country != "US" {
		t.Errorf("expected country 'US', got %q", resp.Data.Country)
	}
	if resp.Data.Year != 2025 {
		t.Errorf("expected year 2025, got %d", resp.Data.Year)
	}
	if len(resp.Data.Holidays) == 0 {
		t.Error("expected non-empty holidays list")
	}
}

func TestHolidays_UK(t *testing.T) {
	r := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/holidays?country=GB&year=2025", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d: %s", w.Code, w.Body.String())
	}

	var resp httpx.Response[HolidayList]
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Data.Country != "GB" {
		t.Errorf("expected country 'GB', got %q", resp.Data.Country)
	}
}

func TestHolidays_MissingCountry(t *testing.T) {
	r := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/holidays?year=2025", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}

func TestHolidays_MissingYear(t *testing.T) {
	r := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/holidays?country=US", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}

func TestHolidays_InvalidCountry(t *testing.T) {
	r := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/holidays?country=INVALID&year=2025", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}

func TestHolidays_InvalidYear(t *testing.T) {
	r := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/holidays?country=US&year=0", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}

func TestHolidays_NoParams(t *testing.T) {
	r := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/holidays", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}

func TestHolidays_LowercaseCountry(t *testing.T) {
	r := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/holidays?country=us&year=2025", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d: %s", w.Code, w.Body.String())
	}

	var resp httpx.Response[HolidayList]
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Data.Country != "US" {
		t.Errorf("expected country 'US', got %q", resp.Data.Country)
	}
}

// ---- batch endpoint tests ----

func TestHolidaysBatch_ValidRequest(t *testing.T) {
	r := setupRouter(t)

	body := `{"queries":[{"country":"US","year":2025},{"country":"GB","year":2025}]}`
	req := httptest.NewRequest(http.MethodPost, "/holidays/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d: %s", w.Code, w.Body.String())
	}

	var resp httpx.Response[BatchResponse]
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode batch response: %v", err)
	}

	if resp.Data.Total != 2 {
		t.Errorf("expected total 2, got %d", resp.Data.Total)
	}
	for _, item := range resp.Data.Results {
		if !item.Found {
			t.Errorf("expected found=true for %s %d", item.Country, item.Year)
		}
	}
}

func TestHolidaysBatch_SetsUsageCountHeader(t *testing.T) {
	r := setupRouter(t)

	body := `{"queries":[{"country":"US","year":2025},{"country":"DE","year":2025}]}`
	req := httptest.NewRequest(http.MethodPost, "/holidays/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
	if got := w.Header().Get("X-Usage-Count"); got != "2" {
		t.Errorf("expected X-Usage-Count: 2, got %q", got)
	}
}

func TestHolidaysBatch_PartialFailure(t *testing.T) {
	r := setupRouter(t)

	// AQ (Antarctica) is a valid ISO 3166-1 alpha-2 code but has no holiday data,
	// so it passes input validation and triggers the found:false path in the service.
	body := `{"queries":[{"country":"US","year":2025},{"country":"AQ","year":2025}]}`
	req := httptest.NewRequest(http.MethodPost, "/holidays/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d: %s", w.Code, w.Body.String())
	}

	var resp httpx.Response[BatchResponse]
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode batch response: %v", err)
	}

	if resp.Data.Results[0].Found != true {
		t.Error("expected Results[0] (US) to be found")
	}
	if resp.Data.Results[1].Found != false {
		t.Error("expected Results[1] (AQ) to be not found")
	}
}

func TestHolidaysBatch_EmptyQueries(t *testing.T) {
	r := setupRouter(t)

	body := `{"queries":[]}`
	req := httptest.NewRequest(http.MethodPost, "/holidays/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected status 422, got %d", w.Code)
	}
}

func TestHolidaysBatch_MissingBody(t *testing.T) {
	r := setupRouter(t)

	req := httptest.NewRequest(http.MethodPost, "/holidays/batch", strings.NewReader(`{}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected status 422, got %d", w.Code)
	}
}
