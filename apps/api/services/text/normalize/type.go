package normalize

import (
	normalizer "github.com/bobadilla-tech/go-email-normalizer"

	"requiems-api/platform/httpx"
)

type EmailNormalizationRequest struct {
	Email string `json:"email" validate:"required"`
}

type EmailNormalization struct {
	Original   string              `json:"original"`
	Normalized string              `json:"normalized"`
	Local      string              `json:"local"`
	Domain     string              `json:"domain"`
	Changes    []normalizer.Change `json:"changes"`
}

type BatchEmailNormalizationRequest struct {
	Emails []string `json:"emails" validate:"required,min=1,max=100,dive,required"`
}

type EmailNormalizationBatchItem struct {
	Original   string              `json:"original"`
	Normalized string              `json:"normalized,omitempty"`
	Local      string              `json:"local,omitempty"`
	Domain     string              `json:"domain,omitempty"`
	Changes    []normalizer.Change `json:"changes,omitempty"`
	Valid      bool                `json:"valid"`
	Message    string              `json:"message,omitempty"`
}

type EmailNormalizationBatchResponse = httpx.BatchResponse[EmailNormalizationBatchItem]

func (EmailNormalization) IsData() {}
