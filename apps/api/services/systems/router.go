package systems

import (
	"log"

	"github.com/bobadilla-tech/go-ip-intelligence/v2/ipi"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"

	"requiems-api/platform/config"
	dataintegrity "requiems-api/services/systems/data_integrity"
	globaldata "requiems-api/services/systems/global_data"
	identityrisk "requiems-api/services/systems/identity_risk"
	"requiems-api/services/systems/payments"
)

// RegisterRoutes mounts all system engine routers under r.
// r is expected to be mounted at /systems by the v1 router.
func RegisterRoutes(r chi.Router, pool *pgxpool.Pool, rdb *redis.Client, cfg config.Config) {
	dataintegrity.RegisterRoutes(r)

	ipiClient, err := ipi.New(
		ipi.WithDatabasePath(cfg.VPNDatabasePath),
		ipi.WithASNDatabasePath(cfg.VPNASNDatabasePath),
		ipi.WithCityDatabasePath(cfg.IPCityDatabasePath),
	)
	if err != nil {
		log.Printf("systems: failed to initialize ip intelligence client; ip-based routes disabled: %v", err)
	}

	identityrisk.RegisterRoutes(r, identityrisk.Deps{
		Pool:      pool,
		IPIClient: ipiClient,
		RDB:       rdb,
	})

	payments.RegisterRoutes(r, payments.Deps{
		Pool:      pool,
		IPIClient: ipiClient,
	})

	globaldata.RegisterRoutes(r, globaldata.Deps{
		IPIClient: ipiClient,
		Cfg:       cfg,
		RDB:       rdb,
	})
}
