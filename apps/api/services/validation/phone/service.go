package phone

import (
	"github.com/nyaruka/phonenumbers"
)

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

// BatchValidateResponse is the response for a batch phone number validation request.
type BatchValidateResponse struct {
	Results []ValidateResponse `json:"results"`
	Total   int                `json:"total"`
}

// Service provides phone number validation logic.
type Service struct{}

// NewService creates a new phone validation Service.
func NewService() *Service { return &Service{} }

// Validate parses and validates a phone number, returning structured metadata.
// When the number cannot be parsed or is not valid, Valid is false and the
// optional fields are omitted.
func (s *Service) Validate(number string) ValidateResponse {
	num, err := phonenumbers.Parse(number, "")
	if err != nil || !phonenumbers.IsValidNumber(num) {
		return ValidateResponse{Number: number, Valid: false}
	}

	numType := phonenumbers.GetNumberType(num)

	var c *Carrier
	if name, err := phonenumbers.GetCarrierForNumber(num, "en"); err == nil && name != "" {
		c = &Carrier{Name: name, Source: "metadata"}
	}

	risk := numberRisk(numType)

	return ValidateResponse{
		Number:    number,
		Valid:     true,
		Country:   phonenumbers.GetRegionCodeForNumber(num),
		Type:      numberType(numType),
		Formatted: phonenumbers.Format(num, phonenumbers.INTERNATIONAL),
		Carrier:   c,
		Risk:      &risk,
	}
}

// numberType converts a phonenumbers type constant to a human-readable string.
func numberType(t phonenumbers.PhoneNumberType) string {
	switch t {
	case phonenumbers.MOBILE:
		return "mobile"
	case phonenumbers.FIXED_LINE:
		return "landline"
	case phonenumbers.FIXED_LINE_OR_MOBILE:
		return "landline_or_mobile"
	case phonenumbers.TOLL_FREE:
		return "toll_free"
	case phonenumbers.PREMIUM_RATE:
		return "premium_rate"
	case phonenumbers.SHARED_COST:
		return "shared_cost"
	case phonenumbers.VOIP:
		return "voip"
	case phonenumbers.PERSONAL_NUMBER:
		return "personal_number"
	case phonenumbers.PAGER:
		return "pager"
	case phonenumbers.UAN:
		return "uan"
	case phonenumbers.VOICEMAIL:
		return "voicemail"
	default:
		return "unknown"
	}
}

// ValidateBatch validates a slice of phone numbers and returns results in the same order.
func (s *Service) ValidateBatch(numbers []string) []ValidateResponse {
	results := make([]ValidateResponse, len(numbers))
	for i, n := range numbers {
		results[i] = s.Validate(n)
	}
	return results
}

// numberRisk returns VOIP and virtual risk flags for a given phone number type.
func numberRisk(t phonenumbers.PhoneNumberType) Risk {
	switch t {
	case phonenumbers.VOIP:
		return Risk{IsVoIP: true, IsVirtual: true}
	case phonenumbers.PERSONAL_NUMBER, phonenumbers.UAN,
		phonenumbers.PAGER, phonenumbers.VOICEMAIL:
		return Risk{IsVirtual: true}
	default:
		return Risk{}
	}
}
