package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
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

func TestOpenAPISpec_DescribesErrorBilling(t *testing.T) {
	t.Parallel()

	var spec struct {
		Info struct {
			Description string `json:"description"`
		} `json:"info"`
	}
	require.NoError(t, json.Unmarshal(openapiSpec, &spec))

	assert.Contains(t, spec.Info.Description, "Client (4xx) and server (5xx) errors are not charged")
	assert.Contains(t, spec.Info.Description, "unknown routes or with an unsupported method are not recorded")
}

// The MCP server is generated from its own copy of the spec; it must not drift
// from the one this API serves.
func TestOpenAPISpec_MatchesTheMCPCopy(t *testing.T) {
	t.Parallel()

	mcp, err := os.ReadFile("../../mcp/openapi.json")
	if err != nil {
		if os.Getenv("CI") == "" {
			t.Skipf("MCP spec is not in this checkout: %v", err)
		}

		require.NoError(t, err)
	}

	assert.True(t, string(mcp) == string(openapiSpec), "apps/mcp/openapi.json differs from apps/api/app/openapi.json")
}
