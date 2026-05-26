package app

import (
	"context"
	"net/http"
	"time"

	"requiems-api/platform/httpx"
)

type dbPinger interface {
	Ping(ctx context.Context) error
}

type healthzResponse struct {
	Status string `json:"status"`
}

func Healthz(pool dbPinger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		if err := pool.Ping(ctx); err != nil {
			httpx.Error(w, http.StatusServiceUnavailable, "db_unavailable", "database ping failed")
			return
		}

		httpx.JSON(w, http.StatusOK, healthzResponse{Status: "ok"})
	}
}
