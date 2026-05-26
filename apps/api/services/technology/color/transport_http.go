package color //nolint:revive // package name matches the service domain it implements

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request holds the validated query parameters for GET /convert/color.
type Request struct {
	From  string `query:"from"  validate:"required,oneof=hex rgb hsl cmyk"`
	To    string `query:"to"    validate:"required,oneof=hex rgb hsl cmyk"`
	Value string `query:"value" validate:"required"`
}

// RegisterRoutes mounts the color conversion handler on the given router.
// Path is relative to the parent mount point.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/color", httpx.HandleGet(func(ctx context.Context, req Request) (Response, error) {
		return svc.Convert(req.From, req.To, req.Value)
	}))
}
