package inputvalidate

import (
	"context"
	"math"
	"sync"

	"requiems-api/services/text/sentiment"
	"requiems-api/services/validation/email"
	"requiems-api/services/validation/phone"
	"requiems-api/services/validation/profanity"
)

// Request defines the input for the validate endpoint.
// At least one of Email, Phone, or Text must be present.
type Request struct {
	Email string `json:"email" validate:"required_without_all=Phone Text"`
	Phone string `json:"phone" validate:"required_without_all=Email Text"`
	Text  string `json:"text"  validate:"required_without_all=Email Phone"`
}

// CarrierInfo holds the name of the carrier associated with a phone number.
type CarrierInfo struct {
	Name string `json:"name"`
}

// PhoneRisk describes line-type risk signals for a phone number.
type PhoneRisk struct {
	IsVoIP    bool `json:"is_voip"`
	IsVirtual bool `json:"is_virtual"`
}

// EmailResult is the normalized output of the email validation service,
// enriched with quality signals and flags computed by this system.
type EmailResult struct {
	Valid        bool     `json:"valid"`
	Normalized   *string  `json:"normalized"`
	Original     string   `json:"original"`
	SyntaxValid  bool     `json:"syntax_valid"`
	MXValid      bool     `json:"mx_valid"`
	Disposable   bool     `json:"disposable"`
	Suggestion   *string  `json:"suggestion"`
	Flags        []string `json:"flags"`
	QualityScore float64  `json:"quality_score"`
}

// PhoneResult is the normalized output of the phone validation service,
// enriched with quality signals and flags computed by this system.
type PhoneResult struct {
	Valid        bool         `json:"valid"`
	Normalized   string       `json:"normalized,omitempty"`
	Country      string       `json:"country,omitempty"`
	Type         string       `json:"type,omitempty"`
	Carrier      *CarrierInfo `json:"carrier"`
	Risk         PhoneRisk    `json:"risk"`
	Flags        []string     `json:"flags"`
	QualityScore float64      `json:"quality_score"`
}

// TextResult is the combined output of the profanity and sentiment services.
type TextResult struct {
	IsSafe        bool     `json:"is_safe"`
	ToxicityScore *float64 `json:"toxicity_score"`
	Sentiment     string   `json:"sentiment"`
	Flags         []string `json:"flags"`
}

// Response is the unified output of the validate endpoint.
// Fields are nil when the corresponding input was not provided —
// null in JSON signals "not requested", not "invalid".
type Response struct {
	Email               *EmailResult `json:"email"`
	Phone               *PhoneResult `json:"phone"`
	Text                *TextResult  `json:"text"`
	OverallQualityScore float64      `json:"overall_quality_score"`
}

// EmailService is the interface for the email validation composed service.
type EmailService interface {
	ValidateEmail(ctx context.Context, email string) email.Validation
}

// PhoneService is the interface for the phone validation composed service.
type PhoneService interface {
	Validate(number string) phone.ValidateResponse
}

// ProfanityService is the interface for the profanity check composed service.
type ProfanityService interface {
	Check(ctx context.Context, text string) profanity.Result
}

// SentimentService is the interface for the sentiment analysis composed service.
type SentimentService interface {
	Analyze(text string) sentiment.Result
}

type Service struct {
	emailSvc     EmailService
	phoneSvc     PhoneService
	profanitySvc ProfanityService
	sentimentSvc SentimentService
}

// NewService constructs a Service with all four composed service dependencies injected.
// Each dependency is defined as an interface to allow stubbing in tests.
func NewService(
	emailSvc EmailService,
	phoneSvc PhoneService,
	profanitySvc ProfanityService,
	sentimentSvc SentimentService,
) *Service {
	return &Service{
		emailSvc:     emailSvc,
		phoneSvc:     phoneSvc,
		profanitySvc: profanitySvc,
		sentimentSvc: sentimentSvc,
	}
}

// Validate composes the email, phone, profanity, and sentiment services into a single call.
// Only the services needed for the provided fields are invoked — all active calls run in parallel.
// The overall_quality_score is a weighted mean of the present signals only,
// so the denominator adjusts based on which fields were provided.
func (s *Service) Validate(ctx context.Context, emailAddress, phoneNumber, text string) Response {

	var emailFinalResult EmailResult
	var phoneFinalResult PhoneResult
	var textFinalResult TextResult

	overallScore := 0.0
	weightedSum := 0.0
	totalWeight := 0.0
	var wg sync.WaitGroup

	// Raw results from each composed service, populated concurrently.
	var (
		emailResult     email.Validation
		phoneResult     phone.ValidateResponse
		profanityResult profanity.Result
		sentimentResult sentiment.Result
	)

	// Fan out: launch only the goroutines needed based on which fields are present.
	if emailAddress != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			emailResult = s.emailSvc.ValidateEmail(ctx, emailAddress)
		}()
	}

	if phoneNumber != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			phoneResult = s.phoneSvc.Validate(phoneNumber)
		}()
	}

	if text != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			profanityResult = s.profanitySvc.Check(ctx, text)
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			sentimentResult = s.sentimentSvc.Analyze(text)
		}()
	}

	wg.Wait()

	// --- Email scoring ---
	// Weight: 0.5. Deductions applied for each quality signal.
	// Floor at 0.0 — score cannot go negative.
	if emailAddress != "" {

		emailScore := 1.0

		if !emailResult.SyntaxValid {
			// Invalid syntax is disqualifying — full deduction, skip remaining checks.
			emailFinalResult.Flags = append(emailFinalResult.Flags, "email_syntax_invalid")
			emailScore -= 1.0
		} else {
			if emailResult.MxValid == false {
				// Domain has no mail server — email is undeliverable.
				emailFinalResult.Flags = append(emailFinalResult.Flags, "email_mx_invalid")
				emailFinalResult.Flags = append(emailFinalResult.Flags, "email_invalid")
				emailScore -= 0.6
			}

			if emailResult.Disposable {
				// Disposable/temporary address — high fraud signal.
				emailFinalResult.Flags = append(emailFinalResult.Flags, "email_disposable")
				emailScore -= 0.5
			}

			if emailResult.Suggestion != nil && len(*emailResult.Suggestion) > 0 {
				// Likely typo in domain (e.g. gmial.com) — low confidence signal.
				emailFinalResult.Flags = append(emailFinalResult.Flags, "email_has_suggestion")
				emailScore -= 0.1
			}
		}

		emailScore = math.Max(0, emailScore)

		emailFinalResult.Valid = emailResult.Valid
		emailFinalResult.Normalized = emailResult.Normalized
		emailFinalResult.Original = emailAddress
		emailFinalResult.SyntaxValid = emailResult.SyntaxValid
		emailFinalResult.MXValid = emailResult.MxValid
		emailFinalResult.Disposable = emailResult.Disposable
		emailFinalResult.Suggestion = emailResult.Suggestion
		emailFinalResult.QualityScore = emailScore

		weightedSum += emailScore * 0.5
		totalWeight += 0.5
	}

	// --- Phone scoring ---
	if phoneNumber != "" {
		phoneScore := 1.0

		if phoneResult.Valid {

			if phoneResult.Risk != nil {
				if phoneResult.Risk.IsVirtual || phoneResult.Risk.IsVoIP {
					// VoIP and virtual share a single deduction — apply once.
					if phoneResult.Risk.IsVirtual {
						phoneFinalResult.Flags = append(phoneFinalResult.Flags, "phone_virtual")
					}

					if phoneResult.Risk.IsVoIP {
						phoneFinalResult.Flags = append(phoneFinalResult.Flags, "phone_voip")
					}
					phoneScore -= 0.3
				}
			}

			if phoneResult.Type == "unknown" {
				// Cannot determine line type — minor confidence penalty.
				phoneFinalResult.Flags = append(phoneFinalResult.Flags, "phone_unknown_type")
				phoneScore -= 0.1

			}

			if phoneResult.Type == "landline" {
				// Landline is less common for contact forms — slight penalty.
				phoneScore -= 0.05
			}

			phoneFinalResult.Normalized = phoneResult.Formatted
			phoneFinalResult.Country = phoneResult.Country
			phoneFinalResult.Type = phoneResult.Type
			if phoneResult.Carrier != nil && phoneResult.Carrier.Name != "" {
				phoneFinalResult.Carrier = &CarrierInfo{Name: phoneResult.Carrier.Name}
			}
			if phoneResult.Risk != nil {
				phoneFinalResult.Risk.IsVirtual = phoneResult.Risk.IsVirtual
				phoneFinalResult.Risk.IsVoIP = phoneResult.Risk.IsVoIP
			}

		} else {
			// Invalid number — full deduction.
			phoneFinalResult.Flags = append(phoneFinalResult.Flags, "phone_invalid")
			phoneScore -= 1.0
		}

		phoneScore = math.Max(0, phoneScore)

		phoneFinalResult.Valid = phoneResult.Valid
		phoneFinalResult.QualityScore = phoneScore

		totalWeight += 0.4
		weightedSum += phoneScore * 0.4
	}

	// --- Text scoring ---
	// Weight: 0.1. Scoring is deferred until ToxicityScore is available from the sentiment service.
	// text validation runs but does not contribute to overall_quality_score yet.
	if text != "" {
		textFinalResult.IsSafe = true
		if profanityResult.HasProfanity {
			textFinalResult.IsSafe = false
			textFinalResult.Flags = append(textFinalResult.Flags, "text_profanity")
		}
		textFinalResult.Sentiment = sentimentResult.Sentiment
	}

	// Normalize weighted sum by total active weight.
	// Denominator adjusts based on which fields were present,
	if totalWeight > 0 {
		overallScore = weightedSum / totalWeight
	}

	// Return nil for absent fields — null in JSON signals "not requested",
	return Response{
		Email: func() *EmailResult {
			if emailAddress != "" {
				return &emailFinalResult
			}
			return nil
		}(),
		Phone: func() *PhoneResult {
			if phoneNumber != "" {
				return &phoneFinalResult
			}
			return nil
		}(),
		Text: func() *TextResult {
			if text != "" {
				return &textFinalResult
			}
			return nil
		}(),
		OverallQualityScore: overallScore,
	}
}
