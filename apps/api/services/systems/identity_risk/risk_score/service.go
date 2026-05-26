package riskscore

import (
	"context"

	"requiems-api/services/systems/identity_risk/internal/scorer"
)

type Service struct {
	email  scorer.EmailChecker
	phone  scorer.PhoneChecker
	vpn    scorer.VPNChecker
	ipInfo scorer.IPInfoChecker
}

func NewService(e scorer.EmailChecker, p scorer.PhoneChecker, v scorer.VPNChecker, i scorer.IPInfoChecker) *Service {
	return &Service{email: e, phone: p, vpn: v, ipInfo: i}
}

type Result struct {
	RiskScore  float64  `json:"risk_score"`
	IsSafe     bool     `json:"is_safe"`
	Confidence float64  `json:"confidence"`
	Flags      []string `json:"flags"`
}

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
