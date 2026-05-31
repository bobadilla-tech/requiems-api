package vpn

import (
	"context"
	"errors"
	"net"
	"sync"

	"github.com/bobadilla-tech/go-ip-intelligence/v2/ipi"
)

// IPCheckResponse is the JSON payload returned by the VPN/proxy detection endpoint.
type IPCheckResponse struct {
	// IP is the analysed address.
	IP string `json:"ip"`
	// IsVPN is true when the address belongs to a known VPN provider.
	IsVPN bool `json:"is_vpn"`
	// IsProxy is true when the address is a known public or web proxy.
	IsProxy bool `json:"is_proxy"`
	// IsTor is true when the address is a known Tor exit node, detected via the
	// IP2Proxy database or the Tor Project's DNSBL.
	IsTor bool `json:"is_tor"`
	// IsHosting is true when the address belongs to a data-centre or hosting
	// provider range (DCH in IP2Proxy terminology).
	IsHosting bool `json:"is_hosting"`
	// Score is the raw threat score used to derive Threat.
	// Tor contributes 3, VPN or Proxy each contribute 2, Hosting contributes 1.
	Score int `json:"score"`
	// Threat is the threat level derived from Score:
	// 0 → None, 1 → Low, 2–3 → Medium, 4–5 → High, ≥6 → Critical.
	Threat ipi.ThreatLevel `json:"threat"`
	// FraudScore is populated when using an IP2Proxy database of tier PX5 or
	// higher. It ranges from 0 (no fraud risk) to 100 (high fraud risk). Zero
	// means the value is unavailable for the current database tier.
	FraudScore int `json:"fraud_score"`
	// AsnOrg is the name of the organisation that owns the Autonomous System
	// containing the address (e.g. "DIGITALOCEAN-ASN"). Empty when the ASN
	// lookup returns no record.
	AsnOrg string `json:"asn_org"`
}
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

const maxWorkers = 10

func (s *Service) CheckBatch(ctx context.Context, ips []string) ([]IPCheckResponse, error) {
	results := make([]IPCheckResponse, len(ips))

	type itemResult struct {
		index  int
		result IPCheckResponse
	}

	ch := make(chan itemResult, len(ips))
	sem := make(chan struct{}, maxWorkers)
	var wg sync.WaitGroup

	for i, rawIP := range ips {
		wg.Add(1)
		sem <- struct{}{}
		go func(index int, raw string) {
			defer wg.Done()
			defer func() { <-sem }()

			ip := net.ParseIP(raw)
			if ip == nil {
				ch <- itemResult{index: index, result: IPCheckResponse{IP: raw}}
				return
			}

			result, err := s.CheckIP(ctx, ip)
			if err != nil {
				if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
					ch <- itemResult{index: index, result: IPCheckResponse{}}
					return
				}
				ch <- itemResult{index: index, result: IPCheckResponse{IP: raw}}
				return
			}

			ch <- itemResult{index: index, result: result}
		}(i, rawIP)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for item := range ch {
		results[item.index] = item.result
	}

	return results, nil
}
