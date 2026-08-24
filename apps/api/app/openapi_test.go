package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpenAPISpec(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(http.MethodGet, "/openapi.json", http.NoBody)
	w := httptest.NewRecorder()
	OpenAPISpec().ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var spec map[string]any
	err := json.NewDecoder(w.Body).Decode(&spec)
	require.NoError(t, err)
	assert.Equal(t, "3.0.3", spec["openapi"])
}
