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

type Service struct {
	c *ipi.Client
}

func NewService(c *ipi.Client) *Service {
	if c == nil {
		return nil
	}
	return &Service{c: c}
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
