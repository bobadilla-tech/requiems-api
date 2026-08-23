package info

import (
	"context"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// BatchInfoRequest is the input for the batch IP geolocation endpoint.
type BatchInfoRequest struct {
	IPs []string `json:"ips" validate:"required,min=1,max=50,dive,ip"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	handler := httpx.Guard(svc, func(w http.ResponseWriter, r *http.Request) {
		ipStr := chi.URLParam(r, "ip")
		if ipStr == "" {
			ipStr = httpx.ClientIP(r)
		}

		if net.ParseIP(ipStr) == nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", "invalid IP address")
			return
		}

		result, err := svc.CheckInfo(r.Context(), ipStr)
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "internal error")
			return
		}

		httpx.JSON(w, http.StatusOK, result)
	})

	r.Get("/ip/{ip}", handler)
	r.Get("/ip", handler)

	r.Post("/ip/info/batch", httpx.Guard(svc, httpx.HandleBatch(
		func(ctx context.Context, req BatchInfoRequest) (httpx.BatchResponse[BatchIPInfoItem], error) {
			return httpx.BatchResponse[BatchIPInfoItem]{Results: svc.CheckInfoBatch(ctx, req.IPs)}, nil
		},
	)))
}
