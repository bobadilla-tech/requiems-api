package disposable

// CheckEmailRequest is the body for a single-email disposable check.
type CheckEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

// CheckEmailResponse represents the response for a single email check
type CheckEmailResponse struct {
	Email        string `json:"email"`
	IsDisposable bool   `json:"is_disposable"`
	Domain       string `json:"domain,omitempty"`
}

// BatchCheckRequest is the body for checking multiple emails at once.
type BatchCheckRequest struct {
	Emails []string `json:"emails" validate:"required,min=1,max=100,dive,email"`
}

// DomainCheckResponse represents the response for a domain check
type DomainCheckResponse struct {
	Domain       string `json:"domain"`
	IsDisposable bool   `json:"is_disposable"`
}

// DomainsListResponse represents the response for listing all domains
type DomainsListResponse struct {
	Domains []string `json:"domains"`
	Total   int      `json:"total"`
	Page    int      `json:"page"`
	PerPage int      `json:"per_page"`
	HasMore bool     `json:"has_more"`
}

type DomainsListQuery struct {
	Page    int `query:"page"     validate:"min=1"`
	PerPage int `query:"per_page" validate:"min=1"`
}

// StatsResponse represents statistics about disposable domains
type StatsResponse struct {
	TotalDomains int `json:"total_domains"`
}

func (CheckEmailResponse) IsData()  {}
func (DomainCheckResponse) IsData() {}
func (DomainsListResponse) IsData() {}
func (StatsResponse) IsData()       {}
