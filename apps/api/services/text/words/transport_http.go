package words

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
	"requiems-api/platform/svcerr"
)

// RegisterRoutes mounts words handlers on the given router.
// Paths are relative to the parent mount point.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/words/random", func(w http.ResponseWriter, r *http.Request) {
		wrd, err := svc.Random(r.Context())
		if err != nil {
			if se, ok := errors.AsType[*svcerr.Error](err); ok {
				httpx.Error(w, svcerr.HTTPStatus(se), se.Code, se.Message)
				return
			}
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "internal server error")
			return
		}

		httpx.JSON(w, http.StatusOK, wrd)
	})

	r.Get("/dictionary/{word}", func(w http.ResponseWriter, r *http.Request) {
		word := chi.URLParam(r, "word")

		entry, err := svc.Define(word)
		if err != nil {
			if se, ok := errors.AsType[*svcerr.Error](err); ok {
				httpx.Error(w, svcerr.HTTPStatus(se), se.Code, se.Message)
				return
			}
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "internal server error")
			return
		}

		httpx.JSON(w, http.StatusOK, entry)
	})

	r.Post("/words/batch", httpx.HandleBatch(
		func(ctx context.Context, req BatchRequest) (BatchResponse, int, error) {

			resp, err := svc.BatchDefine(ctx, req)

			if err != nil {
				return BatchResponse{}, 0, err
			}

			return resp, len(req.Items), nil
		},
	))
}
