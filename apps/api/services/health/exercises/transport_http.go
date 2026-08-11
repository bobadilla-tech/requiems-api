package exercises

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
	"requiems-api/platform/svcerr"
)

// ListParams holds the query parameters accepted by the list and random endpoints.
type ListParams struct {
	BodyPart  string `query:"body_part"`
	Equipment string `query:"equipment"`
	Muscle    string `query:"muscle"`
	Search    string `query:"search"`
	Page      int    `query:"page"     validate:"min=1"`
	PerPage   int    `query:"per_page" validate:"min=1,max=100"`
}

// BatchGetRequest is the body for fetching multiple exercises by ID.
type BatchGetRequest struct {
	IDs []int `json:"ids" validate:"required,min=1,max=50,dive,min=1"`
}

// BatchExerciseResponse is the response for a batch exercise lookup.
type BatchExerciseResponse = httpx.BatchResponse[Exercise]

// exerciseQuerier is the interface consumed by HTTP handlers, allowing stub
// injection in tests without a live database.
type exerciseQuerier interface {
	List(ctx context.Context, p ListParams) (ExerciseList, error)
	Get(ctx context.Context, id int) (Exercise, error)
	Random(ctx context.Context, p ListParams) (Exercise, error)
	GetBatch(ctx context.Context, ids []int) ([]Exercise, error)
	BodyParts(ctx context.Context) (StringList, error)
	Equipment(ctx context.Context) (StringList, error)
	Muscles(ctx context.Context) (StringList, error)
}

// RegisterRoutes mounts all exercise endpoints onto r.
func RegisterRoutes(r chi.Router, svc *Service) {
	registerExerciseRoutes(r, svc)
}

// registerExerciseRoutes wires the exerciseQuerier interface to the router.
// Kept unexported so tests can inject a stub.
// Note: /exercises/random must be registered before /exercises/{id} so chi
// matches the literal segment first.
func registerExerciseRoutes(r chi.Router, q exerciseQuerier) {
	r.Get("/exercises", handleListExercises(q))
	r.Get("/exercises/random", handleRandomExercise(q))
	r.Get("/exercises/{id}", handleExerciseByID(q))
	r.Get("/body-parts", handleBodyParts(q))
	r.Get("/equipment", handleEquipment(q))
	r.Get("/muscles", handleMuscles(q))
	r.Post("/exercises/batch", handleExercisesBatch(q))
}

// handleListExercises godoc
//
//	@Summary		List Exercises
//	@Description	Returns a paginated list of exercises. All filter parameters are optional and combinable.
//	@Tags			fitness-exercises
//	@Produce		json
//	@Param			body_part	query		string	false	"Body part filter (e.g. chest). Valid values via /v1/health/body-parts"
//	@Param			equipment	query		string	false	"Equipment filter (e.g. barbell). Valid values via /v1/health/equipment"
//	@Param			muscle		query		string	false	"Muscle filter (e.g. biceps). Valid values via /v1/health/muscles"
//	@Param			search		query		string	false	"Full-text search on exercise name"
//	@Param			page		query		integer	false	"Page number (default 1)"
//	@Param			per_page	query		integer	false	"Results per page, 1–100 (default 20)"
//	@Success		200			{object}	httpx.Response[ExerciseList]
//	@Failure		400			{object}	httpx.ErrorResponse
//	@Failure		500			{object}	httpx.ErrorResponse
//	@Router			/v1/health/exercises [get]
func handleListExercises(q exerciseQuerier) http.HandlerFunc {
	return httpx.HandleGet(func(ctx context.Context, params ListParams) (ExerciseList, error) {
		return q.List(ctx, params)
	}, ListParams{Page: 1, PerPage: 20})
}

// handleRandomExercise godoc
//
//	@Summary		Random Exercise
//	@Description	Returns a single randomly selected exercise, accepting the same filters as the list endpoint.
//	@Tags			fitness-exercises
//	@Produce		json
//	@Param			body_part	query		string	false	"Body part filter (e.g. chest)"
//	@Param			equipment	query		string	false	"Equipment filter (e.g. barbell)"
//	@Param			muscle		query		string	false	"Muscle filter (e.g. biceps)"
//	@Param			search		query		string	false	"Full-text search on exercise name"
//	@Success		200			{object}	httpx.Response[Exercise]
//	@Failure		404			{object}	httpx.ErrorResponse
//	@Failure		500			{object}	httpx.ErrorResponse
//	@Router			/v1/health/exercises/random [get]
func handleRandomExercise(q exerciseQuerier) http.HandlerFunc {
	return httpx.HandleGet(func(ctx context.Context, params ListParams) (Exercise, error) {
		return q.Random(ctx, params)
	}, ListParams{Page: 1, PerPage: 20})
}

// handleExerciseByID godoc
//
//	@Summary		Exercise by ID
//	@Description	Returns a single exercise by its numeric ID.
//	@Tags			fitness-exercises
//	@Produce		json
//	@Param			id	path		integer	true	"Numeric exercise ID"
//	@Success		200	{object}	httpx.Response[Exercise]
//	@Failure		400	{object}	httpx.ErrorResponse
//	@Failure		404	{object}	httpx.ErrorResponse
//	@Failure		500	{object}	httpx.ErrorResponse
//	@Router			/v1/health/exercises/{id} [get]
func handleExerciseByID(q exerciseQuerier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		raw := chi.URLParam(r, "id")
		id, err := strconv.Atoi(raw)
		if err != nil || id <= 0 {
			httpx.Error(w, http.StatusBadRequest, "bad_request", "id must be a positive integer")
			return
		}

		exercise, err := q.Get(r.Context(), id)
		if err != nil {
			if se, ok := errors.AsType[*svcerr.Error](err); ok {
				httpx.Error(w, svcerr.HTTPStatus(se), se.Code, se.Message)
				return
			}
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "failed to fetch exercise")
			return
		}

		httpx.JSON(w, http.StatusOK, exercise)
	}
}

// handleBodyParts godoc
//
//	@Summary		List Body Parts
//	@Description	Returns a sorted list of all distinct body part values.
//	@Tags			fitness-exercises
//	@Produce		json
//	@Success		200	{object}	httpx.Response[StringList]
//	@Failure		500	{object}	httpx.ErrorResponse
//	@Router			/v1/health/body-parts [get]
func handleBodyParts(q exerciseQuerier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := q.BodyParts(r.Context())
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "failed to fetch body parts")
			return
		}
		httpx.JSON(w, http.StatusOK, result)
	}
}

// handleEquipment godoc
//
//	@Summary		List Equipment
//	@Description	Returns a sorted list of all distinct equipment types.
//	@Tags			fitness-exercises
//	@Produce		json
//	@Success		200	{object}	httpx.Response[StringList]
//	@Failure		500	{object}	httpx.ErrorResponse
//	@Router			/v1/health/equipment [get]
func handleEquipment(q exerciseQuerier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := q.Equipment(r.Context())
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "failed to fetch equipment")
			return
		}
		httpx.JSON(w, http.StatusOK, result)
	}
}

// handleMuscles godoc
//
//	@Summary		List Muscles
//	@Description	Returns a sorted list of all distinct muscle names.
//	@Tags			fitness-exercises
//	@Produce		json
//	@Success		200	{object}	httpx.Response[StringList]
//	@Failure		500	{object}	httpx.ErrorResponse
//	@Router			/v1/health/muscles [get]
func handleMuscles(q exerciseQuerier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := q.Muscles(r.Context())
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "failed to fetch muscles")
			return
		}
		httpx.JSON(w, http.StatusOK, result)
	}
}

// handleExercisesBatch godoc
//
//	@Summary		Batch Get Exercises
//	@Description	Fetches up to 50 exercises by numeric IDs. Non-existent IDs are silently skipped.
//	@Tags			fitness-exercises
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchGetRequest	true	"List of exercise IDs"
//	@Success		200		{object}	httpx.Response[BatchExerciseResponse]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/health/exercises/batch [post]
func handleExercisesBatch(q exerciseQuerier) http.HandlerFunc {
	return httpx.HandleBatch(
		func(ctx context.Context, req BatchGetRequest) (BatchExerciseResponse, error) {
			results, err := q.GetBatch(ctx, req.IDs)
			if err != nil {
				return BatchExerciseResponse{}, err
			}
			return BatchExerciseResponse{Results: results}, nil
		},
	)
}
