package lorem

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

// Servicio falso
type mockService struct {
	result Lorem
}

// Implementa interfaz, y devuelve m.result
func (m *mockService) Generate(paragraphs, sentences int) Lorem {
	return m.result
}

// aux, arma el router
func newRouter(svc Generator) chi.Router {
	r := chi.NewRouter()
	RegisterRoutes(r, svc)
	return r
}

// slice de casos del tests
func TestGetLorem(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		mockResult Lorem
		wantStatus int
	}{
		{
			name:       "happy path",
			url:        "/lorem?paragraphs=2&sentences=3",
			mockResult: Lorem{Text: "hello", Paragraphs: 2, WordCount: 10},
			wantStatus: http.StatusOK,
		},
		{
			name:       "defaults",
			url:        "/lorem",
			mockResult: Lorem{Text: "default", Paragraphs: 1, WordCount: 5},
			wantStatus: http.StatusOK,
		},
		{
			name:       "invalid param too large",
			url:        "/lorem?paragraphs=999",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "invalid param not number",
			url:        "/lorem?paragraphs=abc",
			wantStatus: http.StatusBadRequest,
		},
	}

	// recorre cada caso y ejecuta una req falsa y captura con rec
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &mockService{result: tt.mockResult}

			req := httptest.NewRequest(http.MethodGet, tt.url, http.NoBody)
			rec := httptest.NewRecorder()

			// simula llamada http
			newRouter(svc).ServeHTTP(rec, req)
			// validacion
			if rec.Code != tt.wantStatus {
				t.Errorf("esperado %d, obtuve %d", tt.wantStatus, rec.Code)
			}
		})
	}
}
