package scorer

import (
	"context"
	"math"
	"net"
	"sync"

	ipinfo "requiems-api/services/networking/ip/info"
	ipvpn "requiems-api/services/networking/ip/vpn"
	"requiems-api/services/validation/email"
	"requiems-api/services/validation/phone"
)

type EmailChecker interface {
	ValidateEmail(ctx context.Context, addr string) email.Validation
}

type PhoneChecker interface {
	Validate(number string) phone.ValidateResponse
}

type VPNChecker interface {
	CheckIP(ctx context.Context, ip net.IP) (ipvpn.IPCheckResponse, error)
}

type IPInfoChecker interface {
	CheckInfo(ctx context.Context, ip string) (ipinfo.LookupResponse, error)
}

type Resolved struct {
	Signals Signals

	EmailResult *email.Validation
	PhoneResult *phone.ValidateResponse
	VPNResult   *ipvpn.IPCheckResponse
	IPResult    *ipinfo.LookupResponse
}

func Resolve(
	ctx context.Context,
	emailSvc EmailChecker,
	phoneSvc PhoneChecker,
	vpnSvc VPNChecker,
	ipInfoSvc IPInfoChecker,
	emailAddr, phoneNum, ipAddr string,
) Resolved {
	var (
		mu  sync.Mutex
		wg  sync.WaitGroup
		out Resolved
	)

	if emailAddr != "" {
		wg.Go(func() {
			r := emailSvc.ValidateEmail(ctx, emailAddr)
			mu.Lock()
			out.EmailResult = &r
			out.Signals.EmailPresent = true
			out.Signals.EmailDisposable = r.Disposable
			out.Signals.EmailInvalid = !r.Valid
			out.Signals.EmailNoMX = !r.MxValid
			mu.Unlock()
		})
	}

	if phoneNum != "" {
		wg.Go(func() {
			r := phoneSvc.Validate(phoneNum)
			mu.Lock()
			out.PhoneResult = &r
			out.Signals.PhonePresent = true
			out.Signals.PhoneInvalid = !r.Valid
			if r.Risk != nil {
				out.Signals.PhoneVoIP = r.Risk.IsVoIP
				out.Signals.PhoneVirtual = r.Risk.IsVirtual
			}
			out.Signals.PhoneCountry = r.Country
			mu.Unlock()
		})
	}

	parsedIP := net.ParseIP(ipAddr)
	if ipAddr != "" && parsedIP != nil {
		if vpnSvc != nil {
			wg.Go(func() {
				r, err := vpnSvc.CheckIP(ctx, parsedIP)
				if err != nil {
					return
				}
				mu.Lock()
				out.VPNResult = &r
				out.Signals.IPPresent = true
				out.Signals.IsTOR = r.IsTor
				out.Signals.IsProxy = r.IsProxy
				out.Signals.IsVPN = r.IsVPN
				out.Signals.IsHosting = r.IsHosting
				out.Signals.FraudScore = r.FraudScore
				mu.Unlock()
			})
		}
		if ipInfoSvc != nil {
			wg.Go(func() {
				r, err := ipInfoSvc.CheckInfo(ctx, ipAddr)
				if err != nil {
					return
				}
				mu.Lock()
				out.IPResult = &r
				out.Signals.IPCountry = r.CountryCode
				out.Signals.IPPresent = true
				mu.Unlock()
			})
		}
	}

	wg.Wait()
	return out
}

type Signals struct {
	EmailPresent    bool
	EmailDisposable bool
	EmailInvalid    bool
	EmailNoMX       bool

	PhonePresent bool
	PhoneInvalid bool
	PhoneVoIP    bool
	PhoneVirtual bool
	PhoneCountry string

	IPPresent  bool
	IsTOR      bool
	IsProxy    bool
	IsVPN      bool
	IsHosting  bool
	FraudScore int
	IPCountry  string
}

type ScoreResult struct {
	RiskScore  float64
	Confidence float64
	IsSafe     bool
	Flags      []string
}

const totalSignals = 3

func Compute(s Signals) ScoreResult {
	score := 0.0
	flags := make([]string, 0, 4)

	if s.EmailPresent {
		if s.EmailInvalid {
			score += 0.40
			flags = append(flags, "email_invalid")
		}
		if s.EmailNoMX {
			score += 0.25
			flags = append(flags, "email_no_mx")
		}
		if s.EmailDisposable {
			score += 0.30
			flags = append(flags, "disposable_email")
		}
	}

	if s.PhonePresent {
		switch {
		case s.PhoneInvalid:
			score += 0.25
			flags = append(flags, "phone_invalid")
		case s.PhoneVoIP:
			score += 0.15
			flags = append(flags, "phone_voip")
		case s.PhoneVirtual:
			score += 0.15
			flags = append(flags, "phone_virtual")
		}
		if s.IPPresent && s.PhoneCountry != "" && s.IPCountry != "" && s.PhoneCountry != s.IPCountry {
			score += 0.10
			flags = append(flags, "geo_mismatch_phone_ip")
		}
	}

	if s.IPPresent {
		if s.IsTOR {
			score += 0.40
			flags = append(flags, "tor_detected")
		}
		if s.IsProxy {
			score += 0.25
			flags = append(flags, "proxy_detected")
		}
		if s.IsVPN {
			score += 0.20
			flags = append(flags, "vpn_detected")
		}
		if s.IsHosting {
			score += 0.10
			flags = append(flags, "hosting_ip")
		}
		score += float64(s.FraudScore) / 100.0 * 0.30
	}

	score = math.Min(score, 1.0)
	score = math.Round(score*100) / 100

	present := 0
	if s.EmailPresent {
		present++
	}
	if s.PhonePresent {
		present++
	}
	if s.IPPresent {
		present++
	}
	confidence := math.Round(float64(present)/float64(totalSignals)*100) / 100

	isSafe := score < 0.5 && confidence > 0.6 && !s.IsTOR && !s.IsProxy

	return ScoreResult{
		RiskScore:  score,
		Confidence: confidence,
		IsSafe:     isSafe,
		Flags:      flags,
	}
}
