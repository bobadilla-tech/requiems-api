package email

import "requiems-api/platform/httpx"

// Request holds the JSON body for email validation.
type Request struct {
	Email string `json:"email" validate:"required"`
}

// Validation is the full validation result for an email address.
type Validation struct {
	Email       *string `json:"email"`
	Valid       bool    `json:"valid"`
	SyntaxValid bool    `json:"syntax_valid"`
	MxValid     bool    `json:"mx_valid"`
	Disposable  bool    `json:"disposable"`
	Normalized  *string `json:"normalized"`
	Domain      *string `json:"domain"`
	Suggestion  *string `json:"suggestion"`
}

func (Validation) IsData() {}

// BatchRequest is the body for validating multiple emails at once.
type BatchRequest struct {
	Emails []string `json:"emails" validate:"required,min=1,max=50,dive,required"`
}

// BatchItem holds the result for a single email in a batch request.
// Each email is always processed — invalid ones return valid: false.
type BatchItem struct {
	Email       *string `json:"email"`
	Valid       bool    `json:"valid"`
	SyntaxValid bool    `json:"syntax_valid"`
	MXValid     bool    `json:"mx_valid"`
	Disposable  bool    `json:"disposable"`
	Normalized  *string `json:"normalized,omitempty"`
	Domain      *string `json:"domain,omitempty"`
	Suggestion  *string `json:"suggestion,omitempty"`
}

// BatchResponse is the response payload for POST /v1/validation/email/batch.
type BatchResponse = httpx.BatchResponse[BatchItem]
