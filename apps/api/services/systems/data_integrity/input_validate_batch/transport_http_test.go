package inputvalidatebatch

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
	inputvalidate "requiems-api/services/systems/data_integrity/input_validate"
)

func setupRouter(inputValidateSvc InputValidateService) chi.Router {
	r := chi.NewRouter()
	svc := NewService(inputValidateSvc)
	RegisterRoutes(r, svc)
	return r
}

type stubInputValidateService struct {
	responses map[string]inputvalidate.Response
}

func (s *stubInputValidateService) Validate(ctx context.Context, email, phone, text string) inputvalidate.Response {
	return s.responses[email]
}

func strPtr(s string) *string {
	return &s
}

type stubInputValidateServiceWithTimeout struct {
	responses map[string]inputvalidate.Response
	blockOn   string // email que va a causar timeout
}

func (s *stubInputValidateServiceWithTimeout) Validate(ctx context.Context, email, phone, text string) inputvalidate.Response {
	if email == s.blockOn {
		<-ctx.Done()
		return inputvalidate.Response{}
	}
	return s.responses[email]
}

func validEmailResult(email string) inputvalidate.Response {
	normalized := email
	country := "US"
	phoneNormalized := "+14155550001"

	return inputvalidate.Response{
		Email: &inputvalidate.EmailResult{
			Valid:        true,
			Normalized:   &normalized,
			Original:     email,
			SyntaxValid:  true,
			MXValid:      true,
			Disposable:   false,
			Flags:        []string{},
			QualityScore: 0.98,
		},
		Phone: &inputvalidate.PhoneResult{
			Valid:        true,
			Normalized:   &phoneNormalized,
			Country:      &country,
			Flags:        []string{},
			QualityScore: 0.95,
		},
		OverallQualityScore: 0.97,
	}
}

func post(t *testing.T, r chi.Router, body string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodPost, "/input/validate/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestValidateBatch_TwoValidItems(t *testing.T) {
	t.Parallel()

	stub := &stubInputValidateService{
		responses: map[string]inputvalidate.Response{
			"alice@example.com": validEmailResult("alice@example.com"),
			"bob@example.com":   validEmailResult("bob@example.com"),
		},
	}

	r := setupRouter(stub)

	w := post(t, r, `{
    "items": [
        { "email": "alice@example.com", "phone": "+14155550001" },
        { "email": "bob@example.com",   "phone": "+14155550002" }
    ]
}`)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[BatchResponse]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))

	assert.Equal(t, 2, resp.Data.Total)
	assert.Equal(t, 2, resp.Data.ValidCount)
	assert.Equal(t, 0, resp.Data.InvalidCount)
	assert.Greater(t, resp.Data.AverageQualityScore, 0.9)
}

func invalidEmailResult(email string) inputvalidate.Response {
	country := "US"
	phoneNormalized := "+14155550003"

	return inputvalidate.Response{
		Email: &inputvalidate.EmailResult{
			Valid:        false,
			Normalized:   nil,
			Original:     email,
			SyntaxValid:  false,
			MXValid:      false,
			Disposable:   false,
			Flags:        []string{"invalid_syntax"},
			QualityScore: 0.0,
		},
		Phone: &inputvalidate.PhoneResult{
			Valid:        true,
			Normalized:   &phoneNormalized,
			Country:      &country,
			Flags:        []string{},
			QualityScore: 0.95,
		},
		OverallQualityScore: 0.0,
	}
}

func TestValidateBatch_MixValidAndInvalid(t *testing.T) {
	t.Parallel()

	stub := &stubInputValidateService{
		responses: map[string]inputvalidate.Response{
			"alice@example.com": validEmailResult("alice@example.com"),
			"invalid-email":     invalidEmailResult("invalid-email"),
		},
	}

	r := setupRouter(stub)

	w := post(t, r, `{
		"items": [
			{ "email": "alice@example.com", "phone": "+14155550001" },
			{ "email": "invalid-email",     "phone": "+14155550003" }
		]
	}`)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[BatchResponse]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))

	assert.Equal(t, 2, resp.Data.Total)
	assert.Equal(t, 1, resp.Data.ValidCount)
	assert.Equal(t, 1, resp.Data.InvalidCount)

	// per-item valid differs
	assert.True(t, resp.Data.Results[0].Email.Valid)
	assert.False(t, resp.Data.Results[1].Email.Valid)
}

func TestValidateBatch_OneItemTimeout(t *testing.T) {
	t.Parallel()

	stub := &stubInputValidateServiceWithTimeout{
		blockOn: "slow@example.com",
		responses: map[string]inputvalidate.Response{
			"alice@example.com": validEmailResult("alice@example.com"),
			"bob@example.com":   validEmailResult("bob@example.com"),
		},
	}

	r := setupRouter(stub)

	w := post(t, r, `{
        "items": [
            { "email": "alice@example.com", "phone": "+14155550001" },
            { "email": "slow@example.com",  "phone": "+14155550002" },
            { "email": "bob@example.com",   "phone": "+14155550003" }
        ]
    }`)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[BatchResponse]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))

	assert.Equal(t, 3, resp.Data.Total)
	assert.Equal(t, 2, resp.Data.ValidCount)
	assert.Equal(t, 1, resp.Data.InvalidCount)

	assert.NotNil(t, resp.Data.Results[1].Error)

	assert.Nil(t, resp.Data.Results[0].Error)
	assert.Nil(t, resp.Data.Results[2].Error)
}

func TestValidateBatch_Over50Items(t *testing.T) {
	t.Parallel()

	stub := &stubInputValidateService{
		responses: map[string]inputvalidate.Response{},
	}

	r := setupRouter(stub)

	items := make([]Item, 51)
	for i := range items {
		items[i] = Item{
			Email: "contact@example.com",
			Phone: "+14155550001",
		}
	}

	body, _ := json.Marshal(Request{
		Items: items,
	})

	w := post(t, r, string(body))

	require.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestValidateBatch_EmptyItems(t *testing.T) {
	t.Parallel()

	stub := &stubInputValidateService{
		responses: map[string]inputvalidate.Response{},
	}

	r := setupRouter(stub)

	w := post(t, r, `{}`)

	require.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestValidateBatch_SetsUsageCountHeader(t *testing.T) {
	stub := &stubInputValidateService{
		responses: map[string]inputvalidate.Response{
			"alice@example.com": validEmailResult("alice@example.com"),
			"invalid-email":     invalidEmailResult("invalid-email"),
		},
	}

	r := setupRouter(stub)

	w := post(t, r, `{
		"items": [
			{ "email": "alice@example.com", "phone": "+14155550001" },
			{ "email": "invalid-email",     "phone": "+14155550003" }
		]
	}`)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	if got := w.Header().Get("X-Usage-Count"); got != "2" {
		t.Errorf("Expected X-Usage-Count: 2, got %q", got)
	}
}
