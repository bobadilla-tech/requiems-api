package svcerr

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPStatus_Nil(t *testing.T) {
	t.Parallel()
	assert.Equal(t, http.StatusInternalServerError, HTTPStatus(nil))
}

func TestHTTPStatus_AllKinds(t *testing.T) {
	t.Parallel()
	cases := []struct {
		kind Kind
		want int
	}{
		{KindNotFound, http.StatusNotFound},
		{KindInvalid, http.StatusBadRequest},
		{KindUnknown, http.StatusUnprocessableEntity},
		{KindUpstream, http.StatusServiceUnavailable},
	}
	for _, tc := range cases {
		e := &Error{Kind: tc.kind}
		assert.Equal(t, tc.want, HTTPStatus(e))
	}
}

func TestHTTPStatus_UnknownKind(t *testing.T) {
	t.Parallel()
	e := &Error{Kind: Kind(999)}
	assert.Equal(t, http.StatusInternalServerError, HTTPStatus(e))
}

func TestError_ErrorMethod(t *testing.T) {
	t.Parallel()
	e := &Error{Message: "something went wrong"}
	assert.Equal(t, "something went wrong", e.Error())
}

func TestNotFound_Fields(t *testing.T) {
	t.Parallel()
	e := NotFound("not_found", "resource missing")
	assert.Equal(t, KindNotFound, e.Kind)
	assert.Equal(t, "not_found", e.Code)
	assert.Equal(t, "resource missing", e.Message)
}

func TestInvalid_Fields(t *testing.T) {
	t.Parallel()
	e := Invalid("bad_input", "input is bad")
	assert.Equal(t, KindInvalid, e.Kind)
	assert.Equal(t, "bad_input", e.Code)
	assert.Equal(t, "input is bad", e.Message)
}

func TestUnknown_Fields(t *testing.T) {
	t.Parallel()
	e := Unknown("unprocessable", "cannot process")
	assert.Equal(t, KindUnknown, e.Kind)
	assert.Equal(t, "unprocessable", e.Code)
	assert.Equal(t, "cannot process", e.Message)
}

func TestUpstream_Fields(t *testing.T) {
	t.Parallel()
	e := Upstream("upstream_error", "service down")
	assert.Equal(t, KindUpstream, e.Kind)
	assert.Equal(t, "upstream_error", e.Code)
	assert.Equal(t, "service down", e.Message)
}
