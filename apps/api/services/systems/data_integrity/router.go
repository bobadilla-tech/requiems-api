package dataintegrity

import (
	"github.com/go-chi/chi/v5"

	"requiems-api/services/networking/domain"
	"requiems-api/services/networking/whois"
	contentmoderate "requiems-api/services/systems/data_integrity/content_moderate"
	domaintrust "requiems-api/services/systems/data_integrity/domain_trust"
	inputvalidate "requiems-api/services/systems/data_integrity/input_validate"
	textnormalize "requiems-api/services/systems/data_integrity/text_normalize"
	"requiems-api/services/text/detectlanguage"
	"requiems-api/services/text/sentiment"
	"requiems-api/services/validation/email"
	"requiems-api/services/validation/phone"
	"requiems-api/services/validation/profanity"
)

func RegisterRoutes(r chi.Router) {
	textnormalize.RegisterRoutes(r, textnormalize.NewService())

	detectlanguageSvc := detectlanguage.NewService()
	profanitySvc := profanity.NewService()
	sentimentSvc := sentiment.NewService()
	contentmoderateSvc := contentmoderate.NewService(detectlanguageSvc, profanitySvc, sentimentSvc)
	contentmoderate.RegisterRoutes(r, contentmoderateSvc)

	domainSvc := domain.NewService()
	whoIsSvc := whois.NewService()
	domainTrustSvc := domaintrust.NewService(whoIsSvc, domainSvc)
	domaintrust.RegisterRoutes(r, domainTrustSvc)

	emailSvc := email.NewService()
	phoneSvc := phone.NewService()
	inputValidateSvc := inputvalidate.NewService(emailSvc, phoneSvc, profanitySvc, sentimentSvc)
	inputvalidate.RegisterRoutes(r, inputValidateSvc)
}
