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

// exerciseQuerier is the interface consumed by HTTP handlers, allowing stub
// injection in tests without a live database.
type exerciseQuerier interface {
	List(ctx context.Context, p ListParams) (ExerciseList, error)
	Get(ctx context.Context, id int) (Exercise, error)
	Random(ctx context.Context, p ListParams) (Exercise, error)
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
	r.Get("/exercises", func(w http.ResponseWriter, r *http.Request) {
		params := ListParams{Page: 1, PerPage: 20}

		if err := httpx.BindQuery(r, &params); err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		result, err := q.List(r.Context(), params)
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "failed to fetch exercises")
			return
		}

		httpx.JSON(w, http.StatusOK, result)
	})

	r.Get("/exercises/random", func(w http.ResponseWriter, r *http.Request) {
		params := ListParams{Page: 1, PerPage: 20}

		if err := httpx.BindQuery(r, &params); err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		exercise, err := q.Random(r.Context(), params)
		if err != nil {
			var se *svcerr.Error
			if errors.As(err, &se) {
				httpx.Error(w, svcerr.HTTPStatus(se), se.Code, se.Message)
				return
			}
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "failed to fetch random exercise")
			return
		}

		httpx.JSON(w, http.StatusOK, exercise)
	})

	r.Get("/exercises/{id}", func(w http.ResponseWriter, r *http.Request) {
		raw := chi.URLParam(r, "id")
		id, err := strconv.Atoi(raw)
		if err != nil || id <= 0 {
			httpx.Error(w, http.StatusBadRequest, "bad_request", "id must be a positive integer")
			return
		}

		exercise, err := q.Get(r.Context(), id)
		if err != nil {
			var se *svcerr.Error
			if errors.As(err, &se) {
				httpx.Error(w, svcerr.HTTPStatus(se), se.Code, se.Message)
				return
			}
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "failed to fetch exercise")
			return
		}

		httpx.JSON(w, http.StatusOK, exercise)
	})

	r.Get("/body-parts", func(w http.ResponseWriter, r *http.Request) {
		result, err := q.BodyParts(r.Context())
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "failed to fetch body parts")
			return
		}
		httpx.JSON(w, http.StatusOK, result)
	})

	r.Get("/equipment", func(w http.ResponseWriter, r *http.Request) {
		result, err := q.Equipment(r.Context())
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "failed to fetch equipment")
			return
		}
		httpx.JSON(w, http.StatusOK, result)
	})

	r.Get("/muscles", func(w http.ResponseWriter, r *http.Request) {
		result, err := q.Muscles(r.Context())
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "failed to fetch muscles")
			return
		}
		httpx.JSON(w, http.StatusOK, result)
	})
}
