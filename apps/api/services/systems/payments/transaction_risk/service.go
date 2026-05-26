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

type BINLooker interface {
	Lookup(ctx context.Context, raw string) (bin.LookupResponse, error)
}

type VPNChecker interface {
	CheckIP(ctx context.Context, ip net.IP) (ipvpn.IPCheckResponse, error)
}

type IPInfoChecker interface {
	CheckInfo(ctx context.Context, ip string) (ipinfo.LookupResponse, error)
}

type Service struct {
	bin    BINLooker
	vpn    VPNChecker
	ipInfo IPInfoChecker
}

func NewService(b BINLooker, v VPNChecker, i IPInfoChecker) *Service {
	return &Service{bin: b, vpn: v, ipInfo: i}
}

type Signals struct {
	IPCountry       string  `json:"ip_country"`
	BillingCountry  *string `json:"billing_country"` // null when not provided
	BINCountry      string  `json:"bin_country"`
	VPNDetected     bool    `json:"vpn_detected"`
	IsProxy         bool    `json:"is_proxy"`
	IsTOR           bool    `json:"is_tor"`
	CountryMismatch bool    `json:"country_mismatch"`
}

type Result struct {
	RiskScore float64  `json:"risk_score"`
	IsSafe    bool     `json:"is_safe"`
	Flags     []string `json:"flags"`
	Signals   Signals  `json:"signals"`
}

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

type riskLookups struct {
	bin  binOut
	vpn  vpnOut
	info infoOut
}

func (s *Service) Score(ctx context.Context, req Request) (Result, error) {
	lookups := s.runLookups(ctx, req)
	countries := normalizeCountries(req, lookups)

	vpnDetected, proxyDetected, torDetected := vpnSignals(lookups.vpn)

	flags := make([]string, 0, 5)
	countryScore, countryMismatch := scoreCountrySignals(countries, &flags)
	networkScore := scoreNetworkSignals(lookups.vpn, vpnDetected, proxyDetected, torDetected, &flags)
	score := roundRiskScore(countryScore + networkScore + scoreHighValueVPN(req, vpnDetected, &flags))

	return Result{
		RiskScore: score,
		IsSafe:    score < 0.5 && !torDetected && !proxyDetected,
		Flags:     flags,
		Signals: Signals{
			IPCountry:       countries.ip,
			BillingCountry:  billingCountryPtr(req.BillingCountry),
			BINCountry:      countries.bin,
			VPNDetected:     vpnDetected,
			IsProxy:         proxyDetected,
			IsTOR:           torDetected,
			CountryMismatch: countryMismatch,
		},
	}, nil
}

func (s *Service) runLookups(ctx context.Context, req Request) riskLookups {
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

	return riskLookups{
		bin:  <-binCh,
		vpn:  <-vpnCh,
		info: <-infoCh,
	}
}

type countries struct {
	ip      string
	bin     string
	billing string
}

func normalizeCountries(req Request, lookups riskLookups) countries {
	out := countries{}
	if lookups.info.err == nil {
		out.ip = strings.ToUpper(lookups.info.r.CountryCode)
	}
	if lookups.bin.err == nil {
		out.bin = strings.ToUpper(lookups.bin.r.CountryCode)
	}
	if req.BillingCountry != "" {
		out.billing = strings.ToUpper(req.BillingCountry)
	}
	return out
}

func vpnSignals(vpnResult vpnOut) (bool, bool, bool) {
	return vpnResult.err == nil && vpnResult.r.IsVPN,
		vpnResult.err == nil && vpnResult.r.IsProxy,
		vpnResult.err == nil && vpnResult.r.IsTor
}

func scoreCountrySignals(c countries, flags *[]string) (float64, bool) {
	score := 0.0
	countryMismatch := false
	if c.billing != "" {
		if c.ip != "" && c.ip != c.billing {
			score += 0.35
			*flags = append(*flags, "ip_country_mismatch")
			countryMismatch = true
		}
		if c.bin != "" && c.bin != c.billing {
			score += 0.25
			*flags = append(*flags, "bin_country_mismatch")
			countryMismatch = true
		}
	} else if c.ip != "" && c.bin != "" && c.ip != c.bin {
		countryMismatch = true
	}

	if countryMismatch {
		*flags = append(*flags, "country_mismatch")
	}
	return score, countryMismatch
}

func scoreNetworkSignals(
	vpnResult vpnOut,
	vpnDetected bool,
	proxyDetected bool,
	torDetected bool,
	flags *[]string,
) float64 {
	score := 0.0
	if torDetected {
		score += 0.40
		*flags = append(*flags, "tor_detected")
	}
	if proxyDetected {
		score += 0.30
		*flags = append(*flags, "proxy_detected")
	}
	if vpnDetected {
		score += 0.20
		*flags = append(*flags, "vpn_detected")
	}

	if vpnResult.err == nil {
		score += float64(vpnResult.r.FraudScore) / 100.0 * 0.25
	}
	return score
}

func scoreHighValueVPN(req Request, vpnDetected bool, flags *[]string) float64 {
	if vpnDetected && req.AmountUSD != nil && *req.AmountUSD > 500 {
		*flags = append(*flags, "high_value_vpn")
		return 0.15
	}
	return 0
}

func roundRiskScore(score float64) float64 {
	return math.Round(math.Min(score, 1.0)*100) / 100
}

func billingCountryPtr(country string) *string {
	if country == "" {
		return nil
	}
	normalized := strings.ToUpper(country)
	return &normalized
}
