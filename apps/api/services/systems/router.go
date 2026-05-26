package systems

import (
	"github.com/go-chi/chi/v5"

	dataintegrity "requiems-api/services/systems/data_integrity"
)

// RegisterRoutes mounts all system engine routers under r.
// r is expected to be mounted at /systems by the v1 router.
func RegisterRoutes(r chi.Router) {
	dataintegrity.RegisterRoutes(r)
}
