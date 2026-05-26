package identityrisk

import (
	"github.com/bobadilla-tech/go-ip-intelligence/v2/ipi"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"requiems-api/services/networking/domain"
	ipinfo "requiems-api/services/networking/ip/info"
	ipvpn "requiems-api/services/networking/ip/vpn"
	"requiems-api/services/networking/mx"
	"requiems-api/services/networking/whois"
	"requiems-api/services/systems/identity_risk/internal/scorer"
	riskscore "requiems-api/services/systems/identity_risk/risk_score"
	signupprotect "requiems-api/services/systems/identity_risk/signup_protect"
	userverify "requiems-api/services/systems/identity_risk/user_verify"
	"requiems-api/services/validation/email"
	"requiems-api/services/validation/phone"
)

// Deps holds the external service clients needed by the identity risk system.
type Deps struct {
	Pool      *pgxpool.Pool
	IPIClient *ipi.Client
}

// RegisterRoutes mounts all Identity & Risk system endpoints on r.
func RegisterRoutes(r chi.Router, deps Deps) {
	emailSvc := email.NewService()
	phoneSvc := phone.NewService()

	// Use interface-typed vars so nil is a true nil interface (not a non-nil
	// interface wrapping a nil concrete pointer), preventing panics in the
	// scorer when an IP address is supplied but the IPI client is unavailable.
	var vpnSvc scorer.VPNChecker
	var infoSvc scorer.IPInfoChecker
	if deps.IPIClient != nil {
		vpnSvc = ipvpn.NewService(deps.IPIClient)
		infoSvc = ipinfo.NewService(deps.IPIClient)
	}

	riskSvc := riskscore.NewService(emailSvc, phoneSvc, vpnSvc, infoSvc)
	riskscore.RegisterRoutes(r, riskSvc)

	signupSvc := signupprotect.NewService(emailSvc, phoneSvc, vpnSvc, infoSvc)
	signupprotect.RegisterRoutes(r, signupSvc)

	whoisSvc := whois.NewService()
	domainSvc := domain.NewService()
	mxSvc := mx.NewService()
	verifySvc := userverify.NewService(emailSvc, whoisSvc, domainSvc, mxSvc, vpnSvc)
	userverify.RegisterRoutes(r, verifySvc)
}
