package convformat_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
	convformat "requiems-api/services/technology/format"
)

func setupRouter() chi.Router {
	r := chi.NewRouter()
	convformat.RegisterRoutes(r, convformat.NewService())
	return r
}

func TestFormat_HappyPath_JSONToYAML(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"from":"json","to":"yaml","content":"{\"name\":\"Alice\",\"age\":30}"}`
	req := httptest.NewRequest(http.MethodPost, "/format", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[convformat.Response]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.True(t, strings.Contains(resp.Data.Result, "Alice"), "expected YAML with 'Alice', got %q", resp.Data.Result)
}

func TestFormat_HappyPath_CSVToJSON(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"from":"csv","to":"json","content":"name,age\nAlice,30\n"}`
	req := httptest.NewRequest(http.MethodPost, "/format", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[convformat.Response]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.True(t, strings.Contains(resp.Data.Result, "Alice"), "expected JSON with 'Alice', got %q", resp.Data.Result)
}

func TestFormat_InvalidFromFormat(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"from":"txt","to":"json","content":"hello"}`
	req := httptest.NewRequest(http.MethodPost, "/format", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestFormat_MissingBody(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/format", http.NoBody)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestFormat_MalformedInput(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"from":"json","to":"yaml","content":"{invalid json"}`
	req := httptest.NewRequest(http.MethodPost, "/format", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestFormat_MissingFields(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"from":"json","to":"yaml"}`
	req := httptest.NewRequest(http.MethodPost, "/format", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestFormat_Batch_HappyPath(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"items":[{"from":"json","to":"yaml","content":"{\"name\":\"Alice\"}"},{"from":"yaml","to":"json","content":"name: Bob\n"}]}`
	req := httptest.NewRequest(http.MethodPost, "/format/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[httpx.BatchResponse[convformat.BatchFormatItem]]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Len(t, resp.Data.Results, 2)
	assert.Equal(t, 2, resp.Data.Total)
	assert.True(t, strings.Contains(resp.Data.Results[0].Result, "Alice"), "expected YAML with 'Alice'")
	assert.Empty(t, resp.Data.Results[0].Error)
}

func TestFormat_Batch_EmptyItems422(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/format/batch", strings.NewReader(`{"items":[]}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
