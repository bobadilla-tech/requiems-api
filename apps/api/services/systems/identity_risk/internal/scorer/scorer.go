package scorer

import (
	"context"
	"fmt"
	"math"
	"net"
	"sync"
	"time"

	ipinfo "requiems-api/services/networking/ip/info"
	ipvpn "requiems-api/services/networking/ip/vpn"
	"requiems-api/services/validation/email"
	"requiems-api/services/validation/phone"

	"github.com/redis/go-redis/v9"
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

const (
	emailMXWeight         = 0.95 // emailMXWeight reflects ~95% reliability of MX record validation.
	emailDisposableWeight = 0.85 // emailDisposableWeight is slightly lower as some disposable providers use legitimate domains.
	voipWeight            = 0.70 // voipWeight is lower due to inconsistent detection across carriers and regions.
	vitualWeight          = 0.70 // virtualWeight mirrors voipWeight — virtual numbers share similar detection uncertainty.
	ipFraudWeight         = 0.80 // ipFraudWeight reflects external provider accuracy, which varies by data freshness.
	proxyWeight           = 0.90 // proxyWeight is high — proxy detection is well-established and reliable.
	torWeight             = 0.90 // torWeight is high — TOR exit node lists are maintained and highly accurate.
	velocityWeight        = 0.75 // velocityWeight reflects strong fraud correlation for high-velocity IPs, tempered by thresholds that are heuristic and pending tuning against real traffic.
)

func Resolve(
	ctx context.Context,
	emailSvc EmailChecker,
	phoneSvc PhoneChecker,
	vpnSvc VPNChecker,
	ipInfoSvc IPInfoChecker,
	emailAddr, phoneNum, ipAddr string,
	rdb *redis.Client,
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
				out.Signals.IPRiskChecked = true
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

		wg.Go(func() {
			key1h := fmt.Sprintf("identity_risk:ip:%s:1h", ipAddr)
			key24h := fmt.Sprintf("identity_risk:ip:%s:24h", ipAddr)

			count1h, err1 := incrVelocityKey(ctx, rdb, key1h, time.Hour)
			count24h, err2 := incrVelocityKey(ctx, rdb, key24h, 24*time.Hour)
			if err1 != nil {
				return
			}

			mu.Lock()
			out.Signals.VelocityChecked = true
			out.Signals.VelocityCount1h = count1h
			if err2 == nil {
				out.Signals.VelocityCount24h = count24h
			}
			mu.Unlock()

		})
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

	IPPresent     bool
	IPRiskChecked bool
	IsTOR         bool
	IsProxy       bool
	IsVPN         bool
	IsHosting     bool
	FraudScore    int
	IPCountry     string

	VelocityChecked  bool
	VelocityCount1h  int64
	VelocityCount24h int64
}

type ScoreResult struct {
	RiskScore  float64
	Confidence float64
	IsSafe     bool
	Flags      []string
}

func Compute(s Signals) ScoreResult {
	flags := make([]string, 0, 4)
	score := scoreEmail(s, &flags) + scorePhone(s, &flags) + scoreIP(s, &flags) + scoreVelocity(s, &flags)
	score = roundScore(score)

	confidence := signalConfidence(s)
	isSafe := score < 0.5 && confidence > 0.6 && !s.IsTOR && !s.IsProxy

	return ScoreResult{
		RiskScore:  score,
		Confidence: confidence,
		IsSafe:     isSafe,
		Flags:      flags,
	}
}

func scoreEmail(s Signals, flags *[]string) float64 {
	if s.EmailPresent {
		score := 0.0
		if s.EmailInvalid {
			score += 0.40
			*flags = append(*flags, "email_invalid")
		}
		if s.EmailNoMX {
			score += 0.25
			*flags = append(*flags, "email_no_mx")
		}
		if s.EmailDisposable {
			score += 0.30
			*flags = append(*flags, "disposable_email")
		}
		return score
	}
	return 0
}

func scorePhone(s Signals, flags *[]string) float64 {
	if s.PhonePresent {
		score := 0.0
		switch {
		case s.PhoneInvalid:
			score += 0.25
			*flags = append(*flags, "phone_invalid")
		case s.PhoneVoIP:
			score += 0.15
			*flags = append(*flags, "phone_voip")
		case s.PhoneVirtual:
			score += 0.15
			*flags = append(*flags, "phone_virtual")
		}
		if s.IPPresent && s.PhoneCountry != "" && s.IPCountry != "" && s.PhoneCountry != s.IPCountry {
			score += 0.10
			*flags = append(*flags, "geo_mismatch_phone_ip")
		}
		return score
	}
	return 0
}

func scoreIP(s Signals, flags *[]string) float64 {
	if s.IPPresent {
		score := 0.0
		if s.IsTOR {
			score += 0.40
			*flags = append(*flags, "tor_detected")
		}
		if s.IsProxy {
			score += 0.25
			*flags = append(*flags, "proxy_detected")
		}
		if s.IsVPN {
			score += 0.20
			*flags = append(*flags, "vpn_detected")
		}
		if s.IsHosting {
			score += 0.10
			*flags = append(*flags, "hosting_ip")
		}
		score += float64(s.FraudScore) / 100.0 * 0.30
		return score
	}
	return 0
}

func roundScore(score float64) float64 {
	return math.Round(math.Min(score, 1.0)*100) / 100
}

func signalConfidence(s Signals) float64 {
	numerator := 0.0
	denominator := 0.0

	if s.EmailPresent {
		denominator += emailMXWeight + emailDisposableWeight

		if !s.EmailNoMX {
			numerator += emailMXWeight
		}

		if !s.EmailDisposable {
			numerator += emailDisposableWeight
		}
	}
	if s.PhonePresent {

		denominator += voipWeight + vitualWeight

		if !s.PhoneVoIP {
			numerator += voipWeight
		}

		if !s.PhoneVirtual {
			numerator += vitualWeight
		}
	}
	if s.IPRiskChecked {
		denominator += ipFraudWeight + torWeight + proxyWeight

		if !s.IsProxy {
			numerator += proxyWeight
		}

		if !s.IsTOR {
			numerator += torWeight
		}

		if s.FraudScore == 0 {
			numerator += ipFraudWeight
		}

	}

	if s.VelocityChecked {
		denominator += velocityWeight
		if s.VelocityCount1h <= 10 {
			numerator += velocityWeight
		}
	}

	if denominator == 0 {
		return 0
	}

	return math.Round(numerator/denominator*100) / 100
}

func incrVelocityKey(ctx context.Context, rdb *redis.Client, key string, ttl time.Duration) (int64, error) {
	ok, err := rdb.SetNX(ctx, key, 1, ttl).Result()

	if err != nil {
		return 0, err
	}

	if ok {
		return 1, nil
	}

	count, err := rdb.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	return count, nil
}

func scoreVelocity(s Signals, flags *[]string) float64 {
	if s.VelocityChecked {
		switch {
		case s.VelocityCount1h > 50:
			*flags = append(*flags, "velocity_high")
			return 0.25
		case s.VelocityCount1h > 10:
			*flags = append(*flags, "velocity_medium")
			return 0.10
		default:
			return 0
		}
	}
	return 0
}
