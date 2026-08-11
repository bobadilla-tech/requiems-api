package randomuser

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// BatchGenerateRequest is the body for generating multiple random users at once.
type BatchGenerateRequest struct {
	Count int `json:"count" validate:"required,min=1,max=50"`
}

// BatchGenerateResponse is the response for a batch random user generation request.
type BatchGenerateResponse = httpx.BatchResponse[User]

// RegisterRoutes mounts the random-user handler on the given router.
// Paths are relative to the parent mount point (e.g. /v1/technology).
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/random-user", handleRandomUser(svc))
	r.Post("/random-user/batch", handleRandomUserBatch(svc))
}

// handleRandomUser godoc
//
//	@Summary		Get Random User
//	@Description	Returns a randomly generated fake user profile.
//	@Tags			random-user
//	@Produce		json
//	@Success		200	{object}	httpx.Response[User]
//	@Failure		500	{object}	httpx.ErrorResponse
//	@Router			/v1/technology/random-user [get]
func handleRandomUser(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpx.JSON(w, http.StatusOK, svc.Generate())
	}
}

// handleRandomUserBatch godoc
//
//	@Summary		Batch Generate Users
//	@Description	Returns multiple random users (max 50). Each call consumes `count` units of quota.
//	@Tags			random-user
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchGenerateRequest	true	"Number of users to generate"
//	@Success		200		{object}	httpx.Response[BatchGenerateResponse]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Failure		500		{object}	httpx.ErrorResponse
//	@Router			/v1/technology/random-user/batch [post]
func handleRandomUserBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(_ context.Context, req BatchGenerateRequest) (BatchGenerateResponse, error) {
			users, err := svc.GenerateBatch(req.Count)
			if err != nil {
				return BatchGenerateResponse{}, err
			}
			return BatchGenerateResponse{Results: users}, nil
		},
	)
}
