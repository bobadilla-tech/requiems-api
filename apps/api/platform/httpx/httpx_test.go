package httpx_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
)

type testData struct {
	Value string `json:"value"`
}

func (testData) IsData() {}

func TestJSON_WritesSuccessEnvelope(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	httpx.JSON(w, http.StatusCreated, testData{Value: "hello"})

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var resp httpx.Response[testData]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Equal(t, "hello", resp.Data.Value)
	assert.NotEmpty(t, resp.Metadata.Timestamp)
}

func TestError_WritesErrorEnvelope(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	httpx.Error(w, http.StatusNotFound, "not_found", "resource not found")

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var resp httpx.ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Equal(t, "not_found", resp.Error)
	assert.Equal(t, "resource not found", resp.Message)
	assert.NotEmpty(t, resp.Metadata.Timestamp)
}
