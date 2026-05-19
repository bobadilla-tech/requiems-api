package mx

import (
	"context"
	"net"
	"sort"
)

// Record represents a single MX record entry.
type Record struct {
	Host     string `json:"host"`
	Priority uint16 `json:"priority"`
}

// LookupResponse is the JSON payload returned by the MX lookup endpoint.
type LookupResponse struct {
	Domain  string   `json:"domain"`
	Records []Record `json:"records"`
}

func (LookupResponse) IsData() {}

// BatchLookupItem represents the result of a single domain lookup in a batch.
type BatchLookupItem struct {
	Domain string         `json:"domain"`
	Found  bool           `json:"found"`
	Error  string         `json:"error,omitempty"`
	Data   LookupResponse `json:"data,omitempty"`
}

// Service performs MX DNS record lookups.
type Service struct{}

// NewService returns a new MX lookup Service.
func NewService() *Service {
	return &Service{}
}

// Lookup queries the MX records for the given domain and returns them sorted
// by priority (ascending — lower value means higher priority).
func (s *Service) Lookup(ctx context.Context, domain string) (LookupResponse, error) {
	records, err := net.DefaultResolver.LookupMX(ctx, domain)
	if err != nil {
		return LookupResponse{}, err
	}

	sorted := make([]Record, 0, len(records))
	for _, mx := range records {
		sorted = append(sorted, Record{
			Host:     mx.Host,
			Priority: mx.Pref,
		})
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Priority < sorted[j].Priority
	})

	return LookupResponse{
		Domain:  domain,
		Records: sorted,
	}, nil
}

func (s *Service) LookupBatch(ctx context.Context, domains []string) []BatchLookupItem {
	results := make([]BatchLookupItem, len(domains))

	for i, d := range domains {

		result, err := s.Lookup(ctx, d)

		if err != nil {
			results[i] = BatchLookupItem{Domain: d, Found: false, Error: err.Error()}
			continue
		}

		results[i] = BatchLookupItem{
			Domain: d,
			Found:  true,
			Data:   result,
		}
	}

	return results
}
