package vpn

import (
	"context"
	"net"

	"github.com/bobadilla-tech/go-ip-intelligence/v2/ipi"
)

// IPCheckResponse is the JSON payload returned by the VPN/proxy detection endpoint.
type IPCheckResponse struct {
	IP        string          `json:"ip"`
	IsVPN     bool            `json:"is_vpn"`
	IsProxy   bool            `json:"is_proxy"`
	IsTor     bool            `json:"is_tor"`
	IsHosting bool            `json:"is_hosting"`
	Score     int             `json:"score"`
	Threat    ipi.ThreatLevel `json:"threat"`
	FraudScore int            `json:"fraud_score"`
	AsnOrg    string          `json:"asn_org"`
}

func (IPCheckResponse) IsData() {}

type Service struct {
	c *ipi.Client
}

func NewService(c *ipi.Client) *Service {
	if c == nil {
		return nil
	}
	return &Service{
		c: c,
	}
}

func (s *Service) CheckIP(ctx context.Context, ip net.IP) (IPCheckResponse, error) {
	result, err := s.c.Check(ctx, ip)
	if err != nil {
		return IPCheckResponse{}, err
	}
	return IPCheckResponse{
		IP:         ip.String(),
		IsVPN:      result.IsVPN,
		IsProxy:    result.IsProxy,
		IsTor:      result.IsTor,
		IsHosting:  result.IsHosting,
		Score:      result.Score,
		Threat:     result.Threat,
		FraudScore: result.FraudScore,
		AsnOrg:     result.AsnOrg,
	}, nil
}
