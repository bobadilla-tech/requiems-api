package domaintrust

import (
	"context"
	"math"
	"time"

	"requiems-api/services/networking/domain"
	"requiems-api/services/networking/whois"
)

// WhoIs holds registration metadata for a domain retrieved from WHOIS lookup.
// This field may be omitted from the response if WHOIS is unavailable for TLD.
type WhoIs struct {
	Registrar string   `json:"registrar"`
	CreatedAt string   `json:"created_at"`
	ExpiresAt string   `json:"expires_at"`
	AgeDays   int      `json:"age_days"`
	Status    []string `json:"status"`
}

// Dns holds the result of the DNS health check for a domain.
// Used to determine whether the domain is operational and resolves correctly.
type Dns struct {
	HasARecords  bool `json:"has_a_records"`
	HasMxRecords bool `json:"has_mx_records"`
	HasNsRecords bool `json:"has_ns_records"`
	Available    bool `json:"available"`
}

// MxRecords represents a single mail exchange record for the domain.
// Multiple MX records may be present, each with a different priority.
type MxRecords struct {
	HostName string `json:"host_name"`
	Priority uint16 `json:"priority"`
}

// Response is the top-level structure returned by GET /v1/domain/trust/{domain}.
// It composes results from the WHOIS, DNS, and MX services into a single trust assessment.
type Response struct {
	Domain     string      `json:"domain"`
	TrustScore float64     `json:"trust_score"`
	TrustLevel string      `json:"trust_level"`
	WhoIs      *WhoIs      `json:"who_is,omitempty"`
	Dns        Dns         `json:"dns"`
	MxRecords  []MxRecords `json:"mx_records"`
	Flags      []string    `json:"flags"`
}

// WhoIsService defines the contract for retrieving WHOIS registration data for a domain.
type WhoIsService interface {
	Lookup(_ context.Context, domain string) (whois.LookupResponse, error)
}

// DomainService defines the contract for retrieving DNS information for a domain.
type DomainService interface {
	GetInfo(ctx context.Context, domainName string) domain.InfoResponse
}

// Service evaluates domains.
type Service struct {
	whoIs  WhoIsService
	domain DomainService
}

// NewService returns a new Service.
func NewService(whois WhoIsService, domain DomainService) *Service {
	return &Service{
		whoIs:  whois,
		domain: domain,
	}
}

// Evaluate analyzes a domain and returns a trust assessment based on DNS, WHOIS,
// and MX record data. The trust score starts at 1.0 and is reduced by penalties
// depending on the issues found during evaluation.
func (s *Service) Evaluate(ctx context.Context, domain string) Response {

	// initialize response with default values.
	var result Response
	result.Domain = domain
	result.TrustScore = 1.0
	result.Flags = []string{}
	result.MxRecords = []MxRecords{}

	// fetch DNS and availability info for the domain.
	domainResult := s.domain.GetInfo(ctx, domain)

	result.Dns.Available = domainResult.Available

	// if the domain is not registered, return immediately with a zero trust score.
	if domainResult.Available {
		result.TrustScore = 0.0
		result.TrustLevel = "low"
		result.Flags = append(result.Flags, "domain_not_registered")
		return result
	}

	// penalize if no A records are found — domain may not resolve to any server.
	if len(domainResult.DNS.A) == 0 {
		result.TrustScore -= 0.2
		result.Flags = append(result.Flags, "no_a_records")
	} else {
		result.Dns.HasARecords = true
	}

	// penalize if no MX records are found — domain cannot receive emails.
	if len(domainResult.DNS.MX) == 0 {
		result.TrustScore -= 0.35
		result.Flags = append(result.Flags, "no_mx")
	} else {
		result.Dns.HasMxRecords = true

		// map each MX record to the response structure.
		for _, mx := range domainResult.DNS.MX {
			result.MxRecords = append(result.MxRecords, MxRecords{
				HostName: mx.Host,
				Priority: mx.Priority,
			})
		}
	}

	// flag if no NS records are found — domain may lack proper nameserver delegation.
	if len(domainResult.DNS.NS) == 0 {
		result.Flags = append(result.Flags, "no_ns_records")
	} else {
		result.Dns.HasNsRecords = true
	}

	// fetch WHOIS registration data for the domain.
	whoisResult, whoisErr := s.whoIs.Lookup(ctx, domain)

	// if WHOIS lookup fails, flag it and skip all WHOIS-based evaluations.
	if whoisErr != nil {
		result.Flags = append(result.Flags, "whois_unavailable")
	} else {
		// populate WHOIS metadata from the lookup result.
		result.WhoIs = &WhoIs{}
		result.WhoIs.Registrar = whoisResult.Registrar
		result.WhoIs.CreatedAt = whoisResult.CreatedDate
		result.WhoIs.ExpiresAt = whoisResult.ExpiryDate

		expiresAt, expiresErr := time.Parse(time.RFC3339, whoisResult.ExpiryDate)

		// penalize if the domain is expiring within 14 days.
		if expiresErr == nil {
			if daysUntilExpiry(expiresAt) < 14 {
				result.TrustScore -= 0.15
				result.Flags = append(result.Flags, "expiring_soon")
			}
		}

		// calculate domain age and apply penalties based on how long it has been registered.
		// if an error occurs, neither age calculation nor penalties are applied.
		createdAt, err := time.Parse(time.RFC3339, whoisResult.CreatedDate)
		if err == nil {
			result.WhoIs.AgeDays = ageDays(createdAt)

			// heavily penalize very new domains (under 30 days) — high risk of being fraudulent.
			if result.WhoIs.AgeDays < 30 {
				result.TrustScore -= 0.5
				result.Flags = append(result.Flags, "new_domain")
			}

			// moderately penalize young domains (30–180 days) — still building trust.
			if result.WhoIs.AgeDays >= 30 && result.WhoIs.AgeDays <= 180 {
				result.TrustScore -= 0.25
				result.Flags = append(result.Flags, "young_domain")
			}
		}

		result.WhoIs.Status = whoisResult.Status
	}

	// ensure trust score never goes below zero.
	result.TrustScore = math.Max(result.TrustScore, 0.0)

	// assign a human-readable trust level based on the final score.
	switch {
	case result.TrustScore >= 0.75:
		result.TrustLevel = "high"
	case result.TrustScore >= 0.4:
		result.TrustLevel = "medium"
	default:
		result.TrustLevel = "low"

	}

	return result
}

// ageDays returns the number of days elapsed since the domain was created.
func ageDays(createdAt time.Time) int {
	return int(time.Since(createdAt).Hours() / 24)
}

// daysUntilExpiry returns the number of days remaining until the domain expires.
func daysUntilExpiry(expiresAt time.Time) int {
	return int(time.Until(expiresAt).Hours() / 24)
}
