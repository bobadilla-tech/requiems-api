package vpn

import (
	"context"
	"net"

	"github.com/bobadilla-tech/go-ip-intelligence/v2/ipi"
)

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

func (s *Service) CheckBatch(ctx context.Context, ips []string) (BatchResponse, error) {
	results := make([]IPCheckResponse, len(ips))

	type item struct {
		index  int
		result IPCheckResponse
	}

	ch := make(chan item, len(ips))

	for i, rawIP := range ips {
		go func(index int, raw string) {
			ip := net.ParseIP(raw)
			if ip == nil {
				ch <- item{
					index: index,
					result: IPCheckResponse{
						IP: raw,
					},
				}
				return
			}

			result, err := s.CheckIP(ctx, ip)
			if err != nil {
				ch <- item{
					index: index,
					result: IPCheckResponse{
						IP: raw,
					},
				}
				return
			}

			ch <- item{
				index:  index,
				result: result,
			}
		}(i, rawIP)
	}

	for range ips {
		item := <-ch
		results[item.index] = item.result
	}

	return BatchResponse{
		Results: results,
	}, nil
}
