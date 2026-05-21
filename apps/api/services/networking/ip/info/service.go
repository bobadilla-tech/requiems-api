package info

import (
	"context"
	"net"

	"github.com/bobadilla-tech/go-ip-intelligence/v2/ipi"
)

// LookupResponse is the JSON payload returned by the IP geolocation endpoint.
type LookupResponse struct {
	IP          string `json:"ip"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	City        string `json:"city"`
	ISP         string `json:"isp"`
	IsVPN       bool   `json:"is_vpn"`
}

// BatchIPInfoItem is the per-item result returned by CheckInfoBatch.
type BatchIPInfoItem struct {
	IP     string          `json:"ip"`
	Result *LookupResponse `json:"result,omitempty"`
	Error  string          `json:"error,omitempty"`
}

type Service struct {
	c *ipi.Client
}

func NewService(c *ipi.Client) *Service {
	if c == nil {
		return nil
	}
	return &Service{c: c}
}

// CheckInfoBatch looks up IP info for each address and returns results in input order.
// Per-item errors are absorbed in-band.
func (s *Service) CheckInfoBatch(ctx context.Context, ips []string) []BatchIPInfoItem {
	results := make([]BatchIPInfoItem, len(ips))
	for i, ip := range ips {
		r, err := s.CheckInfo(ctx, ip)
		if err != nil {
			results[i] = BatchIPInfoItem{IP: ip, Error: err.Error()}
		} else {
			results[i] = BatchIPInfoItem{IP: ip, Result: &r}
		}
	}
	return results
}

func (s *Service) CheckInfo(ctx context.Context, ip string) (LookupResponse, error) {
	result, err := s.c.CheckString(ctx, ip)
	if err != nil {
		return LookupResponse{}, err
	}
	return LookupResponse{
		IP:          net.IP(result.IP).String(),
		Country:     result.Country,
		CountryCode: result.CountryCode,
		City:        result.City,
		ISP:         result.AsnOrg,
		IsVPN:       result.IsVPN,
	}, nil
}
