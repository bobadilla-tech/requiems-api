package lorem

import (
	"net/http"
	"net/http/httptest"
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
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "invalid param not number",
			url:        "/lorem?paragraphs=abc",
			wantStatus: http.StatusBadRequest,
		},
	}

	// Run each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			svc := &mockService{result: tt.mockResult}

			req := httptest.NewRequest(http.MethodGet, tt.url, http.NoBody)
			rec := httptest.NewRecorder()

			newRouter(svc).ServeHTTP(rec, req)
			assert.Equal(t, tt.wantStatus, rec.Code)
		})
	}
}
