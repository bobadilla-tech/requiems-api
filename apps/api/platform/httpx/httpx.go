package httpx

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// BatchResponse is the standard envelope for batch endpoints. HandleBatch
// reads Results to set X-Usage-Count and auto-populates Total before writing.
type BatchResponse[T any] struct {
	Results []T `json:"results"`
	Total   int `json:"total"`
}

// Included in every response.
type Metadata struct {
	Timestamp string `json:"timestamp"`
	TraceID   string `json:"trace_id,omitempty"`
}

// Standard success envelope
type Response[T any] struct {
	Data     T        `json:"data"`
	Metadata Metadata `json:"metadata"`
}

// Standard error envelope.
type ErrorResponse struct {
	Error    string       `json:"error"`
	Message  string       `json:"message,omitempty"`
	Fields   []FieldError `json:"fields,omitempty"`
	Metadata Metadata     `json:"metadata"`
}

// Writes a 200-class success response wrapped in {"data": ..., "metadata": ...}.
func JSON[T any](w http.ResponseWriter, status int, v T) {
	write(w, status, Response[T]{
		Data: v,
		Metadata: Metadata{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		},
	})
}

// Writes a JSON error response with a machine-readable code and a human-readable message.
func Error(w http.ResponseWriter, status int, code, message string) {
	write(w, status, ErrorResponse{
		Error:   code,
		Message: message,
		Metadata: Metadata{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		},
	})
}

// ValidationError writes 422 for a BindAndValidate / BindQuery validation failure.
func ValidationError(w http.ResponseWriter, vf *ValidationFailure) {
	if vf == nil {
		Error(w, http.StatusUnprocessableEntity, "validation_failed", "Validation failed.")
		return
	}

	writeValidationError(w, vf.Fields)
}

// writeValidationError writes the same 422 envelope as httpx.Handle for validation_failed.
func writeValidationError(w http.ResponseWriter, fields []FieldError) {
	if len(fields) == 0 {
		write(w, http.StatusUnprocessableEntity, ErrorResponse{
			Error: "validation_failed",
			Metadata: Metadata{
				Timestamp: time.Now().UTC().Format(time.RFC3339),
			},
		})
		return
	}

	write(w, http.StatusUnprocessableEntity, ErrorResponse{
		Error:  "validation_failed",
		Fields: fields,
		Metadata: Metadata{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		},
	})
}

func write(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("httpx: failed to encode JSON response: %v", err)
	}
}
