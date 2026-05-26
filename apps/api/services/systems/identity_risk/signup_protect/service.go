package signupprotect

import (
	"context"

	"requiems-api/services/systems/identity_risk/internal/scorer"
)

// Service composes multiple signals into a full signup protection response.
type Service struct {
	email  scorer.EmailChecker
	phone  scorer.PhoneChecker
	vpn    scorer.VPNChecker
	ipInfo scorer.IPInfoChecker
}

// NewService returns a new signup protect Service.
func NewService(e scorer.EmailChecker, p scorer.PhoneChecker, v scorer.VPNChecker, i scorer.IPInfoChecker) *Service {
	return &Service{email: e, phone: p, vpn: v, ipInfo: i}
}

// Request is the input for POST /signup/protect.
type Request struct {
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	IPAddress string `json:"ip_address"`
}

// EmailSignal is the per-signal email breakdown.
type EmailSignal struct {
	Valid      bool    `json:"valid"`
	Disposable bool    `json:"disposable"`
	MXValid    bool    `json:"mx_valid"`
	Suggestion *string `json:"suggestion"`
}

// PhoneSignal is the per-signal phone breakdown.
type PhoneSignal struct {
	Valid     bool   `json:"valid"`
	Country   string `json:"country"`
	IsVoIP    bool   `json:"is_voip"`
	IsVirtual bool   `json:"is_virtual"`
}

// IPSignal is the per-signal IP breakdown.
type IPSignal struct {
	CountryCode string  `json:"country_code"`
	IsVPN       bool    `json:"is_vpn"`
	IsProxy     bool    `json:"is_proxy"`
	IsTOR       bool    `json:"is_tor"`
	IsHosting   bool    `json:"is_hosting"`
	FraudScore  float64 `json:"fraud_score"` // normalised 0–1
}

// Signals groups the per-signal breakdowns; omitted fields are null.
type Signals struct {
	Email *EmailSignal `json:"email"`
	Phone *PhoneSignal `json:"phone"`
	IP    *IPSignal    `json:"ip"`
}

// Result is the full signup protection response.
type Result struct {
	RiskScore  float64  `json:"risk_score"`
	IsSafe     bool     `json:"is_safe"`
	Confidence float64  `json:"confidence"`
	Flags      []string `json:"flags"`
	Signals    Signals  `json:"signals"`
}

// Protect fans out to the relevant services and returns the full breakdown.
func (s *Service) Protect(ctx context.Context, req Request) (Result, error) {
	resolved := scorer.Resolve(ctx, s.email, s.phone, s.vpn, s.ipInfo,
		req.Email, req.Phone, req.IPAddress)
	r := scorer.Compute(resolved.Signals)

	flags := r.Flags
	if flags == nil {
		flags = []string{}
	}

	var sigs Signals

	if resolved.EmailResult != nil {
		sigs.Email = &EmailSignal{
			Valid:      resolved.EmailResult.Valid,
			Disposable: resolved.EmailResult.Disposable,
			MXValid:    resolved.EmailResult.MxValid,
			Suggestion: resolved.EmailResult.Suggestion,
		}
	}

	if resolved.PhoneResult != nil {
		ps := &PhoneSignal{
			Valid:   resolved.PhoneResult.Valid,
			Country: resolved.PhoneResult.Country,
		}
		if resolved.PhoneResult.Risk != nil {
			ps.IsVoIP = resolved.PhoneResult.Risk.IsVoIP
			ps.IsVirtual = resolved.PhoneResult.Risk.IsVirtual
		}
		sigs.Phone = ps
	}

	if resolved.VPNResult != nil || resolved.IPResult != nil {
		ip := &IPSignal{}
		if resolved.VPNResult != nil {
			ip.IsVPN = resolved.VPNResult.IsVPN
			ip.IsProxy = resolved.VPNResult.IsProxy
			ip.IsTOR = resolved.VPNResult.IsTor
			ip.IsHosting = resolved.VPNResult.IsHosting
			ip.FraudScore = float64(resolved.VPNResult.FraudScore) / 100.0
		}
		if resolved.IPResult != nil {
			ip.CountryCode = resolved.IPResult.CountryCode
		}
		sigs.IP = ip
	}

	return Result{
		RiskScore:  r.RiskScore,
		IsSafe:     r.IsSafe,
		Confidence: r.Confidence,
		Flags:      flags,
		Signals:    sigs,
	}, nil
}
