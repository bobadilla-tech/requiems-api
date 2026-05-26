package transactionrisk

import (
	"context"
	"math"
	"net"
	"strings"
	"sync"

	"requiems-api/services/finance/bin"
	ipinfo "requiems-api/services/networking/ip/info"
	ipvpn "requiems-api/services/networking/ip/vpn"
)

// -- Dependency interfaces ---------------------------------------------------

type BINLooker interface {
	Lookup(ctx context.Context, raw string) (bin.LookupResponse, error)
}

type VPNChecker interface {
	CheckIP(ctx context.Context, ip net.IP) (ipvpn.IPCheckResponse, error)
}

type IPInfoChecker interface {
	CheckInfo(ctx context.Context, ip string) (ipinfo.LookupResponse, error)
}

// -- Service -----------------------------------------------------------------

// Service scores a payment transaction for fraud risk.
type Service struct {
	bin    BINLooker
	vpn    VPNChecker
	ipInfo IPInfoChecker
}

// NewService returns a new transaction risk Service.
func NewService(b BINLooker, v VPNChecker, i IPInfoChecker) *Service {
	return &Service{bin: b, vpn: v, ipInfo: i}
}

// Request is the input for POST /transaction/risk.
type Request struct {
	CardBIN        string   `json:"card_bin"    validate:"required"`
	IPAddress      string   `json:"ip_address"  validate:"required"`
	BillingCountry string   `json:"billing_country"`
	AmountUSD      *float64 `json:"amount_usd"`
}

// Signals is the resolved signal breakdown returned in the response.
type Signals struct {
	IPCountry       string  `json:"ip_country"`
	BillingCountry  *string `json:"billing_country"` // null when not provided
	BINCountry      string  `json:"bin_country"`
	VPNDetected     bool    `json:"vpn_detected"`
	IsProxy         bool    `json:"is_proxy"`
	IsTOR           bool    `json:"is_tor"`
	CountryMismatch bool    `json:"country_mismatch"`
}

// Result is the transaction risk response.
type Result struct {
	RiskScore float64  `json:"risk_score"`
	IsSafe    bool     `json:"is_safe"`
	Flags     []string `json:"flags"`
	Signals   Signals  `json:"signals"`
}

// Score fans out to BIN lookup, VPN check, and IP geolocation in parallel
// and returns the transaction risk score.
func (s *Service) Score(ctx context.Context, req Request) (Result, error) {
	type binOut struct {
		r   bin.LookupResponse
		err error
	}
	type vpnOut struct {
		r   ipvpn.IPCheckResponse
		err error
	}
	type infoOut struct {
		r   ipinfo.LookupResponse
		err error
	}

	binCh := make(chan binOut, 1)
	vpnCh := make(chan vpnOut, 1)
	infoCh := make(chan infoOut, 1)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		r, err := s.bin.Lookup(ctx, req.CardBIN)
		binCh <- binOut{r, err}
	}()

	parsedIP := net.ParseIP(req.IPAddress)
	go func() {
		defer wg.Done()
		if parsedIP == nil {
			vpnCh <- vpnOut{}
			return
		}
		r, err := s.vpn.CheckIP(ctx, parsedIP)
		vpnCh <- vpnOut{r, err}
	}()

	go func() {
		defer wg.Done()
		r, err := s.ipInfo.CheckInfo(ctx, req.IPAddress)
		infoCh <- infoOut{r, err}
	}()

	wg.Wait()

	binResult := <-binCh
	vpnResult := <-vpnCh
	infoResult := <-infoCh

	score := 0.0
	flags := make([]string, 0, 5)

	ipCountry := ""
	if infoResult.err == nil {
		ipCountry = strings.ToUpper(infoResult.r.CountryCode)
	}

	binCountry := ""
	if binResult.err == nil {
		binCountry = strings.ToUpper(binResult.r.CountryCode)
	}

	billingCC := ""
	if req.BillingCountry != "" {
		billingCC = strings.ToUpper(req.BillingCountry)
	}

	vpnDetected := vpnResult.err == nil && vpnResult.r.IsVPN
	proxyDetected := vpnResult.err == nil && vpnResult.r.IsProxy
	torDetected := vpnResult.err == nil && vpnResult.r.IsTor

	// Country mismatch checks.
	countryMismatch := false
	if billingCC != "" {
		if ipCountry != "" && ipCountry != billingCC {
			score += 0.35
			flags = append(flags, "ip_country_mismatch")
			countryMismatch = true
		}
		if binCountry != "" && binCountry != billingCC {
			score += 0.25
			flags = append(flags, "bin_country_mismatch")
			countryMismatch = true
		}
	} else if ipCountry != "" && binCountry != "" && ipCountry != binCountry {
		// No billing country — compare BIN vs IP.
		countryMismatch = true
	}

	if countryMismatch {
		flags = append(flags, "country_mismatch")
	}

	// VPN/proxy/TOR.
	if torDetected {
		score += 0.40
		flags = append(flags, "tor_detected")
	}
	if proxyDetected {
		score += 0.30
		flags = append(flags, "proxy_detected")
	}
	if vpnDetected {
		score += 0.20
		flags = append(flags, "vpn_detected")
	}

	// Fraud score contribution.
	if vpnResult.err == nil {
		score += float64(vpnResult.r.FraudScore) / 100.0 * 0.25
	}

	// High-value + VPN bonus.
	if vpnDetected && req.AmountUSD != nil && *req.AmountUSD > 500 {
		score += 0.15
		flags = append(flags, "high_value_vpn")
	}

	score = math.Min(score, 1.0)
	score = math.Round(score*100) / 100

	var billingPtr *string
	if req.BillingCountry != "" {
		b := strings.ToUpper(req.BillingCountry)
		billingPtr = &b
	}

	return Result{
		RiskScore: score,
		IsSafe:    score < 0.5,
		Flags:     flags,
		Signals: Signals{
			IPCountry:       ipCountry,
			BillingCountry:  billingPtr,
			BINCountry:      binCountry,
			VPNDetected:     vpnDetected,
			IsProxy:         proxyDetected,
			IsTOR:           torDetected,
			CountryMismatch: countryMismatch,
		},
	}, nil
}
