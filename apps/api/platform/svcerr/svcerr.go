package svcerr

import "net/http"

// Kind classifies a service error so the HTTP layer can map it to a status code.
type Kind int

const (
	KindNotFound Kind = iota // 404
	KindInvalid              // 400
	KindUnknown              // 422
	KindUpstream             // 503
)

// A transport-agnostic service error. Services return this type;
// the HTTP layer maps Kind to an HTTP status code.
type Error struct {
	Kind    Kind
	Code    string
	Message string
}

func (e *Error) Error() string { return e.Message }

// Returns the HTTP status code that corresponds to the error kind.
func HTTPStatus(e *Error) int {
	switch e.Kind {
	case KindNotFound:
		return http.StatusNotFound
	case KindInvalid:
		return http.StatusBadRequest
	case KindUnknown:
		return http.StatusUnprocessableEntity
	case KindUpstream:
		return http.StatusServiceUnavailable
	}
	return http.StatusInternalServerError
}

func NotFound(code, msg string) *Error { return &Error{Kind: KindNotFound, Code: code, Message: msg} }
func Invalid(code, msg string) *Error  { return &Error{Kind: KindInvalid, Code: code, Message: msg} }
func Unknown(code, msg string) *Error  { return &Error{Kind: KindUnknown, Code: code, Message: msg} }
func Upstream(code, msg string) *Error { return &Error{Kind: KindUpstream, Code: code, Message: msg} }
