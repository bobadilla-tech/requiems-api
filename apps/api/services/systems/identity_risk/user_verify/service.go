package userverify

import (
	"context"
	"math"
	"net"
	"strings"
	"sync"
	"time"

	"requiems-api/services/networking/domain"
	ipvpn "requiems-api/services/networking/ip/vpn"
	"requiems-api/services/networking/mx"
	"requiems-api/services/networking/whois"
	"requiems-api/services/validation/email"
)

// -- Dependency interfaces ---------------------------------------------------

type EmailChecker interface {
	ValidateEmail(ctx context.Context, addr string) email.Validation
}

type WHOISLooker interface {
	Lookup(ctx context.Context, d string) (whois.LookupResponse, error)
}

type DomainInfoGetter interface {
	GetInfo(ctx context.Context, d string) domain.InfoResponse
}

type MXLooker interface {
	Lookup(ctx context.Context, d string) (mx.LookupResponse, error)
}

type VPNChecker interface {
	CheckIP(ctx context.Context, ip net.IP) (ipvpn.IPCheckResponse, error)
}

// -- Service -----------------------------------------------------------------

// Service verifies a user's identity signals.
type Service struct {
	email  EmailChecker
	whois  WHOISLooker
	domain DomainInfoGetter
	mx     MXLooker
	vpn    VPNChecker // optional; only used when ip_address provided
}

// NewService returns a new user verify Service.
func NewService(e EmailChecker, w WHOISLooker, d DomainInfoGetter, m MXLooker, v VPNChecker) *Service {
	return &Service{email: e, whois: w, domain: d, mx: m, vpn: v}
}

// Request is the input for POST /user/verify.
type Request struct {
	Email     string `json:"email" validate:"required"`
	IPAddress string `json:"ip_address"`
}

// EmailSignal is the email signal breakdown.
type EmailSignal struct {
	Valid      bool `json:"valid"`
	Disposable bool `json:"disposable"`
	MXValid    bool `json:"mx_valid"`
}

// DomainSignal is the domain signal breakdown.
type DomainSignal struct {
	AgeDays     int  `json:"age_days"` // -1 if unknown
	HasMX       bool `json:"has_mx"`
	HasARecords bool `json:"has_a_records"`
	Available   bool `json:"available"`
}

// IPSignal is the IP signal breakdown.
type IPSignal struct {
	IsVPN      bool    `json:"is_vpn"`
	IsProxy    bool    `json:"is_proxy"`
	FraudScore float64 `json:"fraud_score"` // normalised 0–1
}

// Signals groups the per-signal breakdowns.
type Signals struct {
	Email  EmailSignal  `json:"email"`
	Domain DomainSignal `json:"domain"`
	IP     *IPSignal    `json:"ip"`
}

// Result is the full user verification response.
type Result struct {
	Verified   bool     `json:"verified"`
	Confidence float64  `json:"confidence"`
	RiskScore  float64  `json:"risk_score"`
	Flags      []string `json:"flags"`
	Signals    Signals  `json:"signals"`
}

// Verify fans out to the relevant services and returns the verification result.
func (s *Service) Verify(ctx context.Context, req Request) (Result, error) {
	emailResult := s.email.ValidateEmail(ctx, req.Email)

	// Extract domain from validated email.
	emailDomain := ""
	if emailResult.Domain != nil {
		emailDomain = *emailResult.Domain
	} else if idx := strings.LastIndex(req.Email, "@"); idx >= 0 {
		emailDomain = req.Email[idx+1:]
	}

	type whoisOut struct {
		r   whois.LookupResponse
		err error
	}
	type domainOut struct{ r domain.InfoResponse }
	type mxOut struct {
		r   mx.LookupResponse
		err error
	}
	type vpnOut struct {
		r   ipvpn.IPCheckResponse
		err error
	}

	whoisCh := make(chan whoisOut, 1)
	domainCh := make(chan domainOut, 1)
	mxCh := make(chan mxOut, 1)
	vpnCh := make(chan vpnOut, 1)

	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer wg.Done()
		r, err := s.whois.Lookup(ctx, emailDomain)
		whoisCh <- whoisOut{r, err}
	}()
	go func() {
		defer wg.Done()
		r := s.domain.GetInfo(ctx, emailDomain)
		domainCh <- domainOut{r}
	}()
	go func() {
		defer wg.Done()
		r, err := s.mx.Lookup(ctx, emailDomain)
		mxCh <- mxOut{r, err}
	}()

	parsedIP := net.ParseIP(req.IPAddress)
	if req.IPAddress != "" && parsedIP != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			r, err := s.vpn.CheckIP(ctx, parsedIP)
			vpnCh <- vpnOut{r, err}
		}()
	} else {
		vpnCh <- vpnOut{}
	}

	wg.Wait()

	whoisResult := <-whoisCh
	domainResult := <-domainCh
	mxResult := <-mxCh
	vpnResult := <-vpnCh

	// Build signals and compute risk score.
	flags := make([]string, 0, 4)
	score := 0.0
	servicesResolved := 2 // email + domain always count
	servicesTotal := 2
	if req.IPAddress != "" && parsedIP != nil {
		servicesTotal++
	}

	// Email signals.
	emailSig := EmailSignal{
		Valid:      emailResult.Valid,
		Disposable: emailResult.Disposable,
		MXValid:    emailResult.MxValid,
	}
	if !emailResult.Valid {
		score += 0.50
		flags = append(flags, "email_invalid")
	}
	if emailResult.Disposable {
		score += 0.30
		flags = append(flags, "disposable_email")
	}

	// Domain signals from WHOIS + domain info + MX.
	ageDays := -1
	if whoisResult.err == nil && whoisResult.r.CreatedDate != "" {
		if age, ok := parseDomainAge(whoisResult.r.CreatedDate); ok {
			ageDays = age
			if ageDays < 30 {
				score += 0.25
				flags = append(flags, "young_domain")
			} else if ageDays < 180 {
				score += 0.10
				flags = append(flags, "young_domain")
			}
		}
	} else if whoisResult.err != nil {
		flags = append(flags, "whois_unavailable")
		servicesResolved-- // WHOIS didn't resolve
		servicesResolved++ // subtract from total later — don't penalise confidence
	}

	hasMX := mxResult.err == nil && len(mxResult.r.Records) > 0
	hasA := len(domainResult.r.DNS.A) > 0
	available := domainResult.r.Available

	domainSig := DomainSignal{
		AgeDays:     ageDays,
		HasMX:       hasMX,
		HasARecords: hasA,
		Available:   available,
	}

	if !hasMX {
		score += 0.30
		flags = append(flags, "no_mx")
	}
	if available {
		score += 0.50
		flags = append(flags, "domain_not_registered")
	}

	// IP signals.
	var ipSig *IPSignal
	if req.IPAddress != "" && parsedIP != nil && vpnResult.err == nil {
		ipSig = &IPSignal{
			IsVPN:      vpnResult.r.IsVPN,
			IsProxy:    vpnResult.r.IsProxy,
			FraudScore: float64(vpnResult.r.FraudScore) / 100.0,
		}
		if vpnResult.r.IsVPN || vpnResult.r.IsProxy || vpnResult.r.IsTor {
			score += 0.20
			flags = append(flags, "ip_risk")
		}
		servicesResolved++
	}

	score = math.Min(score, 1.0)
	score = math.Round(score*100) / 100
	confidence := math.Round(float64(servicesResolved)/float64(servicesTotal)*100) / 100

	verified := score < 0.3 && confidence > 0.5 && emailResult.Valid && hasMX && !available

	return Result{
		Verified:   verified,
		Confidence: confidence,
		RiskScore:  score,
		Flags:      flags,
		Signals: Signals{
			Email:  emailSig,
			Domain: domainSig,
			IP:     ipSig,
		},
	}, nil
}

// parseDomainAge parses common WHOIS date formats and returns age in days.
func parseDomainAge(dateStr string) (int, bool) {
	formats := []string{
		time.RFC3339,
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05-07:00",
		"2006-01-02",
		"02-Jan-2006",
		"January 2, 2006",
		"2006.01.02",
	}
	for _, f := range formats {
		if t, err := time.Parse(f, dateStr); err == nil {
			days := int(time.Since(t).Hours() / 24)
			if days < 0 {
				days = 0
			}
			return days, true
		}
	}
	return -1, false
}
