package emoji

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// SearchRequest holds the query parameter for emoji search.
type SearchRequest struct {
	Query string `query:"q" validate:"required,min=1,max=100"`
}

// RegisterRoutes mounts emoji handlers on the given router.
// Paths are relative to the parent mount point (/v1/entertainment).
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/emoji/random", handleRandomEmoji(svc))
	r.Get("/emoji/search", handleEmojiSearch(svc))
	r.Get("/emoji/{name}", handleEmojiByName(svc))
}

// handleRandomEmoji godoc
//
//	@Summary		Get Random Emoji
//	@Description	Returns a randomly selected emoji with its full metadata.
//	@Tags			emoji
//	@Produce		json
//	@Success		200	{object}	httpx.Response[Emoji]
//	@Router			/v1/entertainment/emoji/random [get]
func handleRandomEmoji(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpx.JSON(w, http.StatusOK, svc.Random())
	}
}

// handleEmojiSearch godoc
//
//	@Summary		Search Emoji
//	@Description	Search emojis whose name or category contains the query string (case-insensitive).
//	@Tags			emoji
//	@Produce		json
//	@Param			q	query		string	true	"Search term"
//	@Success		200	{object}	httpx.Response[List]
//	@Failure		400	{object}	httpx.ErrorResponse
//	@Router			/v1/entertainment/emoji/search [get]
func handleEmojiSearch(svc *Service) http.HandlerFunc {
	return httpx.HandleGet(func(ctx context.Context, req SearchRequest) (List, error) {
		return svc.Search(req.Query), nil
	})
}

// handleEmojiByName godoc
//
//	@Summary		Get Emoji by Name
//	@Description	Returns a specific emoji by its CLDR snake_case name (case-insensitive).
//	@Tags			emoji
//	@Produce		json
//	@Param			name	path		string	true	"Emoji name (e.g. grinning_face, thumbs_up)"
//	@Success		200		{object}	httpx.Response[Emoji]
//	@Failure		404		{object}	httpx.ErrorResponse
//	@Router			/v1/entertainment/emoji/{name} [get]
func handleEmojiByName(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := strings.ToLower(chi.URLParam(r, "name"))
		e, ok := svc.GetByName(name)
		if !ok {
			httpx.Error(w, http.StatusNotFound, "not_found", "emoji not found")
			return
		}

		httpx.JSON(w, http.StatusOK, e)
	}
}
