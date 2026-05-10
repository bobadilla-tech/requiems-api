package disposable

import (
	"strings"

	disposable "github.com/bobadilla-tech/is-email-disposable"

	"requiems-api/platform/svcerr"
)

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
func (s *Service) CheckBatch(emails []string) BatchCheckResponse {
	results := make([]CheckEmailResponse, 0, len(emails))

	for _, email := range emails {
		result := s.CheckEmail(email)
		results = append(results, result)
	}

	return BatchCheckResponse{
		Results: results,
		Total:   len(results),
	}
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

	start := (page - 1) * perPage
	if start >= total {
		return DomainsListResponse{}, svcerr.NotFound("page_out_of_range", "page exceeds total number of available pages")
	}

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
