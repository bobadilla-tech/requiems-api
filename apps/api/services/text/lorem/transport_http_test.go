package lorem

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

type mockService struct {
	result Lorem
}

func (m *mockService) Generate(paragraphs, sentences int) Lorem {
	return m.result
}

func newRouter(svc Generator) chi.Router {
	r := chi.NewRouter()
	RegisterRoutes(r, svc)
	return r
}

func TestGetLorem(t *testing.T) {
	t.Parallel()
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
			wantStatus: http.StatusUnprocessableEntity,
		},
		{
			name:       "invalid param not number",
			url:        "/lorem?paragraphs=abc",
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			svc := &mockService{result: tt.mockResult}

			// Volvimos a poner http.NoBody acá como pidió el revisor
			req := httptest.NewRequest(http.MethodGet, tt.url, http.NoBody)
			rec := httptest.NewRecorder()

			newRouter(svc).ServeHTTP(rec, req)
			assert.Equal(t, tt.wantStatus, rec.Code)
		})
	}
}

func TestPostBatchLorem(t *testing.T) {
	t.Parallel()

	svc := &mockService{
		result: Lorem{Text: "batch text", Paragraphs: 1, WordCount: 2},
	}
	r := newRouter(svc)

	// Generamos un JSON con 51 ítems para probar el límite máximo que exige el RFC
	var oversizeItems []string
	for range 51 {
		oversizeItems = append(oversizeItems, `{"paragraphs": 1}`)
	}
	oversizeJSON := `{"items": [` + strings.Join(oversizeItems, ",") + `]}`

	tests := []struct {
		name       string
		body       string
		wantStatus int
	}{
		{
			name: "happy path - valid batch",
			// ¡Acá agregamos el {} al final para el coverage del 100%!
			body:       `{"items": [{"paragraphs": 2, "sentences": 3}, {"paragraphs": 1}, {}]}`,
			wantStatus: http.StatusOK,
		},
		{
			name:       "validation error - empty array",
			body:       `{"items": []}`,
			wantStatus: http.StatusUnprocessableEntity,
		},
		{
			name:       "validation error - missing items key",
			body:       `{}`,
			wantStatus: http.StatusUnprocessableEntity,
		},
		{
			name:       "validation error - oversize batch",
			body:       oversizeJSON,
			wantStatus: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/lorem/batch", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			r.ServeHTTP(rec, req)

			assert.Equal(t, tt.wantStatus, rec.Code)
		})
	}
}
