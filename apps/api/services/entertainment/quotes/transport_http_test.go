package quotes

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

// reuse mockRow from service_test.go

type httpMockQuerier struct {
	row pgx.Row
}

func (m *httpMockQuerier) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return m.row
}

func TestTransport_HappyPath(t *testing.T) {
	mockRow := &mockRow{
		scanFn: func(dest ...any) error {
			*dest[0].(*int) = 1
			*dest[1].(*string) = "Hello"
			*dest[2].(*string) = "Test"
			return nil
		},
	}

	svc := &Service{
		db: &httpMockQuerier{row: mockRow},
	}

	r := chi.NewRouter()
	RegisterRoutes(r, svc)

	req := httptest.NewRequest("GET", "/quotes/random", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Result().StatusCode)
	}
}

func TestTransport_Error(t *testing.T) {
	mockRow := &mockRow{
		scanFn: func(dest ...any) error {
			return pgx.ErrNoRows
		},
	}

	svc := &Service{
		db: &httpMockQuerier{row: mockRow},
	}

	r := chi.NewRouter()
	RegisterRoutes(r, svc)

	req := httptest.NewRequest("GET", "/quotes/random", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Result().StatusCode != http.StatusServiceUnavailable {
		t.Errorf("expected 503, got %d", w.Result().StatusCode)
	}
}

func TestTransport_MethodNotAllowed(t *testing.T) {
	svc := &Service{}

	r := chi.NewRouter()
	RegisterRoutes(r, svc)

	req := httptest.NewRequest("POST", "/quotes/random", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Result().StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("expected 405")
	}
}