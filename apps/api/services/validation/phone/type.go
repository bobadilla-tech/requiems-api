package phone

import "requiems-api/platform/httpx"

// ValidateRequest holds query parameters for the phone validation endpoint.
type ValidateRequest struct {
	Number string `query:"number" validate:"required"`
}

// Carrier holds carrier name and detection source for a phone number.
type Carrier struct {
	Name   string `json:"name,omitempty"`
	Source string `json:"source,omitempty"`
}

// Risk holds VOIP and virtual number risk flags for a phone number.
type Risk struct {
	IsVoIP    bool `json:"is_voip"`
	IsVirtual bool `json:"is_virtual"`
}

// ValidateResponse is the response for a phone number validation request.
type ValidateResponse struct {
	Number    string   `json:"number"`
	Valid     bool     `json:"valid"`
	Country   string   `json:"country,omitempty"`
	Type      string   `json:"type,omitempty"`
	Formatted string   `json:"formatted,omitempty"`
	Carrier   *Carrier `json:"carrier,omitempty"`
	Risk      *Risk    `json:"risk,omitempty"`
}

func (ValidateResponse) IsData() {}

// BatchValidateRequest is the body for validating multiple phone numbers at once.
type BatchValidateRequest struct {
	Numbers []string `json:"numbers" validate:"required,min=1,max=50,dive,required"`
}

// BatchValidateResponse is the response for a batch phone number validation request.
type BatchValidateResponse = httpx.BatchResponse[ValidateResponse]
