package whois

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

// LookupResponse is the JSON payload returned by the WHOIS endpoint.
type LookupResponse struct {
	Domain      string   `json:"domain"`
	Registrar   string   `json:"registrar,omitempty"`
	NameServers []string `json:"name_servers,omitempty"`
	Status      []string `json:"status,omitempty"`
	CreatedDate string   `json:"created_date,omitempty"`
	UpdatedDate string   `json:"updated_date,omitempty"`
	ExpiryDate  string   `json:"expiry_date,omitempty"`
	DNSSec      bool     `json:"dnssec"`
}

// BatchLookupItem represents a single entry in a WHOIS batch response.
type BatchLookupItem struct {
	Domain string         `json:"domain"`
	Found  bool           `json:"found"`
	Error  string         `json:"error,omitempty"`
	Data   LookupResponse `json:"data"`
}

// Does raw WHOIS queries.
type Querier interface {
	Whois(domain string, servers ...string) (string, error)
}

// Performs WHOIS lookups.
type Service struct {
	q Querier
}

// Creates a new WHOIS Service using the default whois client.
func NewService() *Service {
	return &Service{q: whois.DefaultClient}
}

// Returned when no WHOIS record is found for the domain.
var ErrDomainNotFound = errors.New("domain not found")

// Lookup queries WHOIS information for the given domain.
func (s *Service) Lookup(_ context.Context, domain string) (LookupResponse, error) {
	raw, err := s.q.Whois(domain)

	if err != nil {
		return LookupResponse{}, err
	}

	info, err := whoisparser.Parse(raw)

	if err != nil {
		if errors.Is(err, whoisparser.ErrNotFoundDomain) {
			return LookupResponse{}, ErrDomainNotFound
		}

		return LookupResponse{}, err
	}

	resp := LookupResponse{Domain: domain}

	if info.Domain != nil {
		resp.NameServers = info.Domain.NameServers
		resp.Status = info.Domain.Status
		resp.CreatedDate = info.Domain.CreatedDate
		resp.UpdatedDate = info.Domain.UpdatedDate
		resp.ExpiryDate = info.Domain.ExpirationDate
		resp.DNSSec = info.Domain.DNSSec
	}

	if info.Registrar != nil {
		resp.Registrar = info.Registrar.Name
	}

	return resp, nil
}

func (s *Service) LookupBatch(ctx context.Context, domains []string) ([]BatchLookupItem, error) {
	const (
		maxWorkers     = 10
		perItemTimeout = 3 * time.Second
	)

	results := make([]BatchLookupItem, len(domains))

	sem := make(chan struct{}, maxWorkers)

	var wg sync.WaitGroup

	for i, domain := range domains {
		wg.Add(1)

		sem <- struct{}{}

		go func(i int, domain string) {
			defer wg.Done()
			defer func() { <-sem }()

			itemCtx, cancel := context.WithTimeout(ctx, perItemTimeout)
			defer cancel()

			resp, err := s.Lookup(itemCtx, domain)
			if err != nil {
				results[i] = BatchLookupItem{
					Domain: domain,
					Found:  false,
					Error:  err.Error(),
				}

				return
			}

			results[i] = BatchLookupItem{
				Domain: domain,
				Found:  true,
				Data:   resp,
			}
		}(i, domain)
	}

	wg.Wait()

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
