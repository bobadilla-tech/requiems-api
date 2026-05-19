package httpx_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
	"requiems-api/platform/svcerr"
)

// handleReq / handleRes are minimal request and response types used across
// Handle and HandleBatch tests.
type handleReq struct {
	Name string `json:"name" validate:"required"`
}

type handleRes struct {
	Greeting string `json:"greeting"`
}

func TestHandle_HappyPath(t *testing.T) {
	t.Parallel()

	h := httpx.Handle(func(_ context.Context, req handleReq) (handleRes, error) {
		return handleRes{Greeting: "hello " + req.Name}, nil
	})

	body := strings.NewReader(`{"name":"world"}`)
	r := httptest.NewRequest(http.MethodPost, "/", body)
	w := httptest.NewRecorder()
	h(w, r)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[handleRes]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Equal(t, "hello world", resp.Data.Greeting)
}

func TestHandle_MalformedJSON_Returns400(t *testing.T) {
	t.Parallel()

	h := httpx.Handle(func(_ context.Context, req handleReq) (handleRes, error) {
		return handleRes{}, nil
	})

	body := strings.NewReader(`not-json`)
	r := httptest.NewRequest(http.MethodPost, "/", body)
	w := httptest.NewRecorder()
	h(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHandle_ValidationFailure_Returns422(t *testing.T) {
	t.Parallel()

	h := httpx.Handle(func(_ context.Context, req handleReq) (handleRes, error) {
		return handleRes{}, nil
	})

	// name is required but absent
	body := strings.NewReader(`{}`)
	r := httptest.NewRequest(http.MethodPost, "/", body)
	w := httptest.NewRecorder()
	h(w, r)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)

	var resp httpx.ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Equal(t, "validation_failed", resp.Error)
	assert.NotEmpty(t, resp.Fields)
}

func TestHandle_SvcError_MapsStatus(t *testing.T) {
	t.Parallel()

	h := httpx.Handle(func(_ context.Context, req handleReq) (handleRes, error) {
		return handleRes{}, svcerr.NotFound("not_found", "resource not found")
	})

	body := strings.NewReader(`{"name":"x"}`)
	r := httptest.NewRequest(http.MethodPost, "/", body)
	w := httptest.NewRecorder()
	h(w, r)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var resp httpx.ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Equal(t, "not_found", resp.Error)
}

func TestHandle_InternalError_Returns500(t *testing.T) {
	t.Parallel()

	h := httpx.Handle(func(_ context.Context, req handleReq) (handleRes, error) {
		return handleRes{}, errors.New("boom")
	})

	body := strings.NewReader(`{"name":"x"}`)
	r := httptest.NewRequest(http.MethodPost, "/", body)
	w := httptest.NewRecorder()
	h(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var resp httpx.ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Equal(t, "internal_error", resp.Error)
}

func TestHandleBatch_HappyPath_SetsUsageCountHeader(t *testing.T) {
	t.Parallel()

	h := httpx.HandleBatch(func(_ context.Context, req handleReq) (httpx.BatchResponse[handleRes], error) {
		items := []handleRes{{Greeting: "a"}, {Greeting: "b"}, {Greeting: "c"}, {Greeting: "d"}, {Greeting: "e"}}
		return httpx.BatchResponse[handleRes]{Results: items}, nil
	})

	body := strings.NewReader(`{"name":"batch"}`)
	r := httptest.NewRequest(http.MethodPost, "/", body)
	w := httptest.NewRecorder()
	h(w, r)

	require.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "5", w.Header().Get("X-Usage-Count"))
}

func TestHandleBatch_AutoSetsTotal(t *testing.T) {
	t.Parallel()

	h := httpx.HandleBatch(func(_ context.Context, req handleReq) (httpx.BatchResponse[handleRes], error) {
		return httpx.BatchResponse[handleRes]{Results: []handleRes{{Greeting: "x"}, {Greeting: "y"}}}, nil
	})

	body := strings.NewReader(`{"name":"batch"}`)
	r := httptest.NewRequest(http.MethodPost, "/", body)
	w := httptest.NewRecorder()
	h(w, r)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[httpx.BatchResponse[handleRes]]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Equal(t, 2, resp.Data.Total)
}

func TestHandleBatch_ValidationFailure_Returns422(t *testing.T) {
	t.Parallel()

	h := httpx.HandleBatch(func(_ context.Context, req handleReq) (httpx.BatchResponse[handleRes], error) {
		return httpx.BatchResponse[handleRes]{}, nil
	})

	body := strings.NewReader(`{}`)
	r := httptest.NewRequest(http.MethodPost, "/", body)
	w := httptest.NewRecorder()
	h(w, r)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestHandleBatch_SvcError_MapsStatus(t *testing.T) {
	t.Parallel()

	h := httpx.HandleBatch(func(_ context.Context, req handleReq) (httpx.BatchResponse[handleRes], error) {
		return httpx.BatchResponse[handleRes]{}, svcerr.Upstream("upstream_error", "upstream down")
	})

	body := strings.NewReader(`{"name":"x"}`)
	r := httptest.NewRequest(http.MethodPost, "/", body)
	w := httptest.NewRecorder()
	h(w, r)

	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
}

func TestHandleBatch_InternalError_Returns500(t *testing.T) {
	t.Parallel()

	h := httpx.HandleBatch(func(_ context.Context, req handleReq) (httpx.BatchResponse[handleRes], error) {
		return httpx.BatchResponse[handleRes]{}, errors.New("unexpected")
	})

	body := strings.NewReader(`{"name":"x"}`)
	r := httptest.NewRequest(http.MethodPost, "/", body)
	w := httptest.NewRecorder()
	h(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestHandle_BodyTooLarge_Returns400(t *testing.T) {
	t.Parallel()

	h := httpx.Handle(func(_ context.Context, req handleReq) (handleRes, error) {
		return handleRes{}, nil
	})

	big := strings.NewReader(`{"name":"` + strings.Repeat("x", 1<<20+1) + `"}`)
	r := httptest.NewRequest(http.MethodPost, "/", big)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var resp httpx.ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Contains(t, resp.Message, "too large")
}
