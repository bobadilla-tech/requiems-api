package dataintegrity

import (
	"github.com/go-chi/chi/v5"

	textnormalize "requiems-api/services/systems/data_integrity/text_normalize"
)

func RegisterRoutes(r chi.Router) {
	textnormalize.RegisterRoutes(r, textnormalize.NewService())
}
