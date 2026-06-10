package inputvalidate

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
	"requiems-api/services/text/sentiment"
	"requiems-api/services/validation/email"
	"requiems-api/services/validation/phone"
	"requiems-api/services/validation/profanity"
)

func setupRouter(
	emailSvc EmailService,
	phoneSvc PhoneService,
	profanitySvc ProfanityService,
	sentimentSvc SentimentService,
) chi.Router {
	r := chi.NewRouter()
	svc := NewService(emailSvc, phoneSvc, profanitySvc, sentimentSvc)
	RegisterRoutes(r, svc)
	return r
}

// --- Email stub ---
type stubEmailService struct {
	emailResult email.Validation
}

func (s *stubEmailService) ValidateEmail(ctx context.Context, email string) email.Validation {
	return s.emailResult
}

// --- Phone stub ---
type stubPhoneService struct {
	phoneResult phone.ValidateResponse
}

func (s *stubPhoneService) Validate(number string) phone.ValidateResponse {
	return s.phoneResult
}

// --- Profanity stub ---
type stubProfanityService struct {
	result profanity.Result
}

func (s *stubProfanityService) Check(ctx context.Context, text string) profanity.Result {
	return s.result
}

// --- Sentiment stub ---
type stubSentimentService struct {
	result sentiment.Result
}

func (s *stubSentimentService) Analyze(text string) sentiment.Result {
	return s.result
}

func ptr(s string) *string {
	return &s
}

func post(t *testing.T, r chi.Router, body string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodPost, "/input/validate", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestInputValidate_Validate_ValidEmailAndPhone_ReturnsHighScore(t *testing.T) {
	t.Parallel()
	emailSvc := &stubEmailService{
		emailResult: email.Validation{
			Valid:       true,
			Normalized:  ptr("user@gmail.com"),
			SyntaxValid: true,
			MxValid:     true,
			Disposable:  false,
			Suggestion:  nil,
		},
	}

	phoneSvc := &stubPhoneService{
		phoneResult: phone.ValidateResponse{
			Valid:     true,
			Formatted: "+4915123456789",
			Country:   "DE",
			Type:      "mobile",
			Carrier:   &phone.Carrier{Name: "Telekom"},
			Risk:      &phone.Risk{IsVoIP: false, IsVirtual: false},
		},
	}

	r := setupRouter(emailSvc, phoneSvc, &stubProfanityService{}, &stubSentimentService{})
	w := post(t, r, `{"email": "user@gmail.com", "phone": "+4915123456789"}`)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Response]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))

	assert.True(t, resp.Data.Email.Valid)
	assert.True(t, resp.Data.Phone.Valid)
	assert.Nil(t, resp.Data.Text) // text not requested → null
	assert.Greater(t, resp.Data.OverallQualityScore, 0.9)
}

func TestInputValidate_Validate_DisposableEmail_ReturnsLowScore(t *testing.T) {
	t.Parallel()
	emailSvc := &stubEmailService{
		emailResult: email.Validation{
			Valid:       true,
			Normalized:  ptr("user@mailinator.com"),
			SyntaxValid: true,
			MxValid:     true,
			Disposable:  true,
			Suggestion:  nil,
		},
	}

	r := setupRouter(emailSvc, &stubPhoneService{}, &stubProfanityService{}, &stubSentimentService{})
	w := post(t, r, `{"email": "user@mailinator.com"}`)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Response]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))

	assert.True(t, resp.Data.Email.Disposable)
	assert.Contains(t, resp.Data.Email.Flags, "email_disposable")
	assert.LessOrEqual(t, resp.Data.Email.QualityScore, 0.5)
}

func TestInputValidate_Validate_InvalidEmailSyntax_ReturnsZeroScore(t *testing.T) {
	t.Parallel()
	emailSvc := &stubEmailService{
		emailResult: email.Validation{
			Valid:       false,
			Normalized:  nil,
			SyntaxValid: false,
			MxValid:     false,
			Disposable:  false,
			Suggestion:  nil,
		},
	}

	r := setupRouter(emailSvc, &stubPhoneService{}, &stubProfanityService{}, &stubSentimentService{})
	w := post(t, r, `{"email": "not-an-email"}`)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Response]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))

	assert.False(t, resp.Data.Email.Valid)
	assert.Contains(t, resp.Data.Email.Flags, "email_syntax_invalid")
	assert.Equal(t, 0.0, resp.Data.Email.QualityScore)
}

func TestInputValidate_Validate_EmailOnly_PhoneIsNullAndScoreBasedOnEmailOnly(t *testing.T) {
	t.Parallel()
	emailSvc := &stubEmailService{
		emailResult: email.Validation{
			Valid:       true,
			Normalized:  ptr("user@gmail.com"),
			SyntaxValid: true,
			MxValid:     true,
			Disposable:  false,
			Suggestion:  nil,
		},
	}

	r := setupRouter(emailSvc, &stubPhoneService{}, &stubProfanityService{}, &stubSentimentService{})
	w := post(t, r, `{"email": "user@gmail.com"}`)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Response]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))

	assert.Nil(t, resp.Data.Phone)                      // phone not requested → null
	assert.Nil(t, resp.Data.Text)                       // text not requested → null
	assert.Equal(t, 1.0, resp.Data.OverallQualityScore) // email score 1.0, weight 0.5/0.5 = 1.0
}

func TestInputValidate_Validate_EmailWithTypoSuggestion_ReturnsSuggestionAndFlag(t *testing.T) {
	t.Parallel()
	emailSvc := &stubEmailService{
		emailResult: email.Validation{
			Valid:       true,
			Normalized:  ptr("user@gmail.com"),
			SyntaxValid: true,
			MxValid:     true,
			Disposable:  false,
			Suggestion:  ptr("gmail.com"),
		},
	}

	r := setupRouter(emailSvc, &stubPhoneService{}, &stubProfanityService{}, &stubSentimentService{})
	w := post(t, r, `{"email": "user@gmial.com"}`)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Response]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))

	assert.NotNil(t, resp.Data.Email.Suggestion)
	assert.Equal(t, "gmail.com", *resp.Data.Email.Suggestion)
	assert.Contains(t, resp.Data.Email.Flags, "email_has_suggestion")
	assert.Equal(t, 0.9, resp.Data.Email.QualityScore) // 1.0 - 0.1 typo deduction
}

func TestInputValidate_Validate_VoIPPhone_ReturnsVoIPFlagAndReducedScore(t *testing.T) {
	t.Parallel()
	phoneSvc := &stubPhoneService{
		phoneResult: phone.ValidateResponse{
			Valid:     true,
			Formatted: "+14155552671",
			Country:   "US",
			Type:      "voip",
			Carrier:   nil,
			Risk:      &phone.Risk{IsVoIP: true, IsVirtual: false},
		},
	}

	r := setupRouter(&stubEmailService{}, phoneSvc, &stubProfanityService{}, &stubSentimentService{})
	w := post(t, r, `{"phone": "+14155552671"}`)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Response]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))

	assert.True(t, resp.Data.Phone.Risk.IsVoIP)
	assert.Contains(t, resp.Data.Phone.Flags, "phone_voip")
	assert.Equal(t, 0.7, resp.Data.Phone.QualityScore) // 1.0 - 0.3 voip deduction
}

func TestInputValidate_Validate_NoFieldsPresent_Returns422(t *testing.T) {
	t.Parallel()

	r := setupRouter(&stubEmailService{}, &stubPhoneService{}, &stubProfanityService{}, &stubSentimentService{})
	w := post(t, r, `{}`)

	require.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
