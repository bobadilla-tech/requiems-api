package disposable

import (
	"strings"

	disposable "github.com/bobadilla-tech/is-email-disposable"

	"requiems-api/platform/svcerr"
)

// CheckEmailResponse represents the response for a single email check.
type CheckEmailResponse struct {
	Email        string `json:"email"`
	IsDisposable bool   `json:"is_disposable"`
	Domain       string `json:"domain,omitempty"`
}

// DomainCheckResponse represents the response for a domain check.
type DomainCheckResponse struct {
	Domain       string `json:"domain"`
	IsDisposable bool   `json:"is_disposable"`
}

// DomainsListResponse represents the response for listing all domains.
type DomainsListResponse struct {
	Domains []string `json:"domains"`
	Total   int      `json:"total"`
	Page    int      `json:"page"`
	PerPage int      `json:"per_page"`
	HasMore bool     `json:"has_more"`
}

// StatsResponse represents statistics about disposable domains.
type StatsResponse struct {
	TotalDomains int `json:"total_domains"`
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

// CheckEmail checks if a single email is disposable
func (s *Service) CheckEmail(email string) CheckEmailResponse {
	isDisposable := disposable.IsDisposable(email)

	domain := ExtractDomainFromEmail(email)

	return CheckEmailResponse{
		Email:        email,
		IsDisposable: isDisposable,
		Domain:       domain,
	}
}

// CheckBatch checks multiple emails for disposability
func (s *Service) CheckBatch(emails []string) []CheckEmailResponse {
	results := make([]CheckEmailResponse, 0, len(emails))

	for _, email := range emails {
		result := s.CheckEmail(email)
		results = append(results, result)
	}

	return results
}

// CheckDomain checks if a domain is disposable
func (s *Service) CheckDomain(domain string) DomainCheckResponse {
	isDisposable := disposable.IsDisposableDomain(domain)

	return DomainCheckResponse{
		Domain:       domain,
		IsDisposable: isDisposable,
	}
}

// GetDomains returns paginated list of disposable domains
func (s *Service) GetDomains(page, perPage int) (DomainsListResponse, error) {
	if page < 1 {
		return DomainsListResponse{}, svcerr.Invalid("bad_request", "page must be at least 1")
	}
	if perPage < 1 || perPage > 1000 {
		return DomainsListResponse{}, svcerr.Invalid("bad_request", "per_page must be between 1 and 1000")
	}

	allDomains := disposable.GetAllDomains()
	total := len(allDomains)

	maxPages := (total + perPage - 1) / perPage
	if page > maxPages {
		return DomainsListResponse{}, svcerr.NotFound("page_out_of_range", "page exceeds total number of available pages")
	}

	start := (page - 1) * perPage

	end := start + perPage
	if end > total {
		end = total
	}

	return DomainsListResponse{
		Domains: allDomains[start:end],
		Total:   total,
		Page:    page,
		PerPage: perPage,
		HasMore: end < total,
	}, nil
}

// GetStats returns statistics about disposable domains
func (s *Service) GetStats() StatsResponse {
	return StatsResponse{
		TotalDomains: disposable.Count(),
	}
}

func ExtractDomainFromEmail(email string) string {
	parts := strings.Split(email, "@")

	if len(parts) != 2 {
		return ""
	}
	return parts[1]
}
