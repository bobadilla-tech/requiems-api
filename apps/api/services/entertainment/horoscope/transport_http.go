package horoscope

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// BatchRequest is the body for generating multiple horoscope readings in the same call.
type BatchRequest struct {
	Signs []string `json:"signs" validate:"required,min=1,max=12,dive,required,oneof=aries taurus gemini cancer leo virgo libra scorpio sagittarius capricorn aquarius pisces"`
}

// RegisterRoutes mounts horoscope handlers on the given router.
// Paths are relative to the parent mount point.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/horoscope/{sign}", handleDailyHoroscope(svc))
	r.Post("/horoscope/batch", handleHoroscopeBatch(svc))
}

// handleDailyHoroscope godoc
//
//	@Summary		Get Daily Horoscope
//	@Description	Returns a daily horoscope reading for the specified zodiac sign.
//	@Tags			horoscope
//	@Produce		json
//	@Param			sign	path		string	true	"Zodiac sign (case-insensitive): aries … pisces"
//	@Success		200		{object}	httpx.Response[Horoscope]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		500		{object}	httpx.ErrorResponse
//	@Router			/v1/entertainment/horoscope/{sign} [get]
func handleDailyHoroscope(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sign := strings.ToLower(chi.URLParam(r, "sign"))
		if !IsValidSign(sign) {
			httpx.Error(w, http.StatusBadRequest, "bad_request", "invalid zodiac sign")
			return
		}

		h, err := svc.Daily(sign)
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "failed to generate horoscope")
			return
		}

		httpx.JSON(w, http.StatusOK, h)
	}
}

// handleHoroscopeBatch godoc
//
//	@Summary		Batch Daily Horoscopes
//	@Description	Returns up to 12 daily horoscopes.
//	@Tags			horoscope
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchRequest	true	"List of zodiac signs"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[Horoscope]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/entertainment/horoscope/batch [post]
func handleHoroscopeBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(ctx context.Context, req BatchRequest) (httpx.BatchResponse[Horoscope], error) {
			results, err := svc.DailyBatch(req.Signs)
			if err != nil {
				return httpx.BatchResponse[Horoscope]{}, err
			}
			return httpx.BatchResponse[Horoscope]{Results: results}, nil
		},
	)
}
