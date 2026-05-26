package riskscore

import (
	"context"

	"requiems-api/services/systems/identity_risk/internal/scorer"
)

// Service computes a risk score without per-signal breakdown.
type Service struct {
	email  scorer.EmailChecker
	phone  scorer.PhoneChecker
	vpn    scorer.VPNChecker
	ipInfo scorer.IPInfoChecker
}

// NewService returns a new risk score Service.
func NewService(e scorer.EmailChecker, p scorer.PhoneChecker, v scorer.VPNChecker, i scorer.IPInfoChecker) *Service {
	return &Service{email: e, phone: p, vpn: v, ipInfo: i}
}

// Request is the input for POST /risk/score.
// At least one field must be present (validated at transport layer).
type Request struct {
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	IPAddress string `json:"ip_address"`
}

// Result is the response — identical computation to signup_protect but
// without the signals breakdown object.
type Result struct {
	RiskScore  float64  `json:"risk_score"`
	IsSafe     bool     `json:"is_safe"`
	Confidence float64  `json:"confidence"`
	Flags      []string `json:"flags"`
}

// Score fans out to the relevant services in parallel and returns the risk score.
func (s *Service) Score(ctx context.Context, req Request) (Result, error) {
	resolved := scorer.Resolve(ctx, s.email, s.phone, s.vpn, s.ipInfo,
		req.Email, req.Phone, req.IPAddress)
	r := scorer.Compute(resolved.Signals)
	flags := r.Flags
	if flags == nil {
		flags = []string{}
	}
	return Result{
		RiskScore:  r.RiskScore,
		IsSafe:     r.IsSafe,
		Confidence: r.Confidence,
		Flags:      flags,
	}, nil
}
