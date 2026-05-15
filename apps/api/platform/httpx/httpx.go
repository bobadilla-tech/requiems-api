package httpx

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Data is a marker interface for types that can be used as API response payloads.
// Add IsData() to your response struct to use it with httpx.JSON.
type Data interface {
	IsData()
}

// Metadata is included in every response.
type Metadata struct {
	Timestamp string `json:"timestamp"`
	TraceID   string `json:"trace_id,omitempty"`
}

// Response is the standard success envelope: {"data": ..., "metadata": ...}
type Response[T Data] struct {
	Data     T        `json:"data"`
	Metadata Metadata `json:"metadata"`
}

// ErrorResponse is the standard error envelope.
// Fields is populated only for validation errors (error: "validation_failed").
type ErrorResponse struct {
	Error    string       `json:"error"`
	Message  string       `json:"message,omitempty"`
	Fields   []FieldError `json:"fields,omitempty"`
	Metadata Metadata     `json:"metadata"`
}

// JSON writes a 200-class success response wrapped in {"data": ..., "metadata": ...}.
func JSON[T Data](w http.ResponseWriter, status int, v T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(Response[T]{
		Data:     v,
		Metadata: Metadata{Timestamp: time.Now().UTC().Format(time.RFC3339)},
	}); err != nil {
		log.Printf("httpx: failed to encode JSON response: %v", err)
	}
}

// Error writes a JSON error response with a machine-readable code and a
// human-readable message.
//
//	httpx.Error(w, http.StatusBadRequest, "bad_request", "invalid email format")
func Error(w http.ResponseWriter, status int, code, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(ErrorResponse{
		Error:    code,
		Message:  message,
		Metadata: Metadata{Timestamp: time.Now().UTC().Format(time.RFC3339)},
	}); err != nil {
		log.Printf("httpx: failed to encode error response: %v", err)
	}
}

// BatchResponse is the standard envelope for batch endpoints.
// Total is set automatically by HandleBatch from len(Results).
type BatchResponse[T any] struct {
	Results []T `json:"results"`
	Total   int `json:"total"`
}

func (BatchResponse[T]) IsData() {}

// writeValidationError writes a 422 Unprocessable Entity with a structured
// list of field-level constraint violations.
func writeValidationError(w http.ResponseWriter, fields []FieldError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity)
	if err := json.NewEncoder(w).Encode(ErrorResponse{
		Error:    "validation_failed",
		Fields:   fields,
		Metadata: Metadata{Timestamp: time.Now().UTC().Format(time.RFC3339)},
	}); err != nil {
		log.Printf("httpx: failed to encode validation error response: %v", err)
	}
}
