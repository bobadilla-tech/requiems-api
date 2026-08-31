package scorer

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
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
		nil,
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
		nil,
	)

	assert.True(t, resolved.Signals.PhonePresent)
	assert.False(t, resolved.Signals.PhoneInvalid)
	assert.Equal(t, "US", resolved.Signals.PhoneCountry)
}

func TestResolve_IPRiskChecked_WhenVPNSucceeds(t *testing.T) {
	ctx := context.Background()

	mr := miniredis.RunT(t)
	rdb := redis.NewClient(&redis.Options{Addr: mr.Addr()})

	resolved := Resolve(
		ctx,
		&stubEmail{r: cleanEmail()},
		&stubPhone{},
		&stubVPN{r: ipvpn.IPCheckResponse{IsProxy: false, IsTor: false, FraudScore: 0}},
		&stubIPInfo{},
		"user@example.com", "", "8.8.8.8",
		rdb,
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
func TestResolve_Velocity_FirstCheck_BothCountersAreOne(t *testing.T) {
	ctx := context.Background()
	mr := miniredis.RunT(t)
	rdb := redis.NewClient(&redis.Options{Addr: mr.Addr()})

	resolved := Resolve(
		ctx,
		&stubEmail{r: cleanEmail()},
		&stubPhone{},
		&stubVPN{r: ipvpn.IPCheckResponse{}},
		&stubIPInfo{},
		"user@example.com", "", "8.8.8.8",
		rdb,
	)

	assert.True(t, resolved.Signals.VelocityChecked)
	assert.Equal(t, int64(1), resolved.Signals.VelocityCount1h)
	assert.Equal(t, int64(1), resolved.Signals.VelocityCount24h)
}

func TestResolve_Velocity_1hAnd24hAreIndependentCounters(t *testing.T) {
	ctx := context.Background()
	mr := miniredis.RunT(t)
	rdb := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	mr.Set("identity_risk:ip:8.8.8.8:1h", "5")
	mr.Set("identity_risk:ip:8.8.8.8:24h", "40")

	resolved := Resolve(
		ctx,
		&stubEmail{r: cleanEmail()},
		&stubPhone{},
		&stubVPN{r: ipvpn.IPCheckResponse{}},
		&stubIPInfo{},
		"user@example.com", "", "8.8.8.8",
		rdb,
	)

	// each key increments from its own prior value, they don't collide
	assert.Equal(t, int64(6), resolved.Signals.VelocityCount1h)
	assert.Equal(t, int64(41), resolved.Signals.VelocityCount24h)
}

func TestResolve_Velocity_IncrementsAcrossMultipleCalls(t *testing.T) {
	ctx := context.Background()
	mr := miniredis.RunT(t)
	rdb := redis.NewClient(&redis.Options{Addr: mr.Addr()})

	var resolved Resolved
	for range 3 {
		resolved = Resolve(ctx, &stubEmail{r: cleanEmail()}, &stubPhone{}, &stubVPN{r: ipvpn.IPCheckResponse{}}, &stubIPInfo{}, "user@example.com", "", "8.8.8.8", rdb)
	}

	assert.Equal(t, int64(3), resolved.Signals.VelocityCount1h)
	assert.Equal(t, int64(3), resolved.Signals.VelocityCount24h)
}

func TestResolve_Velocity_KeysHaveDistinctTTLs(t *testing.T) {
	ctx := context.Background()
	mr := miniredis.RunT(t)
	rdb := redis.NewClient(&redis.Options{Addr: mr.Addr()})

	Resolve(ctx, &stubEmail{r: cleanEmail()}, &stubPhone{}, &stubVPN{r: ipvpn.IPCheckResponse{}}, &stubIPInfo{}, "user@example.com", "", "8.8.8.8", rdb)

	ttl1h := mr.TTL("identity_risk:ip:8.8.8.8:1h")
	ttl24h := mr.TTL("identity_risk:ip:8.8.8.8:24h")

	assert.True(t, ttl1h > 0 && ttl1h <= time.Hour)
	assert.True(t, ttl24h > time.Hour && ttl24h <= 24*time.Hour)
}

func TestResolve_Velocity_DifferentIPsHaveIndependentCounters(t *testing.T) {
	ctx := context.Background()
	mr := miniredis.RunT(t)
	rdb := redis.NewClient(&redis.Options{Addr: mr.Addr()})

	Resolve(ctx, &stubEmail{r: cleanEmail()}, &stubPhone{}, &stubVPN{r: ipvpn.IPCheckResponse{}}, &stubIPInfo{}, "user@example.com", "", "1.1.1.1", rdb)
	Resolve(ctx, &stubEmail{r: cleanEmail()}, &stubPhone{}, &stubVPN{r: ipvpn.IPCheckResponse{}}, &stubIPInfo{}, "user@example.com", "", "1.1.1.1", rdb)
	resolved := Resolve(ctx, &stubEmail{r: cleanEmail()}, &stubPhone{}, &stubVPN{r: ipvpn.IPCheckResponse{}}, &stubIPInfo{}, "user@example.com", "", "2.2.2.2", rdb)

	assert.Equal(t, int64(1), resolved.Signals.VelocityCount1h) // new IP, unaffected by 1.1.1.1's history
}

func TestResolve_Velocity_NotCheckedWhenNoIP(t *testing.T) {
	ctx := context.Background()

	resolved := Resolve(
		ctx,
		&stubEmail{r: cleanEmail()},
		&stubPhone{},
		nil,
		nil,
		"user@example.com", "", "",
		nil,
	)

	assert.False(t, resolved.Signals.VelocityChecked)
	assert.Equal(t, int64(0), resolved.Signals.VelocityCount1h)
	assert.Equal(t, int64(0), resolved.Signals.VelocityCount24h)
}

func TestCompute_Velocity_LowTier_NoScoreNoFlag(t *testing.T) {
	s := Signals{VelocityChecked: true, VelocityCount1h: 10}

	result := Compute(s)

	assert.Equal(t, 0.0, result.RiskScore)
	assert.Empty(t, result.Flags)
}

func TestCompute_Velocity_MediumTier_At11(t *testing.T) {
	s := Signals{VelocityChecked: true, VelocityCount1h: 11}

	result := Compute(s)

	assert.Equal(t, 0.10, result.RiskScore)
	assert.Contains(t, result.Flags, "velocity_medium")
}

func TestCompute_Velocity_MediumTier_At50_StillMedium(t *testing.T) {
	// scoreVelocity uses > 50 for high, so exactly 50 still falls in medium (> 10)
	s := Signals{VelocityChecked: true, VelocityCount1h: 50}

	result := Compute(s)

	assert.Equal(t, 0.10, result.RiskScore)
	assert.Contains(t, result.Flags, "velocity_medium")
	assert.NotContains(t, result.Flags, "velocity_high")
}

func TestCompute_Velocity_HighTier_Above50(t *testing.T) {
	s := Signals{VelocityChecked: true, VelocityCount1h: 51}

	result := Compute(s)

	assert.Equal(t, 0.25, result.RiskScore)
	assert.Contains(t, result.Flags, "velocity_high")
}

func TestCompute_Velocity_NotCheckedContributesZero(t *testing.T) {
	s := Signals{VelocityChecked: false, VelocityCount1h: 999} // should be ignored when not checked

	result := Compute(s)

	assert.Equal(t, 0.0, result.RiskScore)
	assert.Empty(t, result.Flags)
}

func TestSignalConfidence_VelocityLow_FullWeight(t *testing.T) {
	s := Signals{VelocityChecked: true, VelocityCount1h: 5}

	got := signalConfidence(s)

	assert.Equal(t, 1.0, got) // only signal present, and it's "clean" (<=10)
}

func TestSignalConfidence_VelocityHigh_ZeroWeight(t *testing.T) {
	s := Signals{VelocityChecked: true, VelocityCount1h: 51}

	got := signalConfidence(s)

	assert.Equal(t, 0.0, got) // only signal present, and it's well over the threshold (>10)
}

func TestSignalConfidence_VelocityNotChecked_ExcludedFromDenominator(t *testing.T) {
	withoutVelocity := Signals{EmailPresent: true}

	// VelocityCount1h alone (without VelocityChecked=true) must not affect the result
	assert.Equal(t, signalConfidence(withoutVelocity), signalConfidence(Signals{EmailPresent: true, VelocityCount1h: 999}))
}

func TestSignalConfidence_VelocityCheckedAndClean_MatchesUnchecked(t *testing.T) {
	withoutVelocity := Signals{EmailPresent: true}
	withVelocityCheckedAndClean := Signals{EmailPresent: true, VelocityChecked: true, VelocityCount1h: 1}

	assert.Equal(t, signalConfidence(withoutVelocity), signalConfidence(withVelocityCheckedAndClean))
}

func TestSignalConfidence_VelocityCheckedAndDirty_LowersConfidence(t *testing.T) {
	withoutVelocity := Signals{EmailPresent: true}
	withVelocityCheckedAndDirty := Signals{EmailPresent: true, VelocityChecked: true, VelocityCount1h: 999}

	assert.Less(t, signalConfidence(withVelocityCheckedAndDirty), signalConfidence(withoutVelocity))
}
