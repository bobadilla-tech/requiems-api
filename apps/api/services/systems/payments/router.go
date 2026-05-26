package payments

import (
	"github.com/bobadilla-tech/go-ip-intelligence/v2/ipi"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"requiems-api/services/finance/bin"
	"requiems-api/services/finance/iban"
	"requiems-api/services/finance/swift"
	ipinfo "requiems-api/services/networking/ip/info"
	ipvpn "requiems-api/services/networking/ip/vpn"
	"requiems-api/services/systems/payments/payment_validate"
	"requiems-api/services/systems/payments/transaction_risk"
)

// Deps holds the external clients needed by the payments system.
type Deps struct {
	Pool      *pgxpool.Pool
	IPIClient *ipi.Client
}

// RegisterRoutes mounts all Payments Intelligence system endpoints on r.
func RegisterRoutes(r chi.Router, deps Deps) {
	binSvc := bin.NewService(deps.Pool)
	ibanSvc := iban.NewService(deps.Pool)
	swiftSvc := swift.NewService(deps.Pool)

	validateSvc := paymentvalidate.NewService(binSvc, ibanSvc, swiftSvc)
	paymentvalidate.RegisterRoutes(r, validateSvc)

	if deps.IPIClient != nil {
		vpnSvc := ipvpn.NewService(deps.IPIClient)
		infoSvc := ipinfo.NewService(deps.IPIClient)
		riskSvc := transactionrisk.NewService(binSvc, vpnSvc, infoSvc)
		transactionrisk.RegisterRoutes(r, riskSvc)
	}
}
