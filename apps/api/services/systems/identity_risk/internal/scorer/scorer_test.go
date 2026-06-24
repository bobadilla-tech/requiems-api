package scorer

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"

	ipinfo "requiems-api/services/networking/ip/info"
	ipvpn "requiems-api/services/networking/ip/vpn"
	"requiems-api/services/validation/email"
	"requiems-api/services/validation/phone"
)

type stubEmail struct{ r email.Validation }

func (s *stubEmail) ValidateEmail(_ context.Context, _ string) email.Validation { return s.r }

type stubPhone struct{ r phone.ValidateResponse }

func (s *stubPhone) Validate(_ string) phone.ValidateResponse { return s.r }

type stubVPN struct {
	r   ipvpn.IPCheckResponse
	err error
}

func (s *stubVPN) CheckIP(_ context.Context, _ net.IP) (ipvpn.IPCheckResponse, error) {
	return s.r, s.err
}

type stubIPInfo struct {
	r   ipinfo.LookupResponse
	err error
}

func (s *stubIPInfo) CheckInfo(_ context.Context, _ string) (ipinfo.LookupResponse, error) {
	return s.r, s.err
}

func cleanEmail() email.Validation {
	v, mx := true, true
	normalized := "user@example.com"
	domain := "example.com"
	return email.Validation{Valid: v, SyntaxValid: true, MxValid: mx, Disposable: false, Normalized: &normalized, Domain: &domain}
}

func TestResolve_CleanEmail_SetsExpectedSignals(t *testing.T) {
	ctx := context.Background()

	resolved := Resolve(
		ctx,
		&stubEmail{r: cleanEmail()},
		&stubPhone{},
		nil,
		nil,
		"user@example.com", "", "",
	)

	assert.True(t, resolved.Signals.EmailPresent)
	assert.False(t, resolved.Signals.EmailNoMX)
	assert.False(t, resolved.Signals.EmailDisposable)
	assert.False(t, resolved.Signals.EmailInvalid)
	assert.False(t, resolved.Signals.IPRiskChecked)
}

func TestResolve_PhonePresent_SetsExpectedSignals(t *testing.T) {
	ctx := context.Background()

	resolved := Resolve(
		ctx,
		&stubEmail{r: cleanEmail()},
		&stubPhone{r: phone.ValidateResponse{Valid: true, Country: "US"}},
		nil,
		nil,
		"user@example.com", "+13235551234", "",
	)

	assert.True(t, resolved.Signals.PhonePresent)
	assert.False(t, resolved.Signals.PhoneInvalid)
	assert.Equal(t, "US", resolved.Signals.PhoneCountry)
}

func TestResolve_IPRiskChecked_WhenVPNSucceeds(t *testing.T) {
	ctx := context.Background()

	resolved := Resolve(
		ctx,
		&stubEmail{r: cleanEmail()},
		&stubPhone{},
		&stubVPN{r: ipvpn.IPCheckResponse{IsProxy: false, IsTor: false, FraudScore: 0}},
		&stubIPInfo{},
		"user@example.com", "", "8.8.8.8",
	)

	assert.True(t, resolved.Signals.IPRiskChecked)
	assert.False(t, resolved.Signals.IsProxy)
	assert.False(t, resolved.Signals.IsTOR)
	assert.Equal(t, 0, resolved.Signals.FraudScore)
}

func TestSignalConfidence_NoSignals(t *testing.T) {
	s := Signals{}

	got := signalConfidence(s)

	assert.Equal(t, 0.0, got)
}

func TestSignalConfidence_SingleHighQualitySignal(t *testing.T) {
	s := Signals{
		EmailPresent:    true,
		EmailDisposable: false,
		EmailNoMX:       false,
	}

	got := signalConfidence(s)

	assert.Equal(t, 1.0, got)
}

func TestSignalConfidence_EmailAllRisk(t *testing.T) {
	s := Signals{
		EmailPresent:    true,
		EmailNoMX:       true,
		EmailDisposable: true,
	}

	got := signalConfidence(s)

	assert.Equal(t, 0.0, got)
}

func TestSignalConfidence_AllSignalsClean(t *testing.T) {
	s := Signals{
		EmailPresent:    true,
		EmailNoMX:       false,
		EmailDisposable: false,

		PhonePresent: true,
		PhoneVoIP:    false,
		PhoneVirtual: false,

		IPPresent:  true,
		IsProxy:    false,
		IsTOR:      false,
		FraudScore: 0,
	}

	got := signalConfidence(s)

	assert.Equal(t, 1.0, got)
}

func TestSignalConfidence_Mixed(t *testing.T) {
	s := Signals{
		EmailPresent:    true,
		EmailNoMX:       false,
		EmailDisposable: true,

		PhonePresent: true,
		PhoneVoIP:    false,
		PhoneVirtual: false,

		IPPresent:  true,
		IsProxy:    true,
		IsTOR:      false,
		FraudScore: 0,
	}

	got := signalConfidence(s)

	assert.Greater(t, got, 0.0)
	assert.Less(t, got, 1.0)
}

func TestCompute_ConfidenceMatchesSignalConfidence(t *testing.T) {
	s := Signals{
		EmailPresent:    true,
		EmailNoMX:       false,
		EmailDisposable: false,
	}

	result := Compute(s)
	expected := signalConfidence(s)

	assert.Equal(t, expected, result.Confidence)
}

func TestCompute_IsSafe_LowScoreHighConfidence(t *testing.T) {
	s := Signals{
		EmailPresent:    true,
		EmailNoMX:       false,
		EmailDisposable: false,
	}

	result := Compute(s)

	assert.True(t, result.IsSafe)
}

func TestCompute_IsSafe_FalseWhenTOROrProxy(t *testing.T) {
	tor := Signals{
		EmailPresent:    true,
		EmailNoMX:       false,
		EmailDisposable: false,
		IPPresent:       true,
		IsTOR:           true,
		FraudScore:      0,
	}

	proxy := Signals{
		EmailPresent:    true,
		EmailNoMX:       false,
		EmailDisposable: false,
		IPPresent:       true,
		IsProxy:         true,
		FraudScore:      0,
	}

	assert.False(t, Compute(tor).IsSafe)
	assert.False(t, Compute(proxy).IsSafe)
}

func TestSignalConfidence_IPAllClean(t *testing.T) {
	s := Signals{
		IPPresent:     true,
		IPRiskChecked: true,
		IsProxy:       false,
		IsTOR:         false,
		FraudScore:    0,
	}

	got := signalConfidence(s)

	assert.Equal(t, 1.0, got)
}

func TestSignalConfidence_IPAllRisk(t *testing.T) {
	s := Signals{
		IPPresent:     true,
		IPRiskChecked: true,
		IsProxy:       true,
		IsTOR:         true,
		FraudScore:    100,
	}

	got := signalConfidence(s)

	assert.Equal(t, 0.0, got)
}
