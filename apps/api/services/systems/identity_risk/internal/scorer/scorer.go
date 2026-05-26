// Package scorer contains the shared risk-signal resolution and risk-score
// computation used by both risk_score and signup_protect. Keeping the weight
// table and fan-out logic here prevents drift between the two endpoints.
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

// -- Dependency interfaces --------------------------------------------------

// EmailChecker validates an email address.
type EmailChecker interface {
	ValidateEmail(ctx context.Context, addr string) email.Validation
}

// PhoneChecker validates a phone number.
type PhoneChecker interface {
	Validate(number string) phone.ValidateResponse
}

// VPNChecker checks an IP for VPN/proxy/TOR signals.
type VPNChecker interface {
	CheckIP(ctx context.Context, ip net.IP) (ipvpn.IPCheckResponse, error)
}

// IPInfoChecker looks up geolocation for an IP address.
type IPInfoChecker interface {
	CheckInfo(ctx context.Context, ip string) (ipinfo.LookupResponse, error)
}

// -- Resolution ---------------------------------------------------------------

// Resolved holds the raw service results plus the derived Signals used for
// scoring. signup_protect uses the raw fields for its signal breakdown;
// risk_score only uses Signals.
type Resolved struct {
	Signals Signals

	EmailResult *email.Validation
	PhoneResult *phone.ValidateResponse
	VPNResult   *ipvpn.IPCheckResponse
	IPResult    *ipinfo.LookupResponse
}

// Resolve fans out to whichever services correspond to the non-empty fields and
// returns a Resolved bundle. An empty string for a field means "not provided".
// Service errors are treated as "no signal" to keep the endpoint non-blocking.
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
		wg.Add(1)
		go func() {
			defer wg.Done()
			r := emailSvc.ValidateEmail(ctx, emailAddr)
			mu.Lock()
			out.EmailResult = &r
			out.Signals.EmailPresent = true
			out.Signals.EmailDisposable = r.Disposable
			out.Signals.EmailInvalid = !r.Valid
			out.Signals.EmailNoMX = !r.MxValid
			mu.Unlock()
		}()
	}

	if phoneNum != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
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
		}()
	}

	parsedIP := net.ParseIP(ipAddr)
	if ipAddr != "" && parsedIP != nil {
		wg.Add(2)

		go func() {
			defer wg.Done()
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
		}()

		go func() {
			defer wg.Done()
			r, err := ipInfoSvc.CheckInfo(ctx, ipAddr)
			if err != nil {
				return
			}
			mu.Lock()
			out.IPResult = &r
			out.Signals.IPCountry = r.CountryCode
			out.Signals.IPPresent = true
			mu.Unlock()
		}()
	}

	wg.Wait()
	return out
}

// -- Scoring ------------------------------------------------------------------

// Signals is the normalised boolean/numeric set fed into Compute.
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
	FraudScore int // 0–100; Compute normalises to 0–1
	IPCountry  string
}

// ScoreResult is the computed output.
type ScoreResult struct {
	RiskScore  float64
	Confidence float64
	IsSafe     bool
	Flags      []string
}

const totalSignals = 3 // email, phone, ip

// Compute calculates the risk score, confidence, and derived is_safe from the
// provided signals. All weights are additive and capped at 1.0.
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
		if s.PhoneInvalid {
			score += 0.25
			flags = append(flags, "phone_invalid")
		} else if s.PhoneVoIP {
			score += 0.15
			flags = append(flags, "phone_voip")
		} else if s.PhoneVirtual {
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

	isSafe := score < 0.5 && confidence > 0.6

	return ScoreResult{
		RiskScore:  score,
		Confidence: confidence,
		IsSafe:     isSafe,
		Flags:      flags,
	}
}
