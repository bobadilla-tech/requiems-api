package dataintegrity

import (
	"github.com/go-chi/chi/v5"

	contentmoderate "requiems-api/services/systems/data_integrity/content_moderate"
	textnormalize "requiems-api/services/systems/data_integrity/text_normalize"
	"requiems-api/services/text/detectlanguage"
	"requiems-api/services/text/sentiment"
	"requiems-api/services/validation/profanity"
)

func RegisterRoutes(r chi.Router) {
	textnormalize.RegisterRoutes(r, textnormalize.NewService())

	detectlanguageSvc := detectlanguage.NewService()
	profanitySvc := profanity.NewService()
	sentimentSvc := sentiment.NewService()
	contentmoderateSvc := contentmoderate.NewService(detectlanguageSvc, profanitySvc, sentimentSvc)
	contentmoderate.RegisterRoutes(r, contentmoderateSvc)
}
