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

type dummyService struct{}

func TestGuard_NilService_Returns500(t *testing.T) {
	t.Parallel()

	var svc *dummyService

	handler := httpx.Guard(svc, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
	w := httptest.NewRecorder()
	handler(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code)

	var resp httpx.ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "internal_error", resp.Error)
}

func TestGuard_NonNilService_CallsHandler(t *testing.T) {
	t.Parallel()

	svc := &dummyService{}

	handler := httpx.Guard(svc, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
	w := httptest.NewRecorder()
	handler(w, req)

	require.Equal(t, http.StatusOK, w.Code)
}
