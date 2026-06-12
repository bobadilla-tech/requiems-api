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
	Email string `json:"email" validate:"required_without_all=Phone Text" normalize:"trim"`
	Phone string `json:"phone" validate:"required_without_all=Email Text" normalize:"trim"`
	Text  string `json:"text"  validate:"required_without_all=Email Phone" normalize:"trim"`
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
	Normalized   *string      `json:"normalized"`
	Country      *string      `json:"country"`
	Type         *string      `json:"type"`
	Carrier      *CarrierInfo `json:"carrier"`
	Risk         *PhoneRisk   `json:"risk"`
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

	weightedSum := 0.0
	totalWeight := 0.0

	var emailFinalResult *EmailResult
	var phoneFinalResult *PhoneResult
	var textFinalResult *TextResult

	// Weight: 0.5. Deductions applied for each quality signal.
	// Floor at 0.0 — score cannot go negative.
	if emailAddress != "" {
		result, score := buildEmailResult(emailAddress, emailResult)
		emailFinalResult = &result
		weightedSum += score * 0.5
		totalWeight += 0.5
	}

	// --- Phone scoring ---
	if phoneNumber != "" {
		result, score := buildPhoneResult(phoneResult)
		phoneFinalResult = &result
		totalWeight += 0.4
		weightedSum += score * 0.4
	}

	// --- Text scoring ---
	// Weight: 0.1. Scoring is deferred until ToxicityScore is available from the sentiment service.
	// text validation runs but does not contribute to overall_quality_score yet.
	// Note: text-only requests will return overall_quality_score of 0.0 until scoring is enabled.
	if text != "" {
		result := buildTextResult(profanityResult, sentimentResult)
		textFinalResult = &result
	}

	overallScore := 0.0

	// Normalize weighted sum by total active weight.
	// Denominator adjusts based on which fields were present,
	if totalWeight > 0 {
		overallScore = weightedSum / totalWeight
	}

	// Return nil for absent fields — null in JSON signals "not requested",
	return Response{
		Email:               emailFinalResult,
		Phone:               phoneFinalResult,
		Text:                textFinalResult,
		OverallQualityScore: overallScore,
	}
}

func buildEmailResult(emailAddress string, e email.Validation) (result EmailResult, score float64) {

	score = 1.0
	var flags []string

	if !e.SyntaxValid {
		// Invalid syntax is disqualifying — full deduction, skip remaining checks.
		flags = append(flags, "email_syntax_invalid")
		score -= 1.0
	} else {
		if !e.MxValid {
			// Domain has no mail server — email is undeliverable.
			flags = append(flags, "email_mx_invalid")
			flags = append(flags, "email_invalid")
			score -= 0.6
		}

		if e.Disposable {
			// Disposable/temporary address — high fraud signal.
			flags = append(flags, "email_disposable")
			score -= 0.5
		}

		if e.Suggestion != nil && *e.Suggestion != "" {
			// Likely typo in domain (e.g. gmial.com) — low confidence signal.
			flags = append(flags, "email_has_suggestion")
			score -= 0.1
		}
	}

	return EmailResult{
		Valid:        e.Valid,
		Normalized:   e.Normalized,
		Original:     emailAddress,
		SyntaxValid:  e.SyntaxValid,
		MXValid:      e.MxValid,
		Disposable:   e.Disposable,
		Suggestion:   e.Suggestion,
		Flags:        flags,
		QualityScore: math.Max(0, score),
	}, math.Max(0, score)

}

func buildPhoneResult(p phone.ValidateResponse) (result PhoneResult, score float64) {
	score = 1.0
	var flags []string
	result = PhoneResult{Valid: p.Valid}

	if p.Valid {

		if p.Risk != nil {
			if p.Risk.IsVirtual || p.Risk.IsVoIP {
				// VoIP and virtual share a single deduction — apply once.
				if p.Risk.IsVirtual {
					flags = append(flags, "phone_virtual")
				}

				if p.Risk.IsVoIP {
					flags = append(flags, "phone_voip")
				}
				score -= 0.3
			}

			result.Risk = &PhoneRisk{IsVoIP: p.Risk.IsVoIP, IsVirtual: p.Risk.IsVirtual}
		}

		if p.Type == "unknown" {
			// Cannot determine line type — minor confidence penalty.
			flags = append(flags, "phone_unknown_type")
			score -= 0.1

		}

		if p.Type == "landline" {
			// Landline is less common for contact forms — slight penalty.
			flags = append(flags, "phone_landline")
			score -= 0.05
		}

		result.Normalized = &p.Formatted
		result.Country = &p.Country
		result.Type = &p.Type
		
		if p.Carrier != nil && p.Carrier.Name != "" {
			result.Carrier = &CarrierInfo{Name: p.Carrier.Name}
		}

	} else {
		// Invalid number — full deduction.
		flags = append(flags, "phone_invalid")
		score -= 1.0
	}

	result.Flags = flags
	result.QualityScore = math.Max(0, score)
	return result, math.Max(0, score)
}

func buildTextResult(p profanity.Result, sent sentiment.Result) TextResult {
	result := TextResult{
		IsSafe:    true,
		Sentiment: sent.Sentiment,
	}

	if p.HasProfanity {
		result.IsSafe = false
		result.Flags = append(result.Flags, "text_profanity")
	}

	return result
}
