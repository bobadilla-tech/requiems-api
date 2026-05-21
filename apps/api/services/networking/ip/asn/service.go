package asn

import (
	"context"
	"net"

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

// BatchASNItem is the per-item result returned by CheckASNBatch.
type BatchASNItem struct {
	IP     string                `json:"ip"`
	Result *IPAddressASNResponse `json:"result,omitempty"`
	Error  string                `json:"error,omitempty"`
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

// CheckASNBatch looks up ASN info for each IP and returns results in input order.
// Private/reserved IPs return a zero-value result with no error, matching single-item behaviour.
// Per-item errors are absorbed in-band.
func (s *Service) CheckASNBatch(ctx context.Context, ips []string) []BatchASNItem {
	results := make([]BatchASNItem, len(ips))
	for i, ip := range ips {
		if parsed := net.ParseIP(ip); parsed != nil &&
			(parsed.IsPrivate() || parsed.IsLoopback() || parsed.IsLinkLocalUnicast()) {
			results[i] = BatchASNItem{IP: ip, Result: &IPAddressASNResponse{IP: ip}}
			continue
		}
		r, err := s.CheckASN(ctx, ip)
		if err != nil {
			results[i] = BatchASNItem{IP: ip, Error: err.Error()}
		} else {
			results[i] = BatchASNItem{IP: ip, Result: &r}
		}
	}
	return results
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
