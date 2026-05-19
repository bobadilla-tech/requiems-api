package asn

import (
	"context"

	"github.com/bobadilla-tech/go-ip-intelligence/v2/ipi"
)

// IPAddressASNResponse is the JSON payload returned by the ASN lookup endpoint.
type IPAddressASNResponse struct {
	IP     string `json:"ip"`
	ASN    string `json:"asn"`
	Org    string `json:"org"`
	ISP    string `json:"isp"`
	Domain string `json:"domain"`
	Route  string `json:"route"`
	Type   string `json:"type"`
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

func (s *Service) CheckASN(ctx context.Context, ip string) (IPAddressASNResponse, error) {
	info, err := s.c.CheckASNString(ctx, ip)
	
	if err != nil {
		return IPAddressASNResponse{}, err
	}

	return IPAddressASNResponse{
		IP:     info.IP,
		ASN:    info.ASN,
		Org:    info.Org,
		ISP:    info.ISP,
		Domain: info.Domain,
		Route:  info.Route,
		Type:   info.Type,
	}, nil
}
