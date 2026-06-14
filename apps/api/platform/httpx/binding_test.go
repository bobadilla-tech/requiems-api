package httpx_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
)

// bindReq is a sample struct used to exercise BindAndValidate.
type bindReq struct {
	Email string `json:"email" validate:"required,email"`
	Count int    `json:"count" validate:"min=1"`
}

type bindTrimReq struct {
	Email string `json:"email" validate:"required,email" normalize:"trim"`
}

func TestBindAndValidate_ValidJSON(t *testing.T) {
	t.Parallel()

	body := strings.NewReader(`{"email":"user@example.com","count":3}`)
	r := httptest.NewRequest("POST", "/", body)

	var req bindReq
	err := httpx.BindAndValidate(r, &req)
	require.NoError(t, err)
	assert.Equal(t, "user@example.com", req.Email)
	assert.Equal(t, 3, req.Count)
}

func TestBindAndValidate_MalformedJSON(t *testing.T) {
	t.Parallel()

	body := strings.NewReader(`not-json`)
	r := httptest.NewRequest("POST", "/", body)

	var req bindReq
	err := httpx.BindAndValidate(r, &req)
	require.Error(t, err)
}

func TestBindAndValidate_UnknownFields(t *testing.T) {
	t.Parallel()

	body := strings.NewReader(`{"email":"user@example.com","count":1,"unknown":"field"}`)
	r := httptest.NewRequest("POST", "/", body)

	var req bindReq
	err := httpx.BindAndValidate(r, &req)
	require.Error(t, err)
}

func TestBindAndValidate_ValidationFailure(t *testing.T) {
	t.Parallel()

	// count is below min=1
	body := strings.NewReader(`{"email":"user@example.com","count":0}`)
	r := httptest.NewRequest("POST", "/", body)

	var req bindReq
	err := httpx.BindAndValidate(r, &req)
	require.Error(t, err)

	var vf *httpx.ValidationFailure
	require.ErrorAs(t, err, &vf)
	assert.NotEmpty(t, vf.Fields)
}

func TestBindAndValidate_NormalizeTrim(t *testing.T) {
	t.Parallel()

	body := strings.NewReader(`{"email":"  user@example.com  "}`)
	r := httptest.NewRequest("POST", "/", body)

	var req bindTrimReq
	err := httpx.BindAndValidate(r, &req)
	require.NoError(t, err)
	assert.Equal(t, "user@example.com", req.Email)
}

// queryReq is used to exercise BindQuery with several field types.
type queryReq struct {
	Name    string    `query:"name"    validate:"required"`
	Age     int       `query:"age"`
	Score   float64   `query:"score"`
	Active  bool      `query:"active"`
	Since   time.Time `query:"since"`
	Ignored string    // no query tag – must be skipped
}

type queryTrimReq struct {
	Name string `query:"name" validate:"required" normalize:"trim"`
}

func TestBindQuery_StringField(t *testing.T) {
	t.Parallel()

	r := httptest.NewRequest("GET", "/?name=alice", http.NoBody)

	var req queryReq
	err := httpx.BindQuery(r, &req)
	require.NoError(t, err)
	assert.Equal(t, "alice", req.Name)
}

func TestBindQuery_IntField(t *testing.T) {
	t.Parallel()

	r := httptest.NewRequest("GET", "/?name=x&age=30", http.NoBody)

	var req queryReq
	err := httpx.BindQuery(r, &req)
	require.NoError(t, err)
	assert.Equal(t, 30, req.Age)
}

func TestBindQuery_FloatField(t *testing.T) {
	t.Parallel()

	r := httptest.NewRequest("GET", "/?name=x&score=9.5", http.NoBody)

	var req queryReq
	err := httpx.BindQuery(r, &req)
	require.NoError(t, err)
	assert.Equal(t, 9.5, req.Score)
}

func TestBindQuery_BoolField(t *testing.T) {
	t.Parallel()

	r := httptest.NewRequest("GET", "/?name=x&active=true", http.NoBody)

	var req queryReq
	err := httpx.BindQuery(r, &req)
	require.NoError(t, err)
	assert.True(t, req.Active, "active: want true, got false")
}

func TestBindQuery_TimeField(t *testing.T) {
	t.Parallel()

	r := httptest.NewRequest("GET", "/?name=x&since=2024-06-15", http.NoBody)

	var req queryReq
	err := httpx.BindQuery(r, &req)
	require.NoError(t, err)

	want := time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC)
	assert.True(t, req.Since.Equal(want), "since: want %v, got %v", want, req.Since)
}

func TestBindQuery_InvalidInt(t *testing.T) {
	t.Parallel()

	r := httptest.NewRequest("GET", "/?name=x&age=notanint", http.NoBody)

	var req queryReq
	err := httpx.BindQuery(r, &req)
	require.Error(t, err)
}

func TestBindQuery_InvalidFloat(t *testing.T) {
	t.Parallel()

	r := httptest.NewRequest("GET", "/?name=x&score=notafloat", http.NoBody)

	var req queryReq
	err := httpx.BindQuery(r, &req)
	require.Error(t, err)
}

func TestBindQuery_InvalidBool(t *testing.T) {
	t.Parallel()

	r := httptest.NewRequest("GET", "/?name=x&active=notabool", http.NoBody)

	var req queryReq
	err := httpx.BindQuery(r, &req)
	require.Error(t, err)
}

func TestBindQuery_InvalidTime(t *testing.T) {
	t.Parallel()

	r := httptest.NewRequest("GET", "/?name=x&since=not-a-date", http.NoBody)

	var req queryReq
	err := httpx.BindQuery(r, &req)
	require.Error(t, err)
}

func TestBindQuery_ValidationFailure(t *testing.T) {
	t.Parallel()

	// name is required but missing
	r := httptest.NewRequest("GET", "/", http.NoBody)

	var req queryReq
	err := httpx.BindQuery(r, &req)
	require.Error(t, err)

	var vf *httpx.ValidationFailure
	require.ErrorAs(t, err, &vf)
}

func TestBindQuery_NonPointerDst(t *testing.T) {
	t.Parallel()

	r := httptest.NewRequest("GET", "/", http.NoBody)

	var req queryReq
	err := httpx.BindQuery(r, req)
	require.Error(t, err)
}

func TestBindQuery_NormalizeTrim(t *testing.T) {
	t.Parallel()

	r := httptest.NewRequest("GET", "/?name=%20%20alice%20%20", http.NoBody)

	var req queryTrimReq
	err := httpx.BindQuery(r, &req)
	require.NoError(t, err)
	assert.Equal(t, "alice", req.Name)
}
