package iban

import (
	"context"

	"requiems-api/platform/httpx"
)

// ParseResponse is the response payload for GET /v1/finance/iban/{iban}.
type ParseResponse struct {
	IBAN     string `json:"iban"`
	Valid    bool   `json:"valid"`
	Country  string `json:"country"`
	BankCode string `json:"bank_code"`
	Account  string `json:"account"`
}

func (ParseResponse) IsData() {}

// Validator is the interface used by the HTTP transport layer, allowing
// transport tests to inject a stub without requiring a database connection.
type Validator interface {
	Parse(ctx context.Context, raw string) (ParseResponse, error)
	ParseBatch(ctx context.Context, numbers []string) (BatchParseResponse, error)
}

// BatchParseRequest is the body for validating multiples IBAN at once.
type BatchParseRequest struct {
	Numbers []string `json:"numbers" validate:"required,min=1,max=50,dive,required"`
}

// BatchParseResponse is the response for a batch IBAN parse request.
type BatchParseResponse = httpx.BatchResponse[ParseResponse]
