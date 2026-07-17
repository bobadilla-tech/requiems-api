package text

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"requiems-api/platform/config"
	"requiems-api/services/text/detectlanguage"
	"requiems-api/services/text/lorem"
	"requiems-api/services/text/normalize"
	"requiems-api/services/text/sentiment"
	"requiems-api/services/text/similarity"
	"requiems-api/services/text/spellcheck"
	"requiems-api/services/text/thesaurus"
	"requiems-api/services/text/words"
)

func RegisterRoutes(r chi.Router, pool *pgxpool.Pool, cfg config.Config) {
	wordsSvc := words.NewService(pool)
	words.RegisterRoutes(r, wordsSvc)

	loremSvc := lorem.NewService()
	lorem.RegisterRoutes(r, loremSvc)

	spellcheckSvc := spellcheck.NewService(cfg.LanguageToolURL)
	spellcheck.RegisterRoutes(r, spellcheckSvc)

	thesaurusSvc := thesaurus.NewService()
	thesaurus.RegisterRoutes(r, thesaurusSvc)

	detectlanguageSvc := detectlanguage.NewService()
	detectlanguage.RegisterRoutes(r, detectlanguageSvc)

	sentimentSvc, err := sentiment.NewService()
	if err != nil {
		log.Fatalf("failed to init sentiment service: %v", err)
	}

	sentiment.RegisterRoutes(r, sentimentSvc)

	similaritySvc := similarity.NewService()
	similarity.RegisterRoutes(r, similaritySvc)

	normalizeSvc := normalize.NewService()
	normalize.RegisterRoutes(r, normalizeSvc)
}
